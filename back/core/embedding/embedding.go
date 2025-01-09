package embedding

type Embedding struct {
	Word   string
	Dim    int
	Vector []float64
}

type Embeddings []Embedding

func (e Embeddings) Find(word string) (*Embedding, bool) {
	for i := range e {
		if e[i].Word == word {
			return &e[i], true
		}
	}
	return nil, false
}
