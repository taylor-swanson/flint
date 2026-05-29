// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package field

import "hash/fnv"

type Index struct {
	Name    string
	Pattern string
	ECSRef  string
	Version string
	Stack   string
}

func (t *Index) Hash() uint64 {
	h := fnv.New64a()

	_, _ = h.Write([]byte(t.Name))
	_, _ = h.Write([]byte(t.Pattern))
	_, _ = h.Write([]byte(t.ECSRef))
	_, _ = h.Write([]byte(t.Version))
	_, _ = h.Write([]byte(t.Stack))

	return h.Sum64()
}
