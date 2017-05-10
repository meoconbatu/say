package say

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
	//	{100, "hundred"},
}

func Say(number uint64) (output string) {
	var b uint64
	if number < 20 {
		return numNames[number]
	}
	for _, s := range spellOuts {
		a := number / s.uint64
		b = number % s.uint64
		if a > 0 {
			if output != "" {
				output = output + " "
			}
			output = output + Say100(int(a)) + " " + s.string
			number = number % s.uint64
		}
	}
	if b > 0 {
		if output != "" {
			output = output + " "
		}
		output = output + Say100(int(b))

	}
	return
}
func Say100(number int) (output string) {
	a := int(number) / 100
	if a > 0 {
		output = numNames[a] + " hundred"
	}
	if b := int(number) - 100*a; b > 0 {
		if output != "" {
			output = output + " "
		}
		output = output + SayBasicCase(b)
	}
	return
}
func SayBasicCase(number int) (output string) {
	if number < 20 {
		return numNames[number]
	}
	a := number / 10
	b := number - a*10
	output = tensNames[a]
	if b > 0 {
		output = output + "-" + numNames[b]
	}
	return
}
