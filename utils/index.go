package utils

// Index is an inverted index. It maps tokens to document IDs.
type Index map[string][]int

// add adds documents to the Index.
func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		//for each return the tokens and their corresponding document IDs
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// search queries the Index for the given text.
func (idx Index) Search(text string) []int {
	var r []int
	//for each token in the text, find the corresponding document IDs
	for _, token := range analyze(text) {
		//if the token exists in the index, find the intersection of the document IDs
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}
