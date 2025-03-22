package main

import (
	"fmt"
	"strings"
	"unicode"
)

type ParsingState string

const ( // États possibles: normal, inQuote, inDoubleQuote
	Normal      ParsingState = "normal"
	SingleQuote ParsingState = "singleQuote"
	DoubleQuote ParsingState = "doubleQuote"
)

func ParseShellWords(line string) (string, []string, error) {
	var words []string
	var word strings.Builder

	var state = Normal
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
		if char == '\\' {
			if state != SingleQuote && state != DoubleQuote {
				escapeNext = true
				continue
			}
		}

		// Traitement selon l'état actuel
		switch state {
		case Normal:
			if char == '"' {
				state = DoubleQuote
			} else if char == '\'' {
				state = SingleQuote
			} else if unicode.IsSpace(char) {
				// Fin d'un mot
				if word.Len() > 0 {
					words = append(words, word.String())
					word.Reset()
				}
			} else {
				word.WriteRune(char)
			}

		case SingleQuote:
			if char == '\'' {
				state = Normal
			} else {
				word.WriteRune(char)
			}

		case DoubleQuote:
			if char == '"' {
				state = "normal"
			} else {
				word.WriteRune(char)
			}
		}
	}

	// Vérifier si nous avons terminé dans un état valide
	if state != Normal {
		return "", nil, fmt.Errorf("guillemets non fermés")
	}

	// Ajouter le dernier mot s'il existe
	if word.Len() > 0 {
		words = append(words, word.String())
	}

	return words[0], words[1:], nil
}
