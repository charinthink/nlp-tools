package utils

import (
	"github.com/charinthink/thai-sentence/back/core/algorithm/consinesimilarity"
	"github.com/charinthink/thai-sentence/back/core/embedding"
	"github.com/charinthink/thai-sentence/back/core/neighbor"
	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
)

type similarity struct {
	Embs embedding.Embeddings
}

type Similarity interface {
	Similarity(wording string, limitSimilarity int) neighbor.Neighbors
}

func NewSimilarity(modelName string, parallelConfig parallelconfig.Config) (Similarity, error) {
	result, err := LoadModel(modelName, parallelConfig)
	if err != nil {
		return nil, err
	}
	return &similarity{result}, nil
}

func (t *similarity) Similarity(wording string, limitSimilarity int) neighbor.Neighbors {
	embedWording, _ := t.Embs.Find(wording)
	wordSimilaritys := make(neighbor.Neighbors, limitSimilarity)
	low := .0
	for _, v := range t.Embs {
		if len(v.Vector) < 2 {
			continue
		}
		score := consinesimilarity.Cal(v.Vector, embedWording.Vector)
		if score > low {
			tempValue := neighbor.Neighbor{Word: v.Word, Similarity: score}
			for i := range wordSimilaritys {
				if score > wordSimilaritys[i].Similarity {
					// Shift element
					copy(wordSimilaritys[i+1:], wordSimilaritys[i:])
					wordSimilaritys[i] = tempValue
					break
				}
			}
			low = wordSimilaritys[len(wordSimilaritys)-1].Similarity
		}
	}
	return wordSimilaritys
}
