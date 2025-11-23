package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

var studenti = make(map[string]Student)

func DodajOceno(vpisna string, ocena int) {
	if ocena < 1 || ocena > 10 {
		fmt.Println("Ocena mora biti med 1 in 10")
		return
	}

	s, ok := studenti[vpisna]
	if !ok {
		fmt.Println("Študent ne obstaja.")
		return
	}

	s.Ocene = append(s.Ocene, ocena)
	studenti[vpisna] = s
}

func IzpisVsehOcen(vpisna string) []int {
	s, ok := studenti[vpisna]
	if !ok {
		return nil
	}
	return s.Ocene
}

func IzpisiKoncniUspeh(minOcena, maxOcena, stOcen int) {
	for _, s := range studenti {
		avg := povprecje(s.Ocene)
		opis := ""

		if avg >= 9 {
			opis = "Odličen študent"
		} else if avg >= 6 {
			opis = "Povprečen študent"
		} else {
			opis = "Neuspešen študent"
		}

		fmt.Printf("%s %s -> povprečje: %.2f (%s)\n", s.Ime, s.Priimek, avg, opis)
	}
}

func povprecje(ocene []int) float64 {
	if len(ocene) == 0 {
		return 0
	}
	sum := 0.0
	for _, o := range ocene {
		sum += float64(o)
	}
	return sum / float64(len(ocene))
}

func DodajStudenta(vpisna, ime, priimek string) {
	studenti[vpisna] = Student{
		Ime:     ime,
		Priimek: priimek,
		Ocene:   []int{},
	}
}
