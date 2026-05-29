# GO 
___
Cours du 26/05/2026

## Langage GO 

Compilé, comme C++, C# 
→ Syntaxe complexe

!= Interprété Python
→ Syntaxe facile
Couche intermédiaire
Performance dégradée

Désavantage langage compilé :
- +complexe
- -rigide
- re-compilation lors de la modification
- fuite de mémoire
-> Risques

## Langages Objets

Java (garbage collector), C/C++/C#

## Objectifs du GO

- Bas niveau
- Web
- Mobile
-> meilleures performances

## Plus d'avantages du GO (le cours est bien structuré)

- Simplicité
    - Lire
    - Écrire
- Rapidité
    - Éxecution
    - Compilation
    -> Contraintes
- Concurrence

    
Dans les fichiers .go
`import fmt`

Rigidité pour simplification

Pragmatisme : une seule façon de faire
→ Stabilité, compatibilité
GO Routines → plusieurs taches, multi-threading

GO → 25 mots clés (35 Python, ~80 Java)
~ 5M développers

Utilisé par Docker, Kubernetes, Terraform, Prometheus...
Utilisation :
1. Micro-services et API REST
2. CLI
3. IaaS

## TP premier pas

```go
package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Pow(64, 8))
	fmt.Println(time.Now().Date())
	hourMinute := strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute())
	fmt.Println(hourMinute)
	currentDate := time.Now()
	fmt.Println(currentDate)
	birthday := time.Date(2000, time.December, 28, 0, 0, 0, 0, time.Local)
	fmt.Println(birthday)
	fmt.Println(dateDiff(currentDate, birthday).Year())
}

func dateDiff(a, b time.Time) (c time.Time) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	minute := int(m2 - m1)
	second := int(s2 - s1)

	// Normalize negative values
	if second < 0 {
		second += 60
		minute--
	}
	if minute < 0 {
		minute += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}
```

```terminaloutput
/home/poloine/.cache/JetBrains/GoLand2026.1/tmp/GoLand/___go_build_hello_go #gosetup
Hello World
2.81474976710656e+14
2026 May 26
17:44
2026-05-26 17:44:14.895925027 +0200 CEST m=+0.000062365
2000-12-28 00:00:00 +0100 CET
25

Process finished with the exit code 0
```

## Documentation go run & go build

```bash
go run
```
Compile et exécute le code source dans le module "main" en une seule étape.

```bash
go build [-o output] [build flags] [packages]
```
Compile les packages et leurs dépendances, mais ne les exécute pas.
- `-o output` : Spécifie le nom du fichier exécutable généré.
- `build flags` : Options supplémentaires pour la compilation (ex: `-v` pour afficher les packages compilés).
- `packages` : Les packages à compiler (ex: `./...` pour compiler tous les packages dans le répertoire courant et ses sous-répertoires).

`go build -ldflags -s -w` : Supprime les symboles de débogage et les informations de la table des symboles pour réduire la taille de l'exécutable.

## Les types en GO

### Entiers

- int8 : -128 à 127
- int16 : -32,768 à 32,767
- int32 : -2,147,483,648 à 2,147,483647
- int64 : -9,223,372,036,854,775,808 à 9,223,372,036,854,775,807
- uint8 : 0 à 255
- uint16 : 0 à 65,535
- uint32 : 0 à 4,294,967,295
- uint64 : 0 à 18,446,744,073,709,551

`int` et `uint` : Taille dépendante de l'architecture (32 ou 64 bits).

### Flottants

- float32 : 6-9 chiffres significatifs, plage de ±1.5e-45 à ±3.4e38
- float64 : 15-17 chiffres significatifs, plage de ±5e-324 à ±1.7e308

### Booleans
- bool : true ou false

```go
var bool1 = true
```

### String

```go
var text = "Hello, World!"
```
-> Concaténation, addition de string, interpolation
Aussi indexable (`text[0] -> 'H'`) et itérable

### Spéciaux (Alias)

- `byte` : alias pour `uint8`, utilisé pour représenter des données binaires ou des caractères ASCII.
- `rune` : alias pour `int32`, utilisé pour représenter des caractères Unicode

### Valeurs nulles

`nil` : Représente une valeur nulle pour les types de référence (pointeurs, slices, maps, channels, interfaces, fonctions)

Pour les autres types :
- Entiers : 0
- Flottants : 0.0
- Booleans : false
- Strings : "" (chaîne vide)

### Constantes

```go
const Pi = 3.14
```
- Doivent être assignées à une valeur constante au moment de la déclaration
- Ne peuvent pas être modifiées après leur déclaration (duh)

### Iota

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```
- `iota` est un compteur qui commence à 0 et s'incrémente de 1 pour chaque ligne dans une déclaration de constantes
- Utile pour créer des énumérations
- Peut être utilisé pour générer des valeurs basées sur des expressions (`1 << iota` pour des puissances de 2)

___
Cours du 29/05/2026

### Initialisation des variables

Une variable déclarée en Go est automatiquement initialisée à sa valeur zéro (zero value) :
```go
var i int     // i est initialisé à 0
var f float64 // f est initialisé à 0.0
var b bool    // b est initialisé à false
var s string  // s est initialisé à "" (chaîne vide)
var p *int    // p est initialisé à nil (pointeur nul)
var slice []int // slice est initialisé à nil (slice nul)
var funcVar func() // funcVar est initialisé à nil (fonction nulle)

type Personne struct {
    Nom string
    Age int
}
var personne Personne // personne est initialisé à Personne{Nom: "", Age: 0}

// Utile pour les patterns idiomatiques
var compteur int
for i := 0; i < 10; i++ {
    compteur++
}
```

Très peu de NullPointerException par rapport à Java/C

### Déclaration des variables

3 façons de déclarer :
```go
var x int = 10 // var avec type et valeur explicites
var ville = "Paris" // var avec inférence de type (Go éduit le type)
y := 20 // Déclaration courte et initialisation implicite (type inféré, UNIQUEMENT dans les fonctions)
```

Plusieurs déclarations à la fois possible :
```go
var a, b, c int = 1, 2, 3
x, y := 4, 5
```

Échange de valeurs sans variable temporaire :
```go
a, b = b, a
```

Variable avec plusieurs déclarations groupées :
```go
var (
    serveur string = "localhost"
    port    int    = 8080
)
```

### Constantes

Constantes typées et non typées :
```go
const Pi = 3.14 // Constante non typée (peut être utilisée comme float32 ou float64)
const E float64 = 2.71828 // Constante typée (float64)
```

Iota : générateur d'énumérations
```go
type NiveauLog int
const (
    DEBUG NiveauLog = iota // 0
    INFO                   // 1
    WARNING                // 2
    ERROR                  // 3
	CRITICAL               // 4
)
```

Iota avec calcul - tailles de stockage
```go
const (
    Octet = 1
    Ko = 1 << (10 * (iota + 1)) // 1 << 10 (1024)
    Mo = 1 << (10 * (iota + 1)) // 1 << 20 (1,048,576)
    Go = 1 << (10 * (iota + 1)) // 1 << 30 (1,073,741,824)
    To = 1 << (10 * (iota + 1)) // 1 << 40 (1,099,511,627,776)
)
```

### Retours multiples
Pattern idiomatique (résultat, erreur)
```go
func diviser(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division par zéro")
    }
    return a / b, nil
}
resultat, err := diviser(10, 2)
if err != nil {
    fmt.Println("Erreur:", err)
} else {
    fmt.Println("Résultat:", resultat)
}
```

Ignorer une valeur de retour avec `_`
```go
resultat, _ := diviser(10, 0) // Ignorer l'erreur
```

Retours nommés (pour les fonctions et documentation)
```go
func calculer(a, b int) (somme int, produit int) {
    somme = a + b
    produit = a * b
    return // Retourne les valeurs nommées
}
s, p := calculer(3, 4)
fmt.Println("Somme:", s, "Produit:", p)
```

### Fonctions variadiques
```go
func somme(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
fmt.Println(somme(1, 2, 3)) // 6
fmt.Println(somme(4, 5))    // 9
```

Une fonction est une valeur comme int ou string
```go
var f func(int) int
f = func(x int) int {
    return x * x
}
fmt.Println(f(5)) // 25
```

Cas pratique : middleware HTTP, tris personnalisés, callbacks
```go
personnes := []string {"Alice", "Bob", "Charlie"}
sort.Slice(personnes, func(i, j int) bool {
    return personnes[i] < personnes[j] // Tri alphabétique
})
fmt.Println(personnes) // [Alice Bob Charlie]
```

### Boucles
For classique
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

While :
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

For range (itération sur slices, maps, strings)
```go
personnes := []string {"Alice", "Bob", "Charlie"}
for index, nom := range personnes {
    fmt.Printf("Index: %d, Nom: %s\n", index, nom)
}
```
```go
ages := map[string]int {
    "Alice": 30,
    "Bob": 25,
    "Charlie": 35,
}
for nom, age := range ages {
    fmt.Printf("Nom: %s, Age: %d\n", nom, age)
}
```

For infini, break pour sortir
```go
for {
    fmt.Println("Boucle infinie")
    if condition {
        break
    }
}
```

### Switch case
```go
jour := "lundi"
switch jour {
case "lundi", "mardi", "mercredi", "jeudi", "vendredi":
    fmt.Println("C'est un jour de semaine")
case "samedi", "dimanche":
    fmt.Println("C'est un jour de week-end")
default:
    fmt.Println("Jour inconnu")
}
```

Switch sans expression (équivalent à if-else)
```go
age := 25
switch {
case age < 18:  
    fmt.Println("Mineur")
case age >= 18 && age < 65:
    fmt.Println("Adulte")
case age >= 65:
    fmt.Println("Senior")
}
```

Fallthrough pour exécuter le case suivant (rarement utilisé)
```go
switch score {
case 90, 100:
    fmt.Println("Excellent")
    fallthrough
case 80, 89:        
    fmt.Println("Très bien")
case 70, 79:
    fmt.Println("Bien")
default:
    fmt.Println("À améliorer")
}
```

### Tableaux et Slices

#### Tableau (Array) - Taille FIXE

Déclaration :
```go
var scores [8]int = [8]int{95, 72, 88, 45, 67, 91, 55, 78}
var entiers [5]int // Tableau de 5 entiers initialisés à 0
```

- Taille fixe : Définie à la compilation, immuable
- Typage strict : `[5]int` ≠ `[6]int` (types différents)
- Passage par copie : Inefficace pour les grandes données
- Pas d'append() : Impossible d'ajouter des éléments
- Compatible avec les slices : Un slice peut être créé depuis un tableau

```go
arr := [3]int{1, 2, 3}
fmt.Println(len(arr), cap(arr)) // 3 3
// arr = append(arr, 4) ❌ ERREUR: arrays don't have append
```

Cas d'usage : Données de taille connue et fixe (ex: coordonnées XY, dès [6]int)

#### Slice - Taille DYNAMIQUE

Déclaration :
```go
scores := []int{95, 72, 88, 45, 67, 91, 55, 78} // Création avec valeurs
var vide []int // Slice nul (nil)
nombres := make([]int, 5) // Slice de 5 éléments (valeurs zéro)
nombres := make([]int, 5, 10) // Slice: len=5, cap=10
```

- Taille dynamique : Peut croître avec `append()`
- Typage unifié : `[]int` = même type quelque soit la taille
- Passage par référence : Efficace (pointeur interne vers les données)
- append() possible : Ajoute des éléments dynamiquement
- len() et cap() : `len()` = taille actuelle, `cap()` = capacité allouée

```go
scores := []int{95, 72}
fmt.Println(len(scores), cap(scores)) // 2 2
scores = append(scores, 88)
fmt.Println(len(scores), cap(scores)) // 3 4 (capacité doublée)
```

#### Slicing - Extraction d'une portion

Créer un slice depuis un tableau ou un autre slice :
```go
arr := [8]int{95, 72, 88, 45, 67, 91, 55, 78}
slicePartiel := arr[2:5]  // Indices 2, 3, 4 (5 exclus)
fmt.Println(slicePartiel) // [88 45 67]
fmt.Println(len(slicePartiel), cap(slicePartiel)) // 3 6
```

Variations de slicing :
```go
scores := []int{10, 20, 30, 40, 50}
scores[1:3] // [20 30]
scores[:3] // [10 20 30]
scores[2:] // [30 40 50]
scores[:] // [10 20 30 40 50]
```

Important : Un slice créé depuis un tableau partage la même mémoire que le tableau original
```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:3] // [2 3]
slice[0] = 999 // Modifie le slice
fmt.Println(arr) // [1 999 3 4 5] Le tableau est aussi modifié
```

#### Opérations courantes sur les slices

```go
// Créer un slice
slice := []int{10, 20, 30}

// Ajouter éléments
slice = append(slice, 40) // [10 20 30 40]
slice = append(slice, 50, 60) // [10 20 30 40 50 60]

// Fusionner deux slices
autres := []int{70, 80}
slice = append(slice, autres...) // ... déplie le slice

// Copier un slice
copie := make([]int, len(slice))
copy(copie, slice)

// Parcourir avec range
for index, valeur := range slice {
    fmt.Printf("Index: %d, Valeur: %d\n", index, valeur)
}
```

### Structures

Définition type structure :
```go
type Produit struct {
    Nom  string
	Prix float64
	Stock int
	Actif bool
	Categorie []string
}
```

3 façons d'initialiser une struct :
```go
// 1. Avec des valeurs explicites
p1 := Produit{
    Nom: "Laptop",
    Prix: 999.99,
    Stock: 10,
    Actif: true,
    Categorie: []string{"Électronique", "Informatique"},
}
// 2. Avec des valeurs positionnelles (fragile si l'on ajoute des champs)
p2 := Produit{"Smartphone", 499.99, 20, true, []string{"Électronique", "Téléphonie"}}

// 3. Avec des champs individuels (valeurs par défaut à la base)
var p3 Produit
p3.Nom = "Table"
p3.Prix = 199.99
p3.Stock = 5
p3.Actif = true
p3.Categorie = []string{"Meubles"}
```

### Méthodes : donner des comportements aux structs
```go
func (p Produit) Afficher() {
    fmt.Printf("Produit: %s, Prix: %.2f, Stock: %d\n", p.Nom, p.Prix, p.Stock)
}
p1.Afficher() // Produit: Laptop, Prix: 999.99, Stock: 10
```

Value receiver : travaille sur une copie de la struct (ne modifie pas l'original)
```go
func (p Produit) AugmenterPrix(pourcentage float64) {
    p.Prix += p.Prix * pourcentage / 100
}
p1.AugmenterPrix(10)
p1.Afficher() // Produit: Laptop, Prix: 999.99, Stock: 10 (prix non modifié)
```

Pointer receiver : travaille sur un pointeur vers la struct (modifie l'original)
```go
func (p *Produit) AugmenterPrix(pourcentage float64) {
    p.Prix += p.Prix * pourcentage / 100
}
p1.AugmenterPrix(10)
p1.Afficher() // Produit: Laptop, Prix: 1099.99, Stock: 10 (prix modifié)
```

Utilisation - Go gère automatiquement & et *
```go
p1.AugmenterPrix(10) // Go convertit automatiquement en pointeur
```

### Visibilité

La visibilité est détermminée par la casse du nom :
- Exporté (public) : Commence par une majuscule (ex: `Produit`, `Afficher()`)
- Non exporté (privé) : Commence par une minuscule (ex: `produit`, `afficher()`)

Les éléments exportés sont accessibles depuis d'autres packages

### Defer

Permet de différer l'exécution d'une fonction jusqu'à la fin de la fonction courante
```go
func lireFichier(chemin string) {
    fichier, err := os.Open(chemin)
    if err != nil {
        fmt.Println("Erreur d'ouverture:", err)
        return
    }
    defer fichier.Close() // Assure la fermeture du fichier à la fin de la fonction
    // Lire le fichier...
}
```

Ordre LIFO, important pour les ressources dépendantes
```go
func exemple() {
    fmt.Println("Début")
    defer fmt.Println("Fin") // S'exécutera à la fin de la fonction
    defer fmt.Println("Milieu") // S'exécutera avant "Fin"
    fmt.Println("Traitement")
}
```

Cas classique : sync.Mutex
```go
var mu sync.Mutex
func sectionCritique() {
    mu.Lock() // Verrouille le mutex
    defer mu.Unlock() // Assure le déverrouillage même en cas d'erreur
    // Code critique...
}
```

### Packages standards à connaitre

`strings` : Manipulation de chaînes de caractères
- `strings.Contains(s, substr)` : Vérifie si `s` contient `substr`
- `strings.HasPrefix(s, prefix)` : Vérifie si `s` commence par `prefix`
- `strings.HasSuffix(s, suffix)` : Vérifie si `s` se termine par `suffix`
- `strings.Split(s, sep)` : Divise `s` en un slice de sous-chaînes séparées par `sep`
- `strings.Join(slice, sep)` : Concatène les éléments d'un slice en une chaîne séparée par `sep`
- `strings.ToUpper(s)` : Convertit `s` en majuscules
- `strings.ToLower(s)` : Convertit `s` en minuscules
- `strings.TrimSpace(s)` : Supprime les espaces en début et fin de `s`

`strconv` : Conversion de types
- `strconv.Atoi(s)` : Convertit une chaîne en entier (string to int)
- `strconv.Itoa(i)` : Convertit un entier en chaîne (int to string)
- `strconv.ParseFloat(s, bitSize)` : Convertit une chaîne en float (string to float)
- `strconv.FormatFloat(f, fmt, prec, bitSize)` : Convertit un float en chaîne (float to string)
- `strconv.ParseBool(s)` : Convertit une chaîne en booléen (string to bool)
- `strconv.FormatBool(b)` : Convertit un booléen en chaîne (bool to string)

`math` : Fonctions mathématiques
- `math.Pow(x, y)` : Calcule `x` à la puissance `y`
- `math.Sqrt(x)` : Calcule la racine carrée de `x`
- `math.Abs(x)` : Calcule la valeur absolue de `x`
- `math.Max(x, y)` : Retourne le maximum entre `x` et `y`
- `math.Min(x, y)` : Retourne le minimum entre `x` et `y`
- `math.Round(x)` : Arrondit `x` à l'entier le plus proche
- `math.Floor(x)` : Arrondit `x` à l'entier inférieur
- `math.Ceil(x)` : Arrondit `x` à l'entier supérieur

`sort` : Tri
- `sort.Ints(slice)` : Trie un slice d'entiers
- `sort.Strings(slice)` : Trie un slice de chaînes de caractères
- `sort.Slice(slice, lessFunc)` : Trie un slice avec une fonction de comparaison personnalisée

`time` : Manipulation du temps
- `time.Now()` : Retourne l'heure actuelle
- `time.Date(year, month, day, hour, min, sec, nsec, loc)` : Crée une valeur de temps à partir de composants individuels
- `time.Parse(layout, value)` : Analyse une chaîne de caractères en une valeur de temps selon un format spécifié
- `time.Format(layout)` : Formate une valeur de temps en une chaîne selon un format spécifié
- `time.Add(d)` : Ajoute une durée `d` à une valeur de temps
- `time.Sub(t)` : Calcule la différence entre deux valeurs de temps
- `time.Sleep(d)` : Met en pause l'exécution pendant une durée `d`

![img.png](annexes-notes/a-savoir-fin-second-cours.png)