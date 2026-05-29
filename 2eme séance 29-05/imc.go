package main

import (
	"fmt"
	"math"
)

const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

const (
	Nom    = "Berthillot"
	Prénom = "Antoine"
)

func main() {
	var poids, taille float64 = 74.5, 1.72

	imc := poids / math.Pow(taille, 2)
	fmt.Println(imc)
	fmt.Printf("%.2f\n", imc)

	switch {
	case imc < IMCMaigreur:
		fmt.Println("Maigreur")
	case imc >= IMCMaigreur && imc < IMCNormal:
		fmt.Println("Normal")
	case imc >= IMCNormal && imc < IMCSurpoids:
		fmt.Println("Surpoids")
	default:
		fmt.Println("Obésité")
	}
}
