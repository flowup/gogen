## Gogen

[![Master Status](https://travis-ci.org/flowup/gogen.svg?branch=master)](https://travis-ci.org/flowup/gogen)
[![Develop Status](https://travis-ci.org/flowup/gogen.svg?branch=develop)](https://travis-ci.org/flowup/gogen)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/flowup/gogen)

> Warning: This project is under heavy development

Gogen is a library that helps to build Golang code generators with ease. It
produces *builds* from the given `.go` files or packages. Every build contains useful
information about types, functions, methods, constants, etc. This information
can be simply used to build templates based on an already existing code.

Gogen directly parses `.go` files so there is no need of intermediary language.
This allows us to simply integrate Gogen into already existing projects.

## Example

```go
build := gogen.Parse(BasePathToYourFile)
file := build.File(YourFileName)

// Print all functions
for i, f := range file.Functions() {
  fmt.Println(i,": Structure", f.Name(), "Is method:", f.IsMethod())
}

// Print structures with @dao comment tag
for i, s := range file.Structs().Filter("@dao") {
  fmt.Println(i,": Structure", s.Name(), "Num of fields:", len(s.Fields()))
  for j, field := range s.Fields() {
    ft, _ := field.Type()
    fmt.Println("Field no:", j, "Name:", field.Name(), "Type:", ft)
  }
}

```
## Kickstart

Will be added soon