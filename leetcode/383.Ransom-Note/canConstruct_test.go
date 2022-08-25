package ransomnote

import (
	"testing"
)

type canConstructCase struct {
	ransomNote, magazine string
	res                  bool
}

func canConstruct(ransomNote string, magazine string) bool {
	charDict := make(map[rune]int)
	for _, c := range magazine {
		charDict[c]++

	}

	for _, c := range ransomNote {
		charDict[c]--
		if charDict[c] < 0 {
			return false
		}
	}

	return true

}

func TestCountSmaller(t *testing.T) {
	qs := []canConstructCase{
		{
			ransomNote: "a",
			magazine:   "b",
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			res:        true,
		},
	}

	for idx, q := range qs {
		if got := canConstruct(q.ransomNote, q.magazine); got != q.res {
			t.Fatalf("case %v: got %v, expected %v ", idx, got, q.res)
		}

	}
}
