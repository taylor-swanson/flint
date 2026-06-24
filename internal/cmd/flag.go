// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package cmd

import (
	"fmt"
	"strings"
)

type kvFlag struct {
	Key   string
	Value string
}

func (f *kvFlag) String() string {
	return f.Key + ":" + f.Value
}

func (f *kvFlag) Type() string {
	return "kvFlag"
}

func (f *kvFlag) Set(value string) error {
	parts := strings.SplitN(value, ":", 2)

	if len(parts) != 2 {
		return fmt.Errorf("invalid flag format: %q (expect 'key:value')", value)
	}

	f.Key = parts[0]
	f.Value = parts[1]

	return nil
}

type kvListFlag []kvFlag

func (f *kvListFlag) String() string {
	if len(*f) == 0 {
		return ""
	}

	s := make([]string, len(*f))
	for i, v := range *f {
		s[i] = v.String()
	}

	return strings.Join(s, ", ")
}

func (f *kvListFlag) Type() string {
	return "kvListFlag"
}

func (f *kvListFlag) Set(value string) error {
	var kv kvFlag
	if err := kv.Set(value); err != nil {
		return err
	}

	*f = append(*f, kv)

	return nil
}
