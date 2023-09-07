package inter

import "fmt"

type Inter struct {
	local, global *Environnement
}

func NewInter() *Inter {
	it := new(Inter)
	it.PushEnv() // initialize both local and global scopes
	return it
}

func (it Inter) String() string {
	return fmt.Sprintf("--- Interpreter state ---\n\nLocal:\n%s\nGlobal:\n%s", it.local, it.global)
}
