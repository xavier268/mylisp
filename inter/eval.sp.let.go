package inter

// The three binding constructs let, let*, and letrec, give Scheme block structure.
// The syntax of the three constructs is identical, but they differ in the regions
// they establish for their variable bindings.
// In a let expression, the initial values are computed before any of the variables become bound.
// In a let* expression, the evaluations and bindings are sequentially interleaved.
// And in a letrec expression, all the bindings are in effect while the initial values
// are being computed (thus allowing mutually recursive definitions).

func init() {
	Register("let", false, spLet)
}

// ( let ((variable init) ...) expression expression ... )
//
//	The inits are evaluated in the current environment (in some unspecified order),
//	the variables are bound to fresh locations holding the results,
//	the expressions are evaluated sequentially in the extended environment,
//	and the value of the last expression is returned.
//	Each binding of a variable has the expressions as its region.
//	MIT/GNU Scheme allows any of the inits to be omitted, in which case the corresponding variables are unassigned.
//	Note that the following are equivalent:
//	          (let ((variable init) ...) expression expression ...)
//	          ((lambda (variable ...) expression expression ...) init ...)
func spLet(it *Inter, t Term) Term { //  t is (((variable init) ...) expression expression ...)

	newev := it.current.Fork()

	// Vars and inits are evaluated in the current environment (in some unspecified order)
	// Vars are bound to fresh locations holding the results in newev
	if varinits := car(t); varinits != nil { // varinits is ((var1 init1) ( var2 init2) ...)

		for rest := varinits; rest != nil && (rest != Pair{}); rest = cdr(rest) {
			varinit := car(rest) // varinit is (var init)
			if va, ok := car(varinit).(Symbol); ok && (va != Symbol{}) {
				in := cadr(varinit)
				newev.Set(va, it.Eval(in))
			} else {
				return Error{ErrInvalidArgumentListFormat, varinit}
			}
		}
	}

	it.current = newev // now, we evaluate in the new environement we just prepared with the new variable bindings
	defer it.PopEnv()

	var result Term
	for exs := cdr(t); exs != nil && (exs != Pair{}); exs = cdr(exs) { // exs is ( ex1 ex2 ex3 ... )
		ex := car(exs)
		if ex != nil {
			result = it.Eval(ex)
		}
	}
	return result
}
