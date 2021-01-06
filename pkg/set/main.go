package set

import (
	"fmt"
	"sort"
)

type StringSet map[string]bool

func (self *StringSet) Add(s string) {
	(*self)[s] = true
}
func (self *StringSet) Remove(s string) {
	delete(*self, s)
}
func (self *StringSet) Pop() string {
	for k := range *self {
		delete(*self, k)
		return k
	}
	panic("Operation on empty Set")
}
func (self *StringSet) Contains(s string) bool {
	if _, there := (*self)[s]; there {
		return true
	} else {
		return false
	}
}
func (self *StringSet) ToSlice() []string {
	s := make([]string, 0, len(*self))
	for k := range *self {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}
func (self *StringSet) String() string {
	return fmt.Sprint((*self).ToSlice())
}

func NewStringSet(d ...string) StringSet {
	s := make(StringSet, len(d))
	for _, v := range d {
		s[v] = true
	}
	return s
}

type IntSet map[int64]bool

func (self *IntSet) Add(s int64) {
	(*self)[s] = true
}
func (self *IntSet) Remove(s int64) {
	delete(*self, s)
}
func (self *IntSet) Pop() int64 {
	for k := range *self {
		delete(*self, k)
		return k
	}
	panic("Operation on empty Set")
}
func (self *IntSet) Contains(s int64) bool {
	if _, there := (*self)[s]; there {
		return true
	} else {
		return false
	}
}
func (self *IntSet) ToSlice() []int64 {
	s := make([]int64, 0, len(*self))
	for k := range *self {
		s = append(s, k)
	}
	sort.Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
	return s
}
func (self IntSet) String() string {
	return fmt.Sprint(self.ToSlice())
}

func NewIntSet(d ...int64) IntSet {
	s := make(IntSet, len(d))
	for _, v := range d {
		s[v] = true
	}
	return s
}
