package struct2interface

import (
	"io/ioutil"
	"testing"
)

func TestDir(t *testing.T) {
	err := MakeDir("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	output, err := ioutil.ReadFile("./testdata/interface_Method.go")
	if err != nil {
		t.Fatal(err)
	}
	test, err := ioutil.ReadFile("./testdata/test_interface_Method.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(test) {
		t.Fail()
	}

	output, err = ioutil.ReadFile("./testdata/interface_Method1.go")
	if err != nil {
		t.Fatal(err)
	}
	test, err = ioutil.ReadFile("./testdata/test_interface_Method1.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(test) {
		t.Fail()
	}
}
