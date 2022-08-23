package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderadad":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
}