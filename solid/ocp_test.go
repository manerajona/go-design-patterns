package solid

import (
	"testing"
)

var products = []Product{
	{"Apple", green, small},
	{"Tree", green, large},
	{"House", blue, large},
}

func TestFilterByColor_Green(t *testing.T) {
	f := Filter{}
	greenProducts := f.filterByColor(products, green)
	if len(greenProducts) != 2 {
		t.Errorf("Expected 2 green products, got %d", len(greenProducts))
	}
}

func TestFilterBySize_Large(t *testing.T) {
	f := Filter{}
	largeProducts := f.filterBySize(products, large)
	if len(largeProducts) != 2 {
		t.Errorf("Expected 2 large products, got %d", len(largeProducts))
	}
}

func TestFilterBySizeAndColor_LargeGreen(t *testing.T) {
	f := Filter{}
	greenLargeProducts := f.filterBySizeAndColor(products, large, green)
	if len(greenLargeProducts) != 1 {
		t.Errorf("Expected 1 large green product, got %d", len(greenLargeProducts))
	}
}

func TestColorSpecification_Green(t *testing.T) {
	ef := ExtensibleFilter{}
	greenSpec := ColorSpecification{green}
	greenProducts := ef.Filter(products, greenSpec)
	if len(greenProducts) != 2 {
		t.Errorf("Expected 2 green products, got %d", len(greenProducts))
	}
}

func TestAndSpecification_LargeGreen(t *testing.T) {
	ef := ExtensibleFilter{}
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	greenAndLargeSpec := AndSpecification{greenSpec, largeSpec}
	greenLargeProducts := ef.Filter(products, greenAndLargeSpec)
	if len(greenLargeProducts) != 1 {
		t.Errorf("Expected 1 large green product, got %d", len(greenLargeProducts))
	}
}

func TestOrSpecification_LargeGreen(t *testing.T) {
	ef := ExtensibleFilter{}
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	greenOrLargeSpec := OrSpecification{greenSpec, largeSpec}
	greenOrLargeProducts := ef.Filter(products, greenOrLargeSpec)
	if len(greenOrLargeProducts) != 3 {
		t.Errorf("Expected 3 products (green or large), got %d", len(greenOrLargeProducts))
	}
}
