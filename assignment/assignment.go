package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	number1 := float64(x)    //initialize number 1 to float64
	number2 := float64(y)    //initialize number 2 to float64
	sum := number1 + number2 // add number1 and number2
	//initialize overflow to true
	overflow := true
	if sum > math.MaxUint32 { // if sum is overflowing, return sum of the numbers and overflow (true)
		return x + y, overflow
	}
	return x + y, !overflow //if sum is not overflowing, return sum of the numbers and not overflow (false)
}

func CeilNumber(f float64) float64 {
	return math.Ceil(f*4) / 4 //round the number to ceil quarter
}

func AlphabetSoup(s string) string {
	lowerCase := strings.ToLower(s)             //converted string to lower case
	splitString := strings.Split(lowerCase, "") //splitted the string with space
	sort.Strings(splitString)                   //sorted the words
	return strings.Join(splitString, "")        //added the words to a new string

}

func StringMask(s string, n uint) string { // Big-O of O(n)

	list := strings.Split(s, "") //splitted the string with space

	length := len(s) //string length
	star := int(n)   //uint convert to int

	if length == 0 || length == 1 { // if string length is 1 or 0 -> "" or "a"
		return "*"
	}

	for i := 0; i < length; i++ { //loop the length of the string
		// if star number is 0 or greater and equal than string length, return all string to "*"
		// or star number is smaller and equal than i, return string up to i number and star up to star number
		if star == 0 || star >= length || i >= star {
			list[i] = "*" //replace the list array to "*"
		}
	}
	return strings.Join(list, "") //added the words to a new string

}

func WordSplit(arr [2]string) string {
	// second element of array
	words := arr[1]
	//splitting the given string, it returns a slice
	sliceWords := strings.Split(words, ",")

	// first element of array
	word := arr[0]
	//length of the word
	wordSize := len(word)

	//empty array
	newArray := []string{}

	//loops through every element in sliceWords
	for _, v := range sliceWords {
		if strings.Contains(word, v) { //checks if word contains v(sliceWords element)
			//vArr contains v(sliceWords element) // 1- cat 2- hello
			vArr := []string{v}
			//append seperated newArray to vArr, and assign to newArray variable.
			newArray = append(vArr, newArray...) // 1- cat 2- hello cat
		}
	}
	//add commas between the elements of newArray, and assign to newArray2 variable.
	newArray2 := strings.Join(newArray, ",")

	//checks if newArray's length smaller than 2 and word size equals to zero
	if len(newArray) < 2 || wordSize == 0 {
		return "not possible" //return that
	} else { //else
		return newArray2 //return newArray2
	}

}

func VariadicSet(i ...interface{}) []interface{} {
	//make a map function, key -> interface{} value -> bool
	keys := make(map[interface{}]bool)

	list := []interface{}{} //empty interface

	//loops through every element(entry) in i interface
	for _, entry := range i {
		//initialized value to keys[entry] , checks if value is false,
		if _, value := keys[entry]; !value {
			keys[entry] = true // than keys[entry] returned to  true
			//append entry to list.
			list = append(list, entry)
		}
	}
	return list //return list

}
