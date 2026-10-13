package main

import (
	"fmt"
	"os"

	"projet_go_reloaded/internal/correcteur"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Impossible de lire :", err)
		os.Exit(1)
	}

	texte := string(data)
	resultat := correcteur.CorrigerTexte(texte)

	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Écriture du résultat impossible :", err)
		os.Exit(1)
	}
}
