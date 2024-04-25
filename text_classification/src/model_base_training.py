from datasets import load_dataset
from transformers import DistilBertForSequenceClassification, Trainer, TrainingArguments
from .preprocessing.utility import tokenize_function

# Initialize DistilBERT model
model = DistilBertForSequenceClassification.from_pretrained(
    "distilbert-base-uncased", num_labels=7
)

# Load AG News dataset
dataset = load_dataset("ag_news")


tokenized_datasets = dataset.map(tokenize_function, batched=True)

# Training arguments
training_args = TrainingArguments(
    output_dir="models/results_v1",
    evaluation_strategy="epoch",
    learning_rate=2e-5,
    per_device_train_batch_size=16,
    per_device_eval_batch_size=16,
    num_train_epochs=4,
    weight_decay=0.01,
)

# Initialize Trainer
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=tokenized_datasets["train"],
    eval_dataset=tokenized_datasets["test"],
)

# Train the model
trainer.train()
trainer.evaluate()
model.save_pretrained("../models/v1")
