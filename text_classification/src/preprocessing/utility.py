from tqdm import tqdm
from transformers import (
    BartForConditionalGeneration,
    BartTokenizer,
    DistilBertTokenizerFast,
)
import torch
from transformers import TrainerCallback


class EarlyStoppingCallback(TrainerCallback):
    def __init__(self, early_stopping_patience):
        self.early_stopping_patience = early_stopping_patience
        self.best_loss = None
        self.patience_counter = 0

    def on_evaluate(self, args, state, control, **kwargs):
        eval_loss = kwargs["metrics"]["eval_loss"]
        if self.best_loss is None or eval_loss < self.best_loss:
            self.best_loss = eval_loss
            self.patience_counter = 0
        else:
            self.patience_counter += 1
            if self.patience_counter >= self.early_stopping_patience:
                control.should_training_stop = True
                print("Early stopping triggered!")

# Check if CUDA is available
if torch.cuda.is_available():
    device = torch.device("cuda")
    print("Using GPU:", torch.cuda.get_device_name(0))
else:
    device = torch.device("cpu")
    print("Using CPU")

model_name = "facebook/bart-large-cnn"
summarizer_tokenizer = BartTokenizer.from_pretrained(model_name)
summarizer_model = BartForConditionalGeneration.from_pretrained(model_name)
summarizer_model.to(device)


def summarize_batch(texts, max_input_length=1024, max_output_length=450):
    try:
        inputs = summarizer_tokenizer(
            texts,
            max_length=max_input_length,
            return_tensors="pt",
            truncation=True,
            padding=True,
        )

        # Move input tensors to the same device as the model (GPU if available)
        inputs = {k: v.to(device) for k, v in inputs.items()}

        # Generate summaries
        summary_ids = summarizer_model.generate(
            inputs["input_ids"],
            num_beams=3,
            max_length=max_output_length,
            min_length=250,
            length_penalty=2.0,
            early_stopping=True,
            no_repeat_ngram_size=3,
        )

        # Move generated summaries back to the CPU for further processing
        summaries = [
            summarizer_tokenizer.decode(
                g, skip_special_tokens=True, clean_up_tokenization_spaces=False
            )
            for g in summary_ids.cpu()
        ]

        return summaries
    except Exception as e:
        print(f"Error during summarization: {e}")
        exit(1)


batch_size = 16  # Adjust the batch size based on system's capabilities


# Function to process data in batches
def process_in_batches(texts, batch_size):
    summaries = []
    for i in tqdm(
        range(0, len(texts), batch_size),
        desc=f"Summarizing data in batches of {batch_size}",
    ):
        batch_texts = texts[i : i + batch_size]
        batch_summaries = summarize_batch(batch_texts)
        summaries.extend(batch_summaries)
    return summaries


# Initialize tokenizer
tokenizer = DistilBertTokenizerFast.from_pretrained("distilbert-base-uncased")


# Tokenize the custom dataset
def tokenize_function(examples):
    return tokenizer(examples, padding="max_length", truncation=True)


class CustomDataset(torch.utils.data.Dataset):
    def __init__(self, encodings, labels):
        self.encodings = encodings
        self.labels = labels

    def __getitem__(self, idx):
        item = {key: torch.tensor(val[idx]) for key, val in self.encodings.items()}
        item["labels"] = torch.tensor(self.labels[idx])
        return item

    def __len__(self):
        return len(self.labels)


def remove_rows_by_value(df, column_name, value):
    """Remove rows where the column has the specific value."""
    return df[df[column_name] != value]


def remove_pattern_from_text(df, column_name, pattern):
    """Remove a regex pattern from a specified column in the dataframe."""
    df[column_name] = df[column_name].str.replace(pattern, "", regex=True)
    return df


def combine_columns(df, column_names):
    """Combine multiple text columns into a single text column."""
    return df[column_names].agg(" ".join, axis=1)