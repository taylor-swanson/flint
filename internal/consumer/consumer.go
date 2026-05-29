// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package consumer

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

type Consumer interface {
	Consume(state *State) ([]field.Usage, error)
}

// NewFunc is a function that creates a new Consumer with the given config.
type NewFunc = func(indexPattern string, paths []string) (Consumer, error)

var registry = map[string]NewFunc{}

// New creates a new Consumer from the given config.
func New(kind, indexPattern string, paths []string) (Consumer, error) {
	newFn, exists := registry[kind]
	if !exists {
		return nil, fmt.Errorf("consumer: unknown consumer: %q", kind)
	}

	return newFn(indexPattern, paths)
}

// Register will register a consumer with a given name.
func Register(name string, fn NewFunc) {
	if _, exists := registry[name]; exists {
		panic("consumer: '" + name + "' already registered")
	}
	registry[name] = fn
}
