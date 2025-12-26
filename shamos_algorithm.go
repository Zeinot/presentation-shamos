package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Point représente un point dans le plan 2D
type Point struct {
	X, Y float64
}

// distance calcule la distance euclidienne entre deux points
func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

// ==================== APPROCHE NAÏVE (Force Brute) ====================

// pairePlusProcheNaive trouve la paire la plus proche en O(n²)
func pairePlusProcheNaive(points []Point) (Point, Point, float64) {
	n := len(points)
	if n < 2 {
		return Point{}, Point{}, math.Inf(1)
	}

	minDist := math.Inf(1)
	var p1, p2 Point

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(points[i], points[j])
			if d < minDist {
				minDist = d
				p1, p2 = points[i], points[j]
			}
		}
	}

	return p1, p2, minDist
}

// ==================== ALGORITHME DE SHAMOS (Divide & Conquer) ====================

// forceBrute est utilisée pour les petits ensembles (n <= 3)
func forceBrute(points []Point) (Point, Point, float64) {
	n := len(points)
	if n < 2 {
		return Point{}, Point{}, math.Inf(1)
	}

	minDist := math.Inf(1)
	var p1, p2 Point

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(points[i], points[j])
			if d < minDist {
				minDist = d
				p1, p2 = points[i], points[j]
			}
		}
	}

	return p1, p2, minDist
}

// diviser sépare l'ensemble de points en deux moitiés
func diviser(Px, Py []Point) ([]Point, []Point, []Point, []Point) {
	n := len(Px)
	mid := n / 2

	// Point médian pour définir la ligne de séparation
	pointMedian := Px[mid]

	// Diviser Px
	PxGauche := Px[:mid]
	PxDroite := Px[mid:]

	// Diviser Py en fonction de la ligne médiane
	PyGauche := []Point{}
	PyDroite := []Point{}

	for _, p := range Py {
		if p.X <= pointMedian.X {
			PyGauche = append(PyGauche, p)
		} else {
			PyDroite = append(PyDroite, p)
		}
	}

	return PxGauche, PxDroite, PyGauche, PyDroite
}

// combiner trouve la paire la plus proche qui traverse la ligne médiane
func combiner(Px, Py []Point, d float64, p1Best, p2Best Point) (Point, Point, float64) {
	n := len(Px)
	mid := n / 2
	pointMedian := Px[mid]

	// Créer la bande : points à distance < d de la ligne médiane
	bande := []Point{}
	for _, p := range Py {
		if math.Abs(p.X-pointMedian.X) < d {
			bande = append(bande, p)
		}
	}

	// Vérifier les paires dans la bande (au maximum 7 points par itération)
	minDist := d
	bestP1, bestP2 := p1Best, p2Best

	for i := 0; i < len(bande); i++ {
		// Vérifier au maximum 7 points suivants
		for j := i + 1; j < len(bande) && j < i+8; j++ {
			dist := distance(bande[i], bande[j])
			if dist < minDist {
				minDist = dist
				bestP1, bestP2 = bande[i], bande[j]
			}
		}
	}

	return bestP1, bestP2, minDist
}

// pairePlusProcheRec est la fonction récursive de l'algorithme de Shamos
func pairePlusProcheRec(Px, Py []Point) (Point, Point, float64) {
	n := len(Px)

	// Cas de base : utiliser force brute pour petits ensembles
	if n <= 3 {
		return forceBrute(Px)
	}

	// Diviser
	PxG, PxD, PyG, PyD := diviser(Px, Py)

	// Régner (appels récursifs)
	p1G, p2G, dGauche := pairePlusProcheRec(PxG, PyG)
	p1D, p2D, dDroite := pairePlusProcheRec(PxD, PyD)

	// Trouver le minimum des deux côtés
	var dMin float64
	var bestP1, bestP2 Point

	if dGauche < dDroite {
		dMin = dGauche
		bestP1, bestP2 = p1G, p2G
	} else {
		dMin = dDroite
		bestP1, bestP2 = p1D, p2D
	}

	// Combiner : vérifier la bande médiane
	return combiner(Px, Py, dMin, bestP1, bestP2)
}

// PairePlusProche est la fonction principale utilisant l'algorithme de Shamos
// Complexité : O(n log n)
func PairePlusProche(points []Point) (Point, Point, float64) {
	if len(points) < 2 {
		return Point{}, Point{}, math.Inf(1)
	}

	// Prétraitement : trier les points par X et par Y - O(n log n)
	Px := make([]Point, len(points))
	Py := make([]Point, len(points))
	copy(Px, points)
	copy(Py, points)

	sort.Slice(Px, func(i, j int) bool {
		return Px[i].X < Px[j].X
	})

	sort.Slice(Py, func(i, j int) bool {
		return Py[i].Y < Py[j].Y
	})

	return pairePlusProcheRec(Px, Py)
}

// ==================== FONCTIONS UTILITAIRES ====================

// genererPointsAleatoires génère n points aléatoires dans le plan [0, max] x [0, max]
func genererPointsAleatoires(n int, max float64) []Point {
	rand.Seed(time.Now().UnixNano())
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		points[i] = Point{
			X: rand.Float64() * max,
			Y: rand.Float64() * max,
		}
	}

	return points
}

// afficherResultat affiche les résultats de manière formatée
func afficherResultat(p1, p2 Point, dist float64, methode string, duree time.Duration) {
	fmt.Printf("\n=== Résultats avec %s ===\n", methode)
	fmt.Printf("Point 1: (%.2f, %.2f)\n", p1.X, p1.Y)
	fmt.Printf("Point 2: (%.2f, %.2f)\n", p2.X, p2.Y)
	fmt.Printf("Distance minimale: %.4f\n", dist)
	fmt.Printf("Temps d'exécution: %v\n", duree)
}

// ==================== FONCTION MAIN ====================

func main() {
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║   ALGORITHME DE SHAMOS                            ║")
	fmt.Println("║   Paire de Points la Plus Proche                  ║")
	fmt.Println("║   Présenté par: Omar Sarsar                       ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	// Exemple 1 : Petit ensemble de points
	fmt.Println("\n📍 EXEMPLE 1 : Ensemble de 8 points")
	fmt.Println("──────────────────────────────────────────────────")

	pointsPetits := []Point{
		{2, 3}, {12, 30}, {40, 50}, {5, 1},
		{12, 10}, {3, 4}, {15, 20}, {7, 9},
	}

	// Méthode naïve
	debut := time.Now()
	p1, p2, dist := pairePlusProcheNaive(pointsPetits)
	duree := time.Since(debut)
	afficherResultat(p1, p2, dist, "Force Brute O(n²)", duree)

	// Algorithme de Shamos
	debut = time.Now()
	p1, p2, dist = PairePlusProche(pointsPetits)
	duree = time.Since(debut)
	afficherResultat(p1, p2, dist, "Shamos O(n log n)", duree)

	// Exemple 2 : Grand ensemble de points
	fmt.Println("\n\n📍 EXEMPLE 2 : Comparaison avec 10,000 points aléatoires")
	fmt.Println("──────────────────────────────────────────────────")

	pointsGrands := genererPointsAleatoires(10000, 1000.0)

	// Méthode naïve (attention : lent pour grands ensembles!)
	fmt.Println("\n⏱️  Exécution avec Force Brute (peut prendre du temps)...")
	debut = time.Now()
	p1Naive, p2Naive, distNaive := pairePlusProcheNaive(pointsGrands)
	dureeNaive := time.Since(debut)
	afficherResultat(p1Naive, p2Naive, distNaive, "Force Brute O(n²)", dureeNaive)

	// Algorithme de Shamos
	fmt.Println("\n⚡ Exécution avec Shamos...")
	debut = time.Now()
	p1Shamos, p2Shamos, distShamos := PairePlusProche(pointsGrands)
	dureeShamos := time.Since(debut)
	afficherResultat(p1Shamos, p2Shamos, distShamos, "Shamos O(n log n)", dureeShamos)

	// Comparaison des performances
	fmt.Println("\n\n📊 COMPARAISON DES PERFORMANCES")
	fmt.Println("──────────────────────────────────────────────────")
	fmt.Printf("Nombre de points: %d\n", len(pointsGrands))
	fmt.Printf("Temps Force Brute: %v\n", dureeNaive)
	fmt.Printf("Temps Shamos: %v\n", dureeShamos)

	if dureeNaive > dureeShamos {
		acceleration := float64(dureeNaive) / float64(dureeShamos)
		fmt.Printf("\n🚀 Shamos est %.2fx plus rapide!\n", acceleration)
	}

	// Vérification que les résultats sont identiques
	fmt.Println("\n✅ VÉRIFICATION")
	fmt.Println("──────────────────────────────────────────────────")
	diff := math.Abs(distNaive - distShamos)
	if diff < 0.0001 {
		fmt.Println("✓ Les deux algorithmes trouvent la même distance!")
	} else {
		fmt.Printf("✗ Différence détectée: %.6f\n", diff)
	}
}
