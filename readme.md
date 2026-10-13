go-reloaded

Outil en ligne de commande, écrit en Go, qui corrige automatiquement un texte et écrit le résultat dans un nouveau fichier.

Fonctionnalités
(hex) / (bin) : convertit le mot précédent en décimal (FF (hex) → 255).
(up) / (low) / (cap) : change la casse du mot précédent, ou des n mots précédents avec (up, n).
Ponctuation : colle . , ! ? : ; au mot précédent (Hello ,friend ! → Hello, friend!).
Apostrophes : colle les ' ' aux mots entourés (' hello ' → 'hello').
Articles : transforme a en an devant une voyelle ou un h.
Pré-requis
Go
Git
Installation
bash
git clone https://github.com/<votre-compte>/projet_go_reloaded.git
cd projet_go_reloaded
Utilisation
bash
go run . sample.txt result.txt

Lancer les tests :

bash
go test ./...