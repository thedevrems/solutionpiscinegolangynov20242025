package main

import (
	"io" // Pour l'opération d'entrée/sortie
	"os" // Pour l'accès aux arguments, fichiers et flux standard

	"github.com/01-edu/z01" // Package externe pour l'affichage des runes
)

func main() {
	// Vérifie s'il y a moins de 2 arguments (c'est-à-dire seulement le nom du programme)
	if len(os.Args) < 2 {
		// Si aucun fichier n'est passé en argument, copie les données de l'entrée standard (Stdin)
		// vers la sortie standard (Stdout).
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	// Boucle sur chaque nom de fichier passé en argument à partir du deuxième élément (os.Args[1:])
	for _, fileName := range os.Args[1:] {
		// Lit le contenu du fichier spécifié
		content, err := os.ReadFile(fileName)
		if err != nil {
			// En cas d'erreur lors de la lecture, construit un message d'erreur
			nameError := "ERROR: " + err.Error()

			// Affiche le message d'erreur caractère par caractère en utilisant PrintRune du package z01
			for _, r := range nameError {
				z01.PrintRune(r)
			}
			// Imprime un saut de ligne
			z01.PrintRune('\n')

			// Quitte le programme avec un code d'erreur
			os.Exit(1)
		}

		// Si la lecture réussit, affiche le contenu du fichier caractère par caractère
		for _, r := range string(content) {
			z01.PrintRune(r)
		}
	}
}
