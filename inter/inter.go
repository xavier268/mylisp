package inter

import "fmt"

type Inter struct {
	current, global *Environnement // TODO - do we really need global env ?
	OnQuit          func()         // a function hook, called when the interpreter is quitting.
}

func NewInter() *Inter {
	it := new(Inter)
	it.PushEnv() // initialize both current and global scopes
	return it
}

func (it Inter) String() string {
	return fmt.Sprintf("--- Interpreter state ---\n\nCurrent:\n%s\nGlobal:\n%s", it.current, it.global)
}
