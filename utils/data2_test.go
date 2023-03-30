package utils_test

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Slice[T constraints.Ordered] []T

func (p *Slice[T]) Clip() { *p = slices.Clip(*p) }

func (p Slice[T]) Cls() {}
