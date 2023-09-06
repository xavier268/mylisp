package inter

type Inter struct {
	local, global *Scope
}

func NewInter() *Inter {
	it := new(Inter)
	it.PushScope() // initialize both local and global scopes
	return it
}
