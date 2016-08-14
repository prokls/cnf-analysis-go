package sat

import "fmt"

// Lit are literals; variables with a sign

type Lit int32

func (l Lit) Pos() bool {
	return l > 0
}

func (l Lit) Neg() Lit {
	return Lit(-l)
}

// Clause is a clause; disjunction of literals

type Clause []Lit

func (c *Clause) Length() int {
	return len(*c)
}

// CNF is a conjunctive normal form; conjunction of clauses

type CNF struct {
	NbVars    int
	NbClauses int
	Lits      []Lit
}

func NewCNF() *CNF {
	c := new(CNF)
	c.Lits = make([]Lit, 0, 65536)
	return c
}

func (c *CNF) Dump() {
	fmt.Println("p cnf " + string(c.NbVars) + " " + string(c.NbClauses))
	for _, lit := range c.Lits {
		fmt.Printf("%d ", lit)
	}
}
