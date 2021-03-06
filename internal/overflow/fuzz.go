package overflow

import (
    "fmt"
    "net"
    "time"
)

// sleep duration of 1s
const sleep time.Duration = 1000000000

// the main functionality of the fuzz subroutine
func Fuzz(host string, port int, step int, pref, suff string) {
    var err error
    var length int

    for length = step; err == nil; length += step {
        // create n-byte byte array
        data := generateBytes(0x41, length)

        // build payload
        fmt.Println(" > Building payload.")
        payload := createPayload(data, pref, suff)

        // send payload to target service
        fmt.Printf(" > Sending %d-byte payload.\n", len(payload))
        err = sendPayload(host, port, payload)

        time.Sleep(sleep)
    }

    if length - step == step {
        // no payload was sent, an error occurred
        fmt.Printf("\n Error! %s\n", err.(*net.OpError).Err)
        return
    }

    // print buffer information
    fmt.Printf("\n Success! Length of buffer is in range (%d, %d].\n",
        length - step - step, length - step)
}

