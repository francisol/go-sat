# go-sat

Go bindings for multiple popular SAT solvers. Currently supported solvers include:

- [CaDiCaL](https://github.com/arminbiere/cadical) (2.0.0)

```go
package main

## Example
import (
	"fmt"

	"github.com/francisol/go-sat/cadical"
)

func main() {
	solver := cadical.NewCaDiCaLSolver()
	solver.AddClause([]int{1, -2})
	solver.AddClause([]int{-1, 2})
	solver.AddClause([]int{1, 2})
	sat, _ := solver.Solve()
	if sat {
		model, _ := solver.Model()
		fmt.Println(model)
	}
	fmt.Println("ok")
}

```