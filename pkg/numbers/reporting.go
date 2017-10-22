package numbers

import "strconv"

type Reporter struct{}

func (r *Reporter) Report(input int) string {
	p := &Processor{}
	if p.ByThreeAndFive(input) {
		return "FizzBuzz"
	} else if p.ByFive(input) {
		return "Buzz"
	} else if p.ByThree(input) {
		return "Fizz"
	}
	return strconv.Itoa(input)
}
