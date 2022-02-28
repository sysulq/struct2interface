package struct2interface

import (
	"fmt"
	"strings"
	"testing"
)

var (
	src = []byte(`// comment

package testdata

// InterfaceComment
// Method describes the code and documentation
// tied into a method
type MethodInterface interface {
	// Lines return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines() []string
}

// InterfaceComment
// Method1 describes the code and documentation
// tied into a method
type Method1Interface interface {
	// Lines1 return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines1() []string
}
`)
)

func TestMaker(t *testing.T) {
	files := []string{"./testdata/testdata.go"}
	output, err := Make(files, "comment", "pkg", "Interface", "InterfaceComment", true)
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(src) {
		fmt.Println(len(output), len(src))
		fmt.Println("-------------------------")
		fmt.Println(strings.ReplaceAll(string(output), "\n", " "))
		fmt.Println("-------------------------")
		fmt.Println(strings.ReplaceAll(string(src), "\n", " "))
		fmt.Println("-------------------------")
		t.Fail()
	}
}
