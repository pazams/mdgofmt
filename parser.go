package mdgofmt

import (
	"errors"
	"regexp"
)

type codeBlock struct {
	start int
	end   int
}

type parse struct {
	CodeBlocks []codeBlock
	input      []byte
}

// backtick char
const bt byte = byte('`')

// new line regex
var r = regexp.MustCompile("\r\n?|\n")

// Return new parser
func newParse(input []byte) *parse {
	return &parse{
		input:      input,
		CodeBlocks: make([]codeBlock, 0),
	}
}

func (p *parse) parse() error {

	startBlock := 0
	for i := range p.input {

		if i < startBlock {
			continue
		}
		// expecting start of code block
		if startBlock == 0 {
			if isGoCodeBlockStart(p.input[i:]) {
				npos, err := getNextLineLocation(p.input[i:])
				if err != nil {
					return err
				}
				startBlock = i + npos
			}
		} else { // expecting end of code block
			if isGoCodeBlockEnd(p.input[i:]) {
				p.CodeBlocks = append(p.CodeBlocks, codeBlock{startBlock, i})
				startBlock = 0
			}
		}

	}
	if startBlock != 0 {
		return errors.New("no matching close of code block")
	} else {
		return nil
	}

}

func getNextLineLocation(bytes []byte) (int, error) {

	loc := r.FindIndex(bytes)

	if len(loc) == 0 {
		return 0, errors.New("no new line following code block start")
	} else {
		return loc[1], nil
	}
}

func isGoCodeBlockStart(bytes []byte) bool {

	if !isCodeBlockDelim(bytes) {
		return false
	}

	if string(bytes[3:5]) == "go" {
		return true
	} else {
		return false
	}
}

func isGoCodeBlockEnd(bytes []byte) bool {
	if !isCodeBlockDelim(bytes) {
		return false
	}

	return true
}

func isCodeBlockDelim(bytes []byte) bool {

	if len(bytes) < 3 {
		return false
	}
	if bytes[0] == bt && bytes[1] == bt && bytes[2] == bt {
		return true
	} else {
		return false
	}
}
