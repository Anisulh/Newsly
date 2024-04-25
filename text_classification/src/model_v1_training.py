from transformers import DistilBertForSequenceClassification, Trainer, TrainingArguments
from preprocessing.text_classification_preprocessing import (
    train_dataset,
    val_dataset,
)

# Load the model trained on AG News
model = DistilBertForSequenceClassification.from_pretrained(
    "models/v1", num_labels=7
)

# Training arguments
training_args = TrainingArguments(
    output_dir="models/results_v2",
    evaluation_strategy="epoch",
    learning_rate=2e-5,
    per_device_train_batch_size=16,
    per_device_eval_batch_size=16,
    num_train_epochs=5,
    weight_decay=0.01,
)

# Initialize Trainer for fine-tuning
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=train_dataset,
    eval_dataset=val_dataset,
)

# Fine-tune the model
trainer.train()

# Evaluate the model
trainer.evaluate()

model.save_pretrained("models/v2")
