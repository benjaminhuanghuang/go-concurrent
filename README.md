# Go Concurrent


## Key Point
goroutine
channels
Buffered Channels
Range和Close
Select
超时
runtime goroutine


## goroutine
```
```




## Channel
Communicating Sequential Processes



### Use channel
- create
```
make(chan <Type>, size)
```
- send
```
  ch <- 10
```
- receive
```
  x:= <- ch
```
- close
```
  close(ch)
```
