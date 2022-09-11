package main

import "fmt"

func main() {
	//First Style
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FFFF",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}
	fmt.Println(colors)

	//Second Style
	var colours map[string]string
	fmt.Println(colours)

	//Third Style
	colours2 := make(map[int]string)
	colours2[10] = "#FFFFFF"
	fmt.Println(colours2)
	delete(colours2, 10)
	fmt.Println(colours2)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
