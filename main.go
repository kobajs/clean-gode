package main

import "fmt"

func changeName(oldName *string, newName string) {
	*oldName = newName
}

func main() {
	name := "Souto"

	changeName(&name, "Bruno")

	fmt.Println(name)
}
