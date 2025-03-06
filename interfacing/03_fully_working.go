package main

import (
    "fmt"
)

// Interfaces

type Settable interface {
    Set(string, string)
}

type Representable interface {
    Repr() string
}

// Contrete type

type KvPair struct {
    key string
    value string
}

func (self *KvPair) Set(k string, v string) {
    self.key = k
    self.value = v
}

func (self KvPair) Repr() string {
    return fmt.Sprintf("%s -> \"%s\"", self.key, self.value)
}

// Do stuff

func main() {
    data := KvPair{"hello", "hi"}

    // Settable cannot receive a direct KvPair if KvPair implements
    //  using a pointer reference to self
    //  but it _can_ receive a pointer to a KvPair
    Changeit(&data, "ciao", "bye")

    // Representable can receive a plain KvPair
    Sayit(data)

    // Demonstration - we pass in a &KvPair , but Sayit can still
    //  access methods without explicitly dereferencing
    Sayit(&data)
}

func Changeit(item Settable, k string, v string) {
    item.Set(k,v)
}

func Sayit(data Representable) {
    fmt.Printf("%s\n", data.Repr())
}
