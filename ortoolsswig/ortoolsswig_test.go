package ortoolsswig

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	cmpOpts = []cmp.Option{
		cmpopts.EquateApprox(1e-4, 1e-4),
	}
)

func TestNewSolver(t *testing.T) {
	tests := []struct {
		name      string
		got, want interface{}
	}{
		{"test 1", "a", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, tt.got); diff != "" {
				t.Errorf("unexpected diff (-want, +got):\n%s", diff)
			}
		})
	}
}

// TestSolver is based on https://developers.google.com/optimization/lp/glop.
func TestSolver(t *testing.T) {
	solver := NewSolver("LinearProgrammingExample", SolverGLOP_LINEAR_PROGRAMMING)
	x := solver.NumVar(0, SolverInfinity(), "x")
	y := solver.NumVar(0, SolverInfinity(), "y")

	// Constraint 0: x + 2y <= 14.
	constraint0 := solver.Constraint(-SolverInfinity(), float64(14))
	//constraint0 := solver.Constraint()
	constraint0.SetCoefficient(x, 1)
	constraint0.SetCoefficient(y, 2)

	// Constraint 1: 3x - y >= 0.
	constraint1 := solver.Constraint(0.0, SolverInfinity())
	constraint1.SetCoefficient(x, 3)
	constraint1.SetCoefficient(y, -1)

	// Constraint 2: x - y <= 2.
	constraint2 := solver.Constraint(-SolverInfinity(), 2.0)
	constraint2.SetCoefficient(x, 1)
	constraint2.SetCoefficient(y, -1)

	// Objective function: 3x + 4y.
	objective := solver.Objective()
	objective.SetCoefficient(x, 3)
	objective.SetCoefficient(y, 4)
	objective.SetMaximization()

	status := solver.Solve()
	t.Logf("solver status: %v", status)

	opt := 3*x.SolutionValue() + 4*y.SolutionValue()
	t.Logf("optimizal solution: 3 * %v + 4 * %v = %v", x.SolutionValue(), y.SolutionValue(), opt)

	if got, want := solver.NumVariables(), 2; got != want {
		t.Errorf("got %d variables, want %d", got, want)
	}
	if got, want := solver.NumConstraints(), 3; got != want {
		t.Errorf("got %d variables, want %d", got, want)
	}
	if got, want := x.SolutionValue(), 6.0; !cmp.Equal(got, want, cmpOpts...) {
		t.Errorf("got x_opt = %v, want %v", got, want)
	}
	if got, want := y.SolutionValue(), 4.0; !cmp.Equal(got, want, cmpOpts...) {
		t.Errorf("got y_opt = %v, want %v", got, want)
	}
}
