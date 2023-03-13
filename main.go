package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var stored []string

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

// recursion used if more than 3 letters in the user input
func combo(arr string, dataArray []string, index int, mylist []string) {
	perm := permutations(arr)

	//get all permutations of all letters, so 3, then 4, then 5, then so on and add those to the data array, then use binary search on each of them and see if they are in the dictionary
	// then add them to the global answer array
	//permutation from permutation so permutation abcd then permutations of each of its 3 letter permutations
	for i := 0; i < len(perm)-1; i++ {
		dataArray = append(dataArray, perm[i][0:index])
		dataArray = append(dataArray, perm[i])
	}

	for i := 0; i < len(dataArray)-1; i++ {
		if binarySearch(dataArray[i], mylist) {
			stored = append(stored, dataArray[i])
		}
	}
}

func getThem(word string, index int, mylist []string) {
	data := make([]string, index)
	combo(word, data, index, mylist)
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			var result []string
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func binarySearch(spot string, haystack []string) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < spot {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != spot {
		return false
	}

	return true
}
func unique(s []string) []string {
	inTo := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inTo[str]; !ok {
			inTo[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	dictionary, err := readLines("/home/k46t523/go/fin.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range dictionary {
		fmt.Println(i, line)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your letters: ")
	letters, _ := reader.ReadString('\n')
	fmt.Print("Your letters: " + letters)
	letters = letters[:len(letters)-1]
	length1 := len(letters)
	for index := 3; index < length1; index++ {
		getThem(letters, index, dictionary)
	}
	perm := permutations(letters)
	//only 3 letters
	if length1 == 3 {
		for j := 0; j < len(perm)-1; j++ {
			if binarySearch(perm[j], dictionary) {
				stored = append(stored, perm[j])
			}
		}
	}
	if binarySearch(letters, dictionary) {
		stored = append(stored, letters)
	}

	stored = unique(stored)
	fmt.Println("Answers for Wordscape: ")
	fmt.Println(stored)

}
