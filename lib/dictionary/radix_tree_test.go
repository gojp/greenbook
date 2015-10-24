package dictionary

import "testing"

var getTests = [][]rune{
	//[]rune("しけん"),
	//[]rune("てつだう"),
	//[]rune("手伝う"),
	//[]rune("普通"),
	[]rune("ふつう"),
	[]rune("ふつ"),
}

func TestGet(t *testing.T) {
	r := NewRadixTree()
	for i, entry := range getTests {
		r.Insert(entry, EntryID(i))
		got := r.Get(entry)
		if len(got) != 1 {
			t.Fatalf("%q len(got) = %d, want %d", entry, len(got), 1)
		}
		if got[0] != EntryID(i) {
			t.Fatalf("got[0] = %q, want %q", got[0], entry)
		}
	}
}
