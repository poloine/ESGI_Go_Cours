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



