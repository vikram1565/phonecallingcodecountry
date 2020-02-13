package phonecallingcodecountry

import (
	"errors"
	"strings"

	"github.com/vikram1565/phonecallingcodecountry/world"
)

// Phone interface
type Phone interface {
	FindCountry(code string) (string, error)
}

// PhoneWorld - struct
type PhoneWorld struct {
	pw *world.World
}

// NewPhoneCountryWorld - create new world
func NewPhoneCountryWorld(worldName string) *PhoneWorld {
	if worldName == "" {
		panic("world name is empty")
	}
	return &PhoneWorld{
		pw: world.CreateWorld(worldName),
	}
}

// FindCountry - returns code matching country name
func (w *PhoneWorld) FindCountry(phoneNumber string) (string, error) {

	if phoneNumber == "" {
		return "", errors.New("phone number is empty")
	}
	if !strings.HasPrefix(phoneNumber, "+") {
		return "", errors.New("Invalid phone number")
	}
	currentCountry := w.pw.Head
	names := ""
	for currentCountry.Next != nil {
		if strings.HasPrefix(phoneNumber, currentCountry.Code) {
			names += currentCountry.Name + ", "
		}
		currentCountry = currentCountry.Next
	}
	if names == "" {
		return "", errors.New("country not found")
	}
	if strings.HasSuffix(names, ", ") {
		names = names[:len(names)-len(", ")]
	}
	return names, nil
}
