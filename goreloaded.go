package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You should try: go run . <inputfile> <outputfile>")
		return
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]
	originalText := GetTextFromInputFile(inputFileName)
	var formattedText []string
	for i := 0; i < len(originalText); i++ {
		formattedText = append(formattedText, GetFormattedText(originalText[i]))
	}
	for i := 0; i < len(formattedText); i++ {
		PlaceModifiedTextIntoOutputFile(formattedText[i], outputFileName)
	}
	fmt.Println("The program has been executed.")
}

func GetTextFromInputFile(fileName string) []string {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("unable to open the file: %v\n", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var originalText []string
	for fileScanner.Scan() {
		originalText = append(originalText, fileScanner.Text())
	}
	readFile.Close()
	return originalText
}

func PlaceModifiedTextIntoOutputFile(formattedText string, outputFileName string) {
	file, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("unable to write text to a file: %v\n", err)
	}
	defer file.Close()
	_, err2 := file.WriteString(formattedText + "\n")
	if err2 != nil {
		fmt.Println("Could not write text to example.txt")
	}
}

func GetFormattedText(originalText string) string {
	regex := regexp.MustCompile(`\(\s*(low|up|cap|hex|bin)(?:,\s*\d+)?\s*\)`)
	formattedText := ReplaceArticles(originalText)
	formattedText = HandleSingleQuotes(formattedText)
	formattedText = FixPunctuation(formattedText)
	formattedText = ApplyFormattingCommand(formattedText, regex)
	return formattedText
}

func ApplyFormattingCommand(originalText string, regex *regexp.Regexp) string {
	matches := regex.FindAllString(originalText, -1)

	for _, command := range matches {
		countWordsBeforeFormattingCommand := ExtractNumberFromCommand(command)
		segment := regex.Split(originalText, 2)
		wordsOfSegmentedText := strings.Fields(segment[0])

		if countWordsBeforeFormattingCommand >= len(wordsOfSegmentedText) {
			countWordsBeforeFormattingCommand = len(wordsOfSegmentedText)
		}

		for j := len(wordsOfSegmentedText) - 1; j >= len(wordsOfSegmentedText)-countWordsBeforeFormattingCommand; j-- {
			wordsOfSegmentedText[j] = applyCommand(command, wordsOfSegmentedText[j])
		}

		segment[0] = strings.Join(wordsOfSegmentedText, " ")
		originalText = strings.Join(segment, "")
	}

	return originalText
}

func applyCommand(command, wordToFormat string) string {
	switch {
	case regexp.MustCompile(`\(\s*low(?:,\s*\d+)?\s*\)`).MatchString(command):
		return ToLower(wordToFormat)
	case regexp.MustCompile(`\(\s*up(?:,\s*\d+)?\s*\)`).MatchString(command):
		return ToUpper(wordToFormat)
	case regexp.MustCompile(`\(\s*cap(?:,\s*\d+)?\s*\)`).MatchString(command):
		return ToCapitalize(wordToFormat)
	case regexp.MustCompile(`\(\s*hex\s*\)`).MatchString(command):
		return ConvertHexToDecimal(wordToFormat)
	case regexp.MustCompile(`\(\s*bin\s*\)`).MatchString(command):
		return ConvertBinToDecimal(wordToFormat)
	default:
		return wordToFormat
	}
}

func ReplaceArticles(originalText string) string {
	wordsOfOriginalText := strings.Fields(originalText)
	for i := 0; i < len(wordsOfOriginalText)-1; i++ {
		if wordsOfOriginalText[i] == "a" && IsStartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "an"
		}
		if wordsOfOriginalText[i] == "A" && IsStartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "An"
		}
		if wordsOfOriginalText[i] == "an" && !IsStartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "a"
		}
		if wordsOfOriginalText[i] == "An" && !IsStartedWithVowel(wordsOfOriginalText[i+1]) {
			wordsOfOriginalText[i] = "A"
		}
	}
	return strings.Join(wordsOfOriginalText, " ")
}

func FixPunctuation(originalText string) string {
	originalTextRunes := []rune(originalText)
	for i := len(originalTextRunes) - 1; i > 0; i-- {
		if isPunctuation(originalTextRunes[i]) && originalTextRunes[i-1] == ' ' {
			originalTextRunes = deleteSpace(originalTextRunes, i-1)
		}
	}
	for i := 0; i < len(originalTextRunes)-1; i++ {
		if isPunctuation(originalTextRunes[i]) && (!unicode.IsPunct(originalTextRunes[i+1]) && originalTextRunes[i+1] != ' ') {
			originalTextRunes = addSpace(originalTextRunes, i+1)
		}
	}
	return string(originalTextRunes)
}

func deleteSpace(slice []rune, index int) []rune {
	return append(slice[:index], slice[index+1:]...)
}

func addSpace(slice []rune, index int) []rune {
	newArray := append(slice[:index], append([]rune{' '}, slice[index:]...)...)
	return newArray
}

func isPunctuation(r rune) bool {
	punctuations := ".,:;?!"
	for _, punctuation := range punctuations {
		if punctuation == r {
			return true
		}
	}
	return false
}

func HandleSingleQuotes(originalText string) string {
	formattedText := ReplaceApostrophe(originalText)
	singleQuotesRegex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	formattedText = singleQuotesRegex.ReplaceAllString(formattedText, "'$1'")
	return formattedText
}

func ReplaceApostrophe(originalText string) string {
	regexS := regexp.MustCompile(`\b+'s\b`)
	formattedText := regexS.ReplaceAllLiteralString(originalText, "‘s")
	regexT := regexp.MustCompile(`\b+'t\b`)
	formattedText = regexT.ReplaceAllLiteralString(formattedText, "‘t")
	regexLL := regexp.MustCompile(`\b+'ll\b`)
	formattedText = regexLL.ReplaceAllLiteralString(formattedText, "‘ll")
	regexRe := regexp.MustCompile(`\b+'re\b`)
	formattedText = regexRe.ReplaceAllLiteralString(formattedText, "‘re")
	return formattedText
}

func ConvertHexToDecimal(hexNumber string) string {
	decimalNumber, err := strconv.ParseInt(hexNumber, 16, 64)
	// case when word before command is not hex number.
	if err != nil {
		fmt.Println("Unable to convert hex to decimal.")
		return hexNumber
	}
	hexNumber = strconv.Itoa(int(decimalNumber))
	return hexNumber
}

func ConvertBinToDecimal(binNumber string) string {
	decimalNumber, err := strconv.ParseInt(binNumber, 2, 64)
	// case when word before command is not hex number.
	if err != nil {
		fmt.Println("Unable to convert bin to decimal.")
		return binNumber
	}
	binNumber = strconv.Itoa(int(decimalNumber))
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

func ExtractNumberFromCommand(command string) int {
	var numberStr string
	for _, char := range command {
		if unicode.IsDigit(char) {
			numberStr += string(char)
		} else if numberStr != "" {
			break
		}
	}

	if numberStr == "" {
		return 1
	}

	if len(numberStr) > 10 {
		fmt.Printf("Warning: unable to handle too big numbers <3 (%v)\n", numberStr)
		return 0
	}
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0
	}

	return number
}

func IsStartedWithVowel(nextWordAfterArticle string) bool {
	vowels := "aeiouhAEIOUH"
	for _, letter := range vowels {
		if letter == rune(nextWordAfterArticle[0]) {
			return true
		}
	}
	return false
}
