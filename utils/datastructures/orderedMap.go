package datastructures

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// OrderedMap represents an ordered map.
type OrderedMap[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

// NewOrderedMap creates a new OrderedMap instance.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

// Set inserts or updates a key-value pair in the ordered map.
func (om *OrderedMap[K, V]) Set(key K, value V) {
	// If the key already exists, update the value
	if _, ok := om.values[key]; ok {
		om.values[key] = value
		return
	}

	// Otherwise, add the key to the keys slice and set the value in the map
	om.keys = append(om.keys, key)
	om.values[key] = value
}

// Get retrieves the value associated with the given key from the ordered map.
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, ok := om.values[key]
	return value, ok
}

// Keys returns the keys of the ordered map in the order they were inserted.
func (om *OrderedMap[K, V]) Keys() []K {
	return om.keys
}

// Values returns the values of the ordered map in the order they were inserted.
func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, len(om.keys))
	for i, key := range om.keys {
		values[i] = om.values[key]
	}
	return values
}

// Values returns the values of the ordered map in the order they were inserted.
func (om *OrderedMap[K, V]) FromYaml(node yaml.Node) (*OrderedMap[K, V], error) {
	if node.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("pipeline must contain YAML mapping, has %v", node.Kind)
	}

	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		valueNode := node.Content[i+1]

		var k K
		if err := keyNode.Decode(&k); err != nil {
			return nil, fmt.Errorf("error decoding key %v on line %v", keyNode.Tag, keyNode.Line)
		}

		var v V
		if err := valueNode.Decode(&v); err != nil {
			return nil, fmt.Errorf("error decoding value %v on line %v", valueNode.Tag, valueNode.Line)
		}

		om.Set(k, v)
	}
	return om, nil
}
