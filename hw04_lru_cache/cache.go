package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
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

func (c *lruCache) Set(key Key, value interface{}) bool {
	element, exists := c.items[key]

	if exists {
		val := element.Value.(cacheItem)
		val.value = value
		element.Value = val

		c.queue.MoveToFront(element)
	} else {
		element = c.queue.PushFront(cacheItem{key: key, value: value})
	}

	if c.queue.Len() > c.capacity {
		back := c.queue.Back()
		if back != nil {
			c.queue.Remove(back)
			delete(c.items, back.Value.(cacheItem).key)
		}
	}

	if c.capacity > 0 {
		c.items[key] = element
	}

	return exists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := c.items[key]

	if exists {
		result := element.Value.(cacheItem)
		c.queue.MoveToFront(element)
		return result.value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
