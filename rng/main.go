package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(23)
    rng := rand.Int63()
    fmt.Printf("%d\n", rng)
}
