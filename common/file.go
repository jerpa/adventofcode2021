package common

import (
	"bufio"
	"log"
	"os"
)

// ReadFileStrings read file inte string slice
func ReadFileStrings(filename string) []string {
	res := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

// ReadInputFile shortcut for ReadFileStrings("./input")
func ReadInputFile() []string {
	return ReadFileStrings("./input")
}
