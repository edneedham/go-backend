package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Hot Chocolate",
		Price: 1.00,
		SKU:   "abs-asdkf-faf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
