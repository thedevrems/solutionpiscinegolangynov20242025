package main

import (
	"fmt"
	"os"
)

func main() {
	// Vérifie si le nombre d'arguments est différent de 1 mais inférieur ou égal à 2
	if len(os.Args) != 1 && len(os.Args) <= 2 {

		// Ouvre le fichier spécifié en argument (os.Args[1])
		file, err := os.Open(os.Args[1])
		if err != nil {
			// Si une erreur survient lors de l'ouverture, affiche un message d'erreur et quitte le programme
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		// Assure la fermeture du fichier à la fin du programme ou en cas d'erreur
		defer file.Close()

		// Crée un tableau de bytes pour stocker les données lues depuis le fichier
		arr := make([]byte, 100)
		for {
			// Lit jusqu'à 100 octets depuis le fichier et les stocke dans le tableau `arr`
			n, err := file.Read(arr)
			if err != nil {
				// Si une erreur survient (ex. fin de fichier), sort de la boucle
				break
			}
			// Écrit les octets lus (jusqu'à `n`) dans la sortie standard
			os.Stdout.Write(arr[:n])
		}
	} else if len(os.Args) == 1 {
		// Si aucun argument de fichier n'est fourni, affiche un message d'erreur
		fmt.Printf("File name missing\n")
	} else {
		// Si plus d'un argument est fourni (autre que le nom du programme), affiche un message d'erreur
		fmt.Printf("Too many arguments\n")
	}
}
