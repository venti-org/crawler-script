package script

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func NewVM() *otto.Otto {
	vm := otto.New()
	return vm
}

type ScriptBase struct {
	vm   *otto.Otto
	this otto.Value
}

func NewScriptBase(vm *otto.Otto) (*ScriptBase, error) {
	if vm == nil {
		vm = NewVM()
	}
	return &ScriptBase{
		vm:   vm,
		this: otto.NullValue(),
	}, nil
}

func NewScriptBaseWithScript(vm *otto.Otto, script string) (*ScriptBase, error) {
	if s, err := NewScriptBase(vm); err != nil {
		return nil, err
	} else if _, err := s.GetVM().Run(script); err != nil {
		return nil, err
	} else {
		return s, nil
	}
}

func (s *ScriptBase) GetVM() *otto.Otto {
	return s.vm
}

func (s *ScriptBase) Call(name string, args ...interface{}) (otto.Value, error) {
	callback, err := s.vm.Get(name)
	if err != nil {
		return otto.UndefinedValue(), err
	}
	if !callback.IsFunction() {
		return otto.UndefinedValue(), fmt.Errorf(name + " function is not defined")
	}
	return callback.Call(s.this, args...)
}

func (s *ScriptBase) CallIfExists(name string,
	args ...interface{}) (otto.Value, error) {
	callback, err := s.vm.Get(name)
	if err != nil {
		return otto.UndefinedValue(), err
	}
	if !callback.IsFunction() {
		return otto.UndefinedValue(), nil
	}
	return callback.Call(s.this, args...)
}
