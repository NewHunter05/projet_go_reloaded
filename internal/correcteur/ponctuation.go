package correcteur

import "strings"

func corrigerPonctuation(mots []string) []string {
	propre := []string{}

	for i := 0; i < len(mots); i++ {
		reste := strings.TrimLeft(mots[i], ".,!?:;")
		ponct := mots[i][:len(mots[i])-len(reste)]

		if ponct != "" {
			if len(propre) > 0 {
				propre[len(propre)-1] = propre[len(propre)-1] + ponct
			} else {
				reste = mots[i]
			}
		}

		if reste != "" {
			propre = append(propre, reste)
		}
	}

	return propre
}
