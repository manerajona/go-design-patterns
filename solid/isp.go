package main

type Document struct {
}

/*
FAT INTERFACE
*/
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
	// TODO
}

func (m MultiFunctionPrinter) Print(d Document) {
	// TODO
}

func (m MultiFunctionPrinter) Fax(d Document) {
	// TODO
}

func (m MultiFunctionPrinter) Scan(d Document) {
	// TODO
}

type OldFashionedPrinter struct {
	// TODO
}

func (o OldFashionedPrinter) Print(d Document) {
	// TODO
}

// breaks interface segregation principle
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// breaks interface segregation principle
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

/*
SEGREGATED INTERFACES
*/
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct {
	// TODO
}

func (m MyPrinter) Print(d Document) {
	// TODO
}

// combine interfaces
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	// TODO
}

func (p Photocopier) Print(d Document) {
	// TODO
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination + decorator
type MultiFunctionMachine struct {
	Printer
	Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.Printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.Scanner.Scan(d)
}
