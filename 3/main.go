package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func main() {
	m := New()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)
	if m.Exists("two") {
		fmt.Println("Key 'two' exists")
	}
	if value, ok := m.Get("three"); ok {
		fmt.Printf("Value for 'three': %d\n", value)
	}
	copyMap := m.Copy()
	fmt.Printf("Copied map: %v\n", copyMap)
	m.Remove("one")
	fmt.Printf("After removal, 'one' exists: %v\n", m.Exists("one"))
}
func New() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}