package stats

import (
	"fmt"

	"github.com/prokls/cnf-analysis-go/output"
	"github.com/prokls/cnf-analysis-go/sat"
)

// Union-Find data structure

type UFType uint32

type unionFind struct {
	elements []UFType
}

func newUnionFind(size int) *unionFind {
	uf := new(unionFind)
	uf.elements = make([]UFType, size)

	for i := 0; i < size; i++ {
		uf.elements[i] = UFType(i)
	}

	return uf
}

func (uf *unionFind) Find(e UFType) (UFType, error) {
	if len(uf.elements) <= int(e) {
		return 0, fmt.Errorf("%d exceeds %d", e, len(uf.elements))
	} else if e < 0 {
		return 0, fmt.Errorf("union find element must be non-negative, is %d", e)
	}

	c := e
	for {
		p := uf.elements[c]
		if p == c {
			uf.elements[e] = c
			return c, nil
		}
		c = p
	}
}

func (uf *unionFind) Union(a, b UFType) error {
	reprA, err := uf.Find(a)
	if err != nil {
		return err
	}
	reprB, err := uf.Find(b)
	if err != nil {
		return err
	}
	uf.elements[reprA] = reprB
	return nil
}

func (uf *unionFind) Count() (int, error) {
	reps := make(map[UFType]bool)

	for _, elem := range uf.elements {
		rep, err := uf.Find(elem)
		if err != nil {
			return 0, err
		}
		val, _ := reps[rep]
		if !val {
			reps[rep] = true
		}
	}

	return len(reps), nil
}

// literal and variable components

func posEquiv(lit sat.Lit) sat.Lit {
	if lit < 0 {
		return -2*lit - 2
	} else {
		return 2*lit - 1
	}
}

func EvaluateComponents(cnf *sat.CNF, feat *output.Features, fconf *FeatureConfig) error {
	cc := newUnionFind(2 * cnf.NbVars)

	// literal components
	var ref UFType
	readRef := true
	for _, lit := range cnf.Lits {
		if lit == 0 {
			readRef = true
		} else if readRef {
			ref = UFType(posEquiv(lit))
			readRef = false
		} else {
			err := cc.Union(ref, UFType(posEquiv(lit)))
			if err != nil {
				return err
			}
		}
	}

	connLitComps, err := cc.Count()
	if err != nil {
		return err
	}
	// variable components
	for vari := sat.Lit(1); vari <= sat.Lit(cnf.NbVars); vari++ {
		pos := UFType(posEquiv(vari))
		neg := UFType(posEquiv(-vari))
		err = cc.Union(pos, neg)
		if err != nil {
			return err
		}
	}

	connVarComps, err := cc.Count()
	if err != nil {
		return err
	}
	feat.ConnectedLiteralComponentsCount = uint16(connLitComps)
	feat.ConnectedVariableComponentsCount = uint16(connVarComps)
	return nil
}
