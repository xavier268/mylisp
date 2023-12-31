package inter

// Identifiers that are predefined.
// They all evaluate to themselves.
// Builtin should self register in this map using Register function.
var keywordsDefinitions = make(map[string]kwdata, 10)

// All primitive, built in functions or macros must have this signature.
// They will be called with args containing the cdr of the form ( keyword .... ).
// detailed format depends on the function.
type Primitive func(it *Inter, args Term) Term

type kwdata struct {
	evalArgFirst bool      // ie, should we eval the arguments before calling the primitive function ?
	primitive    Primitive // the primitive function to call.
	help         []string  // help text for keyword
}

// Is this string the name of a registered keyword ?
func IsKeyword(s string) bool {
	_, ok := keywordsDefinitions[s]
	return ok
}

// Register a keyword for a built-in primitive, typically during initialization.
func Register(keyword string, evalArgFirst bool, primitive Primitive, help []string) {
	if IsKeyword(keyword) {
		panic("trying to registered twice an existing keyword : " + keyword)
	}
	keywordsDefinitions[keyword] = kwdata{evalArgFirst, primitive, help}
}
