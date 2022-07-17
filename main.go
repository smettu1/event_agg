package main

import (
    "bufio"
    "github.com/events/metrics"
    "log"
    "os"
    "sync"
    "time"
)

func main() {
    scanner := bufio.NewScanner((os.Stdin))
    var  startConsuming bool = false
    var  endConsuming bool = false

    //Create channel to communicate
    msgChan := make(chan string)
    wg := &sync.WaitGroup{}
    sm := &sync.Map{}
    go metrics.ConsumeData(wg,msgChan,sm)
    wg.Wait()
    startTime := time.Now()
    //Loop to read the input stream of data
    for {
        scanner.Scan()
        input := scanner.Text()
        // If the input sends "BEGIN" then start consuming
        if input == "BEGIN" {
            startConsuming = true
            log.Println("Start consuming")
            continue
        } else if input == "END" {
            endConsuming = true
            log.Println("End consuming")
        }
        //start consuming the stream of data
        if startConsuming {
            msgChan <- input
        }
        //end consuming the stream and dump stats
        if endConsuming {
            metrics.PrintMetrics(sm)
            log.Println("Total time", time.Now().Sub(startTime))
            close(msgChan)
            break
        }
    }
}