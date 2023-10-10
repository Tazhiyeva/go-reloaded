package main

import (
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
	textToModify := GetTextFromInputFile(inputFileName)
	modifiedText := ModifyText(textToModify)
	PlaceModifiedTextIntoOutputFile(modifiedText, outputFileName)
}

func GetTextFromInputFile(fileName string) string {
	text, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(text)
}

func PlaceModifiedTextIntoOutputFile(modifiedText string, outputFileName string) {
	createdFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("unable to write text to a file: %v", err)
	}
	createdFile.WriteString(modifiedText)
	createdFile.WriteString("\n")
	createdFile.Sync()
}

func ModifyText(toModify string) string {
	toModify = ApplyConversionOperation(toModify, "(hex)", 16)
	toModify = ApplyConversionOperation(toModify, "(bin)", 2)
	toModify = FixPunctuation(toModify)
	toModify = handleSingleQuotes(toModify)
	toModify = ReplaceArticles(toModify)

	regexOfLow := regexp.MustCompile(`\(low(?:,\s*\d+)?\)`)
	regexOfUp := regexp.MustCompile(`\(up(?:,\s*\d+)?\)`)
	regexOfCap := regexp.MustCompile(`\(cap(?:,\s*\d+)?\)`)

	toModify = ApplyStringManipulationOperation(toModify, regexOfLow, ToLower)
	toModify = ApplyStringManipulationOperation(toModify, regexOfUp, ToUpper)
	toModify = ApplyStringManipulationOperation(toModify, regexOfCap, ToCapitalize)

	return toModify
}

func ApplyConversionOperation(str, operation string, base int) string {
	splittedText := strings.Fields(str)
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == operation {
			decimalNumber, err := strconv.ParseInt(splittedText[i-1], base, 64)
			if err != nil {
				log.Fatalf("unable to convert hex to decimal: %v", err)
			}
			splittedText[i-1] = strconv.Itoa(int(decimalNumber))
			splittedText = DeleteCommandAfterModification(splittedText, i)
		}
	}
	return strings.Join(splittedText, " ")
}

func DeleteCommandAfterModification(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func ApplyStringManipulationOperation(str string, regex *regexp.Regexp, transformFunc func([]string, int) []string) string {
	matchingOperations := regex.FindAllString(str, -1)
	strSplitted := regex.Split(str, -1)

	for i := 0; i < len(strSplitted)-1; i++ {
		countWordsBeforeOperation := ExtractNumberFromOperation(matchingOperations[i])
		temp := strings.Fields(strSplitted[i])

		if countWordsBeforeOperation >= len(temp) {
			countWordsBeforeOperation = len(temp)
		}

		temp = transformFunc(temp, countWordsBeforeOperation)
		strSplitted[i] = strings.Join(temp, " ")
	}

	return strings.Join(strSplitted, " ")
}

func ReplaceArticles(str string) string {
	splittedText := strings.Fields(str)
	for i := 0; i < len(splittedText)-1; i++ {
		if splittedText[i] == "a" && StartedWithVowel(splittedText[i+1]) {
			splittedText[i] = "an"
		}
	}
	return strings.Join(splittedText, " ")
}

func FixPunctuation(toModify string) string {
	regexOfPunc := regexp.MustCompile(`\s*([!?.,:;]+)`)
	result := regexOfPunc.ReplaceAllString(toModify, "$1 ")
	return result
}

func handleSingleQuotes(text string) string {
	singleQuotesRegex := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = singleQuotesRegex.ReplaceAllString(text, "'$1'")

	return text
}

func ToCapitalize(temp []string, countWordsBeforeOperation int) []string {
	for i := len(temp) - 1; i >= len(temp)-countWordsBeforeOperation; i-- {
		temp[i] = strings.Title(temp[i])
	}
	return temp
}

func ToLower(temp []string, countWordsBeforeOperation int) []string {
	for i := len(temp) - 1; i >= len(temp)-countWordsBeforeOperation; i-- {
		temp[i] = strings.ToLower(temp[i])
	}
	return temp
}

func ToUpper(temp []string, countWordsBeforeOperation int) []string {
	for i := len(temp) - 1; i >= len(temp)-countWordsBeforeOperation; i-- {
		temp[i] = strings.ToUpper(temp[i])
	}
	return temp
}

func ExtractNumberFromOperation(str string) int {
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

func StartedWithVowel(str string) bool {
	vowels := "aeiouhAEIOUH"
	for _, letter := range vowels {
		if rune(str[0]) == letter {
			return true
		}
	}

	return false
}
