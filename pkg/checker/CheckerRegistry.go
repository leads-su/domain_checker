package checker

import "fmt"

type CheckerRegistry struct {
	checkers map[string]CheckerFactoryInterface
}

var checkerRegistry = &CheckerRegistry{
	checkers: make(map[string]CheckerFactoryInterface),
}

func GetCheckerRegistry() *CheckerRegistry {
	return checkerRegistry
}

func (registry *CheckerRegistry) Add(name string, checkerFactory CheckerFactoryInterface) *CheckerRegistry {
	registry.checkers[name] = checkerFactory
	return registry
}

func (registry *CheckerRegistry) Get(name string) CheckerFactoryInterface {
	checkerFactory, ok := registry.checkers[name]
	if !ok {
		panic(fmt.Sprintf("not found checker name %s", name))
	}

	return checkerFactory
}

func (registry *CheckerRegistry) ContainsName(name string) bool {
	_, ok := registry.checkers[name]
	return ok
}

func (registry *CheckerRegistry) GetAvailableNames() *[]string {
	i, keys := 0, make([]string, len(registry.checkers))
	for key, _ := range registry.checkers {
		keys[i] = key
		i++
	}
	return &keys
}
