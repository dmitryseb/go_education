package main

func (q *Queue) push(elem int) {
	q.back = append(q.back, elem)
}

func (q *Queue) pop() int {
	if len(q.front)+len(q.back) == 0 {
		return -1
	}
	if len(q.front) == 0 {
		q.front = q.back[:(len(q.back)+1)/2]
		for i := 0; i < len(q.front)/2; i++ {
			q.front[i], q.front[len(q.front)-i-1] = q.front[len(q.front)-i-1], q.front[i]
		}
		q.back = q.back[(len(q.back)+1)/2:]
	}
	res := q.front[len(q.front)-1]
	q.front = q.front[:len(q.front)-1]
	return res
}

type Queue struct {
	back, front []int
}
