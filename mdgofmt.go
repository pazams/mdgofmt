package mdgofmt

import (
	"go/format"
)

// Formats golang code blocks inside markdown (gfm)
func Format(in []byte) ([]byte, error) {

	p := newParse(in)
	err := p.parse()
	if err != nil {
		return nil, err
	}

	out := make([]byte, 0)

	for i, v := range p.CodeBlocks {

		// previous end position
		pEnd := 0
		if i-1 >= 0 {
			pEnd = p.CodeBlocks[i-1].end
		}

		fBytes, err := format.Source(in[v.start:v.end])
		if err != nil {
			return nil, err
		}
		out = append(out, in[pEnd:v.start]...)
		out = append(out, fBytes...)
	}

	if len(p.CodeBlocks) > 0 {
		lEnd := p.CodeBlocks[len(p.CodeBlocks)-1].end
		out = append(out, in[lEnd:len(in)]...)
	}

	return out, nil
}
