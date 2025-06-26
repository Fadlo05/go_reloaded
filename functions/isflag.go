package functions

import (
	"strconv"
	"strings"
)

// isValidFlag vérifie si la chaîne correspond à un format de flag valide sans utiliser de regexp.
func isFlag(flag string) bool {
	// Supprimer les espaces superflus
	flag = strings.TrimSpace(flag)

	// Vérifier que la chaîne commence par '(' et se termine par ')'
	if !strings.HasPrefix(flag, "(") || !strings.HasSuffix(flag, ")") {
		return false
	}

	// Supprimer les parenthèses
	flag = flag[1 : len(flag)-1]

	// Diviser la chaîne en deux parties : le type (low, cap, up) et le nombre
	parts := strings.Split(flag, ",")
	if len(parts) != 2 {
		return false
	}

	// Supprimer les espaces superflus autour des parties
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])
	// fmt.Println(parts[0])
	// fmt.Println(parts[1])

	// Vérifier que le type est l'un des suivants
	if parts[0] != "low" && parts[0] != "cap" && parts[0] != "up" {
		return false
	}

	// Vérifier que la deuxième partie est un nombre entier positif
	num, err := strconv.Atoi(parts[1])
	if err != nil || num <= 0 {
		return false
	}

	return true
}
