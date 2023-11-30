package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type element struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.items[key]
	if !ok {
		return val, false
	}
	c.queue.MoveToFront(val)
	return val.Value.(element).value, true
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.items[key]; !ok {
		el := element{
			key:   key,
			value: value,
		}
		head := c.queue.PushFront(el)
		c.items[key] = head
		if c.queue.Len() > c.capacity {
			lastElem := c.queue.Back()
			lastElemKey := lastElem.Value.(element)
			c.queue.Remove(lastElem)
			delete(c.items, lastElemKey.key)
		}
		return false
	} else {
		v.Value = element{
			key:   key,
			value: value,
		}
		c.queue.MoveToFront(v)
		return true
	}
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[Key]*ListItem, c.capacity)
}
