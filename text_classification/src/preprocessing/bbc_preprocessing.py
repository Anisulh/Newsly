import pandas as pd
from sklearn.model_selection import train_test_split
import os
from .utility import (
    combine_columns,
    tokenize_function,
    batch_size,
    process_in_batches,
    CustomDataset,
)

# Load bbc dataset
bbc_df = pd.read_csv("data/BBC_dataset.csv", sep="\t")

# Re-map labels in bbc dataset
bbc_label_remap = {
    "politics": 4,  # Politics
    "sport": 1,  # Sport (common)
    "tech": 5,  # Technology
    "entertainment": 6,  # Entertainment
    "business": 2,  # Business (common)
}

bbc_df["label"] = bbc_df["category"].map(bbc_label_remap)

# Check if the summarized data already exists
summarized_data_path = "data/BBC_summarized_dataset.csv"

if os.path.exists(summarized_data_path):
    bbc_df = pd.read_csv(summarized_data_path)
    #bbc_df["label"] = bbc_df["category"].map(bbc_label_remap)
else:
    text_columns_to_combine = ["title", "content"]
    columns_to_keep = ["label", "summary"]

    # Combine title and content
    bbc_df["text"] = combine_columns(bbc_df, text_columns_to_combine)

    # Apply batch processing to summarize texts
    bbc_df["summary"] = process_in_batches(bbc_df["text"].tolist(), batch_size)
    
    
    # content filename title category text summary label

    bbc_df = bbc_df[columns_to_keep]

    # summary label

    # Save the summarized data
    bbc_df.to_csv(summarized_data_path, index=False)


# Split dataset into training and validation sets
bbc_train_df, bbc_val_df = train_test_split(bbc_df, test_size=0.2)


bbc_train_encodings = tokenize_function(bbc_train_df["summary"].tolist())
bbc_val_encodings = tokenize_function(bbc_val_df["summary"].tolist())


bbc_train_dataset = CustomDataset(bbc_train_encodings, bbc_train_df["label"].tolist())
bbc_val_dataset = CustomDataset(bbc_val_encodings, bbc_val_df["label"].tolist())
