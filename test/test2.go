package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "ONE TWO THREE FOUR (low, 1) FIVE SIX SEVEN EIGHT NINE (low, 2) TEN ELEVEN TWELVE THIRTEEN FOURTEEN FIFTEEN SIXTEEN"
	regexOfLow := regexp.MustCompile(`\(low\s*,\s*\d+\)`)
	matchingOperations := regexOfLow.FindAllString(str, -1)
	strSplitted := regexOfLow.Split(str, -1)

	//потом удалить
	for i := 0; i < len(strSplitted); i++ {
		fmt.Print(strSplitted[i])
	}
	fmt.Println()

	for i := 0; i < len(strSplitted)-1; i++ {
		countWordsBeforeOperation := ExtractNumberFromOperation2(matchingOperations[i])
		temp := strings.Fields(strSplitted[i])
		fmt.Println(matchingOperations[i], countWordsBeforeOperation)
		fmt.Println("temp " + temp[i])
		if countWordsBeforeOperation >= len(temp) {
			countWordsBeforeOperation = len(temp)
		}
		fmt.Println(matchingOperations[i], countWordsBeforeOperation)

		temp = toLower2(temp, countWordsBeforeOperation)
		strSplitted[i] = strings.Join(temp, " ")
	}

	//потом удалить
	for i := 0; i < len(strSplitted); i++ {
		fmt.Print(strSplitted[i] + " ")
	}
	fmt.Println()

}

func toLower2(temp []string, countWordsBeforeOperation int) []string {
	for i := len(temp) - 1; i >= len(temp)-countWordsBeforeOperation; i-- {
		temp[i] = strings.ToLower(temp[i])
		fmt.Println(temp[i])
	}
	return temp
}

func ExtractNumberFromOperation2(str string) int {
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
