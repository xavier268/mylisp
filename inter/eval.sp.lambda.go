package inter

import "fmt"

// syntax ( lambda ( x1 x2 ...)  body1 body2 ...)

// A lambda expression evaluates to a procedure.
// The environment in effect when the lambda expression is evaluated is remembered as part of the procedure;
// it is called the closing environment. When the procedure is later called with some arguments,
// the closing environment is extended by binding the variables in the formal parameter list
// to fresh locations, and the locations are filled with the arguments according to rules about to be given.
// The new environment created by this process is referred to as the invocation environment.

// Once the invocation environment has been constructed, the expressions in the body of
// the lambda expression are evaluated sequentially in it.
// This means that the region of the variables bound by the lambda expression is all
// of the expressions in the body.
// The result of evaluating the last expression in the body
// is returned as the result of the procedure call.

var ErrInvalidLambaFormals = fmt.Errorf("lambda: formals must be a  a list of symbols, possibly empty")

func init() {
	Register("lambda", false, spLambda, []string{
		"syntax ( lambda ( x1 x2 ...)  body1 body2 ...)",
		"A lambda expression evaluates to a procedure.",
		"The environment in effect when the lambda expression is evaluated is remembered as part of the procedure;",
		"it is called the closing environment. When the procedure is later called with some arguments,",
		"the closing environment is extended by binding the variables in the formal parameter list",
		"to fresh locations, and the locations are filled with the arguments according to rules about to be given.",
		"The new environment created by this process is referred to as the invocation environment.",
		"Once the invocation environment has been constructed, the expressions in the body of",
		"the lambda expression are evaluated sequentially in it.",
		"This means that the region of the variables bound by the lambda expression is all",
		"of the expressions in the body.",
		"The result of evaluating the last expression in the body",
		"is returned as the result of the procedure call.",
	})
}

func spLambda(it *Inter, t Term) Term { // t contains ( ( x1 x2 ...)  body1 body2 ...)

	p := Procedure{
		Env:     it.current, // capture closing environnement. Will be forked before use.
		Formals: nil,
		Body:    cdr(t), // body contains ( expr1 expr 2 ...), list of expressions to be evaluated.
	}

	// capture formals
	fr := car(t)                   // fr contains ( x1 x2 ...)  or ( start . rest ) or ()
	if fr == nil || !fr.IsPair() { // empty pair, (),  would be acceptable as formals, but not nil.
		return Error{
			Err:     ErrInvalidLambaFormals,
			Context: fr,
		}
	}
	p.Formals = fr

	return p
}
