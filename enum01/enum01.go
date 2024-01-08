package main

import "fmt"

type Season int64

const (
	Summer Season = iota + 1
	Autumn
	Winter
	Spring
)

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func (s Season) String() string {
	switch s {
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	case Spring:
		return "spring"
	}
	return "unknown"
}

func main() {
	s := Season(Winter)
	fmt.Printf("Season name: %v Season value: %v", s.String(), int64(s))
	//printSeason()
}
