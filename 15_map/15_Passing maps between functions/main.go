package main

import "fmt"


func main() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		"AliceBlue":
		"#f0f8ff",
		"Coral":
		"#ff7F50",
		"DarkGray":
		"#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// Call the function to remove the specified key.
	removeColor(colors, "Coral")

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

// ..

// removeColor removes keys from the specified map.
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
