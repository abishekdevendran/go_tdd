package main

// import "fmt"

func Hello(person, lang string) string {
	prefix := "Hello"
	if person=="" {
		person="World"
	}
	switch lang {
		case "fr":
			prefix = "Bonjour"
		case "de":
			prefix = "Hallo"
		case "en":
		default:
			prefix = "Hello"
	}
	return prefix +", "+person
}