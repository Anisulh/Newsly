import pandas as pd
from sklearn.model_selection import train_test_split
import os
from .utility import tokenize_function, batch_size, process_in_batches, CustomDataset

# Load custom dataset
df = pd.read_csv("data/text-classification-data.csv")

# Re-map labels in custom dataset
label_remap = {
    0: 4,  # Politics
    1: 2,  # Sport (common)
    2: 5,  # Technology
    3: 6,  # Entertainment
    4: 3,  # Business (common)
}
df["label"] = df["Label"].map(label_remap)

# Check if the summarized data already exists
summarized_data_path = "data/summarized_text_classification.csv"
if os.path.exists(summarized_data_path):
    df = pd.read_csv(summarized_data_path)
else:
    df["text"] = df["Text"]

    # Apply batch processing to summarize texts
    df["summary"] = process_in_batches(df["text"].tolist(), batch_size)

    # List of columns to keep
    columns_to_keep = ["label", "summary"]

    df = df[columns_to_keep]

    # Save the summarized data
    df.to_csv(summarized_data_path, index=False)
# Split dataset into training and validation sets
train_df, val_df = train_test_split(df, test_size=0.2)


train_encodings = tokenize_function(train_df["summary"].tolist())
val_encodings = tokenize_function(val_df["summary"].tolist())


train_dataset = CustomDataset(train_encodings, train_df["label"].tolist())
val_dataset = CustomDataset(val_encodings, val_df["label"].tolist())
