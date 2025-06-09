package solid

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

/*
NON-EXTENSIBLE VERSION
*/
type Filter struct {
}

func (f *Filter) filterByColor(products []Product, color Color) (productsByColor []*Product) {
	for idx, product := range products {
		if product.color == color {
			productsByColor = append(productsByColor, &products[idx])
		}
	}
	return
}

func (f *Filter) filterBySize(products []Product, size Size) (productsBySize []*Product) {
	for idx, product := range products {
		if product.size == size {
			productsBySize = append(productsBySize, &products[idx])
		}
	}
	return
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size, color Color) (productsBySizeAndColor []*Product) {
	for idx, product := range products {
		if product.size == size && product.color == color {
			productsBySizeAndColor = append(productsBySizeAndColor, &products[idx])
		}
	}
	return
}

/*
EXTENSIBLE VERSION
*/
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	spec1, spec2 Specification
}

func (andSpec AndSpecification) IsSatisfied(p *Product) bool {
	return andSpec.spec1.IsSatisfied(p) && andSpec.spec2.IsSatisfied(p)
}

type OrSpecification struct {
	spec1, spec2 Specification
}

func (orSpec OrSpecification) IsSatisfied(p *Product) bool {
	return orSpec.spec1.IsSatisfied(p) || orSpec.spec2.IsSatisfied(p)
}

type ExtensibleFilter struct{}

func (f *ExtensibleFilter) Filter(products []Product, spec Specification) (productsBySpec []*Product) {
	for idx, product := range products {
		if spec.IsSatisfied(&product) {
			productsBySpec = append(productsBySpec, &products[idx])
		}
	}
	return
}
