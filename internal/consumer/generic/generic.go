// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package generic

import (
	"fmt"
	"os"

	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/field"
)

const Name = "generic"

var _ consumer.Consumer = new(generic)

type generic struct {
	indexPattern string
	files        []string
}

func (c *generic) Consume(state *consumer.State) ([]field.Usage, error) {
	var usages []field.Usage
	for _, f := range c.files {
		got, err := Parse(state.FieldMap, c.indexPattern, f)
		if err != nil {
			return nil, err
		}

		usages = append(usages, got...)
	}

	return usages, nil
}

func New(indexPattern string, files []string) (consumer.Consumer, error) {
	if indexPattern == "" {
		return nil, fmt.Errorf("%s: an index pattern is required", Name)
	}

	c := generic{
		indexPattern: indexPattern,
	}

	for _, f := range files {
		if stat, err := os.Stat(f); err == nil && !stat.IsDir() {
			c.files = append(c.files, f)
		}
	}

	if len(c.files) == 0 {
		return nil, fmt.Errorf("%s: no files found", Name)
	}

	fmt.Println(c.files)

	return &c, nil
}

func init() {
	consumer.Register(Name, New)
}
