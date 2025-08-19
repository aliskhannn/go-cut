package cut

import (
	"bytes"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := "a,b,c,d\n1,2,3,4\nx,y,z\n"
	var out bytes.Buffer

	cfg := Config{
		Fields:    []int{1, 3},
		Delimiter: ",",
		Separated: false,
	}

	if err := Process(strings.NewReader(input), &out, cfg); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	want := "a,c\n1,3\nx,z\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseFieldArg(t *testing.T) {
	got, err := ParseFieldArg("1,3-5,7")
	if err != nil {
		t.Fatal(err)
	}
	want := []int{1, 3, 4, 5, 7}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
