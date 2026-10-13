package correcteur

import "strings"

func corrigerLigne(ligne string) string {
	mots := strings.Fields(ligne)
	mots = convertirNombres(mots)
	mots = modifierCasse(mots)
	mots = corrigerPonctuation(mots)
	mots = corrigerApostrophes(mots)
	mots = corrigerArticles(mots)
	return strings.Join(mots, " ")
}
