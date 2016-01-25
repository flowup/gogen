## Gogen

[![Build Status](https://travis-ci.org/gophergala2016/gogen.svg?branch=master)](https://travis-ci.org/flowup/gogen) [![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/flowup/gogen)

> Warning: This project is in the development

Gogen is lightweight [**resourceful code generator library**](https://github.com/gophergala2016/gogen/wiki) written entirely in Go. It is based on resource definitions, pipes and executable configuration files.

## Kickstart

To start with Gogen, we'll need resources(models, api schemes, blueprints, ...) and generators that should be used to generate results. We'll then mix everything in one or more configuration files, that will allow us to generate code.

Let's start by creating `config.go` file with already existing resources and generators.

```go
package main

import (
	"github.com/gophergala2016/gogen"
	"github.com/gophergala2016/gogen/generator"
	"github.com/gophergala2016/gogen/model"
)

func main() {
  // Define adds resources for the generators, this way
  // we are adding the user model that is already defined
  // for us to the generator resources
  gogen.Define(resource.User)

  // set output for the generator that is generating models
  model.Generator.SetOutputDir("./model")

  // Pipe is set of generators. More pipes can go in parallel
  // and everything in pipe is executed in serie. This allows
  // generators to have dependencies on each other
  //
  // This is creating our pipe with the ModelGenerator in it
	gogen.Pipe(
		model.Generator,
	)

  // start the generator
  gogen.Generate()
}
```

This [**configuration file**](https://github.com/gophergala2016/gogen/wiki/Configuration-files) will create new folder `model` in the current working directory, with generated user model file.

## Known integrable generators

- [Swagger](http://swagger.io/)
- [Goa](http://goa.design/goagen.html)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)
