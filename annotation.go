package gogen

import (
  "go/ast"
  "unicode"
  "strings"
  "regexp"
)

// Annotation holds information about one annotation, its name
// parameters and their values
type Annotation struct {
  name    string //may be obsolete bc annotations are indexed in annotation map by their names
  values  map[string]string
}

// NewAnnotation will create and return an empty annotation
func NewAnnotation(name string) *Annotation{
  return &Annotation{
    name: name,
    values: make(map[string]string),
  }
}

// GetName will return the name of a annotation
func (t *Annotation) GetName () string {
  return t.name
}

// Has will return if a parameter with name
// sent to function is a parameter of this annotation
func (t *Annotation) Has (name string) bool {
  _, ok := t.values[name]
  return ok
}

// Get will return the value of a parameter
// along with bool value that determines if the parameter was found
func (t *Annotation) Get (name string) (string, bool) {
  retVal, ok := t.values[name]
  return retVal, ok
}

// Set will set a value of a parameter.
// Can be used for creating new parameters
func (t *Annotation) Set (name string, value string) {
  t.values[name] = value
}

// Delete will delete a parameter from a annotation
func (t *Annotation) Delete (name string) {
  delete(t.values, name)
}

// Num will return number of parameters of a annotation
func (t *Annotation) Num () int {
  return len(t.values)
}

// GetAll will return all parameters of a annotation with their values
func (t *Annotation) GetAll () map[string]string {
  return t.values
}

// GetParameterNames will return all parameter names.
func (t *Annotation) GetParameterNames () []string {
  keys := []string{}
  for key := range t.values {
    keys = append(keys, key)
  }

  return keys
}

// AnnotationMap  holds build annotations
// of one function/interface/structure etc
type AnnotationMap struct {
  annotations map[string]*Annotation
}

// NewAnnotationMap will create and return an empty annotation map
func NewAnnotationMap() *AnnotationMap {
  return &AnnotationMap{
    annotations: make(map[string]*Annotation),
  }
}

// Has will check if the map has a annotation with
// name given by parameter
func (t *AnnotationMap) Has (name string) bool {
  _, ok := t.annotations[name]
  return ok
}

// Get will get a value of a annotation with
// key given by parameter
func (t *AnnotationMap) Get (name string) (*Annotation, bool) {
  retVal, ok := t.annotations[name]
  return retVal, ok
}

// Set will set a annotation value to value
// given by parameter
func (t *AnnotationMap) Set (name string, annotation *Annotation) {
  t.annotations[name] = annotation
}

// Delete will delete a annotation with name given
// by parameter from the map
func (t *AnnotationMap) Delete (name string) {
  delete(t.annotations, name)
}

// Num will return number of annotations in the map
func (t *AnnotationMap) Num () int {
  return len(t.annotations)
}

// GetAll will return all annotations in the map
func (t *AnnotationMap) GetAll () map[string]*Annotation {
  return t.annotations
}

// GetAnnotationNames will get names of all annotations in the map
func (t *AnnotationMap) GetAnnotationNames () []string {
  keys := []string{}
  for key := range t.annotations {
    keys = append(keys, key)
  }

  return keys
}

// parseValue parses one value of a annotation,
// returns the remaining string to parse,
// parameter of a annotation and its value
func parseValue(input string) (string, string, string) {
  i := 0
  name := ""
  value := ""

  // skip whitespaces before name
  for i < len(input) {
    if !unicode.IsSpace(rune(input[i])) {
      break
    }
    i++
  }

  // skip the "-" signs at the start of name of the annotation
  for i < len(input) {
    if input[i] != '-' {
      break
    }
    i++
  }

  // read the name of a annotation/parameter
  for i < len(input) {
    if unicode.IsSpace(rune(input[i])) || input[i] == '=' {
      break
    }
    name += string(input[i])
    i++
  }

  // skip whitespaces or a = sign before value
  if i < len(input) && input[i] == '=' {
    i++
  } else {
    for i < len(input) {
      if !unicode.IsSpace(rune(input[i])) {
        break
      }
      i++
    }
  }

  //if there is no value and there is another
  // option starting with "-", or input ends return what you have
  if i == len(input) || input[i] == '-' {
    return input[i:], name, value
  }

  // check if delimiter of value of parameter is space or a " sign
  delimiter := ' '
  if input[i] == '"' {
    delimiter = '"'
    i++
  }

  // read the value of a annotation parameter
  if delimiter == ' ' {
    // if the value is separated by spaces
    for i < len(input) {
      if unicode.IsSpace(rune(input[i])) {
        break
      }
      value += string(input[i])
      i++
    }
  } else {
    // if the value is separated by " signs
    for i < len(input) {
      if input[i] == '"' {
        break
      }
      value += string(input[i])
      i++
    }
    i++
  }

  // return the remaining input and values read from line
  return input[i:], name, value
}

// ParseAnnotations will create a AnnotationMap from comments given by
// parameter
func ParseAnnotations (commentMap ast.CommentMap) *AnnotationMap {
  annotationMap := NewAnnotationMap()
  for _, comment := range commentMap.Comments() {
    // split comment to lines
    lines := strings.Split(comment.Text(), "\n")
    for _, line := range lines {
      // if line does not match this regexp made by Miro do not read annotations from line
      if !regexp.MustCompile(`(\@\S+)(:? ([\S]+))*`).Match([]byte(line)) {
        continue
      }
      line, annotationName, _ := parseValue(line)
      // if there is a annotation on the line read its parameters and their values
      if len(annotationName) > 0 && annotationName[0] == '@' {
        annotation := NewAnnotation(annotationName)
        //fmt.Println("Annotation Name:", annotationName)
        // while there is some input check for parameters
        for line != "" {
          var parName, parVal string
          line, parName, parVal = parseValue(line)
          if parName != "" {
            //fmt.Println("Parameter name:", parName, "Parameter value", parVal)
            annotation.Set(parName, parVal)
          }
        }
        // save annotation to map
        annotationMap.Set(annotationName, annotation)
      }
    }
  }

  return annotationMap
}