package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Lipton",
		Price: 1.00,
		SKU:   "abc-abs-edc",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
