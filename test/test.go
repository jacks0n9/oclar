package main

import (
	"fmt"
	"oclar"
)

func main() {
	s := oclar.Style{"bluefg", "blackbg"}
	fmt.Println(s.Render("Hello"))
}
