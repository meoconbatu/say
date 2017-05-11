package say

import "strings"

const testVersion = 1

var tensNames = []string{
	"",
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var numNames = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}
var spellOuts = []struct {
	uint64
	string
}{
	{1000000000000000000, "quintillion"},
	{1000000000000000, "quadrillion"},
	{1000000000000, "trillion"},
	{1000000000, "billion"},
	{1000000, "million"},
	{1000, "thousand"},
	{1, ""},
}

func Say(number uint64) (output string) {
	if number == 0 {
		return numNames[0]
	}
	for _, s := range spellOuts {
		a := number / s.uint64
		if a > 0 {
			output += " " + Say100(int(a)) + " " + s.string
			number = number % s.uint64
		}
	}
	return strings.Trim(output, " ")
}
func Say100(number int) (output string) {
	a := number / 100
	if a > 0 {
		output = numNames[a] + " hundred"
	}
	if b := number % 100; b > 0 {
		if output != "" {
			output = output + " "
		}
		output += SayBasicCase(b)
	} else if output == "" && b == 0 {
		output += SayBasicCase(b)
	}
	return
}
func SayBasicCase(number int) (output string) {
	if number < 20 {
		return numNames[number]
	}
	a := number / 10
	b := number % 10
	output = tensNames[a]
	if b > 0 {
		output = output + "-" + numNames[b]
	}
	return
}
