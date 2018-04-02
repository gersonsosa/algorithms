package main

import "fmt"

const initalCap = 10

type Queue struct {
	items  []interface{}
	top    int
	lenght int
}

func New() *Queue {
	return new(Queue).Init()
}

func (q Queue) Init() *Queue {
	q.items = make([]interface{}, initialCap)
	q.top, q.lenght = 0
	return q
}

func (q Queue) Enqueue(item interface{}) {
	fmt.Printf("Item value %v", item)
}
