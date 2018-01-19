CC=clang++
CC_FLAGS=-std=c++14 -Wall -fvectorize -fslp-vectorize -fcolor-diagnostics -O2 -stdlib=libc++ -lc++

arraylist-cpp:
	${CC} -o build/arraylist-demo ${CC_FLAGS} DataStruct/ArrayList/cpp/demo.cc
	./build/arraylist-demo

arraylist-go:
	cd DataStruct/ArrayList/go && go test

linkedlist-go:
	cd DataStruct/LinkedList/go && go test

clean:
	rm build/*
