package Lab_0

import "fmt"

type Report struct {
	Title string
	Text  string
}

func (r Report) Display() {
	fmt.Printf("=== %s ===\n%s\n", r.Title, r.Text)
}
