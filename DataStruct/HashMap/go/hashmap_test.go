// Copyright 2018 Itamar Ostricher

package hashmap

import (
  "fmt"
  "strconv"
	"testing"
)

func TestHashMap(t *testing.T) {
	hm := New()
  if hm.Len() != 0 {
    t.Errorf("Expected new HashMap to have length 0, got %d", hm.Len())
  }
  _, err := hm.Get("foo")
  if err == nil {
    t.Errorf("Expected Get on empty HashMap to return err, got no err")
  }
  hm.Set("foo", "bar")
  if hm.Len() != 1 {
    t.Errorf("Expected HashMap to have length 1 after Set, got %d", hm.Len())
  }
  foo, err := hm.Get("foo")
  if err != nil {
    t.Errorf("Expected Get(foo) to return with no error, got %v", err)
  }
  if foo != "bar" {
    t.Errorf("Expected Get(foo) to return bar, got %v", foo)
  }
  hm.Remove("foo")
  hm.Remove("bar")
  if hm.Len() != 0 {
    t.Errorf("Expected HashMap after Remove to have length 0, got %d", hm.Len())
  }
  _, err = hm.Get("foo")
  if err == nil {
    t.Errorf("Expected Get on removed key to return err, got no err")
  }
  iter := hm.StartIter()
  if iter.HasNext() {
    t.Errorf("Expected StartIter on empty HashMap to be empty too, got HasNext")
  }
  for i := 0; i < 100; i++ {
    hm.Set(fmt.Sprintf("key %d", i), fmt.Sprintf("%d", i))
  }
  if hm.Len() != 100 {
    t.Errorf("Expected HashMap to have length 100 after Set-loop, got %d", hm.Len())
  }
  var freq [100]int
  iter = hm.StartIter()
  for iter.HasNext() {
    _, value := iter.Next()
    num, _ := strconv.Atoi(value)
    freq[num]++
  }
  for i := 0; i < 100; i++ {
    if freq[i] != 1 {
      t.Errorf("Expected to see key/value pair for %d exactly once, seen %d times", i, freq[i])
    }
  }
}
