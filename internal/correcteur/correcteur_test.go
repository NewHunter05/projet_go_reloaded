package correcteur

import "testing"

func TestCorrigerTexte(t *testing.T) {
	cas := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{"hex", "The bakery sold FF (hex) croissants today", "The bakery sold 255 croissants today"},
		{"bin", "My cat has 1101 (bin) toys", "My cat has 13 toys"},
		{"up", "Please be quiet (up)", "Please be QUIET"},
		{"low", "My NEIGHBOUR (low) is nice", "My neighbour is nice"},
		{"cap", "We visited paris (cap) last summer", "We visited Paris last summer"},
		{"up avec nombre", "the train is very late (up, 3)", "the train IS VERY LATE"},
		{"low avec nombre", "PLEASE CLOSE THE DOOR (low, 4) now", "please close the door now"},
		{"cap avec nombre", "the lord of the rings (cap, 5) is long", "The Lord Of The Rings is long"},
		{"ponctuation", "Hello ,my friend !", "Hello, my friend!"},
		{"groupe de ponctuation", "Wait ... what !?", "Wait... what!?"},
		{"apostrophes un mot", "She just said ' hello '", "She just said 'hello'"},
		{"apostrophes plusieurs mots", "He shouted ' see you tomorrow ' and left", "He shouted 'see you tomorrow' and left"},
		{"a devient an", "I ate a orange and a hamburger", "I ate an orange and an hamburger"},
		{"plusieurs lignes", "go home (up)\nthe end (cap)", "go HOME\nthe End"},
	}

	for i := 0; i < len(cas); i++ {
		obtenu := CorrigerTexte(cas[i].entree)
		if obtenu != cas[i].attendu {
			t.Errorf("%s\n  entrée  : %q\n  obtenu  : %q\n  attendu : %q", cas[i].nom, cas[i].entree, obtenu, cas[i].attendu)
		}
	}
}
