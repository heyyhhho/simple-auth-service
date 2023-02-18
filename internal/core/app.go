package core

import "fmt"

type Container interface {
	Get(key string) (interface{}, error)
	Set(key string, c interface{})
	MustGet(key string) interface{}
}

type ServiceProvider interface {
	Register(c Container)
}

type BootableProvider interface {
	Boot(c Container) error
}

type App struct {
	items     map[string]interface{}
	instances map[string]interface{}
	providers []ServiceProvider
}

func NewApp() *App {
	return &App{
		items:     make(map[string]interface{}),
		instances: make(map[string]interface{}),
	}
}

func (a *App) Register(serviceProvider ServiceProvider) {
	serviceProvider.Register(a)
	a.providers = append(a.providers, serviceProvider)
}

func (a *App) MustGet(key string) interface{} {
	if val, err := a.Get(key); err != nil {
		panic(err)
	} else {
		return val
	}
}

func (a *App) Set(key string, c interface{}) {
	a.items[key] = c
}

func (a *App) Get(key string) (interface{}, error) {
	item, ok := a.items[key]
	if !ok {
		return nil, fmt.Errorf("identifier '%s' is not defined in Container", key)
	}

	if fn, ok := item.(func(c Container) interface{}); ok {
		if instance, exists := a.instances[key]; exists {
			return instance, nil
		} else {
			a.instances[key] = fn(a)
			return a.instances[key], nil
		}
	} else {
		return item, nil
	}
}

func (a *App) Boot() error {
	for _, provider := range a.providers {
		if p, suitable := provider.(BootableProvider); suitable {
			if err := p.Boot(a); err != nil {
				return err
			}
		}
	}

	return nil
}
