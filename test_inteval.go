package main

import (
    "context"
    "time"

    "collectd.org/api"
    "collectd.org/exec"

    "fmt"
    "github.com/egidijus/go-libzfs"
    "log"
    "strconv"
)

func main() {

return Interval()

func Interval() time.Duration {
    i, err := strconv.ParseFloat(os.Getenv("COLLECTD_INTERVAL"), 64)
    if err != nil {
        log.Printf("unable to determine default interval: %v", err)
        return time.Second * 10
    }

    return time.Duration(i * float64(time.Second))
}

}