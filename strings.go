package main

import "fmt"

func main() {
	text1 := "\"what's that\", he said"
	text2 := `"what's that", he said`
	fmt.Println(text1, "\n", text2)

	book := "The Spirit Level" +
		" by Richard Wilkinson"

	book += " and Kate Pickett"

	fmt.Println(book)

	fmt.Println("Josey" < "Jose", "Josey" == "Jose")
	
}
