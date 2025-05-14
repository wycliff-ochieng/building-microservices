package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		ID:          67,
		Name:        "Fratezi",
		Description: "italian cofeeee",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
