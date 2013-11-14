package mailbox

type queue struct {
	arr []interface{}
	ind int
}

func newQueue() *queue {
	return &queue{arr: make([]interface{}, 0), ind: 0}
}

func (q *queue) Size() int {
	return q.ind
}

func (q *queue) Empty() bool {
	return q.ind == 0
}

func (q *queue) Front() interface{} {
	if q.ind == 0 {
		return nil
	} else {
		return q.arr[q.ind-1]	
	}
}

func (q *queue) Back() interface{} {
	if q.ind == 0 {
		return nil
	} else {
		return q.arr[0]
	}
}

func (q *queue) Push(i interface{}) {
	if len(q.arr) <= q.ind {
		q.arr = append(q.arr, i)	
	} else {
		q.arr[q.ind] = i
	}
	q.ind++
}

func (q *queue) Pop() interface{} {
	q.ind--
	return q.arr[q.ind + 1]
}

func (q *queue) Clear() []interface{} {
	outArr := make([]interface{}, q.ind)
	for ; q.ind >= 0; q.ind-- {
		outArr[q.ind] = q.arr[q.ind]
	}
	return outArr
}