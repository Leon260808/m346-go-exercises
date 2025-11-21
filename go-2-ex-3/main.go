package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "Cloud Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104)
	modules[320] = "Objektorientiert Programmieren"
	modules[117] = "Informatik- und Netzinfrastruktur für ein KMU realisieren"
	modules[320] = "Objektorientierte Programmierung mit C#"
	fmt.Println(modules)
}
