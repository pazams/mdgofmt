[travis-url]: https://travis-ci.org/pazams/mdgofmt
[travis-image]: https://api.travis-ci.org/pazams/mdgofmt.svg
[coveralls-image]: https://coveralls.io/repos/pazams/mdgofmt/badge.svg?branch=master&service=github
[coveralls-url]: https://coveralls.io/r/pazams/mdgofmt

# mdgofmt [![Test coverage][coveralls-image]][coveralls-url] [![Build status][travis-image]][travis-url]
Formats golang code blocks inside markdown (gfm)

## Get Started
#### Installation
```sh
$ go get github.com/pazams/mdgofmt
```
#### Examples
__Add to your project:__
```go

import (
	"fmt"
	"github.com/pazams/mdgofmt"
	"io/ioutil"
)

func main() {

	in, err := ioutil.ReadFile("some file path")
	check(err)

	out, err := mdgofmt.Format(in)
	check(err)
	
	fmt.Println(out)
}

func check(e error) {
	if e != nil {
	  panic(e)
	}
}

```

## Before & After
[before](https://github.com/pazams/mdgofmt/blob/master/testdata/struct.md)

[after](https://github.com/pazams/mdgofmt/blob/master/testdata/struct.expected.md)

## Todo
- improve parser to indentify all possible valid gfm golang code blocks. It currently detects only blocks starting with: 
``` 
```go 
```

## License
MIT
