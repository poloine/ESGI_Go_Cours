package main

import (
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return fmt.Errorf("ID %d déjà existant", p.ID)
		}
	}
	p.Actif = true
	c.Produits = append(c.Produits, p)
	return nil
}

func (c *Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.Produits {
		if p.ID == id && p.Actif {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("Produit ID %d non trouvé", id)
}

func (c *Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit
	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) && p.Actif {
			resultats = append(resultats, p)
		}
	}
	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for i, p := range c.Produits {
		if strings.EqualFold(p.Categorie, categorie) && p.Actif {
			c.Produits[i].Prix = c.Produits[i].Prix * (1 - pct/100)
			count++
		}
	}
	return count
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i, p := range c.Produits {
		if p.ID == id && p.Actif {
			if p.Stock < qte {
				return fmt.Errorf("Stock insuffisant pour %s (disponible: %d)", p.Nom, p.Stock)
			}
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("Produit ID %d non trouvé", id)
}

func (c *Catalogue) Rapport() string {
	totalProduits := 0
	valeurTotal := 0.0

	for _, p := range c.Produits {
		if p.Actif {
			totalProduits++
			valeurTotal += p.Prix * float64(p.Stock)
		}
	}

	return fmt.Sprintf("RAPPORT TECHSHOP\n---\nProduits actifs: %d\nValeur totale du stock: %.2f€",
		totalProduits, valeurTotal)
}

func (c *Catalogue) Afficher() {
	fmt.Println("CATALOGUE TECHSHOP")
	fmt.Println("---")
	if len(c.Produits) == 0 {
		fmt.Println("Aucun produit disponible")
		return
	}
	for _, p := range c.Produits {
		if p.Actif {
			fmt.Printf("id: %d, marque: %s, nom: %s, prix: %.2f€, stock: %d\n", p.ID, p.Marque, p.Nom, p.Prix, p.Stock)
		}
	}
}

func main() {
	catalogue := Catalogue{}

	catalogue.AjouterProduit(Produit{1, "Pixel 10 Pro", "Google", 1099.99, 12, "Smartphone", true})
	catalogue.AjouterProduit(Produit{2, "Inspirion 16", "Dell", 889.99, 9, "PC", true})
	catalogue.AjouterProduit(Produit{3, "XPS 13", "Dell", 1299.99, 13, "PC", true})
	catalogue.AjouterProduit(Produit{4, "Galaxy A55", "Samsung", 299.99, 26, "Smartphone", true})
	catalogue.AjouterProduit(Produit{5, "Ami", "Citroën", 7999.99, 2, "Électroménager", true})
	catalogue.AjouterProduit(Produit{6, "XVC-5HF", "HP", 199.99, 5, "Écran", true})

	for {
		fmt.Println("")
		fmt.Println("MENU TECHSHOP")
		fmt.Println("---")
		fmt.Println("[1] Ajouter un produit")
		fmt.Println("[2] Chercher par catégorie")
		fmt.Println("[3] Appliquer une réduction")
		fmt.Println("[4] Vendre")
		fmt.Println("[5] Rapport")
		fmt.Println("[6] Afficher le catalogue")
		fmt.Println("[0] Quitter")
		fmt.Print("Choix: ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var id int
			var nom, marque, cat string
			var prix float64
			var stock int

			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Nom: ")
			fmt.Scan(&nom)
			fmt.Print("Marque: ")
			fmt.Scan(&marque)
			fmt.Print("Prix: ")
			fmt.Scan(&prix)
			fmt.Print("Stock: ")
			fmt.Scan(&stock)
			fmt.Print("Catégorie: ")
			fmt.Scan(&cat)

			p := Produit{id, nom, marque, prix, stock, cat, true}
			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Produit ajouté")
			}

		case 2:
			var cat string
			fmt.Print("Catégorie à chercher: ")
			fmt.Scan(&cat)

			resultats := catalogue.TrouverParCategorie(cat)
			if len(resultats) == 0 {
				fmt.Println("Aucun produit trouvé")
			} else {
				fmt.Printf("Produits de la catégorie '%s':\n", cat)
				for _, p := range resultats {
					fmt.Printf("id: %d, marque: %s, nom: %s, prix: %.2f€, stock: %d\n", p.ID, p.Marque, p.Nom, p.Prix, p.Stock)
				}
			}

		case 3:
			var cat string
			var pct float64
			fmt.Print("Catégorie pour la réduction: ")
			fmt.Scan(&cat)
			fmt.Print("Pourcentage de réduction: ")
			fmt.Scan(&pct)

			count := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("%d produit(s) réduit(s) de %.0f%%\n", count, pct)

		case 4:
			var id, qte int
			fmt.Print("ID du produit: ")
			fmt.Scan(&id)
			fmt.Print("Quantité: ")
			fmt.Scan(&qte)

			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println(err)
			} else {
				produit, _ := catalogue.TrouverParID(id)
				fmt.Printf("%d x %s vendu(s). Stock restant: %d\n", qte, produit.Nom, produit.Stock)
			}

		case 5:
			fmt.Println("\n" + catalogue.Rapport())

		case 6:
			catalogue.Afficher()

		case 0:
			return

		default:
			fmt.Println("Choix invalide")
		}
	}
}
