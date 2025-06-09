package main

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other useful stuff here
}

type Cardinality struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipsBrowser interface {
	FindChildrenOf(parent Person) (children []*Person)
	FindSiblingsOf(person Person) (siblings []*Person)
}

// low-level module
type Relationships struct {
	cardinalities []Cardinality
}

func (rs *Relationships) FindChildrenOf(parent Person) (children []*Person) {
	for idx, cardinality := range rs.cardinalities {
		if cardinality.relationship == Parent &&
			cardinality.from.name == parent.name {
			children = append(children, rs.cardinalities[idx].to)
		}
	}
	return
}

func (rs *Relationships) FindSiblingsOf(person Person) (siblings []*Person) {
	for idx, cardinality := range rs.cardinalities {
		if cardinality.relationship == Sibling &&
			cardinality.from.name == person.name {
			siblings = append(siblings, rs.cardinalities[idx].to)
		}
	}
	return
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.cardinalities = append(rs.cardinalities, Cardinality{parent, Parent, child})
	rs.cardinalities = append(rs.cardinalities, Cardinality{child, Child, parent})
}

func (rs *Relationships) AddSiblings(siblings ...*Person) {
	for _, source := range siblings {
		for _, target := range siblings {
			if source == target {
				continue
			}
			rs.cardinalities = append(rs.cardinalities, Cardinality{source, Sibling, target})
		}
	}
}

// hign-level module
type Research struct {
	// relationships Relationships // breaks dip
	browser RelationshipsBrowser
}

func (r *Research) GenerateParentReport(parent Person) (report []string) {
	for _, p := range r.browser.FindChildrenOf(parent) {
		report = append(report, parent.name, "has a child called", p.name)
	}
	return
}

func (r *Research) GenerateSiblingReport(sibling Person) (report []string) {
	for _, p := range r.browser.FindSiblingsOf(sibling) {
		report = append(report, sibling.name, "has a sibling called", p.name)
	}
	return
}
