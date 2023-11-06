package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var choice int
var lettre []string

func LettresProposes() {
	a := 0
	var reponse string
	fmt.Println("Proposez une lettre")
	fmt.Println(lettre)
	fmt.Scanln(&reponse)
	for i := 0; i < len(lettre); i++ {
		if reponse == lettre[i] {
			fmt.Println("Erreur Vous avez deja selectionner cette lettre")
			Gameplay()
		}
	}
	lettre = append(lettre, reponse)
	for t := 0; t < len(faute.lettre); t++ {

		if reponse == faute.lettre[t] {
			placement[t] = reponse
		} else {
			a = a + 1
			if a == len(faute.lettre) {
				faute.faute = faute.faute + 1
				Faute()
				break
			}
		}
	}
	Gameplay()
}

func ToLower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		return r + 32
	}
	return r
}

func Motcomplet() {
	var reponse string
	fmt.Println("Entre un mot complet:")
	fmt.Scanln(&reponse)
	if reponse != mot {
		fmt.Println("Faux")
		faute.faute = faute.faute + 2
		Faute()
	} else {
		fmt.Println("Vraie")
		for i := 0; i < len(faute.lettre); i++ {
			placement[i] = faute.lettre[i]
		}
		fmt.Println(placement)
		fmt.Println("Tu as trouver le mot tu as gagner")
	}
}

func Menu() {
	fmt.Println("----------------------------------------")
	fmt.Println("1) Entre un mot complet")
	fmt.Println("2 Proposez une lettre")
	fmt.Println("3) Retour au Menu")
	var choice int
	fmt.Scanln(&choice)
	fmt.Println(choice)
	if choice == 1 {
		Motcomplet()
	} else if choice == 2 {
		LettresProposes()
	} else {
		Menu()
	}
}

func ReadFileContent(filename string) string {
	var test []string
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}
	fmt.Println("Contenu du fichier :")
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		fmt.Println(i)
		println(lines[i])
		test = append(test, lines[i])
	}
	mot_random := rand.Intn(len(test))
	return test[mot_random]
}
