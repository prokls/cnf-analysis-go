package input

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/prokls/cnf-analysis-go/sat"
)

type parsingState struct {
	clauses      int
	variables    int
	lineno       int
	col          int
	mode         int8
	inIgnoreLine bool
	bufIndex     uint8
	wasZero      bool
	wordBuf      [20]byte
}

func isNewline(c byte) bool {
	return c == 10
}

func isWhitespace(c byte) bool {
	return c == 9 || c == 10 || c == 11 || c == 12 || c == 13 || c == 32
}

func unexpected(unexp, exp string, ps *parsingState) error {
	// lineno is incremented before considering the line and zero-based, therefore ps.lineno
	// col is incremented after considering the byte and zero-based, therefore ps.col+1
	return fmt.Errorf("Unexpected '%s', expected %s, at line %d col %d", unexp, exp, ps.lineno, ps.col+1)
}

func withPos(ps *parsingState, msg string, args ...interface{}) error {
	// lineno is incremented before considering the line and zero-based, therefore ps.lineno
	// col is incremented after considering the byte and zero-based, therefore ps.col+1
	return fmt.Errorf(msg+fmt.Sprintf(", line %d, col %d", ps.lineno, ps.col+1), args...)
}

func consumeByte(char byte, cnf *sat.CNF, st *parsingState, conf *ParsingConfig) error {
	st.col += 1

	if isNewline(char) {
		st.inIgnoreLine = false
		st.lineno += 1
		st.col = 0
		return consumeWord(cnf, st, conf, isNewline(char))
	}
	if st.inIgnoreLine {
		return nil
	}
	if !isWhitespace(char) {
		st.wordBuf[st.bufIndex] = char
		st.bufIndex += 1
		return nil
	}

	return consumeWord(cnf, st, conf, isNewline(char))
}

func consumeWord(cnf *sat.CNF, st *parsingState, conf *ParsingConfig, nl bool) error {
	if st.bufIndex == 0 {
		return nil
	}

	word := string(st.wordBuf[:st.bufIndex])
	st.bufIndex = 0
	for i := 0; i < len(conf.IgnoreLines); i++ {
		if word == conf.IgnoreLines[i] {
			if !nl {
				st.inIgnoreLine = true
			}
			return nil
		}
	}

	var integer int
	if st.mode >= 2 {
		i, err := strconv.Atoi(word)
		if err != nil {
			return err
		}
		integer = i
	}

	switch st.mode {
	case 0:
		if word != "p" {
			return unexpected(word, "'p' of CNF header", st)
		}
		st.mode = 1
	case 1:
		if word != "cnf" {
			return unexpected(word, "'cnf' of CNF header", st)
		}
		st.mode = 2
	case 2:
		cnf.NbVars = integer
		st.mode = 3
		if integer >= math.MaxInt32 {
			return withPos(st, "cannot consume more than %d variables", math.MaxInt32)
		}
	case 3:
		cnf.NbClauses = integer
		st.mode = 4
		if integer >= math.MaxInt32 {
			return withPos(st, "cannot consume more than %d clauses", math.MaxInt32)
		}
	case 4:
		cnf.Lits = append(cnf.Lits, sat.Lit(integer))
		if integer != 0 {
			variable := integer
			if variable < 0 {
				variable = -variable
			}
			if variable > st.variables {
				st.variables = variable
			}
			if conf.CheckNbVars {
				if variable >= cnf.NbVars {
					return withPos(st, "%d exceeds variable limit %d", variable, cnf.NbVars)
				}
			}
			st.wasZero = false
		} else {
			st.clauses += 1
			if st.wasZero {
				return withPos(st, "Sorry, empty clauses (ie. clauses with no literals) are not allowed")
			}
			st.wasZero = true
		}
	}

	return nil
}

func ReadCNFFile(fd io.Reader, conf *ParsingConfig) (*sat.CNF, error) {
	var st parsingState
	cnf := sat.NewCNF()

	// verify parameters
	for i := 0; i < len(conf.IgnoreLines); i++ {
		if len(conf.IgnoreLines[i]) >= len(st.wordBuf) {
			return nil, withPos(&st, "line prefix '%s' is too long", conf.IgnoreLines[i])
		}
		if conf.IgnoreLines[i] == "p" {
			return nil, withPos(&st, "p-headers cannot be ignored")
		}
		for j := 0; j < len(conf.IgnoreLines[i]); j++ {
			if isWhitespace(conf.IgnoreLines[i][j]) {
				return nil, withPos(&st, "line prefixes must not contain spaces, '%s' does", conf.IgnoreLines[i])
			}
		}
	}

	// read ~4096 bytes
	buf := make([]byte, os.Getpagesize())
	for {
		n, err := fd.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		for i := 0; i < n; i++ {
			err = consumeByte(buf[i], cnf, &st, conf)
			if err != nil {
				return nil, err
			}
		}
		if err == io.EOF {
			break
		}
	}

	// terminate clause
	consumeByte(byte('\n'), cnf, &st, conf)
	if cnf.Lits[len(cnf.Lits)-1] != 0 {
		return nil, fmt.Errorf("Missing 0 to terminate last clause")
	}

	if conf.CheckNbClauses {
		if cnf.NbClauses != st.clauses {
			return nil, fmt.Errorf("Expected %d clauses, got %d clauses", cnf.NbClauses, st.clauses)
		}
	}

	return cnf, nil
}
