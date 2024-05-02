package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	utils "github.com/NoobPeen/text-search-engine/utils"
)

func main() {
	var dumpPath, query string
	// flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	// flag.StringVar(&query, "q", "Small wild cat", "search query")
	// flag.Parse()

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "", "search query")

	// Use a generic prompt as the default
	// flag.StringVar(&query, "q", "Enter search query here", "search query")

	flag.Parse()

	// Check if the query is still the default or empty and prompt user if needed
	if query == "" || query == "Enter search query here" {
		fmt.Print("Please enter your search query: ")
		_, err := fmt.Scanln(&query)
		if err != nil {
			log.Fatalf("Failed to read search query: %v", err)
		}
	}

	log.Println("Running Full Text Search")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
