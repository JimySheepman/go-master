package behavioral

import "fmt"

type Iteratorer interface {
	HasNext() bool
	Next() interface{}
}

// Collection arayüzü
type Collection interface {
	CreateIterator() Iteratorer
}

// ConcreteCollection yapısı
type ConcreteCollection struct {
	items []string
}

func (c *ConcreteCollection) CreateIterator() Iteratorer {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

// ConcreteIterator yapısı
type ConcreteIterator struct {
	collection *ConcreteCollection
	index      int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.collection.items[i.index]
		i.index++
		return item
	}
	return nil
}

func Iterator() {
	collection := &ConcreteCollection{
		items: []string{"Item1", "Item2", "Item3"},
	}

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
