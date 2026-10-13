package correcteur

import (
	"strconv"
	"strings"
)

func modifierCasse(mots []string) []string {
	for i := 1; i < len(mots); i++ {
		if mots[i] == "(up)" {
			mots[i-1] = strings.ToUpper(mots[i-1])

		}
		if mots[i] == "(low)" {
			mots[i-1] = strings.ToLower(mots[i-1])
		}
		if mots[i] == "(cap)" {
			mots[i-1] = strings.ToUpper(mots[i-1][:1]) + strings.ToLower(mots[i-1][1:])
		}

		if mots[i] == "(up," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := i - 1; j >= i-n && j >= 0; j-- {
					mots[j] = strings.ToUpper(mots[j])
				}
				mots[i] = ""
				mots[i+1] = ""
			}
		}

		if mots[i] == "(low," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := i - 1; j >= i-n && j >= 0; j-- {
					mots[j] = strings.ToLower(mots[j])
				}
				mots[i] = ""
				mots[i+1] = ""
			}
		}

		if mots[i] == "(cap," && i+1 < len(mots) {
			nombreTexte := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(nombreTexte)
			if err == nil {
				for j := i - 1; j >= i-n && j >= 0; j-- {
					mots[j] = strings.ToUpper(mots[j][:1]) + strings.ToLower(mots[j][1:])
				}
				mots[i] = ""
				mots[i+1] = ""
			}
		}
	}

	propre := []string{}
	for i := 0; i < len(mots); i++ {
		if mots[i] != "(up)" && mots[i] != "(low)" && mots[i] != "(cap)" && mots[i] != "" {
			propre = append(propre, mots[i])
		}
	}

	return propre
}
