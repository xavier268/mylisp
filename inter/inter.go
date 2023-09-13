package inter

import "fmt"

type Inter struct {
	current *Environnement //current environnement. Should never be nil.
	OnQuit  func()         // a function hook, called when the interpreter is quitting.
}

func NewInter() *Inter {
	it := new(Inter)
	it.PushEnv() // initialize both current and global scopes
	return it
}

func (it Inter) String() string {
	return fmt.Sprintf("--- Interpreter state ---\n\nEnvironnement:\n%s", it.current)
}
