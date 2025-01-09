package utils

import (
	"reflect"
	"testing"
)

func TestSegmentationWithUnknownWords(t *testing.T) {
	segmentation, _ := NewSegmentation("../../../data/dict/th/th_corpus.txt")
	tests := []struct {
		name     string
		sentence string
		want     []string
	}{
		{
			name:     "sentence with unknown words",
			sentence: "สวัสดีครับผมชื่อXXX",
			want:     []string{"สวัสดี", "ครับผม", "ชื่อ"},
		},
		{
			name:     "sentence with punctuation",
			sentence: "สวัสดี,ครับ!",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "empty string",
			sentence: "",
			want:     []string{},
		},
		{
			name:     "sentence with numbers",
			sentence: "สวัสดี123ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with mixed characters",
			sentence: "สวัสดี@ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with emojis",
			sentence: "สวัสดี😊ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with special characters",
			sentence: "สวัสดี#ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with hyphen",
			sentence: "สวัสดี-ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with multiple spaces",
			sentence: "สวัสดี  ครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with newline",
			sentence: "สวัสดี\nครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with tab",
			sentence: "สวัสดี\tครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with mixed languages",
			sentence: "สวัสดีhelloครับ",
			want:     []string{"สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with repeated words",
			sentence: "สวัสดีสวัสดีครับ",
			want:     []string{"สวัสดี", "สวัสดี", "ครับ"},
		},
		{
			name:     "sentence with long unknown word",
			sentence: "สวัสดีครับผมชื่อXXXXXXXXXXXXXXXXXXXX",
			want:     []string{"สวัสดี", "ครับผม", "ชื่อ"},
		},
		{
			name:     "sentence with mixed punctuation",
			sentence: "สวัสดี,ครับ!ผมชื่อ.",
			want:     []string{"สวัสดี", "ครับผม", "ชื่อ"},
		},
		{
			name:     "sentence with trailing spaces",
			sentence: "สวัสดีครับ   ",
			want:     []string{"สวัสดี", "ครับ"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := segmentation.Segmentation(tt.sentence)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Segmentation() = %v, want %v", got, tt.want)
			}
		})
	}
}