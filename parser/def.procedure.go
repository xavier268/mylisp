package parser

import "fmt"

// A procedure is a Term that cannot be entered directly by the user, but is created by a lambda application.
// Doc says : Procedures are created by evaluating lambda expressions (see Lambda Expressions);
// the lambda may either be explicit or may be implicit as in a “procedure define” (see Definitions).
// Also there are special built-in procedures, called primitive procedures,
// such as car; these procedures are not written in Scheme but in go.
type Procedure struct {
	Primitive func(Term) Term // a primitive procedure if this is not nil
	Env       *Env            // the environement captured by the procedure, or nil if the procedure should run in the current eval environement
	Name      string          // the name of the procedure, if a primitive procedure. Could be "".
	Formals   Term            // typically a list of symbols, but could also ba a symbol to represent the parameter list, or even a pair, or any term made of symbols.
	Body      Term            // typically a list of terms
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
	if p.Primitive != nil {
		return fmt.Sprintf("#<primitive procedure %s>", p.Name)
	} else {
		return fmt.Sprintf("#<procedure %s>\n%s --> %s", p.Name, p.Formals, p.Body)
	}
}
