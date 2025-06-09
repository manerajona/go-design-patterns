package main

import (
	"fmt"
	"os"
	"strings"
)

type Journal struct {
	entries []string
	count   int
}

func (j *Journal) Add(e string) (index int) {
	index = j.count
	j.count++
	j.entries = append(j.entries, fmt.Sprintf("%d: %s", index, e))
	return
}

func (j *Journal) Print() string {
	return strings.Join(j.entries, "\n")
}

// breaks Single Responsibility Pattern
func (j *Journal) PersistToFile(filename string) error {
	return os.WriteFile(filename, []byte(j.Print()), 0644)
}

// correct
func PersistToFile(entries []string, filename string) error {
	content := strings.Join(entries, "\n")
	return os.WriteFile(filename, []byte(content), 0644)
}
