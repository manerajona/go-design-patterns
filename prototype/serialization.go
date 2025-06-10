package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type HeadQuarters struct {
	StreetAddress, City, Country string
}

type Company struct {
	Name      string
	HQ        *HeadQuarters
	Employees []string
}

func (c *Company) DeepCopy() *Company {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(c)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	dCopy := Company{}
	_ = d.Decode(&dCopy)
	return &dCopy
}
