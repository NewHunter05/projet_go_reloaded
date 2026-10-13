package correcteur

func corrigerApostrophes(mots []string) []string {
	propre := []string{}
	ouvert := false
	for i := 0; i < len(mots); i++ {
		if mots[i] == "'" {
			if !ouvert {
				ouvert = true
				if i+1 < len(mots) {
					mots[i+1] = "'" + mots[i+1]
					continue
				}
			} else {
				ouvert = false
				if len(propre) > 0 {
					propre[len(propre)-1] = propre[len(propre)-1] + "'"
					continue
				}
			}
		}
		propre = append(propre, mots[i])
	}
	return propre
}
