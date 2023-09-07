package inter

func IsLitteral(t Term) bool {
	if t == nil {
		return true
	}
	switch a := t.(type) {
	case Bool:
		return true
	case Number:
		return true
	case String:
		return true
	case Error:
		return true
	case Procedure:
		return true
	case Symbol: //  keywords
		if IsKeyword(a.Value) {
			return true // keywords eval to themselves
		}
	case Pair:
		if (a == Pair{}) {
			return true // empty list
		}
	}
	// all other cases ...
	return false
}
