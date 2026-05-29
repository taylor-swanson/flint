// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package file

import (
	"go/token"
	"reflect"
)

type fileAnnotator struct {
	filename string
}

func (a fileAnnotator) annotate(val reflect.Value) {
	if val.CanAddr() && val.CanSet() {
		if p, ok := val.Addr().Interface().(*token.Position); ok {
			p.Filename = a.filename
			return
		}
	}

	switch val.Kind() {
	case reflect.Pointer:
		if !val.IsNil() {
			a.annotate(val.Elem())
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			a.annotate(val.Field(i))
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			a.annotate(val.Index(i))
		}
	case reflect.Map:
		itr := val.MapRange()
		for itr.Next() {
			a.annotate(itr.Value())
		}
	}
}

func Annotate(filename string, v any) {
	fileAnnotator{filename: filename}.annotate(reflect.ValueOf(v))
}
