// Package functionalslices
// fontains public api and most implementations for functional slice operations like map and filter
package functionalslices

import "iter"

func Filter[T any](s []T, cond func(T) bool) (r []T) {
	r = make([]T, 0)
	for _, v := range s {
		if cond(v) {
			r = append(r, v)
		}
	}
	return
}

func Map[A any, B any](s []A, fn func(A) B) (r []B) {
	r = make([]B, len(s))
	for k, v := range s {
		r[k] = fn(v)
	}
	return
}

func FoldR[A any, B any](start A, slice []B, fn func(A, B) A) (r A) {
	r = start
	for _, v := range slice {
		r = fn(r, v)
	}
	return
}

func ParMap[A any, B any](s []A, fn func(A) B, threadlimit int) (r []B) {
	r = make([]B, len(s))
	tlch := make(chan bool, threadlimit)
	for range threadlimit {
		tlch <- true
	}
	for k, v := range s {
		go func() {
			<-tlch
			r[k] = fn(v)
		}()
		tlch <- true
	}
	for range threadlimit {
		<-tlch
	}
	return
}

func MapIter[A any, B any](s []A, fn func(A) B) iter.Seq[B] {
	return func(yild func(B) bool) {
		for _, v := range s {
			if !yild(fn(v)) {
				return
			}
		}
	}
}

func MapIter2[A any, B any](s []A, fn func(A) B) iter.Seq2[int, B] {
	return func(yild func(int, B) bool) {
		for k, v := range s {
			if !yild(k, fn(v)) {
				return
			}
		}
	}
}

func FilterIter[T any](s []T, cond func(T) bool) iter.Seq[T] {
	return func(yild func(T) bool) {
		for _, v := range s {
			if !cond(v) {
				continue
			}
			if !yild(v) {
				return
			}
		}
	}
}

// FilterIter2 teakes the key is the index from the original slice
func FilterIter2[T any](s []T, cond func(T) bool) iter.Seq2[int, T] {
	return func(yild func(int, T) bool) {
		for k, v := range s {
			if !cond(v) {
				continue
			}
			if !yild(k, v) {
				return
			}
		}
	}
}
