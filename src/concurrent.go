package main

import (
    "os"
    "fmt"
    "time"
    "sync"
    "strconv"
)

func count(wg *sync.WaitGroup, feedback chan int64) {
    // Note: channels can be passed by value, but waitgroups
    //  MUST be passed by reference
    defer wg.Done()

    var start = time.Now()
    for i := 0; i < 1_000_000_000; i++ {
        _ = i
    }

    elapsed := time.Now().Sub(start)
    fmt.Printf("Thread ended after %v\n", elapsed)

    feedback <- elapsed.Milliseconds()
}

func getTotal() (int, error) {
    return strconv.Atoi(os.Args[1])
}

func main() {
    total, err := getTotal()
    if err != nil {
        fmt.Println("Could not parse argument")
        os.Exit(1)
    }

    feedback := make(chan int64, total)

    var wg sync.WaitGroup
    wg.Add(total)

    var t1 = time.Now()

    for i := 0; i < total; i++ {
        go count(&wg, feedback)
    }

    fmt.Printf("Kicked off after %v\n", time.Now().Sub(t1))

    wg.Wait()

    var duration_counter int64
    for i := range total {
        _ = i
        duration_counter += <-feedback
    }

    completion_time := time.Now().Sub(t1).Milliseconds()
    fmt.Printf("Completed in %d ms , cumulative duration %d ms, ratio %f\n",
        completion_time,
        duration_counter,

        // Either end the line with "," , or close it with ")"
        //    but MUST use one of these - cannot end a list without comma
        float64(duration_counter)/float64(completion_time),
        )
}
