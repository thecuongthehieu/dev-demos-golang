package main

import (
	"fmt"

	"github.com/draffensperger/golp"
)

//Ref: https://lpsolve.sourceforge.net/5.5/integer.htm

/*
min: -x1 -2 x2 +0.1 x3 +3 x4;
r_1: +x1 +x2 <= 5;
r_2: +2 x1 -x2 >= 0;
r_3: -x1 +3 x2 >= 0;
r_4: +x3 +x4 >= 0.5;
bin x3;
*/

func main() {
	lp := golp.NewLP(0, 4)
	lp.AddConstraintSparse([]golp.Entry{{0, 1.0}, {1, 1.0}}, golp.LE, 5.0)
	lp.AddConstraintSparse([]golp.Entry{{0, 2.0}, {1, -1.0}}, golp.GE, 0.0)
	lp.AddConstraintSparse([]golp.Entry{{0, 1.0}, {1, 3.0}}, golp.GE, 0.0)
	lp.AddConstraintSparse([]golp.Entry{{2, 1.0}, {3, 1.0}}, golp.GE, 0.5)
	lp.SetObjFn([]float64{-1.0, -2.0, 0.1, 3.0})
	lp.SetBinary(2, true)
	lp.Solve()

	fmt.Printf("Objective value: %v\n", lp.Objective())
	vars := lp.Variables()
	fmt.Printf("Variable values:\n")
	for i := 0; i < lp.NumCols(); i++ {
		fmt.Printf("x%v = %v\n", i+1, vars[i])
	}
}
