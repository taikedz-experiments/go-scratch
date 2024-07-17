/* Quick test with using anonymous functions
 *
 */
package main

import "fmt"
import "strings"


func main() {
    var mystrings = []string { "hi", "ciao", "salut" }
    for i,s := range filter(mystrings,
                            func(s string) bool { return strings.Contains(s, "a") }) {
        fmt.Printf("%d : %s\n", i, s)
    }
}


// We want to pass a function as a parameter
//   we must set up a type "alias" so to speak, to use as
//   the function signature identifier
// Is it better style to declare the type at the top of the file,
//   or better to do this and declare it where it is used?
type t_strfilter func(string) bool

func filter(strings []string, f_fn t_strfilter) []string {
    var res []string // Use a variable-length array instead

    for _,s := range strings {
        if f_fn(s) {
            res = append(res, s)
        }
    }

    return res
}
