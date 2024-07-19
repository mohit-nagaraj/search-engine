package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// document represents a Wikipedia abstract dump document.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// loadDocuments loads a Wikipedia abstract dump and returns a slice of documents.
// Dump example: https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
func LoadDocuments(path string) ([]document, error) {
	// Open the dump file present at the given path.
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Close the file when the function returns.
	defer f.Close()
	// Create a gzip reader to read the compressed XML file.
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	// Create an XML decoder to decode the XML file.
	dec := xml.NewDecoder(gz)
	//slice of documents
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	//dec.Decode reads the next XML-encoded value from its input and stores it in the value pointed to by v.
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	// Assign an ID to each document.
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
