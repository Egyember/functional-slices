// Package functionalslices
// fontains public api and most implementations for functional slice operations like map and filter
package functionalslices

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
