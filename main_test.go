package functionalslices

import (
	"testing"
)

func TestFilter(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i
	}
	r := Filter(s, func(i int) bool {
		return i >= 5
	})
	if len(r) != 5 {
		t.Errorf("wrong length")
	}
	for i := 5; i < 10; i++ {
		if r[i-5] != i {
			t.Errorf("wrong value")
		}
	}
}

func TestMap(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = 1
	}
	r := Map(s, func(i int) int {
		if i == 1 {
			return 2
		}
		return 3
	})
	if len(r) != 10 {
		t.Errorf("wrong length")
	}
	for i := range 10 {
		if r[i] != 2 {
			t.Errorf("wrong value")
		}
	}
}

func TestMap_order(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i
	}
	r := Map(s, func(i int) int {
		return i
	})
	for i := range 10 {
		if r[i] != i {
			t.Errorf("wrong value")
		}
	}
}

func TestParMap(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = 1
	}
	r := Map(s, func(i int) int {
		if i == 1 {
			return 2
		}
		return 3
	})
	p := ParMap(s, func(i int) int {
		if i == 1 {
			return 2
		}
		return 3
	}, 4)
	for i := range 10 {
		if r[i] != p[i] {
			t.Errorf("wrong value")
		}
	}
}

func TestParMap_order(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i
	}
	p := ParMap(s, func(i int) int {
		return i
	}, 4)
	for i := range 10 {
		if p[i] != i {
			t.Errorf("wrong value")
		}
	}
}

func TestMapIter(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = 1
	}
	for i := range MapIter(s, func(i int) int {
		if i == 1 {
			return 2
		}
		return 3
	}) {
		if i != 2 {
			t.Errorf("wrong value")
		}
	}
}

func TestMapIter_order(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i
	}
	oldi := 0
	for i := range MapIter(s, func(i int) int {
		return i
	}) {
		if i != oldi {
			t.Errorf("wrong value")
		}
		oldi++
	}
}

func TestMapIter2(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = 1
	}
	for _, i := range MapIter2(s, func(i int) int {
		if i == 1 {
			return 2
		}
		return 3
	}) {
		if i != 2 {
			t.Errorf("wrong value")
		}
	}
}

func TestMapIter2_order(t *testing.T) {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i
	}
	oldi := 0
	for k, i := range MapIter2(s, func(i int) int {
		return i
	}) {
		if i != k {
			t.Errorf("wrong value")
		}
		if k != oldi {
			t.Errorf("wrong key")
		}
		oldi++
	}
}
