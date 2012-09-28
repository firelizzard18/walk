// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type Composite struct {
	Widget     **walk.Composite
	Name       string
	HStretch   int
	VStretch   int
	Row        int
	RowSpan    int
	Column     int
	ColumnSpan int
	Layout     Layout
	Children   []Widget
}

func (c Composite) Create(parent walk.Container) (walk.Widget, error) {
	w, err := walk.NewComposite(parent)
	if err != nil {
		return nil, err
	}

	var succeeded bool
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	if err := initWidget(c, w); err != nil {
		return nil, err
	}

	w.SetName(c.Name)

	if c.Widget != nil {
		*c.Widget = w
	}

	succeeded = true

	return w, nil
}

func (c Composite) LayoutParams() (hStretch, vStretch, row, rowSpan, col, colSpan int) {
	return c.HStretch, c.VStretch, c.Row, c.RowSpan, c.Column, c.ColumnSpan
}

func (c Composite) Layout_() Layout {
	return c.Layout
}

func (c Composite) Children_() []Widget {
	return c.Children
}
