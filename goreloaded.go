package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("You should try: go run . <inputfile> <outputfile>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]
	originalText := GetTextFromInputFile(inputFileName)
	formattedText := TextFormatting(originalText)
	PlaceModifiedTextIntoOutputFile(formattedText, outputFileName)
}

func GetTextFromInputFile(fileName string) string {
	originalText, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(originalText)
}

func PlaceModifiedTextIntoOutputFile(formattedText string, outputFileName string) {
	createdFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("unable to write text to a file: %v", err)
	}
	createdFile.WriteString(formattedText)
	createdFile.WriteString("\n")
	createdFile.Sync()
}

func TextFormatting(originalText string) string {
	originalText = FixPunctuation(originalText)
	originalText = HandleSingleQuotes(originalText)
	originalText = ReplaceArticles(originalText)

	regexOfHex := regexp.MustCompile(`\(hex\)`)
	regexOfBin := regexp.MustCompile(`\(bin\)`)
	regexOfLow := regexp.MustCompile(`\(low(?:,\s*\d+)?\)`)
	regexOfUp := regexp.MustCompile(`\(up(?:,\s*\d+)?\)`)
	regexOfCap := regexp.MustCompile(`\(cap(?:,\s*\d+)?\)`)

	originalText = ApplyFormattingComand(originalText, regexOfHex, ConvertHexToDecimal)
	originalText = ApplyFormattingComand(originalText, regexOfBin, ConvertBinToDecimal)
	originalText = ApplyFormattingComand(originalText, regexOfLow, ToLower)
	originalText = ApplyFormattingComand(originalText, regexOfUp, ToUpper)
	originalText = ApplyFormattingComand(originalText, regexOfCap, ToCapitalize)

	return originalText
}

func ApplyFormattingComand(originalText string, regex *regexp.Regexp, transformFunc func(string) string) string {
	matchingCommands := regex.FindAllString(originalText, -1)
	segmentedOriginalText := regex.Split(originalText, -1)

	for i := 0; i < len(segmentedOriginalText)-1; i++ {
		countWordsBeforeFormattingCommand := ExtractNumberFromCommand(matchingCommands[i])
		wordsOfSegmentedText := strings.Fields(segmentedOriginalText[i])

		if countWordsBeforeFormattingCommand >= len(wordsOfSegmentedText) {
			countWordsBeforeFormattingCommand = len(wordsOfSegmentedText)
		}

		for i := len(wordsOfSegmentedText) - 1; i >= len(wordsOfSegmentedText)-countWordsBeforeFormattingCommand; i-- {
			wordsOfSegmentedText[i] = transformFunc(wordsOfSegmentedText[i])
		}
		segmentedOriginalText[i] = strings.Join(wordsOfSegmentedText, " ")
	}
	return strings.Join(segmentedOriginalText, "")
}

func ReplaceArticles(originalText string) string {
	wordsOfOriginalText := strings.Fields(originalText)
	for i := 0; i < len(wordsOfOriginalText)-1; i++ {
		if wordsOfOriginalText[i] == "a" && StartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "an"
		}
		if wordsOfOriginalText[i] == "A" && StartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "An"
		}
		if wordsOfOriginalText[i] == "an" && !StartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "a"
		}
		if wordsOfOriginalText[i] == "An" && !StartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "A"
		}
	}
	return strings.Join(wordsOfOriginalText, " ")
}

func FixPunctuation(originalText string) string {
	regexOfPunc := regexp.MustCompile(`\s*([!?.,:;]+)`)
	formattedText := regexOfPunc.ReplaceAllString(originalText, "$1 ")
	return formattedText
}

func HandleSingleQuotes(originalText string) string {
	originalText = ReplaceApostrophe(originalText)
	singleQuotesRegex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	formattedText := singleQuotesRegex.ReplaceAllString(originalText, "'$1'")
	return formattedText
}

func ReplaceApostrophe(originalText string) string {
	regex := regexp.MustCompile(`\b+'s\b`)
	formattedText := regex.ReplaceAllLiteralString(originalText, "‘s")
	return formattedText
}

func ConvertHexToDecimal(hexNumber string) string {
	decimalNumber, err := strconv.ParseInt(hexNumber, 16, 64)
	// case when word before command is not hex number.
	if err != nil {
		fmt.Println("Unable to convert hex to decimal.")
	}
	hexNumber = " " + strconv.Itoa(int(decimalNumber)) + " "

	return hexNumber
}

func ConvertBinToDecimal(binNumber string) string {
	decimalNumber, err := strconv.ParseInt(binNumber, 2, 64)
	// case when word before command is not hex number.
	if err != nil {
		fmt.Println("Unable to convert hex to decimal.")
	}
	binNumber = " " + strconv.Itoa(int(decimalNumber)) + " "

	return binNumber
}

func ToCapitalize(wordToFormat string) string {
	return strings.Title(wordToFormat)
}
func ToLower(wordToFormat string) string {
	return strings.ToLower(wordToFormat)
}

func ToUpper(wordToFormat string) string {
	return strings.ToUpper(wordToFormat)
}

func ExtractNumberFromCommand(str string) int {
	var numberStr string
	foundDigit := false

	for _, char := range str {
		if unicode.IsDigit(char) {
			numberStr += string(char)
			foundDigit = true
		} else if foundDigit {
			break
		}
	}

	if numberStr == "" {
		return 1
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0
	}

	return number
}

func StartedWithVowel(nextWordAfterArticle string) bool {
	vowels := "aeiouhAEIOUH"
	for _, letter := range vowels {
		if letter == rune(nextWordAfterArticle[0]) {
			return true
		}
	}

	return false
}
