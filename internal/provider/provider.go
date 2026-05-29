// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package provider

import (
	"context"
	"fmt"

	"github.com/taylor-swanson/flint/internal/field"
)

type State struct {
	Ctx      context.Context
	FieldMap *field.Map
	CacheDir string
}

type Provider interface {
	Provide(state *State) error
}

type NewFunc func(path string) (Provider, error)

var registry = map[string]NewFunc{}

// New creates a new Provider from the given config.
func New(kind, path string) (Provider, error) {
	newFn, exists := registry[kind]
	if !exists {
		return nil, fmt.Errorf("provider: unknown provider: %q", kind)
	}

	return newFn(path)
}

// Register will register a provider with a given name.
func Register(name string, fn NewFunc) {
	if _, exists := registry[name]; exists {
		panic("provider: '" + name + "' already registered")
	}
	registry[name] = fn
}
