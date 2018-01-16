// Copyright 2018 Itamar Ostricher
// Demo my ArrayList implementation

#include <iostream>

#include "ArrayList.h"

template <class T>
void PrintArray(const ArrayList<T>& arr) {
  for (int i=0; i < arr.length(); ++i) {
    std::cout << arr[i] << ", ";
  }
  std::cout << std::endl;
}

int main() {
  ArrayList<int> my_arr;
  for (int i=1; i < 1000; i += 17) {
    my_arr.append(i);
  }
  my_arr[1] = 8;
  PrintArray(my_arr);
  while (my_arr.length()) {
    std::cout << my_arr.pop() << " - ";
  }
  std::cout << std::endl;
  return 0;
}
