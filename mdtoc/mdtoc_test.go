package mdtoc

import (
	"testing"
)

var inputTextValid = `# This is example

<!-- toc -->

## foo

aaa

## bar

bbb
`

var inputTextInvalid1 = `# This is example

## foo

aaa

## bar

bbb
`

var inputTextInvalid2 = `# This is example

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
	expect := `
  * [foo](#foo)
  * [bar](#bar)`

	if toc != expect {
		t.Error(toc)
		t.Error(expect)
	}
}

func TestInsertTOC(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))
	expect := `
  * [foo](#foo)
  * [bar](#bar)`

	if toc != expect {
		t.Error(toc)
		t.Error(expect)
	}

	output, err := mt.InsertTOC(inputTextValid, toc)
	if err != nil {
		t.Error(err)
	}
	expect = `# This is example

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

	if output != expect {
		t.Error("unmatch")
		t.Error(output)
		t.Error(expect)
	}
}

func TestInsertTOCWithInvalidString(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))
	expect := `
  * [foo](#foo)
  * [bar](#bar)`

	if toc != expect {
		t.Error(toc)
		t.Error(expect)
	}

	_, err := mt.InsertTOC(inputTextInvalid1, toc)

	if err.Error() != "Can not find toc_pos comment `<!-- toc -->`" {
		t.Error()
	}

	_, err = mt.InsertTOC(inputTextInvalid2, toc)

	if err.Error() != "Can not find toc end position comment `<!-- toc:end -->`." {
		t.Error()
	}
}
