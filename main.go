package main

import (
	"bootcamp/assignment"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Bootcamp week 5 assignment!")
	assignment.CeilNumber(42.42)
	assignment.AlphabetSoup("hello")
	assignment.StringMask("burcu", 0)
	assignment.WordSplit([2]string{"hellocat", "apple,bat,cat,goodbye,hello,yellow,why"})
	assignment.AddUint32(math.MaxUint32, 2)
	assignment.VariadicSet(4, 2, 5, 4, 2, 4)
}
