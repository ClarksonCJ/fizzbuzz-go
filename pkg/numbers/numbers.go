package numbers

type Processor struct{}

func (p *Processor) ByThree(input int) bool {
	return input%3 == 0
}

func (p *Processor) ByFive(input int) bool {
	return input%5 == 0
}

func (p *Processor) ByThreeAndFive(input int) bool {
	return p.ByThree(input) && p.ByFive(input)
}
