# struct2interface

[![Go](https://github.com/sysulq/struct2interface/actions/workflows/go.yml/badge.svg)](https://github.com/sysulq/struct2interface/actions/workflows/go.yml)

This is a development helper program that generates a Golang interface by inspecting
the structure methods of an existing `.go` file. The primary use case is to generate
interfaces for mockery, so that mockery can generate mocks from those interfaces. This
makes unit testing easier.

## Install

```
go install github.com/sysulq/struct2interface/cmd/struct2interface@latest
```

## Usage

Here is the help output of struct2interface:

```
$ struct2interface --help
Usage:
  struct2interface [flags]

Flags:
  -d, --dir string   Go source file dir to read (default ".")
  -h, --help         help for struct2interface
```

As an example, let's say you wanted to generate an interface for the Method structure
in this sample code:

```
package testdata

// Method describes the code and documentation
// tied into a method
type Method struct {
	Code string
	Docs []string
}

// Lines return a []string consisting of
// the documentation and code appended
// in chronological order
func (m *Method) Lines() []string {
	var lines []string
	lines = append(lines, m.Docs...)
	lines = append(lines, m.Code)
	return lines
}
```

The struct2interface helper program can generate this interface for you:

```
$ struct2interface -d testdata
struct2interface: testdata: wrote testdata/interface_Method.go
struct2interface: testdata: wrote testdata/interface_Method1.go
```
