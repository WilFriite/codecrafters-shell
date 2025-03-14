package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ParseShellWords(line string) (string, []string, error) {
	var words []string
	var word strings.Builder

	state := "normal"   // États possibles: normal, inQuote, inDoubleQuote
	escapeNext := false // Indique si le prochain caractère doit être échappé

	// Parcourir la chaîne caractère par caractère
	for _, char := range line {
		// Traiter l'échappement des caractères
		if escapeNext {
			word.WriteRune(char)
			escapeNext = false
			continue
		}

		// Vérifier le caractère d'échappement
		if char == '\\' && state != "inQuote" {
			escapeNext = true
			continue
		}

		// Traitement selon l'état actuel
		switch state {
		case "normal":
			if char == '"' {
				state = "inDoubleQuote"
			} else if char == '\'' {
				state = "inQuote"
			} else if unicode.IsSpace(char) {
				// Fin d'un mot
				if word.Len() > 0 {
					words = append(words, word.String())
					word.Reset()
				}
			} else {
				word.WriteRune(char)
			}

		case "inQuote":
			if char == '\'' {
				state = "normal"
			} else {
				word.WriteRune(char)
			}

		case "inDoubleQuote":
			if char == '"' {
				state = "normal"
			} else {
				word.WriteRune(char)
			}
		}
	}

	// Vérifier si nous avons terminé dans un état valide
	if state != "normal" {
		return "", nil, fmt.Errorf("guillemets non fermés")
	}

	// Ajouter le dernier mot s'il existe
	if word.Len() > 0 {
		words = append(words, word.String())
	}

	return words[0], words[1:], nil
}
