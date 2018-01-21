// Copyright 2018 Itamar Ostricher
// hashmap package providing a HashMap data structure (AKA map / associative array)
// for educational purposes (as builtin maps are superior)

package hashmap

import (
	"container/list"
	"fmt"
)

const dataLen uint32 = 37

// HashMap struct provides associative array semantics for string key/value pairs
// Implementation provides O(1) set, get & remove operations in average case
// Based on an underlying array of linked lists
type HashMap struct {
	len uint32
	data []*list.List
}

// Internal struct for HashMap entries
type hashMapEntry struct {
	key string
	value string
}

// Internal struct for keeping track of HashMap iteration
// Warning: not stable if the HashMap mutates in the middle of iteration
type hashMapIter struct {
	hm *HashMap
	index uint32
	node *list.Element
}

// Return an initialized HashMap
func New() (hm *HashMap) {
	hm = new(HashMap)
	hm.data = make([]*list.List, dataLen)
	for i := uint32(0); i < dataLen; i++ {
		hm.data[i] = list.New()
	}
	return
}

// Return the current number of elements in the HashMap - O(1)
func (hm HashMap) Len() uint32 {
	return hm.len
}

// Set the HashMap entry with `key` to `value`
// Overwrites existing entry with same `key`, or inserts a new entry
// O(1) in average case, O(N) in worst case (all entries under same hash)
func (hm *HashMap) Set(key, value string) {
	l := hm.findList(key)
	entry, _ := hm.findEntry(key, l)
	if entry != nil {
		entry.value = value
	} else {
		l.PushBack(hashMapEntry{key, value})
		hm.len++
		// TODO: grow data slice when it gets crowded
	}
}

// Return the value of the HashMap entry at `key` (or error if no such entry)
// O(1) in average case, O(N) in worst case (all entries under same hash)
func (hm HashMap) Get(key string) (string, error) {
	l := hm.findList(key)
	entry, _ := hm.findEntry(key, l)
	if entry != nil {
		return entry.value, nil
	}
	return "", fmt.Errorf("No entry with key %s", key)
}

// Remove the HashMap entry at `key`, if it exists (otherwise do nothing)
// O(1) in average case, O(N) in worst case (all entries under same hash)
func (hm *HashMap) Remove(key string) {
	l := hm.findList(key)
	_, node := hm.findEntry(key, l)
	if node != nil {
		l.Remove(node)
		hm.len--
	}
}

// Internal function for returning the relevant linked list for a given key by hash
func (hm HashMap) findList(key string) *list.List {
	h := hm.hashKey(key)
	return hm.data[h]
}

// Internal function for finding the list element within a specific list for a given key
// Returns both the "parsed" hashMapEntry and the list element itself
// Returns (nil, nil) if key doesn't exist in the list
func (hm HashMap) findEntry(key string, l *list.List) (*hashMapEntry, *list.Element) {
	for node := l.Front(); node != nil; node = node.Next() {
		entry := node.Value.(hashMapEntry)
		if entry.key == key {
			return &entry, node
		}
	}
	return nil, nil
}

// Internal function for computing the hash of a string key
func (hm HashMap) hashKey(key string) uint32 {
	var h uint32
	for _, b := range key {
		h = 199 * h + uint32(b)
	}
	return h % uint32(len(hm.data))
}

// Return a new iterator over the HashMap
func (hm HashMap) StartIter() (iter hashMapIter) {
	iter = hashMapIter{&hm, 0, hm.data[0].Front()}
	iter.findNext()
	return
}

// Return true if Next() is expected to succeed
func (iter hashMapIter) HasNext() bool {
	return iter.node != nil
}

// Return the next key/value pair from the underlying HashMap
func (iter *hashMapIter) Next() (key, value string) {
	entry := iter.node.Value.(hashMapEntry)
	key, value = entry.key, entry.value
	iter.node = iter.node.Next()
	iter.findNext()
	return
}

// Internal function for advancing the iterator to the next item
func (iter *hashMapIter) findNext() {
	for iter.node == nil {
		iter.index++
		if iter.index >= uint32(len(iter.hm.data)) {
			break
		}
		iter.node = iter.hm.data[iter.index].Front()
	}
}
