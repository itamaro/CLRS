CC=clang++
CC_FLAGS=-std=c++14 -Wall -fvectorize -fslp-vectorize -fcolor-diagnostics -O2 -stdlib=libc++ -lc++

arraylist-cpp:
	${CC} -o build/arraylist-demo ${CC_FLAGS} DataStruct/ArrayList/cpp/demo.cc
	./build/arraylist-demo

arraylist-go:
	cd DataStruct/ArrayList/go && go test -cover

linkedlist-go:
	cd DataStruct/LinkedList/go && go test -cover

reverse-linkedlist:
	cd Algorithms/ReverseLinkedList/go && go run reverse.go && go test

hashmap-go:
	cd DataStruct/HashMap/go && go test -cover

stack-go:
	cd DataStruct/Stack/go && go test -cover

hanoi-stack-go:
	cd DataStruct/HanoiTower/go && go test -cover

queue-go:
	cd DataStruct/Queue/go && go test -cover

sort-go:
	cd Algorithms/Sort/go && go test -cover

bench-sort-go:
	cd Algorithms/Sort/go && go test -bench .

hanoi-go:
	cd Algorithms/HanoiTowers/go && go run hanoitowers.go

clean:
	rm build/*
