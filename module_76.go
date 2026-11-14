package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.8" )

func Exo3qJRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy12nWubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tF07z7dfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLC7yQ7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8dB76RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMaAekGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6AcnSDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZWMivGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHcNeCYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePGnrTa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsukwX5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPViP3hLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xmx0lyCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi0KcGn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leEEAy6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCZf4hRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpVRxNUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GADINFtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQcrkxz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uz2fzn89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MmdkxqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuNy6nXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbmR2VKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpTO6tpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFzFo3eAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9ZwMbZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ramfSPmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QIJyOksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4R4iPh01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJIqheXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DldIy5G9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qkcsdu0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNSl9Xs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vQoKTfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKrVchdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPJQ3n70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmiJ6omwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAZkLwfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRUTvh1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3tvYZT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yd8SC4jrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIG1bwzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gx5ofKmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04JuK15jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Em38x7YYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnNz7zSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUP9qmSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VECqV5oNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7d4QekaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baFDFTq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3GY1AAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBgxHCMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJEFy6GOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrdXzpyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJeq3HAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfVJKioKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bqnNJSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQymx5qtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOR42EszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGezgQx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zPvN7QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPhoHNrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1N2Ga2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uChjRfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func had6SpxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmncy3asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLlxH0GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1LWU9wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2mEDvUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFV3GnRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEC8T2GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcObtPlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDbWBwerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbGP3fX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRFgiwQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e66BtQNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8lXJq6OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCTjt2B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lw8UTbnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0xSmT1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func To88OcUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Zr0Y3RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHkOp4U0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXNnCDQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FX99wt6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSig9yV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aAhgQOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iM5JLCt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEQ13guZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNYFKTtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7u5yiMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYWTuFybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUNupIDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vmv5UJIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8RVyCxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgLSiPIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24fPQHNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoqYnjUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFts3cOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kEKKekZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNWNMbwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBnD6BWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnKNC7kgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PlKapQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9SN5L7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fq2sUGsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1DlUJNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGMfimxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwY7O2ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFphxvBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pv8MAfQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lEFrMraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jp0g0Y9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEkJD8aiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpzarbtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bVpZEsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfeyvXcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTqfkKAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Doi69rx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x98iWSGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVuoLoXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXnWI3oYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meek0XZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmKcgNDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28x3R0iOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFwc2orBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jq2GDQIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekG2EJ1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifnQqey6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9m6j4TE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvdFGqqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1BSczVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcdwtfXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XGscDroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUGywj9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Si43q6w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBRECoZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8FzMNNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GghK5X3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx7ldWeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tib19h4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WK5DwJGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJWxlf4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyLmbPxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgr1RmVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FY7qvlzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJWfxmqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLcCIjL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABNBe738Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45k1FOx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3UgYZ3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BiGhtUSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiBlzAteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVxBR3fBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Qgnk2k5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjuA7gdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGcOYCIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2YWmAuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qg7Mgw6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK5EsjVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1WjjIZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bt0pugUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ASdjqzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaaQaR4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGeMS5DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpCF4XMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydSVeAVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TYckMxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehiVpCwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxlawL1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6NTQVp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32vxCZclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jbVij6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoHBSejKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uStxlBErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JY1GKbvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGoGY2bjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ho9pXzACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95I5NTilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cI16MfTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RysqOg5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZPLTiMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1WiJqEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSQYcUGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuH2CQJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJvOn2svWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWpJrzCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9URbXtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8Lc9aodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdiOLmwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goTFdsT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRKGV3GwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNLYP7JXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUmhA4GwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NpKaetNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqOaTGbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvlXDL8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnPyNvPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0iLHSUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6B5h6HumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p996C45BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waA5o56AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMBxdvoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZlCMQTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkTuicDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ymKZG8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7T1bvPqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8XRO3YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbt90u4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqfo30RjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGs2gmsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A79jSkESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JLGK0x3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDwMJjoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPk3o5QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Uc25idoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NKaFl34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0zq6ZmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHi1YNffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KigGx2u7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MS9YM3hRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnEzeTzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzZRiEcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pklBQr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfMyGf12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xifdTdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywW8sDusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVu60GSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shMtTaFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y90OIPypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3E0zaGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZhMpoCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJcVSsCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UoTaA2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9rb7dH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLR8cL23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hffm3MFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67O9j3OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g86hL5ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLqW6Fv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AbO1HEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whhVcpAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceg6q7BQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6mKTWl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rPWafskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FADE36ygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLjVsqdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 418ei0P2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsjR17YoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZWv2oxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ygq2u1UuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQDyJdE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DPKkpiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVghJuRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0qP8zZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cM9Ja2oDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2lb5JtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTD7SE75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFnNa1tNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31EhLizVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CxXx1FsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RkpsaZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQY9WPLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFABrAUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9ofhYehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmS89NjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuPmIP0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qyDAF9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2cfNcQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKPIkvQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Re9NcCb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcrO4yAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWfbzMR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxxCxmjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6EPH5CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1uzAzJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yxBBERGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uYrTTBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fIZmlURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7USKtS1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYzpNSDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hahD0zbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfQNHNiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAdG7E0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFGyuttyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gi9W0nqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODAlHmdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTDyMRwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME7FLAzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNKmzbWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb2ZcR6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1gpaw3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmX12Y0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBEiz8ZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwumaWNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqK1HehoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joq4igdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70HUda9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4gt0umgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91I6KZiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtYgVEwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ou6wWjUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpb2jsBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5LArgztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkxOVUMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiKt5kpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xk3gOdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrSyY7uQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIqRdonvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezTTJBekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpTuy3wTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4WMKhdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bTygGKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8uzvFQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyPcLuZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZfuI2huWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QH9XlvB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aM4P57oJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JMchW9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzKpWOPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gtazlce5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoaVleYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pHVJVtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l0WjI7T1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghfT0a0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQCi0xSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioo3RR11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GD8Bmi1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWsb1R9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8l9XO6REWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcwMp1M3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRbVqKxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HW3lLFBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mj4j8hjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm7AFEx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4aEDjgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5WNTNrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cj7z6qoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJnH0jfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uV4uDV8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15xDlvSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g474x5fQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8OdAD3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzkpYDeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cYxN7RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPPR17csWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYRkqcMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jjD4tFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtCm9ewkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4i50SFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LvK3wEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEXEPXevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KazeuHJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TMHt0zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBAwF3HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPYGJyDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLHE3HjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZj8BdboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRBdgQTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUdUfj38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uE0rD91jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEqU4eVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXTn6YouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Inb9fEjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1itxaC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaYgwJm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZqhfAiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSJicnEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gU3UHGWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnthEaIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1ZneASzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioHQ02lWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIIH4Jw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZEIA6ioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgWRfyaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p15Qlvz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMXzpR4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRx73wRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKIiMKtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRRpTwGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvb0MistWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40NlmbRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5BwrVxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bglpKWF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9IFQW7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tQEOsz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSXcAkJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GV1d8QGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8kW5EdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bidtUZbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4uJRCsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQAdAUpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7EETwDShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbUFjoSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MANUAnkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtB38gJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tu39soyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1Zy2eZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjXMuvnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPeeVaFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzTtZ5AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MILMOLZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVBQxDcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDqfMFNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqZDr2P8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EH5sQOk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3lifc9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0q17l3F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8U0Goh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujCnE0FBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xq5UOiH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPcCc3reWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAi8F6ugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byq9XTQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMDquuD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSmNtb8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfN9MGztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MroE4JRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD0YoxQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqb8WuhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTyCsxnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSeqWLEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bv2wHu4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czpx56wKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnTMufxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sT8ZnOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bm942ASDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LOqqnMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTDLLfCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BUublCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrxtFCQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kggvauuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROS3psldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxFTVh59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FT9b50oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRhRPceaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50nx5eErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sR8HRUBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v36ArPLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ci7ejHjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLQ5iMe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQrwZF6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mSyNGFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq2qRl6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXmwuIPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwMvEceUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDRWMS56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkw5feNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zsi5uSvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE1HNX38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsmIW9qzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MixT0JBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lyTmRvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URea5SfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYbublblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCUFkPw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diesjkilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIqJMEdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y14egiMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func al8Fjfx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wP7CKaD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1QN6siTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLtoebXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTrX7LkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rze4IFLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8SxhfjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9dO4R7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mglWSYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ws6QFcLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fm2PS2GfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NEpuEWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72O6OV3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XM27uZogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLdabqJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46pQK66PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQfHtp2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LI8EsLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USYT015sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWPNPB7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdahFJ8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7FdYIClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ud69m0YSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNiQj0yEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjQtDbNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFNdlJJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRATg8F0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4g8IFsaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JL3zPjmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwJzmWznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LtL0PRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtif7b59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjFvucJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D97s0NF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l044k9zNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMevuiuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4e1a6LPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MayBvSp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLY1pkErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5HXcW3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMCkPeQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRKCN4RzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbgCz3n4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZE0EPnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BX2ixvtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDitsfOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6qGwrv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkJCnJTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9shk5yCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7kvMX5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbtPXu4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFJxAkvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyR6gyVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qHvq1SMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRv9M7gFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ye0nIuG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csNUqeivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18OrjLRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skMuNYElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqqCqO76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rx4TZl0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgdcupRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwioqTb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxLOUpUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEqhKn8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1GUw4zUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jM5XyW5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPP477a6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2vwKLI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljun1TtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4Ad7tNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq6Nr7R4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4E8byjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9A4UXT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKRZFNjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLZqIihrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbvIV4mUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRoo8AYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eP2nV5MfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXtt2jAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a66rqHZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Nmx68TWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8MfQtJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3o4SAkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unBDbjHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEjlkNm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2s7RzXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqsm9znqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6YtDV8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGoJAo9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHOwoWtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrymoumXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmqIpIcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpGFzzmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTJ11Se6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5fPKwlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJihLLJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zT5MirgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkoHNdPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func babVrlckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAySwDMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2jZ0EudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aA1762B7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OI769BOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pR71zD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbEjOe7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OA47swuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwvJEiuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Vo4fbSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlTMT8rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbhgzEH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LBnGekyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uO8mjQz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NycmywwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rx9gFGTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzvP4kV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YK4nUdJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVjyJMneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZIIUSp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ma9TerXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdwTzYMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaY7OUPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1A3LGkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nomoPLdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkaMucY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPOEdf5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFO5zyl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9J3UFLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZPivLVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEbzQNMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yThUIByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiQvZGdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVwVqH2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3UBs29DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsZ6b8xsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZW9Dkl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UDZnvafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRFsLhJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcshoDeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5uuV5G7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24wFwglhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4urvH4fSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoRrsrB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MS0S6fbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wp4zOUBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTcnbtOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avwyTslGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ut7kSsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLtaBperWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRGOW5miWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QU6jHGesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C53KJ6pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aNgIzqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TW0A00h3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drxsS4ivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ajjbr00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1IvVnUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PNETzFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XaTHtg0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxb4MZhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnBmNYWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghlGHdqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sW0k8zFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGlsEuJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLuKQS2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URxEuZeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teEXEZfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0R24RX1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQSoQVHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNys3fgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZnROSGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cn6AJbnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9H5oruZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69aZDnHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUZl3c2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPY6GPpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwdhnpMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6U5yoTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKS6kPl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gthmbMUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgCX0PhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaMUWzHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VF1w6SaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hw3hQjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtywYxiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKtwvyNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OuWA6UwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tna325QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSWagLI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4Z0hCWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ah0tzQDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTetpNjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Xa5Tl6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itQ2txFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJ7yruu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9svZhp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lo5ugNpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FHDsSS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FODuY8OFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GotLu5eVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsjduGysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cd2V3hSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CM6TYWeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPgxYhmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7xItx9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R42A51vxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYpWNmQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXdhUKUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fruceyWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5QaGS9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsVwyH0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duUqSi1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0p4UY6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YawJwOZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJDw29lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A13JhlFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48cUF7JKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRlGpuLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsndKIcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbav6bHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZlvNyqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EtCOJSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpnp0U8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJoM4zSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6Lq0Q5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4y3nrxRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2egJuJT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HV75EB7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxljVBNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9tgH5ImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fk53xTeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nf4kLd5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktpdzCUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvAhdNBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7XQqD6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNuQAAv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNVMAJPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXKRpCfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdNNqnA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NE1d4LrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVRx6Mv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfRnWYAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQ4LEoFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONLAC2NyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXbO6hdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHkA4WmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S27ttTFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W03msT6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcztjAVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31wS23F5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NARR5HhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ks1ssox0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUbqqgTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVryhyvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DivtkAK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzlIIa03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqtcQZO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hVMVVYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WxdEk8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cr6uVAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qR626w1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEtXO9w3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEW1idx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YL5EUpVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X65r8BjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIHE2pkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsvaMQ4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojwj1OnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOJjM5zOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8V6g4SrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvTEUyCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwXRUzCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qONAolT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDXm7Bc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akVpDuG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bqQCwV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KoPjEa3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z8wPIyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y21ml7s7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jCnUHV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jq17gW0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezdS2nhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfIncDytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7oceQlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJidOK0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T71JUHMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ORSHs60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qayc4YHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUVKZ8xBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkA01GG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7nk3awKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tom6YLhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrUF8xnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkRQgZLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqHYZ36wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocj4UVUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TO7If8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHhmUmi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZ8eIFJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQXEMJ9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtHHrMKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lJ60KolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0r3Hd1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vQmx7BnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOoxTqNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Mbvrsh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eb23hLxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeJCW7L7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6eHeYElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LN6wnswiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmjzyTK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8qf9nIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlgnQ3kdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvjwRy9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6k4hjqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMxmHjFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9XScnkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hX0N7gfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPhjZTNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUHeqy51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WB6EetHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LZKXhHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XernE1WjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9GsdmK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQP9sf0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQz3VI0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pC5RwRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ig5VAHOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pzKGwK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKLlK5CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Clx0FSD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5fEknO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psYMZr7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWYEEmjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBlirdbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WORWnlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMy26TOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rw7JZyhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ex7csXMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LxadexJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hK6416dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8ExI9E9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3P7B3NMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NL4WqLRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc6WrR9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1BILrqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbHEqRh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9EDfYV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhFXMfKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPG2aQeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjH2LKFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fmica1ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw9KW7ZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OqmH2iwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k9RgztQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Koe2FDoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt79eTP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXuuKwMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8GlICfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIZ0Ji1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRBhg3MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGm6Yit7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxUidHdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJIHwtcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEVWYiYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3UgI6kVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pq9Qfx5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLueqQE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ublhLUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43AAwKb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdjWhVyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiuZ1itpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqknKZMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8h9lg1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHyP8WXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJGL4CnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xrZ1BFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QjY5QD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHSkUDeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7n07mJEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwNzCv8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEWeDdhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQXdtEBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIjK1PpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUSQ2b68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWerL4IgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8BiKziPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAH2SvBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8rSB42bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ShgeL1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNwBUMRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsyURUnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj5P5p56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7O2Oeu7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lD48bceuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3obXNhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idSjfAh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gloNJVGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbdGh3JbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9x7PbZslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCAriKPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aO8kdvCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xt6UhomMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKoEaEcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2r0fBvw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOKbku94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uONI4rPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaTXvvjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZTDsQJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIE6Z5TcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GmsOJqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28EcpgQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fahEd7xVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHuBXHNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zuZz3G1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMJbelKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8MJJTaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYWmqNVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuF77gMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZ1fLf66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUK1SjtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmkGxHQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHd61IVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8rhrbnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nMvg12wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrECKBJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IMIIwcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ot7WqWJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNMtJIAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEiZnL0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysr7bh5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3KCIiFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWQDsaZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNVeNn4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Lb0X34pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHcvZkgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xF1lPsUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSmtLJkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBRnzdBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdO3g9HAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acBiuJrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fW5PxWKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tqg76MdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQFIO0C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ei72SwVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nk0SUYDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4MKkiSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rasTBkl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y29ERumDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iv0oDdQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpRuWvB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dS7juOWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dggdA39aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFubI7kpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SufjPNCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcmWLWQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ww7L6PAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qL8pcTB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

