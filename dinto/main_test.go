package dinto

import (
	"testing"
)

//code to test main's ontology object

var ontology = GenerateOntology("./data/test.owl")

func TestPrefixes(t *testing.T) {
	var expectedPrefixLength = 11
	var actualPrefixLength = len(ontology.Prefixes)

	if actualPrefixLength != expectedPrefixLength {
		t.Error("Ontology Prefix amount is incorrect, expected", expectedPrefixLength, ", but got", actualPrefixLength)
	}

	var expectedPrefixIRI = "http://purl.obolibrary.org/obo#"
	var expectedPrefixName = "obo"
	var prefix = ontology.Prefixes[1]
	var actualPrefixIRI = prefix.IRI
	var actualPrefixName = prefix.Name
	if actualPrefixName != expectedPrefixName {
		t.Error("Ontology Prefix name is incorrect, expected",  expectedPrefixName, ", but got", actualPrefixName)
	}

	if actualPrefixIRI != expectedPrefixIRI {
		t.Error("Ontology Prefix name is incorrect, expected",  expectedPrefixIRI, ", but got", actualPrefixIRI)
	}
	
}

func TestImport(t *testing.T) {
	var expectedImportLength = 1
	var actualImportLength = len(ontology.Imports)

	if actualImportLength != expectedImportLength {
		t.Error("Ontology Imports amount is incorrect, expected", expectedImportLength, ", but got", actualImportLength)
	}

	var expectedImportsValue = "http://purl.obolibrary.org/obo/MF/external/ontology-metadata-slim.owl"
	var imports = ontology.Imports[0]
	var actualImportsValue = imports.Value
	if actualImportsValue != expectedImportsValue {
		t.Error("Ontology Import value is incorrect, expected",  expectedImportsValue, ", but got", actualImportsValue)
	}

}

func TestAnnotations(t *testing.T) {
	var expectedAnnotationsLength = 3
	var actualAnnotationsLength = len(ontology.Annotations)

	if actualAnnotationsLength != expectedAnnotationsLength {
		t.Error("Ontology Annotations amount is incorrect, expected", expectedAnnotationsLength, ", but got", actualAnnotationsLength)
	}

	var annotation = ontology.Annotations[2]
	var literal = annotation.Literal
	var actualIRIString = literal.Value
	var expectedIRIString = "This ontology has been created and is maintained by the Department of Computer Science at University Carlos III of Madrid."

	if actualIRIString != expectedIRIString {
		t.Error("Ontology Annotation Value string is incorrect, expected",  expectedIRIString, ", but got", actualIRIString)
	}

}

func TestDeclarations(t *testing.T) {
	var expectedDeclarationsLength = 5
	var actualDeclarationsLength = len(ontology.Declarations)

	if actualDeclarationsLength != expectedDeclarationsLength {
		t.Error("Ontology Declarations amount is incorrect, expected", expectedDeclarationsLength, ", but got", actualDeclarationsLength)
	}

	var expectedClassAbbreviatedIRI AbbreviatedIRI = "obo2:CHEBI_100147"
	var decl = ontology.Declarations[0]
	var class = decl.Class
	var actualClassAbbreviatedIRI = class.AbbreviatedIRI

	if actualClassAbbreviatedIRI != expectedClassAbbreviatedIRI {
		t.Error("Ontology Declaration Class Abbreviated IRI string is incorrect, expected",  expectedClassAbbreviatedIRI, ", but got", actualClassAbbreviatedIRI)
	}
}

func TestSubclasses(t *testing.T) {
	var expectedSubclassesLength = 2
	var actualSubclassesLength = len(ontology.Subclasses)

	if actualSubclassesLength != expectedSubclassesLength {
		t.Error("Ontology Subclasses amount is incorrect, expected", expectedSubclassesLength, ", but got", actualSubclassesLength)
	}

	var expectedSubClassAbbreviatedIRI AbbreviatedIRI = "obo2:OAE_0000287"
	var subclass = ontology.Subclasses[1]
	var class = subclass.Classes[0]
	var actualSubClassAbbreviatedIRI = class.AbbreviatedIRI

	if actualSubClassAbbreviatedIRI != expectedSubClassAbbreviatedIRI {
		t.Error("Ontology SubClass Abbreviated IRI string is incorrect, expected",  expectedSubClassAbbreviatedIRI, ", but got", actualSubClassAbbreviatedIRI)
	}
}

func TestAnnotationAssertions(t *testing.T) {
	var expectedAnnotationAssertionsLength = 1
	var actualAnnotationAssertionsLength = len(ontology.AnnotationAssertions)

	if actualAnnotationAssertionsLength != expectedAnnotationAssertionsLength {
		t.Error("Ontology Annotation Assertions amount is incorrect, expected", expectedAnnotationAssertionsLength, ", but got", actualAnnotationAssertionsLength)
	}

	var expectedAnnotationAssertAbbrevIRI AbbreviatedIRI = "obo2:DINTO_1812"
	var AnnotationAssertion = ontology.AnnotationAssertions[0]
	var actualAnnotationAssertAbbrevIRI = AnnotationAssertion.AbbreviatedIRI
	var literal = AnnotationAssertion.Literal
	var actualLiteralValue = literal.Value
	var expectedLiteralValue = "multidrug and toxin extrusion protein 1"
	if actualAnnotationAssertAbbrevIRI != expectedAnnotationAssertAbbrevIRI {
		t.Error("Ontology Annotation Assertion Abbreviated IRI string is incorrect, expected",  expectedAnnotationAssertAbbrevIRI, ", but got", actualAnnotationAssertAbbrevIRI)
	}

	if actualLiteralValue != expectedLiteralValue {
		t.Error("Ontology Annotation Assertion Literal value string is incorrect, expected",  expectedLiteralValue, ", but got", actualLiteralValue)
	}


}

func TestSubAnnotationProperties(t *testing.T) {
	var expectedSubAnnotationProperties = 1
	var actualSubAnnotationPropertiesLength = len(ontology.SubAnnotationProperties)

	if actualSubAnnotationPropertiesLength != expectedSubAnnotationProperties {
		t.Error("Ontology SubAnnotationProperties amount is incorrect, expected", expectedSubAnnotationProperties, ", but got", actualSubAnnotationPropertiesLength)
	}

	var subAnnotProperty = ontology.SubAnnotationProperties[0]
	var annotationProp = subAnnotProperty.AnnotationProperties[1]
	var expectedAnnotationPropIRI AbbreviatedIRI = "rdfs:label"
	var actualAnnotationPropIRI = annotationProp.AbbreviatedIRI

	if actualAnnotationPropIRI != expectedAnnotationPropIRI {
		t.Error("Ontology sub-annotation property IRI is incorrect, expected", expectedAnnotationPropIRI, ", but got", actualAnnotationPropIRI)
	}
}
