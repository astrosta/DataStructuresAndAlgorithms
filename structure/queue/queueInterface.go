package queue

type Queue interface {
	Enqueue(value interface{}) bool
	Dequeue() interface{}
	String() string
}
