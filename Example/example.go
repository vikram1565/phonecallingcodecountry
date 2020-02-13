package main

import (
	"fmt"
	"log"

	spinner "github.com/vikram1565/gospinners"
	phone "github.com/vikram1565/phonecallingcodecountry"
)

func main() {

	pw := phone.NewPhoneCountryWorld("Earth")
	country, findError := pw.FindCountry("+91 9744985588")
	if findError != nil {
		log.Fatal(findError)
	}

	si := spinner.New("ArrowSpinner", 5)
	si.StartSpinner()

	fmt.Println("\nCountry Names : ", country)
}
