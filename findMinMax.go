package main

func findMinMax(a []int) (minn int, maxx int) {
	minn = a[0]
	maxx = a[0]
	for i := 1; i < len(a); i++ {
		minn = min(minn, a[i])
		maxx = max(maxx, a[i])
	}
	return
}
