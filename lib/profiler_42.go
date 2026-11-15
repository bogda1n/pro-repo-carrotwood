package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.6" )

func bcXqLehcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPIvcNtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5Z6C1KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8LpWAWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0tx4dNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4APzSu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwmSc2gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PhxBLnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjkRQYLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5fXoF0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRAKmstJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLnESU7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqETTL2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdqOqsAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MAgalEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k7XUfO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkyLRBXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpE8kACcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5g02RKEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzTHuPAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6ngJ8ILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auaLEklbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5EP3ECDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XN6ppVfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81AaYyveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLm18M0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzxwbaqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmX2SvMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9c4T0GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QI6MdGtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lb44oyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0XDwOWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F11jMFgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2Dbdr9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WKn83T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaqwfyCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Stwt1EoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xctOOjHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOn6yMP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func husGpklTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bu2j0Y63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfi9mPTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pJRzQ8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ujc6czwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1vUnoeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuLXD71gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBOkhINdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXuVTWCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmXqJK8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTtZz6OxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mGWfCBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIAKphTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pBZ61yJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJzbgrK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoF8BODnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3TsZF7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dURpy4L4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQ8mDuCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbAearKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVDTQTv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jh65BdpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwDJ2nHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcK6sAshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftCkoNs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VI04URt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3jAU61VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWvjvqFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vv8NWk3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJYLF8WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEt7b1S8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuRPQS1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpNIqBYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwmrICDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcO5xOuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROiAqcmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwtErmuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwBomR8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YS8idoUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFaQFIToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeaHNOuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6EUj1HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foV5Te9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIG3Mpy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyMbaGfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chJGYznaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmt3RFYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSt6NKeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDbxBsS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMrnH2pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYOvyqJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1knLdeIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gh3gatucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmd5lCZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AM75qCLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7107GsdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6hSK2urWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRzTYKXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkSErdWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy8lTQDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49nLsm7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpghrSp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Xy66x1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EoF4RzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq1pbGBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQcBnXk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWGv3RtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07CvC110Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1BaGexYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGYqVZZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o68aZICIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyhT8YbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TFjnGtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gvujmnh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6yiY4ToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ICJZAU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fvlGujCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HKz2cYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHbThFDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkH992zgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myWuW9WoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmNzR4qSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9z9fatTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV8LoYkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNcXOIYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kukyMwnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bW9yM0c9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l77rcIHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8dRMA5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3eD2n0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXzpgEESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5cvIgt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kK0oFmkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOhWCuegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13YRHlQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4qugONkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIX9iqRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wbWW9mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R87dM1BOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzNJx63OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYBFA8GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIVu0fpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63EkTzzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beNiha1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jVwsqyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PQ4V2h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJE2CL3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLmVl08QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l77RDOdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pxY8ybcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivjQvyaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFFTRLGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eyfccm70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23TMZ8awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSIduqRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7mVWi09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxMshDiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onfIqiKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNe0ZCZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BONKFq8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QSp5IPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6BpUisgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1o74LdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HFuxsF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGmZtjpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRiXADVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYgcuQ1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0W2Yx19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgmgphjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjmQvkWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNoxCai3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbNYLJR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0m3hMk63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvXa0X8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hezrN2MdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtKhMgqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08JKhAhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0gcOM9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSuqrXFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMKWxQJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lTUeGC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7xhNJXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQPEpSacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVWVLsbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDLnaugLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doSvGIB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMqPBXRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdKR1wRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLtBVu5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSPcpqp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hB5NEkIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oV2AzDEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tL2XFK1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaeBohJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2e7trMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6Im7lyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TycBkoXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mgx1fu3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdQOuEjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC1wz3hCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPmSvedwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1So7Dy4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPZZ6phxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiCpWkb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yr2oH9ugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiZq1vIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqEcdBDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHOaz0CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAowiI3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omN5bNMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSY2H4yEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iPMXuPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S26EiUgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOfc3udJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twybHMLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLJVMQiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JUMbNr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6R14Anb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICUWnQaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MuxJapAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BX5IgYeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndO966pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcnOezrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 047NEBdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCcbXKWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CT40wWo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6p1AmKK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFLG9sVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHmFD035Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHlLJNl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbp5uzx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNzdTDsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAbMj6RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dusP9qvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCpVTYOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNcraxDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlaCXxI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqJkDxQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Hm8QZeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIUkEovUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEgbYlTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKanjJ9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMMPdky3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2vZKhkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pZiaxWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nq7vsee7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xN5u0y0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3WxWbqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKSY9Jn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA4xSxtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSs4jgXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjE2XoL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voca3L7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbkCaQRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrTVHpeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxSmYdvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihH6oILxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kS6qkloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfWHMsWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvL2POiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4VdRFuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX7VZErgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrqfqAWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9dGnXrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsIDVHw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LghYpEfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhZi8a2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXVUU6TyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uv1oOu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUgEuTXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaOwbYRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDQIKYddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkHqwSSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSWAYMomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNcbfdgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdKjnLbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysUK2fy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKhls8bYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxnIVIUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dp3D3EUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUQP33MXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3CTqCjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5LxmDJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aCNZiFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTMD30EKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjJxYML8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtWw1gs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snCMwCWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tM8jIupPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deZ2GZuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdYjVm8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1dWZLwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUrT0ShPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36KcWIogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPowxdQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 739eTbaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyhvrMYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hns6UzahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUPAxFVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0mdztdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6E1K8KHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPI6APfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwWCWJtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdT8w5g8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17f8fDCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqE9QCxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HwdIwUfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXgEQpk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCe6Lb3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7XxB7uMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKvXGqmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o62KSe6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57MBDDJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhefz5IhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eymbPziuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDQfJDd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ds5JrxaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcT0XLVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtlndcgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owFpvF7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMLykAiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCFeNvjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLfto59LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaLLFwKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Inc6ba7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcOPjfhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJrDQJO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLwqrNveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiTKGpICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 408xjhV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJpobZYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIYthJSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpcQMa1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6ksWb5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgTk5YgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGsyEBmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9fqIvxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxzwFcpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JXPw86gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kb4KA5hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUuKyMzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQApc4enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6vbETr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewGZmWjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KrCpN4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhezzMshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcShVF3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnQlYODZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DerHU40hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDcp5gGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atHJk5l0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1xTx6uoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIFR4MxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X0e3cP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izojoET7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B27aM8cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lufcs1zfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5XfakJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMgdLVfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybVKm5qGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVLHpu0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKhDiWrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bh1x5fbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ym5E9yvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4Qa53G7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YrQlICOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0z6FumaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVaLhOdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUvxQHiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOdB9bHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNUjQz6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8okxKlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hxMtu9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqcUeYGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpaeON14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUKTGOuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5e3y4f5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OL1bmxQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwmJVMKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yx09yqfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1bB0ewIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtcF9iAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bclcCf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHiQ6acqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghEwy6zdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8Ts2luzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dZsdoFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lw5EJhFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0M19q2UYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RluIwQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9M9gDlFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o22Jpy5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysebVx2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZwLNNYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBAqYjM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pe0bwzsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vn6bDuBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLaZw2C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8df2adlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDcDem2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEKDypojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyd5oDbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjJEy4jOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pCMA966Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hjB6InmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyRoe98CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4Uma9WcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43zw5ZwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

