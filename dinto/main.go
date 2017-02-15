package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
)

func main() {
	ontology := GenerateOntology("./data/test.owl")
	fmt.Println("Number of Prefixes:", len(ontology.Prefixes))
	fmt.Println("Number of Imports:", len(ontology.Imports))
	fmt.Println("Number of Annotations:", len(ontology.Annotations))
	fmt.Println("Number of Declarations:", len(ontology.Declarations))
	fmt.Println("Number of Subclasses:", len(ontology.Subclasses))
	fmt.Println("Number of Anootation Assertions:", len(ontology.AnnotationAssertions))
	fmt.Println("Number of Sub Annotation Properties:", len(ontology.SubAnnotationProperties))

}

type Ontology struct {
	XMLName                 xml.Name                `xml:"Ontology"`
	Prefixes                []Prefix                `xml:"Prefix"`
	Imports                 []Import                `xml:"Import"`
	Annotations             []Annotation            `xml:"Annotation"`
	Declarations            []Declaration           `xml:"Declaration"`
	Subclasses              []Subclass              `xml:"SubClassOf"`
	AnnotationAssertions    []AnnotationAssertion   `xml:"AnnotationAssertion"`
	SubAnnotationProperties []SubAnnotationProperty `xml:"SubAnnotationPropertyOf"`
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
	IRI            IRI            `xml:"IRI,attr"`
}

type AbbreviatedIRI string

type IRI string

type Literal struct {
	XMLName     xml.Name `xml:"Literal"`
	DatatypeIRI string   `xml:"datatypeIRI,attr"`
	Lang        string   `xml:"lang,attr"`
	Value       string   `xml:",chardata"`
}

/* Sub Annotation Properties */

type SubAnnotationProperty struct {
	XMLName              xml.Name             `xml:"SubAnnotationPropertyOf"`
	AnnotationProperties []AnnotationProperty `xml:"AnnotationProperty"`
}

/* Declaration */

type Declaration struct {
	XMLName            xml.Name `xml:"Declaration"`
	Class              Class
	ObjectProperty     ObjectProperty
	NamedIndividual    NamedIndividual
	AnnotationProperty AnnotationProperty
	Datatype           Datatype
}

type Class struct {
	XMLName        xml.Name       `xml:"Class"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

type NamedIndividual struct {
	XMLName        xml.Name       `xml:"NamedIndividual"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

type Datatype struct {
	XMLName        xml.Name       `xml:"Datatype"`
	AbbreviatedIRI AbbreviatedIRI `xml:"abbreviatedIRI,attr"`
}

/* Subclass */

type Subclass struct {
	XMLName              xml.Name `xml:"SubClassOf"`
	Classes              []Class  `xml:"Class"`
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

/* Annotation */

type Annotation struct {
	XMLName            xml.Name `xml:"Annotation"`
	AnnotationProperty AnnotationProperty
	Literal            Literal
}

/* Prefix */

type Prefix struct {
	XMLName xml.Name `xml:"Prefix"`
	Name    string   `xml:"name,attr"`
	IRI     string   `xml:"IRI,attr"`
}

type Import struct {
	XMLName xml.Name `xml:"Import"`
	Value   string   `xml:",chardata"`
}

func GenerateOntology(filepath string) Ontology {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error opening file:", filepath, err)
		return Ontology{}
	}

	var ontology Ontology

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Decode(&ontology)
	return ontology
}
