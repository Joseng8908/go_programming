package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Add(6)
	s.Add(7)
	s.Add(8)
	s.Add(9)
	s.Add(10)
	fmt.Println(s.String())
	s.Remove(1)
	fmt.Println(s.String())
	sNew := s.Copy()
	fmt.Println(sNew.String())

}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.words[word]&(1<<bit) == 0 {
		return
	}
	s.words[word] ^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var copyIntSet *IntSet
	copyIntSet = new(IntSet)
	copyIntSet.words = make([]uint64, len(s.words))
	copy(copyIntSet.words, s.words)
	return copyIntSet
}

func (s *IntSet) AddAll(nums ...int) {
	var i int
	for i = 0; i < len(nums); i++ {
		s.Add(nums[i])
	}
}
