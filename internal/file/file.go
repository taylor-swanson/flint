// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package file

import (
	"bytes"
	"go/token"
	"os"
)

type File struct {
	filename   string
	data       []byte
	lineBreaks []int
}

func (f *File) Bytes() []byte {
	return f.data
}

func (f *File) Filename() string {
	return f.filename
}

func (f *File) Position(offset int) token.Position {
	if offset >= len(f.data) {
		panic("offset out of range")
	}

	var line int
	var lineOffset int
	for i := range f.lineBreaks {
		if offset < f.lineBreaks[i] {
			break
		}

		line = i + 1
		lineOffset = f.lineBreaks[i]
	}

	return token.Position{
		Filename: f.filename,
		Offset:   offset,
		Line:     line,
		Column:   offset - lineOffset,
	}
}

func Load(filename string) (*File, error) {
	var err error

	f := File{filename: filename, lineBreaks: []int{-1}}
	f.data, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var offset int
	for {
		i := bytes.IndexByte(f.data[offset:], '\n')
		if i == -1 {
			break
		}

		f.lineBreaks = append(f.lineBreaks, offset+i)
		offset += i + 1
	}

	return &f, nil
}
