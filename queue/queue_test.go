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

func TestQueue_Dequeue(t *testing.T) {
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
		want   fields
	}{
		{
			"should dequeue 2 items",
			fields{[]interface{}{3, 2, 0, 5, 1}, 0, 4, 5},
			args{2},
			fields{[]interface{}{0, 5, 1}, 0, 2, 3},
		},
		{
			"should dequeue 3 items",
			fields{[]interface{}{2, 0, 5, 1}, 0, 3, 5},
			args{3},
			fields{[]interface{}{1}, 0, 1, 1},
		},
		{
			"should dequeue 3 items",
			fields{[]interface{}{0, 5, 1}, 0, 2, 5},
			args{3},
			fields{[]interface{}{}, 0, 0, 0},
		},
		{
			"should dequeue 1 item",
			fields{[]interface{}{0, 5, 1}, 0, 2, 5},
			args{1},
			fields{[]interface{}{5, 1}, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := Queue{
				items:  tt.fields.items,
				front:  tt.fields.front,
				rear:   tt.fields.rear,
				length: tt.fields.length,
			}
			for i := 0; i < tt.args.i; i++ {
				queue.Dequeue()
			}
			if got := queue.items; got != tt.want.items {
				t.Errorf("Queue.Enqueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
