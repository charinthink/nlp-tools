package utils

import (
	"github.com/charinthink/thai-sentence/back/core/algorithm/consinesimilarity"
	"github.com/charinthink/thai-sentence/back/core/embedding"
	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
)

type detectBadWord struct {
	Embs       embedding.Embeddings
	SwearWords []string
	Segmentation Segmentation
}

type DetectBadWord interface {
	DectectBadWord(sentence string, threshold float32) []string
}

func NewDetectBadWord(pathModelEmbs, pathDictSwearWord string, segmentation Segmentation, parallelConfig parallelconfig.Config) (DetectBadWord, error) {
	swearWords, err := LoadDict(pathDictSwearWord)
	if err != nil {
		return nil, err
	}
	embs, err := LoadModel(pathModelEmbs, parallelConfig)
	if err != nil {
		return nil, err
	}
	return &detectBadWord{embs, swearWords, segmentation}, nil
}

func (t *detectBadWord) DectectBadWord(sentence string, threshold float32) []string {
	wordSegmentations := t.Segmentation.Segmentation(sentence)
	var detectWord []string

	for i := range len(wordSegmentations) {
		wordSegmentation := wordSegmentations[i]

		for x := range len(t.SwearWords) {
			swearWord := t.SwearWords[x]
			if wordSegmentation == swearWord {
				detectWord = append(detectWord, wordSegmentation)
				break
			}

			vec1, ok := t.Embs.Find(wordSegmentation)
			if !ok {
				break
			}
			vec2, ok := t.Embs.Find(swearWord)
			if !ok {
				continue
			}
			similarity := consinesimilarity.Cal(vec1.Vector, vec2.Vector)
			if similarity >= float64(threshold) {
				detectWord = append(detectWord, wordSegmentation)
				break
			}
		}
	}
	return detectWord
}
