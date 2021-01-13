package main

import (
	"fmt"

	"github.com/SHsteak/hello-go/morestrings"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	p1 := product{1, "조상현의 연필", 50000}
	p2 := product{2, "조상현의 노트북", 3000000}

	fmt.Print("p1: ")
	fmt.Println(p1)
	fmt.Print("p2: ")
	fmt.Println(p2)

	fmt.Println("총 가격: ", sumPrice(&p1.price, &p2.price), "원")
}

type product struct {
	id    int
	name  string
	price int
}

func sumPrice(p ...*int) int {
	result := 0
	for _, p := range p {
		result += *p
	}
	return result
}
