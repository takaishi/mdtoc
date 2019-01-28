package mdtoc

import (
	"testing"
)

var inputTextValid = `
# This is example

<!-- toc -->

## foo

aaa

## bar

bbb
`

var expectTOC = `
  * [foo](#foo)
  * [bar](#bar)
`

var outputTextValid = `
# This is example

<!-- toc -->
<!-- toc:start -->

  * [foo](#foo)
  * [bar](#bar)


<!-- toc:end -->

## foo

aaa

## bar

bbb
`

var inputTextInvalidWithoutTOCComment = `
# This is example

## foo

aaa

## bar

bbb
`

var inputTextInvalidWithoutTOCEndComment = `
# This is example

<!-- toc -->
<!-- toc:start -->

  * [foo](#foo)
  * [bar](#bar)


## foo

aaa

## bar

bbb
`

func TestGenerateTOC(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))

	if toc != expectTOC {
		t.Error(toc)
		t.Error(expectTOC)
	}
}

func TestInsertTOC(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))

	if toc != expectTOC {
		t.Error(toc)
		t.Error(expectTOC)
	}

	output, err := mt.InsertTOC(inputTextValid, toc)
	if err != nil {
		t.Error(err)
	}

	if output != outputTextValid {
		t.Error("unmatch")
		t.Error(output)
		t.Error(expectTOC)
	}
}

func TestInsertTOCWithInvalidString(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))
	_, err := mt.InsertTOC(inputTextInvalidWithoutTOCComment, toc)
	if err.Error() != "Can not find toc_pos comment `<!-- toc -->`" {
		t.Error()
	}

	_, err = mt.InsertTOC(inputTextInvalidWithoutTOCEndComment, toc)

	if err.Error() != "Can not find toc end position comment `<!-- toc:end -->`." {
		t.Error()
	}
}
