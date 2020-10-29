##  序言

> 在一次远程面试的时候，要求写一个示例中不能出现data race的情况。当时是使用各种sync.Mutex解决
> 驳回第一次  使用锁太多，虽然能绝交竞争的问题，但是锁太多会导致性能下降。 使用channel+mutex解决
> 驳回第二次  代码格式不标准(汗)

看上去好像阔以，但事实是我觉得代码还是有很多问题，我对并发条件的处理很懵逼，我该哪种方式处理呢？ 还是好好看看书吧

## 概念

- dataRace
> 竞争 

```golang
var data int
go func() {
    data++
}()
if data == 0 {
    fmt.Println("Hello World", data)
}
time.Sleep(1 * time.Second)
```
> 当前条件下data不一定总是0,能打印 “hell world” 
- 执行 `go run -race main.go ` 发现 data++ 会出现锁竞争的情况出现

- 原子性
原子性就是一起执行或者一起不执行，换句话就是同生共死

> 书中提及到一个说法 程序的每次处理是否属于原子性操作、取决于当前的上下文的环境决定
`i++` 
1 缓存i
2 i+1
3 赋值i=i=1
当前的行为是分3步走战略，所以可以完美的解释上面第一点提交到的竞争的出现

- 死锁

> 并发写code的时候，一定会遇到的情况就是死锁问题

```golang
type value struct {
  mu sync.Mutex
  value int
}

var wg sync.WaitGroup
printSum:=func ( v 1 , v2 * value) {
    defer wg.Done()
    vl.mu.Lock()
    defer v1.mu.Unlock()
    time.Sleep(2 * tirne.Second)
    v2.mu.Lock()
    defer v2.mu.Unlock()
    fmt.Printf(”surn =% v \n ", vi.v a lu e + v2.value)
)

var a, b valu e
wg.Add(2)
go printSurn(& a, &b)
go printSurn(&b, &a)
wg.Wait()
```
> 开启两个协程对处理a与b,上锁会失败

避免情况
- 互相等待

- 相互排斥

## 核心

- CSP
  
> 并发编程模型，两个独立的并发实体通过共享的通讯 Channel（管道）进行通信的并发模型。
 
> 传闻早期，有位学者提出一个csp理论。提出并发编程应该如同顺序编程的方式一样写

> 每次处理都有标准的输入 以及管道 标准的输出。

> golang基于此设计出了channel以及goroutine，极大增加了golang的并发性能以及使用。


了解下历史，知道所有程序语言设计初衷都有一个很伟大的想法


## channel还是sync.Mutex

如何抉择? 正如我一开始所说,我根本不了解我该用什么，我只知道可以解决问题