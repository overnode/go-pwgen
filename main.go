package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	length := flag.Int("l", 8, "Length of password")
	iterations := flag.Int("n", 1, "Number of generated passwords")
	lc := flag.Bool("lc", true, "Lowercase letters")
	uc := flag.Bool("uc", true, "Uppercase letters")
	num := flag.Bool("num", true, "Numbers")
	special := flag.Bool("symbol", false, "Symbols")
	writeToFile := flag.String("f", "", "Write to specified filename")

	flag.Parse()

	result := make([]string, 0)

	genparam := RandomStringParameters{
		length:    *length,
		lowercase: *lc,
		uppercase: *uc,
		numeric:   *num,
		special:   *special,
	}

	timeStart := time.Now()
	for i := 0; i <= *iterations-1; i++ {
		passwd := GenRandomString(genparam)
		result = append(result, passwd)
	}

	for _, pw := range result {
		fmt.Println(pw)
	}

	if len(*writeToFile) > 0 {
		writeSliceToFile(*writeToFile, &result)
	}
	timeStop := time.Now()

	fmt.Println("---")
	fmt.Printf("Duration %s\n", timeStop.Sub(timeStart))
}

type RandomStringParameters struct {
	length    int
	lowercase bool
	uppercase bool
	numeric   bool
	special   bool
}

func GenRandomString(prm RandomStringParameters) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	bigLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "1234567890"
	symbols := "$%&!#$?-_@"

	var passwordArr string

	if prm.lowercase {
		passwordArr += letters
	}
	if prm.uppercase {
		passwordArr += bigLetters
	}
	if prm.numeric {
		passwordArr += numbers
	}
	if prm.special {
		passwordArr += symbols
	}

	charArray := strings.Split(passwordArr, "")

	var result []string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= prm.length-1; i++ {
		rn := r.Intn(len(charArray))
		result = append(result, charArray[rn])
	}

	return strings.Join(result, "")
}

func writeSliceToFile(filename string, list *[]string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
	for _, element := range *list {
		file.WriteString(element + "\n")
	}

	err = file.Sync()
	if err != nil {
		fmt.Println("Error writing file: ", err)
	} else {
		fmt.Println("\nPasswords written to " + filename)
	}
}
