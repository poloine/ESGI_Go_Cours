package main

import "fmt"

func main() {
	emp1 := Employe{
		Personne: Personne{"Alice", "Berth", 45, "alice.berth@mail.com"},
		Adresse:  Adresse{"123 Rue Foch", "Lyon", "69000"},
		Poste:    "Développeuse",
		Salaire:  44800,
	}

	emp2 := Employe{
		Personne: Personne{"Bob", "Smith", 24, "bob.smith@mail.com"},
		Adresse:  Adresse{"8 Avenue du jardin", "Paris", "75000"},
		Poste:    "Développeur",
		Salaire:  33200,
	}

	etu1 := Etudiant{
		Personne: Personne{"Louis", "Lanonse", 20, "louis.lanonse@mail.fr"},
		Promo:    "L2 Droit",
		Moyenne:  18.3,
	}

	etu2 := Etudiant{
		Personne: Personne{"Ethan", "Bougueur", 20, "ethan.bougueur@mail.fr"},
		Promo:    "L3 Info",
		Moyenne:  12.7,
	}

	fmt.Println(emp1.FicheEmploye())
	fmt.Println(emp2.FicheEmploye())

	emp1.AugmenterSalaire(11)
	fmt.Printf("%.2f€\n", emp1.Salaire)

	fmt.Println(etu1.FicheEtudiant())
	fmt.Println(etu2.FicheEtudiant())
}

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans - %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("EMPLOYÉ %s\nPoste: %s\nSalaire: %.2f€\nAdresse: %s",
		e.Presentation(), e.Poste, e.Salaire, e.Adresse.Format())
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire * (1 + pct/100)
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très Bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez Bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Non validé"
	}
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf("ÉTUDIANT %s\nPromo: %s\nMoyenne: %.2f\nMention: %s",
		e.Presentation(), e.Promo, e.Moyenne, e.MentionObtenue())
}
