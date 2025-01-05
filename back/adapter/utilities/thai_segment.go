package utilities

import (
	"regexp"
	"strings"
)

func ThaiSegmentation(corpusDir, sentence string) ([]string, error) {
	wordCorpus, err := LoadDict(corpusDir)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`[.,!?\s+]`)
	sentence = re.ReplaceAllString(sentence, "")
	sentence = strings.TrimSpace(sentence)
	
	i := 0
	lenSentence := len(sentence)
	resultSegment := make([]string, 0)
	for i < lenSentence {
		maxWord := ""
		for j := i + 1; j <= lenSentence; j++ {
			word := sentence[i:j]
			for _, c := range wordCorpus {
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
	return resultSegment, nil
}
