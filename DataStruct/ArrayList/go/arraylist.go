// Copyright 2018 Itamar Ostricher
// arraylist package providing on ArrayList data structure (AKA vector)
// for educational purposes (as builtin slices are more than sufficient)

package arraylist

const MIN_CAPACITY = 8

// Item type for ArrayList values
type ItemType interface{}

// ArrayList struct is a flexible array (auto-growing, auto-shrinking)
// with random-access read/write semantics,
// and O(1) amortized append/pop operationsץ
type ArrayList struct {
  length uint32
  capacity uint32
  data []ItemType
}

// Return the current length (used items) of the ArrayList - O(1)
func (arr *ArrayList) Len() uint32 {
  return arr.length
}

// Return the item at position `index` - O(1)
func (arr *ArrayList) Get(index uint32) ItemType {
  if index < arr.length {
    return arr.data[index]
  }
  panic("out of bounds")
}

// Set the value of the item in position `index` - O(1)
func (arr *ArrayList) Set(index uint32, newItem ItemType) {
  if index < arr.length {
    arr.data[index] = newItem
  } else {
    panic("out of bounds")
  }
}

// Append a new item at the end of the ArrayList, growing the underlying array if
// current capacity is not enough, in amortized O(1) runtime (O(N) when resized)
func (arr *ArrayList) Append(item ItemType) {
  if arr.length == arr.capacity {
    arr.resize(arr.capacity * 2)
  }
  arr.data[arr.length] = item
  arr.length++
}

// Remove and return an item from the end of the ArrayList, shrinking the underlying
// if it's underutilized, in amortized O(1) runtime (O(N) when shrinking)
func (arr *ArrayList) Pop() ItemType {
  if arr.length > 0 {
    // shrink by half if only a third is used - dampening resize operations
    if arr.length < arr.capacity / 3 {
      arr.resize(arr.capacity / 2)
    }
    arr.length--
    return arr.data[arr.length]
  }
  panic("out of bounds")
}

// Resize the underlying array to `newCapacity` in O(N) runtime and O(N) in memory.
// Doesn't do anything if `newCapcity` is the same as current capacity,
// and doesn't allowing shrinking below MIN_CAPACITY capacity.
func (arr *ArrayList) resize(newCapacity uint32) {
  if newCapacity < MIN_CAPACITY {
    newCapacity = MIN_CAPACITY
  }
  if (newCapacity > arr.capacity) {
    var newData []ItemType = make([]ItemType, newCapacity)
    copy(newData, arr.data)
    arr.data = newData
    arr.capacity = newCapacity
  } else if newCapacity < arr.capacity {
    arr.data = arr.data[:newCapacity]
    arr.capacity = newCapacity
  }
}
