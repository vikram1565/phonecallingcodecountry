package world

// countryphoneCode - struct
type countryphoneCode struct {
	Name string
	Code string
	Next *countryphoneCode
}

// World - struct
type World struct {
	name string
	Head *countryphoneCode
}

// CreateWorld - create world
func CreateWorld(name string) *World {
	w := &World{
		name: name,
	}
	w.addCountry("India", "+91")
	w.addCountry("Hong Kong", "+322")
	return w
}

// addCountry - add country in given world
func (p *World) addCountry(name, code string) {
	countryPhone := &countryphoneCode{
		Name: name,
		Code: code,
	}
	if p.Head == nil {
		p.Head = countryPhone
	} else {
		currentCountry := p.Head
		for currentCountry.Next != nil {
			currentCountry = currentCountry.Next
		}
		currentCountry.Next = countryPhone
	}
	return
}
