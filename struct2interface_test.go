package struct2interface

import (
	"io/ioutil"
	"testing"
)

func TestDir(t *testing.T) {
	err := MakeDir("./testdata/case_single_file")
	if err != nil {
		t.Fatal(err)
	}

	output, err := ioutil.ReadFile("./testdata/case_single_file/interface_Method.go")
	if err != nil {
		t.Fatal(err)
	}
	test, err := ioutil.ReadFile("./testdata/case_single_file/test_interface_Method.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(test) {
		t.Fail()
	}

	output, err = ioutil.ReadFile("./testdata/case_single_file/interface_Method1.go")
	if err != nil {
		t.Fatal(err)
	}
	test, err = ioutil.ReadFile("./testdata/case_single_file/test_interface_Method1.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(test) {
		t.Fail()
	}
}

func TestPackage(t *testing.T) {
	err := MakeDir("./testdata/case_package")
	if err != nil {
		t.Fatal(err)
	}

	output, err := ioutil.ReadFile("./testdata/case_package/interface_PackageMethod.go")
	if err != nil {
		t.Fatal(err)
	}
	test, err := ioutil.ReadFile("./testdata/case_package/test_interface_PackageMethod.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(test) {
		t.Fail()
	}
}
