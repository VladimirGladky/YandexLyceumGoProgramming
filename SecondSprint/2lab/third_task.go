package main

import "sync"

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
	queue []interface{} // здесь хранить элементы очереди
	mutex sync.Mutex
}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	c.queue = append(c.queue, element)
	c.mutex.Unlock()
}

func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.queue) == 0 {
		return nil
	}
	elem := c.queue[0]
	c.queue = c.queue[1:]
	return elem
}
