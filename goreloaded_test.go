package main

import (
	"testing"
)

type conversionTest struct {
	originalText, formattingCommand, expectedOutput string
	base                                            int
}

var conversionTests = []conversionTest{
	conversionTest{
		originalText:      "The (hex) AB (hex) CDEF (hex) 1234 (hex) GHIJKLMN.",
		formattingCommand: "(hex)",
		expectedOutput:    "The 171 52719 4660 GHIJKLMN.",
		base:              16},
	conversionTest{
		originalText:      "Binary numbers: (bin) 1010 (bin) 1100 (bin) 1111.",
		formattingCommand: "(bin)",
		expectedOutput:    "Binary numbers: 10 12 1111.",
		base:              2},
	conversionTest{
		originalText:      "Invalid hex number: XYZ (hex) .",
		formattingCommand: "(hex)",
		expectedOutput:    "Invalid hex number: XYZ.",
		base:              16},
}

func TestApplyConversionCommand(t *testing.T) {
	for _, test := range conversionTests {
		if output := ApplyConversionCommand(test.originalText, test.formattingCommand, test.base); output != test.expectedOutput {
			t.Errorf("Output %q not equal to expected %q", output, test.expectedOutput)
		}
	}
}
