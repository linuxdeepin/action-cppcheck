package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func decodeErrors(fname string) ([]CppCheckError, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()
	var result CppCheckResults
	err = xml.NewDecoder(f).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}
	return result.Errors.Errors, nil
}

type CppCheckResults struct {
	XMLName xml.Name       `xml:"results"`
	Version string         `xml:"version,attr"`
	Errors  CppCheckErrors `xml:"errors"`
}
type CppCheckErrors struct {
	XMLName xml.Name        `xml:"errors"`
	Errors  []CppCheckError `xml:"error"`
}
type CppCheckError struct {
	XMLName  xml.Name          `xml:"error"`
	ID       string            `xml:"id,attr"`
	Severity string            `xml:"severity,attr"`
	Message  string            `xml:"msg,attr"`
	Verbose  string            `xml:"verbose,attr"`
	Location *CppCheckLocation `xml:"location"`
}
type CppCheckLocation struct {
	XMLName xml.Name `xml:"location"`
	File    string   `xml:"file,attr"`
	Line    int      `xml:"line,attr"`
}
