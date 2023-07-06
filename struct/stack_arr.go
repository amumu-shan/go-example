package _struct

import (
	"strconv"
	"strings"
)

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (s *Stack) Push(n int) {
	if s.ix+1 > LIMIT {
		return
	}
	s.data[s.ix] = n
	s.ix++
}
func (s *Stack) Pop() int {
	s.ix--
	return s.data[s.ix]
}
func (s *Stack) String() string {
	builder := strings.Builder{}
	for i, v := range s.data {
		builder.WriteString("[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]")
	}

	return builder.String()
}
