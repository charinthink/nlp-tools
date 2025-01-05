package utilities

import (
	"reflect"
	"testing"
)

func TestThaiSegmentation(t *testing.T) {
	tests := []struct {
		name       string
		corpusDir  string
		sentence   string
		want       []string
		wantErr    bool
	}{
		{
			name:      "basic segmentation",
			corpusDir: "../../../data/dict/th/th_corpus.txt",
			sentence:  "สวัสดีครับ",
			want:      []string{"สวัสดี", "ครับ"},
			wantErr:   false,
		},
		{
			name:      "empty sentence",
			corpusDir: "../../../data/dict/th/th_corpus.txt",
			sentence:  "",
			want:      []string{},
			wantErr:   false,
		},
		{
			name:      "no matching words",
			corpusDir: "../../../data/dict/th/th_corpus.txt",
			sentence:  "xyz",
			want:      []string{},
			wantErr:   false,
		},
		{
			name:      "sentence with punctuation",
			corpusDir: "../../../data/dict/th/th_corpus.txt",
			sentence:  "สวัสดี,ครับ!",
			want:      []string{"สวัสดี", "ครับ"},
			wantErr:   false,
		},
		{
			name:      "invalid corpus directory",
			corpusDir: "../../../data/dict/th/test_corpus.txt",
			sentence:  "สวัสดีครับ",
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ThaiSegmentation(tt.corpusDir, tt.sentence)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThaiSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThaiSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}