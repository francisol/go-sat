package cadical

// #cgo CFLAGS: -I../third_parts/cadical/src -I../third_parts/cadical/ -DNBUILD
// #include "ccadical.h"
import "C"
import (
	"errors"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

type CaDiCaLSolver struct {
	inner *C.CCaDiCaL
	vars  int
}

func NewCaDiCaLSolver() *CaDiCaLSolver {

	return &CaDiCaLSolver{
		inner: C.ccadical_init(),
		vars:  0,
	}
}

func (solver *CaDiCaLSolver) AddClause(clause []int) {
	for _, lit := range clause {
		solver.vars = max(solver.vars, abs(lit))
		C.ccadical_add(solver.inner, C.int(lit))
	}
	C.ccadical_add(solver.inner, 0)
}
func (solver *CaDiCaLSolver) Solve() (bool, error) {
	result := C.ccadical_solve(solver.inner)
	if result == 10 {
		return true, nil
	}
	if result == 20 {
		return false, nil
	}
	return false, errors.New("limit reached or interrupted through 'terminate'")
}

func (solver *CaDiCaLSolver) Model() ([]int, error) {
	result := []int{}
	for i := range solver.vars {
		if C.ccadical_val(solver.inner, C.int(i+1)) > 0 {
			result = append(result, i+1)
		}

	}
	return result, nil
}
