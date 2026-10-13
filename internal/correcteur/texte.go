package correcteur

import "strings"

func CorrigerTexte(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i := 0; i < len(lignes); i++ {
		lignes[i] = corrigerLigne(lignes[i])
	}
	return strings.Join(lignes, "\n")
}
