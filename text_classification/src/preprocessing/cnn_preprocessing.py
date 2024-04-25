import pandas as pd
from sklearn.model_selection import train_test_split
from utility import (
    tokenize_function,
    batch_size,
    process_in_batches,
    CustomDataset,
    remove_pattern_from_text,
    remove_rows_by_value,
    combine_columns,
)
import os

# Load cnn dataset
cnn_df = pd.read_csv("data/CNN_dataset.csv")

# Re-map labels in cnn dataset
cnn_label_remap = {
    "news": 0,  # World
    "politics": 4,  # Politics
    "sport": 2,  # Sport (common)
    "vr": 5,  # Technology
    "entertainment": 6,  # Entertainment
    "business": 3,  # Business (common)
    "health": 4,  # Science
}

# Check if the summarized data already exists
summarized_data_path = "data/CNN_summarized_dataset.csv"
if os.path.exists(summarized_data_path):
    cnn_df = pd.read_csv(summarized_data_path)
else:
    # Define patterns to remove and column names
    patterns_to_remove = {"Headline": r" - CNN", "Article text": r"\(CNN[^)]*\)"}
    values_to_remove = ["style", "travel"]
    text_columns_to_combine = [
        "Headline",
        "Description",
        "Second headline",
        "Article text",
    ]
    columns_to_keep = ["label", "summary"]

    # Processing steps
    for value in values_to_remove:
        cnn_df = remove_rows_by_value(cnn_df, "Category", value)

    for column, pattern in patterns_to_remove.items():
        cnn_df = remove_pattern_from_text(cnn_df, column, pattern)

    cnn_df["label"] = cnn_df["Category"].map(cnn_label_remap)

    # Combining text columns
    cnn_df["text"] = combine_columns(cnn_df, text_columns_to_combine)

    # Summarize texts
    cnn_df["summary"] = process_in_batches(cnn_df["text"].tolist(), batch_size)

    # Keeping only necessary columns
    cnn_df = cnn_df[columns_to_keep]

    cnn_df.to_csv(summarized_data_path, index=False)


# Split dataset into training and validation sets
cnn_train_df, cnn_val_df = train_test_split(cnn_df, test_size=0.2)


cnn_train_encodings = tokenize_function(cnn_train_df["summary"].tolist())
cnn_val_encodings = tokenize_function(cnn_val_df["summary"].tolist())


cnn_train_dataset = CustomDataset(cnn_train_encodings, cnn_train_df["label"].tolist())
cnn_val_dataset = CustomDataset(cnn_val_encodings, cnn_val_df["label"].tolist())
