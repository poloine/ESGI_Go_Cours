package main

import (
	"errors"
	"fmt"
)

type Payeur interface {
	Payer(montant float64) (string, error)
}

var _ Payeur = &CarteCredit{}
var _ Payeur = &Paypal{}
var _ Payeur = &Crypto{}

type CarteCredit struct {
	Numero, Titulaire string
	Solde             float64
}

type Paypal struct {
	Email string
	Solde float64
}

type Crypto struct {
	Adresse, Monnaie string
	Solde            float64
}

func (c CarteCredit) Payer(montant float64) (string, error) {
	nouveauSolde := c.Solde - montant
	if nouveauSolde < 0 {
		return "", errors.New("Solde de la carte " + c.Numero + " négatif")
	}
	return "Transaction CB #" + c.Numero + " confirmé", nil
}

func (p Paypal) Payer(montant float64) (string, error) {
	nouveauSolde := p.Solde - montant
	if nouveauSolde < 0 {
		return "", errors.New("Solde PayPal de " + p.Email + " négatif")
	}
	return "Paiement PayPal de " + fmt.Sprintf("%.2f€", montant) + " vers " + p.Email, nil
}

func (c *Crypto) Payer(montant float64) (string, error) {
	switch c.Monnaie {
	case "BTC":
		// 1 BTC = 50_000€
		// montant en € -> montant en BTC
		nouveauSolde := (c.Solde*50_000 - montant) / 50_000
		if nouveauSolde < 0 {
			return "", errors.New("Solde Crypto de " + c.Adresse + " négatif")
		}
		return "Payement Crypto de " + fmt.Sprintf("%.3f %s", montant/50_000, c.Monnaie) + " vers " + c.Adresse, nil
	default:
		return "", errors.New("Monnaie Crypto non supportée : " + c.Monnaie)
	}
}

func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}

	fmt.Printf("Total des achats : %.2f€\n", total)

	message, erreur := payeur.Payer(total)

	switch p := payeur.(type) {
	case *CarteCredit:
		fmt.Printf("Carte Crédit - %s\n", p.Numero)
	case *Paypal:
		fmt.Printf("PayPal - %s\n", p.Email)
	case *Crypto:
		fmt.Printf("Crypto (%s) - %s\n", p.Monnaie, p.Adresse)
	}

	if erreur != nil {
		fmt.Printf("Erreur : %v\n", erreur)
		return
	}
	fmt.Printf("%s\n", message)
}

func main() {

	articles := []float64{20.0, 15.5, 30.0}

	fmt.Println("--- Test Carte Crédit (suffisant)")
	cc := &CarteCredit{Numero: "4242-4242", Titulaire: "Alice", Solde: 100.0}
	ProcesserPanier(cc, articles)
	fmt.Printf("Solde restant carte: %.2f€\n\n", cc.Solde)

	fmt.Println("--- Test PayPal (insuffisant)")
	pp := &Paypal{Email: "bob@example.com", Solde: 50.0}
	ProcesserPanier(pp, articles)
	fmt.Printf("Solde PayPal: %.2f€\n\n", pp.Solde)

	fmt.Println("--- Test Crypto BTC (suffisant)")
	btc := &Crypto{Adresse: "192.168.1.1", Monnaie: "BTC", Solde: 0.005}
	ProcesserPanier(btc, articles)
	fmt.Printf("Solde Crypto restant (BTC): %.6f\n\n", btc.Solde)

	fmt.Println("--- Test Crypto non supportée")
	eth := &Crypto{Adresse: "0x123", Monnaie: "ETH", Solde: 10.0}
	ProcesserPanier(eth, []float64{5.0})
	fmt.Printf("Solde ETH: %.2f\n", eth.Solde)
}
