package main

import "github.com/kid1110/Short-Link-Worker/pkg"

func main() {
	u := pkg.GenerateByMurmurHash("qsqwersdfawefasdfvjasfhqkwehrlqlwejrlqwerjqkwlhsfakdhcxvhkljksadfl")
	println(u)
}