// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package field

import (
	"encoding/binary"
	"go/token"
	"hash/fnv"
)

// Type represents the Elasticsearch mapping type of field.
type Type string

const (
	TypeAggregateMetricDouble Type = "aggregate_metric_double"
	TypeAlias                 Type = "alias"
	TypeHistogram             Type = "histogram"
	TypeConstantKeyword       Type = "constant_keyword"
	TypeText                  Type = "text"
	TypeMatchOnlyText         Type = "match_only_text"
	TypeKeyword               Type = "keyword"
	TypeLong                  Type = "long"
	TypeInteger               Type = "integer"
	TypeShort                 Type = "short"
	TypeByte                  Type = "byte"
	TypeDouble                Type = "double"
	TypeFloat                 Type = "float"
	TypeHalfFloat             Type = "half_float"
	TypeScaledFloat           Type = "scaled_float"
	TypeDate                  Type = "date"
	TypeDateNanos             Type = "date_nanos"
	TypeBoolean               Type = "boolean"
	TypeBinary                Type = "binary"
	TypeIntegerRange          Type = "integer_range"
	TypeFloatRange            Type = "float_range"
	TypeLongRange             Type = "long_range"
	TypeDoubleRange           Type = "double_range"
	TypeDateRange             Type = "date_range"
	TypeIPRange               Type = "ip_range"
	TypeGroup                 Type = "group"
	TypeGeoPoint              Type = "geo_point"
	TypeObject                Type = "object"
	TypeIP                    Type = "ip"
	TypeNested                Type = "nested"
	TypeFlattened             Type = "flattened"
	TypeWildcard              Type = "wildcard"
	TypeVersion               Type = "version"
	TypeUnsignedLong          Type = "unsigned_long"
	TypeCountedKeyword        Type = "counted_keyword"
	TypeSemanticText          Type = "semantic_text"
	TypeGeoShape              Type = "geo_shape"
)

// Kind represents the kind of field (ECS, vendor, etc.).
type Kind string

const (
	KindECS    Kind = "ecs"
	KindVendor Kind = "vendor"
)

type Field struct {
	Name   string  `json:"name"`
	Kind   Kind    `json:"kind"`
	Type   Type    `json:"type"`
	Target *string `json:"target,omitempty"`
	Array  bool    `json:"array"`
}

func (f *Field) Hash() uint64 {
	h := fnv.New64a()

	_, _ = h.Write([]byte(f.Name))
	_, _ = h.Write([]byte(f.Kind))
	_, _ = h.Write([]byte(f.Type))
	if f.Target != nil {
		_, _ = h.Write([]byte(*f.Target))
	}
	_ = binary.Write(h, binary.LittleEndian, f.Array)

	return h.Sum64()
}

func Join(prefix, name string) string {
	if prefix == "" {
		return name
	}

	return prefix + "." + name
}

// Definition is a declaration of a field in an index.
type Definition struct {
	Field    *Field
	Index    *Index
	Position token.Position
}

func (d *Definition) Hash() uint64 {
	h := fnv.New64a()

	_ = binary.Write(h, binary.LittleEndian, d.Field.Hash())
	_ = binary.Write(h, binary.LittleEndian, d.Index.Hash())

	_, _ = h.Write([]byte(d.Position.Filename))
	_ = binary.Write(h, binary.LittleEndian, int32(d.Position.Line))
	_ = binary.Write(h, binary.LittleEndian, int32(d.Position.Column))

	return h.Sum64()
}
