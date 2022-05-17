package main

import "fmt"

type SampleStruct struct {
	property int
}

func (p *SampleStruct) UpdatePropertyWithoutPointer() {
	p.property = 8
	fmt.Println(p.property, "at", &p.property)
}

func (p *SampleStruct) UpdateProperty() {
	p.property = 8
	fmt.Println(p.property, "at", &p.property)
}

func main() {
	p := SampleStruct{1}
	fmt.Println(p.property, "at", &p.property)
	p.UpdateProperty()
	fmt.Println(p.property)
}
