package correcteur

import "strings"

func corrigerArticles(mots []string) []string {
	for i := 0; i+1 < len(mots); i++ {
		if mots[i] == "a" || mots[i] == "A" {
			lettre := strings.ToLower(mots[i+1][:1])
			if strings.Contains("aeiouh", lettre) {
				mots[i] = mots[i] + "n"
			}
		}

	}
	return mots
}
