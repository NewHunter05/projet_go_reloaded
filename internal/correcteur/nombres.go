package correcteur

import "strconv"

func convertirNombres(mots []string) []string {
	for i := 1; i < len(mots); i++ {
		if mots[i] == "(hex)" {
			n, err := strconv.ParseInt(mots[i-1], 16, 64)
			if err == nil {
				mots[i-1] = strconv.Itoa(int(n))
			}
		}
		if mots[i] == "(bin)" {
			n, err := strconv.ParseInt(mots[i-1], 2, 64)
			if err == nil {
				mots[i-1] = strconv.Itoa(int(n))
			}
		}
	}

	propre := []string{}
	for i := 0; i < len(mots); i++ {
		if mots[i] != "(hex)" && mots[i] != "(bin)" {
			propre = append(propre, mots[i])
		}
	}
	return propre
}
