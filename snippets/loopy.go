package main
import "fmt"

func main() {
    i := 1

    marker:

    fmt.Printf("-> %d\n", i)
    i++

    if i <= 2 {
        goto marker
    } else {
        goto endmarker
    }

    endmarker:
    fmt.Println("Done.")

    var g1 string
    var g2 string
    var g3 string
    fmt.Sscan("Hello hi good day", &g1, &g2, &g3)
    fmt.Printf("Became: %s | %s | %s\n", g1, g2, g3)
}
