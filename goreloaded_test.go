package main

import (
	"testing"
)

type startedWithVowelTest struct {
	nextWordAfterArticle string
	expected             bool
}

type extractNumberFromCommandTest struct {
	command  string
	expected int
}

type convertHexToDecimal struct {
	hexNumber string
	expected  string
}

type convertBinToDecimal struct {
	binNumber string
	expected  string
}

type replaceApostrophe struct {
	originalText string
	expected     string
}

type handleSingleQuotes struct {
	originalText string
	expected     string
}

var startedWithVowelTests = []startedWithVowelTest{
	{nextWordAfterArticle: "apple", expected: true},
	{nextWordAfterArticle: "hour", expected: true},
	{nextWordAfterArticle: "love", expected: false},
	{nextWordAfterArticle: "1eye", expected: false},
}

var extractNumberFromCommandTests = []extractNumberFromCommandTest{
	{command: "(cap)", expected: 1},
	{command: "(cap, 1)", expected: 1},
	{command: "(cap,1)", expected: 1},
	{command: "(cap,10)", expected: 10},
	{command: "(cap , 10)", expected: 10},
	{command: "(cap, hh)", expected: 1},
	{command: "(cap,)", expected: 1},
}

var convertHexToDecimalTests = []convertHexToDecimal{
	{hexNumber: "1E", expected: " 30 "},
	{hexNumber: "EEEE", expected: " 61166 "},
	{hexNumber: "The", expected: "The"},
}

var convertBinToDecimalTests = []convertBinToDecimal{
	{binNumber: "1E", expected: "1E"},
	{binNumber: "101", expected: " 5 "},
	{binNumber: "0000", expected: " 0 "},
}

var replaceApostropheTests = []replaceApostrophe{
	{originalText: "Hello! It's me", expected: "Hello! It‘s me"},
	{originalText: "Somebody's dog", expected: "Somebody‘s dog"},
	{originalText: "We're planning", expected: "We‘re planning"},
	{originalText: "doesn't know", expected: "doesn‘t know"},
	{originalText: "I'll see", expected: "I‘ll see"},
	{originalText: "we‘ll see", expected: "we‘ll see"},
}

var handleSingleQuotesTest = []handleSingleQuotes{
	{originalText: "actually ' hi'", expected: "actually 'hi'"},
	{originalText: "'hi ' ' hi'", expected: "'hi' 'hi'"},
	{originalText: "'  hi' she said", expected: "'hi' she said"},
	{originalText: "it‘s ' smart '", expected: "it‘s 'smart'"},
	{originalText: "hello ' hello hello ' hello", expected: "hello 'hello hello' hello"},
}

func TestIsStartedWithVowel(t *testing.T) {
	for _, test := range startedWithVowelTests {
		if output := IsStartedWithVowel(test.nextWordAfterArticle); output != test.expected {
			t.Errorf("Output is not equal to expected")
		}
	}
}

func TestExtractNumberFromCommand(t *testing.T) {
	for _, test := range extractNumberFromCommandTests {
		if output := ExtractNumberFromCommand(test.command); output != test.expected {
			t.Errorf("Output of %v is %v and not equal to expected %v", test.command, output, test.expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	actual := ToUpper("hello")
	expected := "HELLO"

	if actual != expected {
		t.Errorf("got %q, wanted %q", actual, expected)
	}
}

func TestToLower(t *testing.T) {
	actual := ToLower("HeLLo")
	expected := "hello"

	if actual != expected {
		t.Errorf("got %q, wanted %q", actual, expected)
	}
}

func TestToCapitalize(t *testing.T) {
	actual := ToCapitalize("heLLo")
	expected := "HeLLo"

	if actual != expected {
		t.Errorf("got %q, wanted %q", actual, expected)
	}
}

func TestConvertHexToDecimal(t *testing.T) {
	for _, test := range convertHexToDecimalTests {
		if output := ConvertHexToDecimal(test.hexNumber); output != test.expected {
			t.Errorf("actual: %v, expected: %v", output, test.expected)
		}
	}

}

func TestConvertBinToDecimal(t *testing.T) {
	for _, test := range convertBinToDecimalTests {
		if output := ConvertBinToDecimal(test.binNumber); output != test.expected {
			t.Errorf("actual: %v, expected: %v", output, test.expected)
		}
	}

}

func TestReplaceApostrophe(t *testing.T) {
	for _, test := range replaceApostropheTests {
		if output := ReplaceApostrophe(test.originalText); output != test.expected {
			t.Errorf("actual: %v, expected: %v", output, test.expected)
		}
	}
}

func TestHandleSingleQuotes(t *testing.T) {
	for _, test := range handleSingleQuotesTest {
		if output := HandleSingleQuotes(test.originalText); output != test.expected {
			t.Errorf("actual: %v, expected: %v", output, test.expected)
		}
	}

}
