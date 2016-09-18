package mdgofmt_test

import (
	"bytes"
	"github.com/pazams/mdgofmt"
	"io/ioutil"
	"testing"
)

var markdownValidFiles = []string{
	"struct", "multiple", "foreign-block",
}

var markdownInvalidFiles = []string{
	"invalid-line", "invalid-format", "invalid-block",
}

func TestValidFiles(t *testing.T) {
	for _, v := range markdownValidFiles {

		inF := "testdata/" + v + ".md"
		exF := "testdata/" + v + ".expected.md"

		t.Logf("testing valid files", inF, exF)

		in, err := ioutil.ReadFile(inF)
		check(err)

		out, err := mdgofmt.Format(in)
		check(err)

		ex, err := ioutil.ReadFile(exF)
		check(err)

		if !bytes.Equal(out, ex) {
			t.Errorf("output does not match expected", in, out, ex)
		}

	}
}

func TestInvalidFiles(t *testing.T) {

	for _, v := range markdownInvalidFiles {

		inF := "testdata/" + v + ".md"
		t.Logf("testing invalid file", inF)

		in, err := ioutil.ReadFile(inF)

		if err != nil {
			t.Logf("teeeeeeeeeeeeeeeeeeeeeest error")
		}
		check(err)

		out, err := mdgofmt.Format(in)
		if err == nil {
			t.Errorf("expected error for invalid file", inF)
		}
		_ = out
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
