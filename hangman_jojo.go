package hangman

import (
	"fmt"
    "os"
	"strings"
    "math/rand"
)



var choice int
var lettre []string


func  LettresProposes() {
    var reponse string
    fmt.Println("Proposez une lettre")
    fmt.Println(lettre)
    fmt.Scanln(&reponse)
    for i:= 0; i<len(lettre); i++{
        if reponse == lettre[i] {
            fmt.Println("Erreur Vous avez deja selectionner cette lettre")
            Menu()
        }
    }
    lettre = append(lettre, reponse)
    Menu()
}


func ToLower(r rune) rune {
    if 'A' <= r && r <= 'Z' {
        return r + 32
    }
    return r
}

func Motcomplet(){
    var input string
	var reponse string
	fmt.Println("Entre un mot complet:")
	fmt.Scanln(&input)
	fmt.Scanln(&reponse)
	if reponse != input {
		fmt.Println("Faux")
		fmt.Println("Tu perds 2 points")
	} else {
		fmt.Println("Vraie")
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
		switch choice {
		case 1:
			Motcomplet()
		case 2: 
			LettresProposes()
		default:
		}
		Menu()
}

func ReadFileContent(filename string) string{
    var test []string
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}
	fmt.Println("Contenu du fichier :")
    lines := strings.Split(string(data), "\n")
	for i:=0; i<len(lines); i++ {
        fmt.Println(i)
        println(lines[i])
        test = append(test, lines[i])
    }
    mot_random := rand.Intn(len(test))
    return test[mot_random]
}
