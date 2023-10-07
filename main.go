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
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]
	textToModify := GetTextFromInputFile(inputFileName)
	fmt.Println("textToModify: " + textToModify)
	modifiedText := ModifyText(textToModify)
	//удалить
	// for i := 0; i < len(modifiedText); i++ {
	// 	fmt.Println("modifiedText: " + modifiedText[i])
	// }
	PlaceModifiedTextIntoOutputFile(modifiedText, outputFileName)
}
func GetTextFromInputFile(fileName string) string {
	text, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(text)
}
func PlaceModifiedTextIntoOutputFile(modifiedText []string, outputFileName string) {
	createdFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("unable to write text to a file: %v", err)
	}
	for i := 0; i < len(modifiedText); i++ {
		createdFile.WriteString(modifiedText[i] + " ")
	}
	createdFile.WriteString("\n")
	createdFile.Sync()
}

func ModifyText(toModify string) []string {
	splittedText := strings.Fields(toModify)
	splittedText = strings.Fields(ApplyLowWithNumber(toModify))
	splittedText = ApplyBinCommand(splittedText)
	splittedText = ApplyHexCommand(splittedText)
	splittedText = ApplyUpCommand(splittedText)
	splittedText = ApplyLowCommand(splittedText)
	splittedText = ApplyCapCommand(splittedText)

	return splittedText
}

func ApplyHexCommand(splittedText []string) []string {
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == "(hex)" {
			decimalNumber, err := strconv.ParseInt(splittedText[i-1], 16, 64)
			if err != nil {
				log.Fatalf("unable to convert hex to decimal: %v", err)
			}
			splittedText[i-1] = strconv.Itoa(int(decimalNumber))
			splittedText = deleteCommandAfterModification(splittedText, i)
		}
	}
	return splittedText
}

func ApplyBinCommand(splittedText []string) []string {
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == "(bin)" {
			decimalNumber, err := strconv.ParseInt(splittedText[i-1], 2, 64)
			if err != nil {
				log.Fatalf("unable to convert hex to decimal: %v", err)
			}
			splittedText[i-1] = strconv.Itoa(int(decimalNumber))
			splittedText = deleteCommandAfterModification(splittedText, i)
		}
	}
	return splittedText
}

func ApplyUpCommand(splittedText []string) []string {
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == "(up)" {
			splittedText[i-1] = strings.ToUpper(splittedText[i-1])
			splittedText = deleteCommandAfterModification(splittedText, i)
		}
	}
	return splittedText
}

func ApplyLowCommand(splittedText []string) []string {
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == "(low)" {
			splittedText[i-1] = strings.ToLower(splittedText[i-1])
			splittedText = deleteCommandAfterModification(splittedText, i)
		}
	}
	return splittedText
}

func ApplyCapCommand(splittedText []string) []string {
	for i := 1; i < len(splittedText); i++ {
		if splittedText[i] == "(cap)" {
			splittedText[i-1] = strings.Title(splittedText[i-1])
			splittedText = deleteCommandAfterModification(splittedText, i)
		}
	}
	return splittedText
}

func ApplyLowWithNumber(str string) string {
	regexOfLow := regexp.MustCompile(`\(low\s*,\s*\d+\)`)
	matchingOperations := regexOfLow.FindAllString(str, -1)
	strSplitted := regexOfLow.Split(str, -1)

	//потом удалить
	// for i := 0; i < len(strSplitted); i++ {
	// 	fmt.Print(strSplitted[i])
	// }
	// fmt.Println()

	for i := 0; i < len(strSplitted)-1; i++ {
		countWordsBeforeOperation := ExtractNumberFromOperation(matchingOperations[i])
		temp := strings.Fields(strSplitted[i])
		//fmt.Println(matchingOperations[i], countWordsBeforeOperation)
		//fmt.Println("temp " + temp[i])
		if countWordsBeforeOperation >= len(temp) {
			countWordsBeforeOperation = len(temp)
		}
		fmt.Println(matchingOperations[i], countWordsBeforeOperation)

		temp = toLower(temp, countWordsBeforeOperation)
		strSplitted[i] = strings.Join(temp, " ")
	}

	//потом удалить
	// for i := 0; i < len(strSplitted); i++ {
	// 	fmt.Print(strSplitted[i] + " ")
	// }
	// fmt.Println()
	return strings.Join(strSplitted, " ")
}

func deleteCommandAfterModification(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func toLower(temp []string, countWordsBeforeOperation int) []string {
	for i := len(temp) - 1; i >= len(temp)-countWordsBeforeOperation; i-- {
		temp[i] = strings.ToLower(temp[i])
		fmt.Println(temp[i])
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
		return 0
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0
	}

	return number
}
