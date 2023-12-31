package main

func containsDuplicate(nums []int) bool {
	set := NewSet()
	for _, num := range nums {
		if set.Has(num) {
			return true
		}
		set.Add(num)
	}
	return false
}

type Set struct {
	values map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		values: make(map[int]struct{}),
	}
}

func (s *Set) Add(val int) {
	s.values[val] = struct{}{}
}

func (s *Set) Has(val int) bool {
	_, ok := s.values[val]
	return ok
}

// более короткое

func containsDuplicate2(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = true
	}
	return false
}
