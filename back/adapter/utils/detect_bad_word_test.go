package utils

import (
	"testing"

	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
)

func TestDectectBadWord(t *testing.T) {
	segmentation, _ := NewSegmentation("../../../data/dict/th/th_corpus.txt")
	detectBadWord, _ := NewDetectBadWord(
		"../../../data/model/cc.th.300.vec",
		"../../../data/dict/th/th_swear.txt",
		segmentation,
		parallelconfig.Default())
	tests := []struct {
		name       string
		sentence   string
		swearWords []string
		threshold  float32
		want       []string
	}{
		{
			name:       "detect similar swear word",
			sentence:   "สวัสดีหน้าส้นตีน",
			swearWords: []string{"ควย", "FuckYou", "สัส", "กู", "มึง"},
			threshold:  0.5,
			want:       []string{"ส้นตีน"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectBadWord.DectectBadWord(tt.sentence, tt.threshold)
			if len(got) != len(tt.want) {
				t.Errorf("DectectBadWord() = %v, want %v", got, tt.want)
			}
			for i, word := range got {
				if word != tt.want[i] {
					t.Errorf("DectectBadWord() = %v, want %v", word, tt.want[i])
				}
			}
		})
	}
}
