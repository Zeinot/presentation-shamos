# Présentation sur l'Algorithme de Shamos

**Présentée par :** Omar Sarsar
**Cours :** Algorithmes et Structures de Données

## À propos

Cette présentation interactive explique l'algorithme de Shamos pour résoudre le problème de la paire de points la plus proche en utilisant la stratégie "Diviser pour Régner".

## Historique

L'algorithme de la paire de points la plus proche a été introduit en **1975** par **Michael Ian Shamos** et **Dan Hoey** dans leur article intitulé "Closest-Point Problems". Cet algorithme est un exemple classique de la technique "Diviser pour Régner" en géométrie algorithmique.

## Le Problème

**Entrée :** Un ensemble de *n* points dans le plan 2D
**Sortie :** La paire de points ayant la distance euclidienne minimale

## Complexité

- **Approche naïve (force brute) :** O(n²)
- **Algorithme de Shamos :** O(n log n)

## Comment ouvrir la présentation

### Option 1 : Ouvrir directement dans un navigateur
1. Double-cliquez sur le fichier `presentation.html`
2. Ou faites un clic droit → "Ouvrir avec" → votre navigateur préféré

### Option 2 : Depuis un serveur local (recommandé)

Si vous avez Python installé :

```bash
# Python 3
python -m http.server 8000

# Puis ouvrez dans votre navigateur :
# http://localhost:8000/presentation.html
```

Si vous avez Node.js installé :

```bash
npx http-server

# Puis ouvrez dans votre navigateur l'URL affichée
```

## Navigation dans la présentation

- **Flèches ← →** : Navigation horizontale entre les slides
- **Flèches ↑ ↓** : Navigation verticale (slides imbriquées)
- **Espace** : Slide suivant
- **S** : Mode aperçu de toutes les slides
- **F** : Mode plein écran
- **?** : Afficher l'aide
- **ESC** : Sortir du mode aperçu ou plein écran

## Contenu de la présentation

1. Introduction au problème
2. Applications pratiques
3. Approche naïve (force brute)
4. L'algorithme de Shamos
5. Principe général (Diviser-Régner-Combiner)
6. Étapes détaillées de l'algorithme
7. Analyse de complexité
8. Application du Théorème Master
9. Comparaisons et performances
10. Avantages et limitations
11. Extensions et variantes
12. Exemples visuels
13. Conclusion et références

## Technologies utilisées

- **reveal.js 5.0.4** : Framework de présentation HTML
- **Highlight.js** : Coloration syntaxique du code
- **CSS personnalisé** : Design moderne et élégant

## Ressources complémentaires

### Fichiers inclus
- `presentation.html` : La présentation interactive (avec exemples de code en Go)
- `shamos_algorithm.go` : Implémentation complète en Go
- `README.md` : Ce fichier

## Exécuter le code Go

Pour exécuter l'implémentation Go :

```bash
# Compiler et exécuter
go run shamos_algorithm.go

# Ou compiler d'abord puis exécuter
go build shamos_algorithm.go
./shamos_algorithm       # Sur Linux/Mac
shamos_algorithm.exe     # Sur Windows
```

Le programme compare les performances de l'approche naïve O(n²) avec l'algorithme de Shamos O(n log n) sur différents ensembles de points.

### Références académiques

- Shamos, M. I., & Hoey, D. (1975). Closest-Point Problems. *Proceedings of the 16th Annual Symposium on Foundations of Computer Science*.
- Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2009). *Introduction to Algorithms* (3rd ed.). MIT Press.
- Preparata, F. P., & Shamos, M. I. (1985). *Computational Geometry: An Introduction*. Springer-Verlag.

### Liens utiles

- [Closest pair of points problem - Wikipedia](https://en.wikipedia.org/wiki/Closest_pair_of_points_problem)
- [Finding the nearest pair of points - CP-Algorithms](https://cp-algorithms.com/geometry/nearest_points.html)
- [reveal.js Documentation](https://revealjs.com/)

## Amélioration récentes de l'algorithme

Des recherches récentes ont permis d'optimiser l'algorithme :

- **Ge et al.** : Réduction du nombre de calculs de distances à (3n log n)/2
- **Optimisations récentes** : Utilisation de propriétés géométriques (circle-packing) pour calculer au maximum 7n/2 distances euclidiennes

## Sources

- [Closest pair of points problem - Wikipedia](https://en.wikipedia.org/wiki/Closest_pair_of_points_problem)
- [Closest Pair - Hideous Humpback Freak](https://hideoushumpbackfreak.com/algorithms/algorithms-closest-pair.html)
- [Finding the nearest pair of points - Algorithms for Competitive Programming](https://cp-algorithms.com/geometry/nearest_points.html)
- [Engineering the Divide-and-Conquer Closest Pair Algorithm](https://www.researchgate.net/publication/220586046_Engineering_the_Divide-and-Conquer_Closest_Pair_Algorithm)
- [An Improved Algorithm for Finding the Closest Pair of Points](https://link.springer.com/article/10.1007/s11390-006-0027-7)
- [Multidimensional divide-and-conquer](https://dl.acm.org/doi/10.1145/358841.358850)

## Contact

Pour toute question concernant cette présentation :

**Omar Sarsar**
Cours : Algorithmes et Structures de Données

---

*Présentation créée avec reveal.js - Décembre 2025*
