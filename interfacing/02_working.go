package main

import (
    "fmt"
)

// Interfaces

type Settable interface {
    Set(string, string)
}

// Contrete type

type KvPair struct {
    key *string
    value *string
}

func (self KvPair) Set(k string, v string) {
    *self.key = k
    *self.value = v
}

// Do stuff

func main() {
    ik := "x"
    iv := "y"
    data := KvPair{&ik, &iv}
    Changeit(data, "ciao", "bye")
    fmt.Printf("%s -> \"%s\"\n", *data.key, *data.value)
    // original ik and iv have been changed
    //   as their pointers got dereferenced during assignment
    fmt.Printf("(%s)->(%s)\n", ik, iv)
}

func Changeit(item Settable, k string, v string) {
    item.Set(k,v)
}

