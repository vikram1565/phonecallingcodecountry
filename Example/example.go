package main

import (
	"log"
	phone "phonecallingcodecountry"
)

func main() {

	pw := phone.NewPhoneCountryWorld("Earth")
	country, findError := pw.FindCountry("+919744985588")
	if findError != nil {
		log.Fatal(findError)
	}
	println("Country Name : ", country)
}
