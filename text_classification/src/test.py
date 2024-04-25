from transformers import (
    DistilBertTokenizerFast,
    DistilBertForSequenceClassification,
    BartTokenizer,
    BartForConditionalGeneration,
)
import torch
import time
start_time = time.time()

# Check if CUDA is available
if torch.cuda.is_available():
    device = torch.device("cuda")
    print("Using GPU:", torch.cuda.get_device_name(0))
else:
    device = torch.device("cpu")
    print("Using CPU")

# Load the trained model
model_path = "models/v3"
model = DistilBertForSequenceClassification.from_pretrained(model_path, num_labels=7)

# Initialize the tokenizer
tokenizer = DistilBertTokenizerFast.from_pretrained("distilbert-base-uncased")

# Sample text to classify
sample_text = """WASHINGTON — Dramatic, newly released video shot by a Jan. 6 rioter shows the confrontation between two Republican members of Congress and Capitol rioters trying to breach the main door into the chamber of the U.S. House of Representatives. The videos show rioters staring down the barrels of guns through broken glass as they tried to force their way inside.

The eight-minute video, shot by Capitol rioter Damon Beckley and introduced as evidence ahead of his sentencing, was released to the media coalition late Friday in response to a request from NBC News. Federal prosecutors are seeking more than three years in federal prison for Beckley, and his sentencing has now been rescheduled for Feb. 9.

Rep. Troy Nehls, R-Texas, and then-Rep. Markwayne Mullin, R-Okla., can be seen speaking with rioters through the broken glass.

“You ought to be ashamed of yourself!” Nehls says.

“We’re coming in one way or another!” one rioter says.

“I’ve been in law enforcement in Texas for 30 years, and I’ve never had people act this way,” Nehls says. “I’m ashamed!”

Nehls has changed his tone in the three years since Jan. 6, even calling the death of Ashli Babbitt — the rioter who was shot by a Capitol Police officer as she jumped through a door near the House floor — "murder." In a previous interview with “NBC Nightly News,” Nehls explained his thinking on Jan. 6.

“I am going to stay right here with my brothers and sisters in blue,” Nehls said. “I had my Texas mask on, and he looked at me through that broken glass and he said, ‘You're from Texas, you should be with us,’ and at that point, I said, ‘No, sir, I cannot support what you’re doing, this is criminal.’”


In the video, one rioter tells the officers and representatives on the other side of the door “there’s going to be a bigger Civil War and a lot of bloodshed” if they didn’t overturn the election.

“I drove 14 hours to get here and stood in the cold for three and a half hours to find out that Mike Pence is a f------ traitor, man. And I voted for that f------ dude,” Beckley says in the video. “He could’ve done the right thing and certified those legislators, electors, and we wouldn’t be standing here with a 9 mm pointed at me right now!”

“We’re real American citizens, and we’re sick of this,” another rioter says. “We’re making it known that we’re sick of it.”

“They can only kill so many of us,” says one rioter."""


# Tokenize the sample text
tokens = tokenizer.encode(sample_text, add_special_tokens=True)


# Count the tokens
num_tokens = len(tokens)
print(tokens)
print("Number of tokens:", num_tokens)
summary = sample_text
if num_tokens > 512:
    model_name = "facebook/bart-large-cnn"
    summarizer_tokenizer = BartTokenizer.from_pretrained(model_name)
    summarizer_model = BartForConditionalGeneration.from_pretrained(model_name)
    summarizer_model.to(device)
    summarizer_model.eval()

    

    def summarize(text):
        inputs = summarizer_tokenizer(
            [text], max_length=1024, return_tensors="pt", truncation=True, padding=True
        )
        inputs = inputs.to("cuda")
        print("Summarizing")

        summary_ids = summarizer_model.generate(
            inputs["input_ids"],
            max_length=500,
            min_length=250,
            length_penalty=2.0,
            num_beams=4,
            early_stopping=True,
            no_repeat_ngram_size=3,
        )
        print("Summarized")
        summary = summarizer_tokenizer.decode(summary_ids[0], skip_special_tokens=True)
        return summary

    summary = summarize(sample_text)


print(summary)

# Tokenize the sample text
inputs = tokenizer(
    summary, return_tensors="pt", padding=True, truncation=True, max_length=512
)

# Get predictions
with torch.no_grad():
    outputs = model(**inputs)
    predictions = torch.nn.functional.softmax(outputs.logits, dim=1)

# Get the highest probability label
predicted_label = torch.argmax(predictions).item()

# Print the predicted label
print("Predicted Label:", predicted_label)
print("run time:", time.time() - start_time)