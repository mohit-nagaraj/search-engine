package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/mohit-nagaraj/search-engine/utils"
)

func main() {
	var dumpPath, query string
	//flag is used to parse the command line arguments or use the default values
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()
	log.Println("Running Full Text Search")

	//timer to calculate the time taken to load the documents
	start := time.Now()
	//load the documents from the dump path
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	//timer to calculate the time taken to build the index
	start = time.Now()
	//create an index of the documents
	idx := make(utils.Index)
	//add the documents to the index
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	//timer to calculate the time taken to search the query
	start = time.Now()
	//search the query in the index
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		//print the matched documents
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
