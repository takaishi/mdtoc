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

func TestGenerateWithTOC(t *testing.T) {
	mt := MDToc{File: "", InFile: false, Level: 2}

	toc := mt.GenerateTOC([]byte(inputTextValid))
	expect := `
  * [foo](#foo)
  * [bar](#bar)`

	if toc != expect {
		t.Error(toc)
		t.Error(expect)
	}

	output, err := mt.GenerateWithTOC(inputTextValid, toc)
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
