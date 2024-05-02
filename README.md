# GO-text-search-engine

get the wikipedia dump folder from [HERE](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)

### GO-text-Search-Engine
#### Overview
This Go project is designed to efficiently preprocess, tokenize, and reverse-index Wikipedia dump data, enabling fast and relevant search capabilities. By leveraging advanced text processing techniques and optimized data structures, this project provides an effective solution for querying large datasets in real-time.

#### Features

-> Preprocessing and Tokenizing: Processes the Wikipedia dump to clean and tokenize the text, preparing it for indexing.
-> Reverse Indexing: Implements a reverse index for quick lookup of tokens, enhancing search performance.
-> Search Query Interface: Offers a command-line interface that accepts user input for search queries and retrieves relevant results efficiently

##### To run this project

1)Clone the github repo
'''
git clone https://github.com/NoobPeen/GO-text-search-engine/edit/main/README.md
'''

2)Go the directory of the project and run the following command
'''
go build
'''

3)Usage
'''
go run main.go -q "search query"
'''
