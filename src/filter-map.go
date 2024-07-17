package main
import "fmt"

type t_strmap map[string]string
type tf_strconv func(string) string

func main() {
    var data = []string {"alpha", "beta"}
    var stuff = fmap(data, func(s string) string {return string(s[0:1])})
    fmt.Printf("%s\n", stuff)
}

func fmap(elements []string, getkey tf_strconv) t_strmap {
    var new_map = make(t_strmap, len(elements))

    for _,elem := range elements {
        new_map[getkey(elem)] = elem
    }

    return new_map
}
