package main

import (
	"testing"
)

var ontology = GenerateOntology("./data/test.xml")

func TestPrefixes(t *testing.T) {
	var expectedPrefixLength = 11
	var actualPrefixLength = len(ontology.Prefixes)

	if actualPrefixLength != expectedPrefixLength {
		t.Error("Ontology Prefix amount is incorrect, expected",  expectedPrefixLength, ", but got", actualPrefixLength)
	}
}

func TestImport(t *testing.T) {
	var expectedImportLength = 1
	var actualImportLength = len(ontology.Imports)

	if actualImportLength != expectedImportLength {
		t.Error("Ontology Imports amount is incorrect, expected",  expectedImportLength, ", but got", actualImportLength)
	}
}

func TestAnnotations(t *testing.T) {
	var expectedAnnotationsLength = 3
	var actualAnnotationsLength = len(ontology.Annotations)

	if actualAnnotationsLength != expectedAnnotationsLength {
		t.Error("Ontology Annotations amount is incorrect, expected",  expectedAnnotationsLength, ", but got", actualAnnotationsLength)
	}
}

func TestDeclarations(t *testing.T) {
	var expectedDeclarationsLength = 5
	var actualDeclarationsLength = len(ontology.Declarations)

	if actualDeclarationsLength != expectedDeclarationsLength {
		t.Error("Ontology Declarations amount is incorrect, expected",  expectedDeclarationsLength, ", but got", actualDeclarationsLength)
	}
}

func TestSubclasses(t *testing.T) {
	var expectedSubclassesLength = 2
	var actualSubclassesLength = len(ontology.Subclasses)

	if actualSubclassesLength != expectedSubclassesLength {
		t.Error("Ontology Subclasses amount is incorrect, expected",  expectedSubclassesLength, ", but got", actualSubclassesLength)
	}
}

func TestAnnotationAssertions(t *testing.T) {
	var expectedAnnotationAssertionsLength = 1
	var actualAnnotationAssertionsLength = len(ontology.AnnotationAssertions)

	if actualAnnotationAssertionsLength != expectedAnnotationAssertionsLength {
		t.Error("Ontology Annotation Assertions amount is incorrect, expected",  expectedAnnotationAssertionsLength, ", but got", actualAnnotationAssertionsLength)
	}
}

func TestSubAnnotationProperties(t *testing.T) {
	var expectedSubAnnotationProperties = 1
	var actualSubAnnotationPropertiesLength = len(ontology.SubAnnotationProperties)

	if actualSubAnnotationPropertiesLength != expectedSubAnnotationProperties {
		t.Error("Ontology SubAnnotationProperties amount is incorrect, expected",  expectedSubAnnotationProperties, ", but got", actualSubAnnotationPropertiesLength)
	}
}

