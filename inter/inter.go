package inter

import "fmt"

type Inter struct {
	current, global *Environnement
	quit            bool // quit requested
}

func NewInter() *Inter {
	it := new(Inter)
	it.PushEnv() // initialize both current and global scopes
	return it
}

func (it Inter) String() string {
	return fmt.Sprintf("--- Interpreter state ---\n\nCurrent:\n%s\nGlobal:\n%s", it.current, it.global)
}
