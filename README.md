lynxbat/ringbuffer
==========

A Circular Buffer written in Go.

This is a work in progress.

It will support ring buffer structs of different types:

* Always keep one slot open (default: RingBuffer)
* Fill count - TODO
* Mirror bit - TODO
* Read/Write count - TODO
* Absolute indices - TODO
* Last op - TODO

TODO

    * [ ] Threadsafety
    * [x] Integer support for element value
    * [ ] Bytes support for element value
    * [ ] File-backed for persistence
    * [ ] Read access through API (UNIX? NET?)
    * [ ] Fun command line tools
    * [ ] Tests
    
    
Current benchmarks using an integer element value on a new Macbook Pro Retina:

```
go test -bench . -benchtime 10s -benchmem
PASS
BenchmarkWrite	500000000	        46.4 ns/op	       8 B/op	       1 allocs/op
BenchmarkRead	5000000000	         6.21 ns/op	       0 B/op	       0 allocs/op
ok  	github.com/lynxbat/ringbuffer	59.578s
```

Test script:

There is a command script in `scripts/ringbuffercmd.go` that is useful for seeing how this works while wiki and godoc is under work.



