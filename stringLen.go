package main

func stringLen(s string) int {
	ln := 0
	t := s
	for t != "" {
		t = t[1:]
		ln++
	}
	return ln
}
