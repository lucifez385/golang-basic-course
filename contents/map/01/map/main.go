package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// colors := make(map[string]string)
	// colors["color name"] = "hex code of color"
	// delete(colors, "color name")
	updateMap(colors, "red", "#ee00000")
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexcode := range colors {
		fmt.Println(color, hexcode)
	}
}

func updateMap(colors map[string]string, color string, code string) {
	colors[color] = code
}
