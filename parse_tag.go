package gogen

import (
  "go/ast"
  "unicode"
  "strings"
  "regexp"
)

// Tag holds information about one tag, its name
// parameters and their values
type Tag struct {
  name    string //may be obsolete bc tags are indexed in tag map by their names
  values  map[string]string
}

// NewTag will create and return an empty tag
func NewTag(name string) *Tag{
  return &Tag{
    name: name,
    values: make(map[string]string),
  }
}

// GetName will return the name of a tag
func (t *Tag) GetName () string {
  return t.name
}

// HasParameter will return if a parameter with name
// sent to function is a parameter of this tag
func (t *Tag) Has (name string) bool {
  _, ok := t.values[name]
  return ok
}

// GetParameter will return the value of a parameter
// along with bool value that determines if the parameter was found
func (t *Tag) Get (name string) (string, bool) {
  retVal, ok := t.values[name]
  return retVal, ok
}

// SetParameterValue will set a value of a parameter.
// Can be used for creating new parameters
func (t *Tag) Set (name string, value string) {
  t.values[name] = value
}

// DeleteParameter will delete a parameter from a tag
func (t *Tag) Delete (name string) {
  delete(t.values, name)
}

// NumOfParameters will return number of parameters of a tag
func (t *Tag) Num () int {
  return len(t.values)
}

// GetAllParameters will return all parameters of a tag with their values
func (t *Tag) GetAll () map[string]string {
  return t.values
}

// GetAllParameterNames
func (t *Tag) GetParameterNames () []string {
  keys := []string{}
  for key := range t.values {
    keys = append(keys, key)
  }

  return keys
}

// TagMap  holds build tags
// of one function/interface/structure etc
type TagMap struct {
  tags map[string]*Tag
}

// NewTagMap will create and return an empty tag map
func NewTagMap() *TagMap {
  return &TagMap{
    tags: make(map[string]*Tag),
  }
}

// HasTag will check if the map has a tag with
// name given by parameter
func (t *TagMap) Has (name string) bool {
  _, ok := t.tags[name]
  return ok
}

// GetTag will get a value of a tag with
// key given by parameter
func (t *TagMap) Get (name string) (*Tag, bool) {
  retVal, ok := t.tags[name]
  return retVal, ok
}

// SetTagValue will set a tag value to value
// given by parameter
func (t *TagMap) Set (name string, tag *Tag) {
  t.tags[name] = tag
}

// DeleteTag will delete a tag with name given
// by parameter from the map
func (t *TagMap) Delete (name string) {
  delete(t.tags, name)
}

// NumOfTags will return number of tags in the map
func (t *TagMap) Num () int {
  return len(t.tags)
}

// GetAllTags will return all tags in the map
func (t *TagMap) GetAll () map[string]*Tag {
  return t.tags
}

// GetAllKeys will get names of all tags in the map
func (t *TagMap) GetTagNames () []string {
  keys := []string{}
  for key := range t.tags {
    keys = append(keys, key)
  }

  return keys
}

// parseValue parses one value of a tag,
// returns the remaining string to parse,
// parameter of a tag and its value
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

  // skip the "-" signs at the start of name of the tag
  for i < len(input) {
    if input[i] != '-' {
      break
    }
    i++
  }

  // read the name of a tag/parameter
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

  // read the value of a tag parameter
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

// ParseTags will create a TagMap from comments given by
// parameter
func ParseTags (commentMap ast.CommentMap) *TagMap {
  tagMap := NewTagMap()
  for _, comment := range commentMap.Comments() {
    // split comment to lines
    lines := strings.Split(comment.Text(), "\n")
    for _, line := range lines {
      // if line does not match this regexp made by Miro do not read tags from line
      if !regexp.MustCompile("(?:--(\\w{2,})|-(\\w))\\s*(?:\"([^\"]*)\"|(\\w+)|())\\s*").Match([]byte(line)) {
        continue
      }
      line, tagName, _ := parseValue(line)
      // if there is a tag on the line read its parameters and their values
      if len(tagName) > 0 && tagName[0] == '@' {
        tag := NewTag(tagName)
        //fmt.Println("Tag Name:", tagName)
        // while there is some input check for parameters
        for line != "" {
          var parName, parVal string
          line, parName, parVal = parseValue(line)
          if parName != "" {
            //fmt.Println("Parameter name:", parName, "Parameter value", parVal)
            tag.Set(parName, parVal)
          }
        }
        // save tag to map
        tagMap.Set(tagName, tag)
      }
    }
  }

  return tagMap
}