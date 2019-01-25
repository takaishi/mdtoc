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
}
