from confluent_kafka import Consumer, Producer
import json
from transformers import (
    DistilBertTokenizerFast,
    DistilBertForSequenceClassification,
)
import torch
import yake

# Initializing the YAKE instance
yake_kw = yake.KeywordExtractor()

# Check if CUDA is available
if torch.cuda.is_available():
    device = torch.device("cuda")
    print("Using GPU:", torch.cuda.get_device_name(0))
else:
    device = torch.device("cpu")
    print("Using CPU")

# Load the trained model
model_path = "models/v3"
model = DistilBertForSequenceClassification.from_pretrained(
    model_path, num_labels=7
).eval()

# Initialize the tokenizer
tokenizer = DistilBertTokenizerFast.from_pretrained("distilbert-base-uncased")

# Consume from "news-input"
consumer = Consumer(
    {
        "bootstrap.servers": "localhost:9092",
        "group.id": "news-consumer-group",
        "auto.offset.reset": "earliest",
        "api.version.request": "false"
    }
)

consumer.subscribe(["news-input"])

# Produce to "news-output"
producer = Producer({"bootstrap.servers": "localhost:9092"})

categories_dict = {
    0: "World",
    1: "Sports",
    2: "Business",
    3: "Science",
    4: "Politics",
    5: "Technology",
    6: "Entertainment",
}


def delivery_report(err, msg):
    if err is not None:
        print("Message delivery failed: {}".format(err))
    else:
        print("Message delivered to {} [{}]".format(msg.topic(), msg.partition()))


while True:
    msg = consumer.poll(1.0)
    if msg is None:
        continue
    if msg.error():
        print("Consumer error: {}".format(msg.error()))
        continue

    message = msg.value().decode("utf-8")
    try:
        data = json.loads(message)
        # Now you can access the data like a normal dictionary
        text = data["Title"] + " " + data["Description"]
        print(text)

        # Tokenize the sample text
        tokens = tokenizer.encode(text, add_special_tokens=True)

        # Count the tokens
        num_tokens = len(tokens)
        print("Number of tokens:", num_tokens)

        inputs = tokenizer(
            text, return_tensors="pt", padding=True, truncation=True, max_length=512
        )

        # Get predictions
        with torch.no_grad():
            outputs = model(**inputs)
            predictions = torch.nn.functional.softmax(outputs.logits, dim=1)

        # Get the highest probability label
        predicted_label = torch.argmax(predictions).item()

        # Print the predicted label
        print("Predicted Label:", categories_dict[predicted_label])
        # Extracting keywords
        yake_keywords = yake_kw.extract_keywords(text)
        keywords = []
        for item in yake_keywords:
            keywords.append(item[0])
        print(keywords)

        data["Category"] = categories_dict[predicted_label]
        data["Keywords"] = keywords
        # Serialize the data dictionary to a JSON string
        serialized_data = json.dumps(data)

        # Encode the JSON string to bytes
        encoded_data = serialized_data.encode("utf-8")

        try:
            print(f"Producing message to news-output: {serialized_data}")
            producer.produce(
                "news-output", value=encoded_data, callback=delivery_report
            )
            producer.poll(0)
            producer.flush()
        except Exception as e:
            print(f"Error in producing message: {e}")

    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")

consumer.close()