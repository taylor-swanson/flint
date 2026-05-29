// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package field

import "go/token"

type Usage struct {
	// Name of the field.
	Name string
	// Pattern used for looking up the index.
	Pattern string
	// Position of the usage.
	Position token.Position
	// Matches of field definitions.
	Matches []*Definition
}
