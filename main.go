package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	var sum, min, max int
	min = 101 // Nilai minimum diinisialisasi lebih dari maksimum
	max = -1  // Nilai maksimum diinisialisasi kurang dari minimum

	for _, grade := range s.Grades {
		sum += grade
		if grade < min {
			min = grade
		}
		if grade > max {
			max = grade
		}
	}

	average := float64(sum) / float64(len(s.Grades))
	return average, min, max
}

// debugg
func main() {
	school := School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	}

	avg, min, max := Analysis(school)
	fmt.Printf("Average: %.2f, Min: %d, Max: %d\n", avg, min, max)
}
