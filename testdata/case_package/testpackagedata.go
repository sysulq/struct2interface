package testdata

type PackageMethod struct{}

// the //-style comment test
func (m *PackageMethod) Method1() string {
	return ""
}

type PackageMethod2 struct{}

/*
the /*-style comment test
*/
func (m *PackageMethod2) Method1() string {
	return ""
}
