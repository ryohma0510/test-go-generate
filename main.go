package main

import "fmt"

//go:generate go run github.com/hairyhenderson/gomplate/v3/cmd/gomplate --datasource source=animal.yaml --file enum.tmpl --out animal_gen.go
//go:generate go run github.com/hairyhenderson/gomplate/v3/cmd/gomplate --datasource source=animal.yaml --file proto.tmpl --out animal_gen.proto

func main() {
	fmt.Println("Hello World")
}
