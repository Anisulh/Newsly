# Project Title
## Description

The final model(v4) from these files is based on distilBERT NLP model and has been fine-tuned to be able to categorize texts to fall into the following categories: Science, Technology, Entertainment, World, Sports, Business, Politics.

To achieve this final model, it was testes on various labelled datasets of news articles (BBC, CNN, AG News, Text-Classification)

### Models
v1: AG News Dataset
v2: v1 + Text Classification Dataset
v3: v2 + BBC News Dataset
v4: v3 + CNN Dataset

Due to distilBERT input size being 512 tokens, all texts (except for AG News Dataset) were summarized using BART to reduce the total text size without truncating the texts too much. However truncating was still required as the BART model had a limit of 1024 token input size.

Below are the categories in each dataset and what they map to. Then a global map is implemented
```bash
    # AG News Dataset Classes

    World = 0
    Sports = 1
    Business = 2
    Science/Technology = 3

    # Text Document Classification Dataset Classes
    
    Politics = 0
    Sport = 1
    Technology = 2
    Entertainment = 3
    Business = 4

    # Relabelled Text Document Classification Dataset Classes
    Sports = 1    # common
    Business = 2  # common
    Politics = 4
    Technology = 5
    Entertainment = 6

    # BBC Dataset Classes
    Sports = "sport"    
    Business = "business" 
    Politics = "politics"
    Technology = "tech"
    Entertainment = "entertainment"


    # Relabelled BBC Dataset Classes
    Sports = 1    # common
    Business = 2  # common
    Politics = 4  # common
    Technology = 5  # common
    Entertainment = 6 # common


    # New Dataset Classes
    World = 0
    Sports = 1
    Business = 2
    Science = 3  # Science/Technology renamed to science since technology is a separate class
    Politics = 4 
    Technology = 5 
    Entertainment = 6 

```
As such the model is expected to differentiate science from technology and politics from world.

The model was further trained on the BBC News Dataset, which had a similar classes to the Text Classification Dataset, to improve it classification abilities on larger texts.
`{'train_runtime': 758.0574, 'train_samples_per_second': 11.741, 'train_steps_per_second': 0.739, 'train_loss': 0.10026595762797764, 'epoch': 5.0}` 

## Discussion


## File Structure
```bash
.
├── data
│   ├── ag_news.csv                # AG News dataset for initial training
│   └── custom_dataset.csv         # Custom dataset for fine-tuning
├── models
│   ├── saved_model                # Initial model trained on AG News dataset
│   └── saved_model_fine-tuned      # Model fine-tuned on the custom dataset
├── src
│   ├── preprocessing
│   │    ├── bbc_preprocessing.py
│   │    └── text_classification_preprocessing.py   # Script for data preprocessing
│   ├── model_base_training.py      # Script for training distilBERT on AG News Dataset -> model(v1)
│   ├── model_v1_training.py        # Script for training model(v1) on Text Classification Dataset -> model(v2)
│   └── model_v2_training.py        # Script for training model(v2) on BBC Dataset -> model(v3)
```
## Setup and Installation

```bash
  pip install -r requirements.txt
```

## Usage


```bash
  python src/model_training.py
  python src/model_fine-tuning.py
```

## Data Description

  **The AG News dataset (HuggingFace):** used for the initial training of the model. Contains news articles categorized into different classes.    
 **A Text Document Classification dataset (source: [Kaggle](https://www.kaggle.com/datasets/sunilthite/text-document-classification-dataset?resource=download), `data/df_file.csv`):** for fine-tuning the model further. Contains text data categorized into five classes.

## Models

  `v1`: This is the DistilBERT model trained on the AG News dataset.          
  `v2`: This is the DistilBERT model after further fine-tuning on the Text Document Classification dataset.

## Scripts

 ` data_preprocessing.py`: Used for preprocessing the datasets before feeding them into the model.
  `model_training.py`: Contains code for training the model on the AG News dataset.
  `model_fine-tuning.py`: Contains code for fine-tuning the trained model on the Text Document Classification dataset.

