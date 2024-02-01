package main

import "github.com/saivarshith2000/gomb"

func main() {
	store := gomb.NewStore()
	gomb := gomb.NewGomb(8080, 8181, store) // TODO: Fetch ports from CLI
	errChan := make(chan error)
	gomb.Start(errChan)
}
