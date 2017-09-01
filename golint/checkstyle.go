package main

import "encoding/xml"

// DefaultCheckStyleVersion defines the default "version" attribute on "<checkstyle>" lememnt
var DefaultCheckStyleVersion = "1.0.0"

// CheckStyle represents a <checkstle> xml element.
type CheckStyle struct {
	XMLName xml.Name          `xml:"checkstyle"`
	Version string            `xml:"version,attr"`
	File    []*CheckStyleFile `xml:"file"`
}

// AddFile adds a CheckStyleFile with the given filename.
func (cs *CheckStyle) AddFile(csf *CheckStyleFile) {
	cs.File = append(cs.File, csf)
}

// GetFile gets a CheckStyleFile with the given filename.
func (cs *CheckStyle) GetFile(filename string) (csf *CheckStyleFile, ok bool) {
	for _, file := range cs.File {
		if file.Name == filename {
			csf = file
			ok = true
			return
		}
	}
	return
}

// EnsureFile ensures that a CheckStyleFile with the given name exists
// Returns either an exiting CheckStyleFile (if a file with that name exists)
// or a new CheckStyleFile (if a file with that name does not exists)
func (cs *CheckStyle) EnsureFile(filename string) (csf *CheckStyleFile) {
	csf, ok := cs.GetFile(filename)
	if !ok {
		csf = NewCheckStyleFile(filename)
		cs.AddFile(csf)
	}
	return csf
}

// NewCheckStyle returns a new CheckStyle
func NewCheckStyle() CheckStyle {
	return CheckStyle{Version: DefaultCheckStyleVersion, File: []*CheckStyleFile{}}
}

// CheckStyleFile represents a <file> xml element.
type CheckStyleFile struct {
	XMLName xml.Name           `xml:"file"`
	Name    string             `xml:"name,attr"`
	Error   []*CheckStyleError `xml:"error"`
}

// AddError adds a CheckStyleError to the file.
func (csf *CheckStyleFile) AddError(cse *CheckStyleError) {
	csf.Error = append(csf.Error, cse)
}

// NewCheckStyleFile creates a new CheckStyleFile.
func NewCheckStyleFile(filename string) *CheckStyleFile {
	return &CheckStyleFile{Name: filename, Error: []*CheckStyleError{}}
}

// CheckStyleError represents a <error> xml element
type CheckStyleError struct {
	XMLName xml.Name `xml:"error"`
	Line    int      `xml:"line,attr"`
	Message string   `xml:"message,attr"`
	Source  string   `xml:"source,attr"`
}

// NewCheckStyleError creates a new CheckStyleError
func NewCheckStyleError(line int, message string, source string) *CheckStyleError {
	return &CheckStyleError{Line: line, Message: message, Source: source}
}
