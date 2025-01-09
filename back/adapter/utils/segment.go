package utils

import (
	"regexp"
	"strings"
)

type segmentation struct {
	WordCorpus []string
}

type Segmentation interface {
	Segmentation(sentence string) []string
}

func NewSegmentation(corpusDir string) (Segmentation, error) {
	wordCorpus, err := LoadDict(corpusDir)
	if err != nil {
		return nil, err
	}
	return &segmentation{wordCorpus}, nil
}

func (t *segmentation) Segmentation(sentence string) []string {
	re := regexp.MustCompile(`[.,!?\s+]`)
	sentence = re.ReplaceAllString(sentence, "")
	sentence = strings.TrimSpace(sentence)

	i := 0
	lenSentence := len(sentence)
	resultSegment := []string{}
	for i < lenSentence {
		maxWord := ""
		for j := i + 1; j <= lenSentence; j++ {
			word := sentence[i:j]
			for _, c := range t.WordCorpus {
				if c == word && len(word) > len(maxWord) {
					maxWord = word
					break
				}
			}
		}
		if maxWord == "" {
			i++
		} else {
			i += len(maxWord)
			resultSegment = append(resultSegment, maxWord)
		}
	}
	return resultSegment
}
