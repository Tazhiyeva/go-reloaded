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
type fixPunctuation struct {
	originalText string
	expected     string
}
type replaceArticles struct {
	originalText string
	expected     string
}
type getFormattedText struct {
	originalText string
	expected     string
}

var startedWithVowelTests = []startedWithVowelTest{
	{nextWordAfterArticle: "eight", expected: true},
	{nextWordAfterArticle: "hour", expected: true},
	{nextWordAfterArticle: "five", expected: false},
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
	{hexNumber: "1E", expected: "30"},
	{hexNumber: "EEEE", expected: "61166"},
	{hexNumber: "The", expected: "The"},
}
var convertBinToDecimalTests = []convertBinToDecimal{
	{binNumber: "1E", expected: "1E"},
	{binNumber: "101", expected: "5"},
	{binNumber: "0000", expected: "0"},
}
var replaceApostropheTests = []replaceApostrophe{
	{originalText: "Hello! It's me", expected: "Hello! It‘s me"},
	{originalText: "Somebody's dog", expected: "Somebody‘s dog"},
	{originalText: "We're planning", expected: "We‘re planning"},
	{originalText: "doesn't know", expected: "doesn‘t know"},
	{originalText: "I'll see", expected: "I‘ll see"},
	{originalText: "we‘ll see", expected: "we‘ll see"},
}
var handleSingleQuotesTests = []handleSingleQuotes{
	{originalText: "actually ' hi'", expected: "actually 'hi'"},
	{originalText: "'hi ' ' hi'", expected: "'hi' 'hi'"},
	{originalText: "'  hi' she said", expected: "'hi' she said"},
	{originalText: "it‘s ' smart '", expected: "it‘s 'smart'"},
	{originalText: "hello ' hello hello ' hello", expected: "hello 'hello hello' hello"},
	{originalText: "afagf '     asdf   don't   adfd f     '", expected: "afagf 'asdf   don‘t   adfd f'"},
}
var fixPunctuationTests = []fixPunctuation{
	{originalText: "actually , hi", expected: "actually, hi"},
	{originalText: "actual ,hi", expected: "actual, hi"},
	{originalText: "'hi !? hi'", expected: "'hi!? hi'"},
	{originalText: " , hi she said", expected: ", hi she said"},
	{originalText: "it‘s. smart : ", expected: "it‘s. smart: "},
	{originalText: "hello ? hello hello ! hello.", expected: "hello? hello hello! hello."},
	{originalText: "hello ? , ,     ,    , hello hello ! hello.", expected: "hello?,,,, hello hello! hello."},
}
var replaceArticlesTests = []replaceArticles{
	{originalText: "an apple", expected: "an apple"},
	{originalText: "a book", expected: "a book"},
	{originalText: "an five", expected: "a five"},
	{originalText: "a eight", expected: "an eight"},
	{originalText: "A apple", expected: "An apple"},
	{originalText: "An apple", expected: "An apple"},
	{originalText: "A book", expected: "A book"},
	{originalText: "An book", expected: "A book"},
}
var getFormattedTextTests = []getFormattedText{
	{
		originalText: "The (hex) AB (hex) CDEF (hex) 1234 (hex) GHIJKLMN",
		expected:     "The 171 52719 4660 GHIJKLMN",
	},
	{
		originalText: "Invalid hex number: (hex) XYZ",
		expected:     "Invalid hex number: XYZ",
	},
	{
		originalText: "No hex numbers here",
		expected:     "No hex numbers here",
	},
	{
		originalText: "(hex) ABC 12 3XYZ 4A5",
		expected:     " ABC 12 3XYZ 4A5",
	},
	{
		originalText: "(hex) .",
		expected:     ".",
	},
	{
		originalText: "The (bin) 10 (bin) 1010 (bin) 11110 (bin) GHIJKLMN",
		expected:     "The 2 10 30 GHIJKLMN",
	},
	{
		originalText: "Invalid hex number: (bin) XYZ",
		expected:     "Invalid hex number: XYZ",
	},
	{
		originalText: "No bin numbers here",
		expected:     "No bin numbers here",
	},
	{
		originalText: "(bin) ABC 12 3XYZ 4A5",
		expected:     " ABC 12 3XYZ 4A5",
	},
	{
		originalText: "(bin) .",
		expected:     ".",
	},
	{
		originalText: "HELLO. (low) I am good",
		expected:     "hello. I am good",
	},
	{
		originalText: "HELLO MY BABY, I Missed SO mUch (low, 5)",
		expected:     "HELLO MY baby, i missed so much",
	},
	{
		originalText: "I LIKE TO (low) READ (low) BOOKS",
		expected:     "I LIKE to read BOOKS",
	},
	{
		originalText: "Hello , I LIKE TO (low) READ. (low,2) BOOKS",
		expected:     "Hello, I LIKE to read. BOOKS",
	},
	{
		originalText: "HELLO. I am good",
		expected:     "HELLO. I am good",
	},
	{
		originalText: "HELLO (low, 10)",
		expected:     "hello",
	},
	{
		originalText: "HELLO (low, ABC)",
		expected:     "HELLO (low, ABC)",
	},
	{
		originalText: "HELLO (low, 2.) I am good",
		expected:     "HELLO (low, 2.) I am good",
	},
	{
		originalText: "123 (low) HELLO",
		expected:     "123 HELLO",
	},
	{
		originalText: "123 heheh heheh heheh (low, 99999999999) HELLO",
		expected:     "123 heheh heheh heheh HELLO",
	},
	{
		originalText: "hello. (up) I am good",
		expected:     "HELLO. I am good",
	},
	{
		originalText: "HELLO MY BABY, I Missed SO mUch (up, 5)",
		expected:     "HELLO MY BABY, I MISSED SO MUCH",
	},
	{
		originalText: "I LIKE to (up) read (up) BOOKS",
		expected:     "I LIKE TO READ BOOKS",
	},
	{
		originalText: "Hello , i like to (up) read. (up,2) BOOKS",
		expected:     "Hello, i like TO READ. BOOKS",
	},
	{
		originalText: "hello (up, 10)",
		expected:     "HELLO",
	},
	{
		originalText: "HELLO (up, ABC)",
		expected:     "HELLO (up, ABC)",
	},
	{
		originalText: "HELLO (up, 2.) I am good",
		expected:     "HELLO (up, 2.) I am good",
	},
	{
		originalText: "123 (up) HELLO",
		expected:     "123 HELLO",
	},
	{
		originalText: "hello. (cap) I am good",
		expected:     "Hello. I am good",
	},
	{
		originalText: "HELLO MY BABY, I Missed SO mUch (cap, 5).",
		expected:     "HELLO MY BABY, I Missed SO MUch.",
	},
	{
		originalText: "I LIKE to (cap) read (cap) BOOKS.",
		expected:     "I LIKE To Read BOOKS.",
	},
	{
		originalText: "Hello , i like to (cap) read. (cap,2) BOOKS.",
		expected:     "Hello, i like To Read. BOOKS.",
	},
	{
		originalText: "hello (cap, 10).",
		expected:     "Hello.",
	},
	{
		originalText: "HELLO (cap, ABC).",
		expected:     "HELLO (cap, ABC).",
	},
	{
		originalText: "HELLO (cap, 2.) I am good.",
		expected:     "HELLO (cap, 2.) I am good.",
	},
	{
		originalText: "123 (cap) HELLO.",
		expected:     "123 HELLO.",
	},
	{
		originalText: "one (up) , two , three (cap,2), ' four' , an FIVE, (low) ' six ,seven ?! a eight' 00111 (bin) FFF (hex) The end ... ",
		expected:     "ONE, Two, Three, 'four', a five, 'six, seven?! an eight' 7 4095 The end...",
	},
	{
		originalText: "I was sitting over there ,and then BAMM !!",
		expected:     "I was sitting over there, and then BAMM!!",
	},
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
	for _, test := range handleSingleQuotesTests {
		if output := HandleSingleQuotes(test.originalText); output != test.expected {
			t.Errorf("actual: %v, expected: %v", output, test.expected)
		}
	}
}
func TestFixPunctuation(t *testing.T) {
	for _, test := range fixPunctuationTests {
		if output := FixPunctuation(test.originalText); output != test.expected {
			t.Errorf("\nactual: %v\nexpected: %v", output, test.expected)
		}
	}
}
func TestReplaceArticles(t *testing.T) {
	for _, test := range replaceArticlesTests {
		if output := ReplaceArticles(test.originalText); output != test.expected {
			t.Errorf("\nactual: %v, expected: %v", output, test.expected)
		}
	}
}
func TestGetFormattedText(t *testing.T) {
	for _, test := range getFormattedTextTests {
		if output := GetFormattedText(test.originalText); output != test.expected {
			t.Errorf("%v\nactuaaal: %v \nexpected: %v", test.originalText, output, test.expected)
		}
	}
}
