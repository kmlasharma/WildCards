package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Ontology struct {
	XMLName                 xml.Name                `xml:"Ontology"`
	AnnotationAssertions    []AnnotationAssertion   `xml:"AnnotationAssertion"`
	SubAnnotationProperties []SubAnnotationProperty `xml:"SubAnnotationPropertyOf"`
	Declarations            []Declaration           `xml:"Declaration"`
	Subclasses              []Subclass              `xml:"SubClassOf"`
}

/* AnnotationAssertion */
type AnnotationAssertion struct {
	XMLName            xml.Name `xml:"AnnotationAssertion"`
	AnnotationProperty AnnotationProperty
	AbbreviatedIRI     AbbreviatedIRI
	Literal            Literal
}

type AnnotationProperty struct {
	XMLName        xml.Name       `xml:"AnnotationProperty"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

type AbbreviatedIRI struct {
	XMLName xml.Name `xml:"AbbreviatedIRI"`
	Value   string   `xml:",chardata"`
}

type Literal struct {
	XMLName     xml.Name `xml:"Literal"`
	DatatypeIRI string   `xml:"datatypeIRI,attr"`
	Lang        string   `xml:"lang,attr"`
	Value       string   `xml:",chardata"`
}

/* Sub Annotation Properties */

type SubAnnotationProperty struct {
	XMLName                  xml.Name `xml:"SubAnnotationPropertyOf"`
	ParentAnnotationProperty AnnotationProperty
	ChildAnnotationProperty  AnnotationProperty
}

/* Declaration */

type Declaration struct {
	XMLName xml.Name `xml:"Declaration"`
	Class   Class
}

type Class struct {
	XMLName        xml.Name       `xml:"Class"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

/* Subclass */

type Subclass struct {
	XMLName              xml.Name `xml:"SubclassOf"`
	Class                Class
	ObjectSomeValuesFrom ObjectSomeValuesFrom
}

type ObjectSomeValuesFrom struct {
	XMLName        xml.Name `xml:"ObjectSomeValuesFrom"`
	ObjectProperty ObjectProperty
	Class          Class
}

type ObjectProperty struct {
	XMLName        xml.Name       `xml:"ObjectProperty"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

func GenerateOntology(filepath string) Ontology {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error opening file:", filepath, err)
		return
	}

	var ontology Ontology
	return xml.Unmarshal(data, &ontology)
}

func main() {
	ontology := GenerateOntology("./data/DINTO.owl")
	fmt.Println(ontology)
}
