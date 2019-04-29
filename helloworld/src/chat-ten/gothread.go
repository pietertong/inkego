package chat_ten

import (
    "fmt"
    "runtime"
    "sync"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func GoThread(){
    go say("world")
    say("Hello")
}

var counter int = 0

func gCount(lock *sync.Mutex){
    lock.Lock()
    counter ++
    fmt.Println(counter)
    lock.Unlock()
}

func GoThreadSharMemory()  {
    lock := &sync.Mutex{}
    for i := 0; i < 10; i++{
        go gCount(lock)
    }
    
    for {
        lock.Lock()
        c := counter
        fmt.Println(fmt.Sprintf("lock main %s",c))
        lock.Unlock()
        runtime.Gosched()
        if c >= 10{
            break
        }
    }
}