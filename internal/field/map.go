// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package field

import (
	"fmt"
	"regexp"
	"strings"
)

var reReplacer = strings.NewReplacer("*", ".*", ".", "\\.")

type Map struct {
	fields           map[uint64]*Field
	indices          map[uint64]*Index
	definitions      map[uint64]*Definition
	fieldDefinitions map[*Field][]*Definition
	indexDefinitions map[*Index][]*Definition
}

func NewMap() *Map {
	return &Map{
		fields:           make(map[uint64]*Field),
		indices:          make(map[uint64]*Index),
		definitions:      make(map[uint64]*Definition),
		fieldDefinitions: make(map[*Field][]*Definition),
		indexDefinitions: make(map[*Index][]*Definition),
	}
}

func (m *Map) addField(field *Field) *Field {
	h := field.Hash()
	if existing, ok := m.fields[h]; ok {
		return existing
	}

	m.fields[h] = field
	m.fieldDefinitions[field] = nil

	return field
}

func (m *Map) addIndex(index *Index) *Index {
	h := index.Hash()
	if existing, ok := m.indices[h]; ok {
		return existing
	}

	m.indices[h] = index
	m.indexDefinitions[index] = nil

	return index
}

func (m *Map) AddDefinition(definitions ...Definition) {
	for _, def := range definitions {
		dh := def.Hash()
		if _, ok := m.definitions[dh]; ok {
			continue
		}

		def.Field = m.addField(def.Field)
		def.Index = m.addIndex(def.Index)

		defPtr := &def
		m.definitions[dh] = defPtr
		m.indexDefinitions[def.Index] = append(m.indexDefinitions[def.Index], defPtr)
		m.fieldDefinitions[def.Field] = append(m.fieldDefinitions[def.Field], defPtr)
	}
}

// Lookup the usage of a field by name and index pattern.
func (m *Map) Lookup(name, indexPattern string) (Usage, error) {
	indexPatternCompat := "^" + reReplacer.Replace(indexPattern) + "$"
	indexRe, err := regexp.Compile(indexPatternCompat)
	if err != nil {
		return Usage{}, fmt.Errorf("invalid index pattern: %w", err)
	}

	usage := Usage{
		Name:    name,
		Pattern: indexPattern,
	}

	var tmpls []*Index
	for _, v := range m.indices {
		if indexRe.MatchString(v.Pattern) {
			tmpls = append(tmpls, v)
		}
	}

	for _, tmpl := range tmpls {
		for _, d := range m.indexDefinitions[tmpl] {
			if d.Field.Name == name {
				usage.Matches = append(usage.Matches, d)
				break
			}
		}
	}

	return usage, nil
}
