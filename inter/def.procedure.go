package inter

import "fmt"

// A procedure is a Term that cannot be entered directly by the user, but is created by a lambda application.
// All procedures evaluate their arguments.
// Doc says : Procedures are created by evaluating lambda expressions (see Lambda Expressions);
// the lambda may either be explicit or may be implicit as in a “procedure define” (see Definitions).
type Procedure struct {
	Env     *Environnement // the environement captured by the procedure, or nil if the procedure should run in the current eval environement
	Formals Term           // typically a list of symbols, but could also be a symbol to represent the parameter list, or even a pair, or any term made of symbols.
	Body    Term           // typically a list of terms
}

// IsError implements Term.
func (Procedure) IsError() bool {
	return false
}

var _ Term = Procedure{}

// IsProcedure implements Term.
func (Procedure) IsProcedure() bool {
	return true
}

// IsBool implements Term.
func (Procedure) IsBool() bool {
	return false
}

// IsNumber implements Term.
func (Procedure) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (Procedure) IsPair() bool {
	return false
}

// IsString implements Term.
func (Procedure) IsString() bool {
	return false
}

// IsSymbol implements Term.
func (Procedure) IsSymbol() bool {
	return false
}

// String implements Term.
func (p Procedure) String() string {
	return fmt.Sprintf("#<procedure : >\n%s\n\t---> %s", p.Formals, p.Body)
}
