package queue

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		{"should create a new Queue", &Queue{make([]interface{}, initialCap), 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		items  []interface{}
		front  int
		rear   int
		length int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"should enqueue 6", fields{[]interface{}{3, 2, 0, 5, 1}, 0, 4, 5}, args{6}},
		{"should enqueue 7", fields{[]interface{}{2, 0, 5, 1}, 0, 3, 5}, args{7}},
		{"should enqueue 8", fields{[]interface{}{0, 5, 1}, 0, 2, 5}, args{8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := Queue{
				items:  tt.fields.items,
				front:  tt.fields.front,
				rear:   tt.fields.rear,
				length: tt.fields.length,
			}
			queue.Enqueue(tt.args.i)
			want := queue.items[queue.front]
			if got := queue.Dequeue(); got != want {
				t.Errorf("Queue.Enqueue() = %v, want %v", got, want)
			}
		})
	}
}
