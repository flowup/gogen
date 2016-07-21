package gogen

import (
  "go/ast"
  "unicode"
  "strings"
  "regexp"
)

//
type Tag struct {
  Name    string //may be obsolete bc tags are indexed in tag map by their names
  Values  map[string]string
}

func NewTag() *Tag{
  return &Tag{
    Values: make(map[string]string),
  }
}

type TagMap struct {
  tags map[string]*Tag
}

func NewTagMap() *TagMap {
  return &TagMap{
    tags: make(map[string]*Tag),
  }
}

func (t *TagMap) HasTag(name string) bool {
  _, ok := t.tags[name]
  return ok
}

func (t *TagMap) GetTag(name string) (*Tag, bool) {
  val, ok := t.tags[name]
  return val, ok
}

func (t *TagMap) SetTagValue (name string, tag *Tag) {
  t.tags[name] = tag
}

func (t *TagMap) DeleteTag (name string) {
  delete(t.tags, name)
}

func (t *TagMap) NumOfTags () int {
  return len(t.tags)
}

func (t *TagMap) GetAllTags () map[string]*Tag {
  return t.tags
}

func (t *TagMap) GetAllKeys () []string {
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
    value += "\""
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
    value += string(input[i])
    i++
  }

  // return the remaining input and values read from line
  return input[i:], name, value
}



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
        tag := NewTag()
        tag.Name = tagName
        //fmt.Println("Tag Name:", tagName)
        // while there is some input check for parameters
        for line != "" {
          var parName, parVal string
          line, parName, parVal = parseValue(line)
          if parName != "" {
            //fmt.Println("Parameter name:", parName, "Parameter value", parVal)
            tag.Values[parName] = parVal
          }
        }
        // save tag to map
        tagMap.SetTagValue(tagName, tag)
      }
    }
  }

  return tagMap
}