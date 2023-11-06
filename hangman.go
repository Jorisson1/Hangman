package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Draw struct {
	faute  int
	lettre []string
}

var faute Draw

func Affichage() {
	var choix int
	faute.faute = 0
	fmt.Println("|   █  █  █▀▀█  ██  █  █▀▀▀  █▀▄ ▄▀█  █▀▀█  ██  █   |")
	fmt.Println("|   █■■█  █■■█  █ █ █  █ ▄▄  █  ▀  █  █■■█  █ █ █   |")
	fmt.Println("|   █  █  █  █  █  ██  █▄▄█  █     █  █  █  █  ██   |")
	fmt.Println("|               ~~FRENCH VERSION~~                  |")
	fmt.Println("|===================================================|")
	fmt.Println("|▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄                          |")
	fmt.Println("|    █  ▄▀        ║                                 |")
	fmt.Println("|    █▄▀          ║          1- Play solo           |")
	fmt.Println("|    █           😎                                 |")
	fmt.Println("|    █          |()|         2- Play with           |")
	fmt.Println("|    █           ||            friends              |")
	fmt.Println("|    █                                              |")
	fmt.Println("|    █                                              |")
	fmt.Println("|    █                                              |")
	fmt.Println("|    █                                              |")
	fmt.Println("|  ▄▀█▀▄                                            |")
	fmt.Println("|▄█▄▄█▄▄█▄▄▄▄▄                                      |")
	fmt.Scanln(&choix)
	if choix == 1 {
		Clear()
		Solo()
	}
}

var input string
var reponse string
var placement []string
var table []rune

func Solo() {
	fmt.Println("Entrez un mot :")
	fmt.Scanln(&input)
	lettre_random := rand.Intn(len(input))
	for i := 0; i < len(input); i++ {
		placement = append(placement, "_ ")
		faute.lettre = append(faute.lettre, string(input[i]))
	}
	placement[lettre_random] = faute.lettre[lettre_random]
	Gameplay()
}

func Gameplay() {
	fmt.Println(placement)
	fmt.Println("Entrez une lettre ou un mot :")
	fmt.Scanln(&reponse)
	if len(reponse) == 1 {
		a := 1
		for i := 0; i < len(faute.lettre); i++ {
			if reponse == string(faute.lettre[i]) {
				placement[i] = faute.lettre[i]
			} else {
				a = a + 1
				fmt.Println(len(faute.lettre))
				if a == len(faute.lettre) {
					faute.faute = faute.faute + 1
					Faute()
				}
			}
		}
		Gameplay()

	} else if reponse == input {
		fmt.Println("Win")
	} else if len(reponse) > 1 && reponse != input {
		fmt.Println("problème ici !")
		faute.faute = faute.faute + 2
		Faute()
		Gameplay()
	} else if len(reponse) == 1 && reponse != input {
		fmt.Println("Problème là !")
		faute.faute = faute.faute + 1
		Faute()
		Gameplay()
	} else {
		fmt.Println("Win")
	}
}

func Faute() {
	if faute.faute == 1 {
		Clear()
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println(" ▄▄▄▄█▄▄▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 2 {
		Clear()
		fmt.Println("                ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println(" ▄▄▄▄█▄▄▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 3 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println(" ▄▄▄▄█▄▄▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 4 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █▄           ")
		fmt.Println(" ▄▄▄▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 5 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 6 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 7 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █          ")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 8 {
		Clear()
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █        ()")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 9 {
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █       |()  ")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 10 {
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █       |()|  ")
		fmt.Println("     █           ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 11 {
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █       |()|  ")
		fmt.Println("     █        |   ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
		Gameplay()

	} else if faute.faute == 12 {
		fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄               ")
		fmt.Println("     █▄▀       ║    ")
		fmt.Println("     █        😎  ")
		fmt.Println("     █       |()|  ")
		fmt.Println("     █        ||  ")
		fmt.Println("    ▄█▄           ")
		fmt.Println(" ▄▄█▄█▄█▄▄▄▄▄▄  ")
	}
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
