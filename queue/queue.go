package queue

const initialCap = 10

type Queue struct {
	items  []interface{}
	front  int
	rear   int
	length int
}

func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	q.items = make([]interface{}, initialCap)
	q.length, q.front, q.rear = 0, 0, 0
	return q
}

func (q Queue) Enqueue(item interface{}) {
	q.ensureCapacity()
	q.items[q.rear] = item
	q.rear++
	q.length++
}

func (q Queue) ensureCapacity() {
	if q.length == len(q.items) {
		q.resize(q.length * 2)
	}
}

func (q Queue) resize(size int) {
	newItems := make([]interface{}, size)
	copy(newItems, q.items)
}

func (q Queue) Dequeue() interface{} {
	item := q.items[q.front]
	q.items[q.front] = nil
	q.front++
	q.shrink()
	return item
}

func (q Queue) shrink() {

}
