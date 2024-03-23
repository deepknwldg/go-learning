package main

import (
	"fmt"
)

var packageLevelVar = "I'm outside function"

func main() {
	fmt.Println(packageLevelVar)

	var str string
	var strWithInitialization string = "I'm initialized string"
	var str1, str2 string
	fmt.Println(str, strWithInitialization, str1, str2)

	var strWithoutExplicitType = "I'm string data type"
	strShorthandDeclaration := "I'm string with shorthand declaration" // inside function only
	fmt.Println(strWithoutExplicitType, strShorthandDeclaration)

	const myStr = "I'm string constant"
}
