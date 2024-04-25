from transformers import DistilBertForSequenceClassification, Trainer, TrainingArguments
from preprocessing.bbc_preprocessing import bbc_train_dataset, bbc_val_dataset
from preprocessing.utility import EarlyStoppingCallback



# Load the model trained on AG News
model = DistilBertForSequenceClassification.from_pretrained("models/v2", num_labels=7)

# Initialize the early stopping callback
early_stopping_callback = EarlyStoppingCallback(
    early_stopping_patience=1
)  # Stop if no improvement after 1 epoch

# Training arguments
training_args = TrainingArguments(
    output_dir="models/results_v3",
    evaluation_strategy="epoch",
    learning_rate=5e-6,
    per_device_train_batch_size=16,
    per_device_eval_batch_size=16,
    num_train_epochs=5,
    weight_decay=0.01,
)

# Initialize Trainer for fine-tuning with the early stopping callback
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=bbc_train_dataset,
    eval_dataset=bbc_val_dataset,
    callbacks=[early_stopping_callback],
)


# Fine-tune the model
trainer.train()

# Evaluate the model
trainer.evaluate()

model.save_pretrained("models/v3")
