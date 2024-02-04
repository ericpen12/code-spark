package sync

import "fmt"

type queue struct {
	headTail int64
	data     []int
}

const queueBit = 4

func (q *queue) enQueue(data int) {
	q.data = append(q.data, data)
	q.headTail += 1 << queueBit
}

func (q *queue) deQueue() {
	q.data = q.data[1:]
	head, tail := q.unpack()
	head--
	const mask = 1<<queueBit - 1
	q.headTail = (head << queueBit) | (tail & mask)
}

func (q *queue) unpack() (int64, int64) {
	const mask = 1<<queueBit - 1
	head := (q.headTail >> queueBit) & mask
	tail := q.headTail & mask
	return head, tail
}

func (q *queue) info(opt string) {
	head, tail := q.unpack()
	fmt.Printf("操作： %s headTail: %8b, head: %8b, tail:%8b\n", opt, q.headTail, head, tail)
}
