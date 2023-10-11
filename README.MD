# Zhang's Metric in Golang

## Overview
This project is a personal journey in learning how to read and work with sparse matrices in the NPZ format using Golang. It includes a simple Gin API that scores Zhang's Metric based on a list input in JSON format.

## Getting Started

### Prerequisites
- Go 1.21 

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/HeyLucasLeao/zhangmetric
   ```
2. Install Go packages
    ```sh 
    go get .
    ```

### Usage
1. Setting Dataset

    The Kaggle groceries dataset is a viable option for workaround in this project. However, any historical data that can be transformed into a CSR Matrix can also be utilized. For more information, see the eda_example.ipynb.

    Groceries Dataset: https://www.kaggle.com/datasets/irfanasrullah/groceries

2. Setting .env

    After having the CSR Matrix witch respective ordered column products, insert the filenames inside a .env

3. Request on API
    
    With running the project, send a GET request to http://localhost:8080/score, with a JSON body having the names of product you want to score. Example: 
    
    ```
    [
	"whole milk",
	"cereals",
	"tea",
	"yogurt",
	"waffles"
	]
    `````

### Contact
Lucas Leão - heylucasleao@gmail.com