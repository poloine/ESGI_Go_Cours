package main

import "fmt"

func main() {

	// Tableau : taille fixe
	var scoresTableau [8]int = [8]int{95, 72, 88, 45, 67, 91, 55, 78}
	fmt.Printf("Déclaration: var scoresTableau [8]int\n")
	fmt.Printf("Taille: %d (FIXÉE et IMMUABLE)\n", len(scoresTableau))
	fmt.Printf("Capacité: %d\n\n", cap(scoresTableau))

	// slice : taille dynamique
	scores := []int{95, 72, 88, 45, 67, 91, 55, 78}
	fmt.Printf("Déclaration: scores := []int{...}\n")
	fmt.Printf("Taille: %d\n", len(scores))
	fmt.Printf("Capacité: %d\n", cap(scores))

	scores = append(scores, 100)
	fmt.Printf("append(100): len=%d, cap=%d\n\n", len(scores), cap(scores))

	slicePartiel := scoresTableau[2:5] // indices 2, 3, 4
	fmt.Printf("slicePartiel := scoresTableau[2:5]\n")
	fmt.Printf("Contenu: %v\n", slicePartiel)
	fmt.Printf("Len: %d, Cap: %d\n\n", len(slicePartiel), cap(slicePartiel))

	slicePartiel[0] = 999 // Modifie aussi le tableau original!
	fmt.Printf("Après modification du slice: slicePartiel[0] = 999\n")
	fmt.Printf("Tableau original scoresTableau: %v\n", scoresTableau)

	var expert, avance, intermediaire, novice, debutant int

	// Parcourir le slice avec range
	for index, score := range scores {
		fmt.Printf("Score #%d: %d → ", index+1, score)

		// Classifier avec switch
		switch {
		case score >= 90:
			fmt.Println("Expert")
			expert++
		case score >= 80:
			fmt.Println("Avancé")
			avance++
		case score >= 70:
			fmt.Println("Intermédiaire")
			intermediaire++
		case score >= 60:
			fmt.Println("Novice")
			novice++
		default:
			fmt.Println("Débutant")
			debutant++
		}
	}

	fmt.Printf("   Expert:        %d\n", expert)
	fmt.Printf("   Avancé:        %d\n", avance)
	fmt.Printf("   Intermédiaire: %d\n", intermediaire)
	fmt.Printf("   Novice:        %d\n", novice)
	fmt.Printf("   Débutant:      %d\n", debutant)
	fmt.Printf("   Total:         %d\n\n", len(scores))

	for i := len(scores) - 1; i >= 0; i-- {
		fmt.Printf("   Position %d: %d\n", len(scores)-i, scores[i])
	}
}
