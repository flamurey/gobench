# Install

Run command
```bash
go install github.com/flamurey/gobench@latest
```

# Ho to run

CMD format:

```bash
gobench [-workers=size_of_worker_pool] [-threads=size_of_thread_pool] [-single]
```

To see more details run:
```bash
gobench -h
```

# Results

Intel i7-3770k

```
Running bench with config {WorkersCount:1 ThreadsCount:1} 
Results of light payload
Total:  38438869         Average:      52ns 
Results of hard payload
Total:      7575         Average: 264.026µs 
Running bench with config {WorkersCount:8 ThreadsCount:8} 
Results of light payload
Total: 126238882         Average:      15ns 
Results of hard payload
Total:     29406         Average:  68.013µs  
```