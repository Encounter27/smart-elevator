package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func NewQueue() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q *Queue) Push(element interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, element)
	q.len++
}

func (q *Queue) Pop() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	elm := q.queue[0]
	q.queue = q.queue[1:]
	q.len++

	return elm, nil
}

func (q *Queue) Front() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	return q.queue[0], nil
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}

func (q *Queue) GetAll() []interface{} {
	return q.queue
}
