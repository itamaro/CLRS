// Copyright 2018 Itamar Ostricher

#ifndef _CLRS_ARRAY_LIST_
#define _CLRS_ARRAY_LIST_

#include <algorithm>

// minimum ArrayList capacity
const uint MIN_CAP = 8;

class OutOfBoundError : public std::runtime_error {
 public:
  explicit OutOfBoundError(char const* const message) throw()
    : std::runtime_error(message) { }

  virtual char const* what() const throw() {
    return exception::what();
  }
};

// Templated class implementing an ArrayList -
// a flexible size array with random access and append/pop semantics
template <class T>
class ArrayList {
 private:
  T * _data;
  uint _capacity;
  uint _length;

 public:
  ArrayList()
    : _data(new T[MIN_CAP]),
      _capacity(MIN_CAP),
      _length(0) { }

  ~ArrayList() {
    delete [] _data;
  }

  // Appends an item to the end of the ArrayList
  // Runtime is O(1) amortized (when array is doubled it'll be O(N) for copying)
  void append(const T& item) {
    if (_length == _capacity) {
      _resize(2 * _capacity);
    }
    _data[_length++] = item;
  }

  // Random access const operator - O(1)
  const T& operator[](uint i) const {
    if (i < _length)
      return _data[i];
    throw OutOfBoundError("bad index");
  }

  // Random access mutable operator - O(1)
  T& operator[](uint i) {
    if (i < _length)
      return _data[i];
    throw OutOfBoundError("bad index");
  }

  // Remove and return the last element in the array
  // Runtime is O(1) amortized (when array shrinks it'll be O(N) for copying)
  T pop() {
    if (_length > 0) {
      if (_length < _capacity / 3) {
        _resize(_capacity / 2);
      }
      return _data[--_length];
    }
    throw OutOfBoundError("empty");
  }

  // Return the current length (used items) of the array - O(1)
  uint length() const {
    return _length;
  }

 private:
  // Resize the underlying array to have a new capacity `new_cap`
  // New capacity can be greater than or lesser than current capacity,
  // but must be greater than MIN_CAP - otherwise it doesn't do it.
  // Runtime is O(N) copy operations * O(copy(T))
  void _resize(uint new_cap) {
    // don't do anything if lesser than MIN_CAP or identical to current capacity
    if (new_cap >= MIN_CAP && new_cap != _capacity) {
      T * new_data = new T[new_cap];
      // could have used memcpy for PODs, but this is more general
      std::copy(_data, _data + _length, new_data);
      delete [] _data;
      _data = new_data;
      _capacity = new_cap;
    }
  }
};

#endif  // _CLRS_ARRAY_LIST_
