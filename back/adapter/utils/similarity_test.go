package utils

import (
	"fmt"
	"testing"

	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
)

func TestLookSimilarity(t *testing.T) {
	thaiSimilarity, _ := NewSimilarity("../../../data/model/cc.th.300.vec", parallelconfig.Default())

	tests := []struct {
		name            string
		wording         string
		limitSimilarity int
		want            float64
	}{
		{
			name:            "similarity for สวัสดี",
			wording:         "สวัสดี",
			limitSimilarity: 2,
			want:            0.7,
		},
		{
			name:            "similarity for ครับ",
			wording:         "ครับ",
			limitSimilarity: 2,
			want:            0.5,
		},
		{
			name:            "similarity for ผม",
			wording:         "ผม",
			limitSimilarity: 2,
			want:            0.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := thaiSimilarity.Similarity(tt.wording, tt.limitSimilarity)
			if len(got) != tt.limitSimilarity {
				t.Errorf("Similarity() = %v, want %v", got, tt.want)
			}
			for _, neighbor := range got {
				if neighbor.Similarity < tt.want {
					t.Errorf("Similarity() = %v, want %v", neighbor.Similarity, tt.want)
				}
				fmt.Printf("Word() = %v, Similarity() = %v, want %v \n", neighbor.Word, neighbor.Similarity, tt.want)
			}
		})
	}
}
