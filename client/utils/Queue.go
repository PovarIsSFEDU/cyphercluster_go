package utils

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []Message
	Size     int
}

func (q *Queue) Enqueue(elem Message) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() Message {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return Message{}
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (Message, error) {
	if q.IsEmpty() {
		return Message{}, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
