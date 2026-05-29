package main

import "fmt"

func main() {
	var a, b float64
	var op string
	for {
		fmt.Scan(&a, &b, &op)
		if op == "quit" {
			break
		}
		var res, err = operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println(res)
		}

	}
}

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division par zéro")
		}
		a /= b
		return a, nil
	default:
		return 0, fmt.Errorf("opération inconnue")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	}
	return nil
}
