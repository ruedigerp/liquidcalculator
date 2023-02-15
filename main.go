package main

import "fmt"

func main() {

	var dest, shot, nikotin, aroma, usearoma, base, hnik, nikml float32
	myFunc()
	fmt.Print("Dest mililiter: ")
	fmt.Scanln(&dest)
	fmt.Print("10 ml Nikotin (mg/ml): ")
	fmt.Scanln(&shot)
	fmt.Print("Nikotin (mg): ")
	fmt.Scanln(&nikotin)
	fmt.Print("Aroma (%): ")
	fmt.Scanln(&aroma)

	usearoma = dest / 100 * aroma
	base = dest - usearoma - 15
	hnik = nikotin / 2
	var mlnik float32
	mlnik = shot / 10
	nikml = 10 / mlnik * hnik
	fmt.Println("=======================================")
	fmt.Println("Deine Mischung sitzt sich zusammen aus.")
	fmt.Println("Aroma: ", usearoma, "ml")
	fmt.Println("Base: ", base, "ml")
	fmt.Println("Nikotin: ", nikml, "ml")
}

func myFunc() {
	fmt.Println("Test...")
}
