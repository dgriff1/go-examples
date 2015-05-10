package main

import (
	"fmt"
	"math"
	"os"
)

var lookups = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func LookupNumber(num int) string {
	if ret := lookups[num]; ret != "" {
		return ret
	} else {
		ones := int(math.Mod(float64(num), 10))
		tens := num - ones
		return fmt.Sprintf("%s %s", lookups[tens], lookups[ones])
	}
}

func ConvertDollars(dollars int) string {
	var result string
	for dollars > 0 {
		place := int(math.Mod(float64(dollars), 100)) // will never be a float remainder
		dollars = dollars / 100
		result = fmt.Sprintf(" %d %s ", place, result)
		fmt.Printf("Place is %d dollars %d result %s\n", place, dollars, result)
	}
	return result
}

func ConvertCents(cents int) string {
	return "forty two cents"
}

func ParseString(s string) (int, int) {
	var dollars, dec int
	fmt.Sscanf(s, "%d.%d", &dollars, &dec)
	return dollars, dec
}

func ConvertToStr(s string) string {
	var integer, dec = ParseString(s)
	dollars := ConvertDollars(integer)
	cents := ConvertCents(dec)
	return dollars + " and " + cents
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fmt.Println(ParseString("123.45"))
}
