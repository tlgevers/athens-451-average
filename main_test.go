package main

import (
    "testing"
    "fmt"
)

func TestReadData(t *testing.T) {
    read, err := readData()
    if err != nil {
        t.Error(err)
    }
    var i interface{} = read
    f, ok := i.([]byte)
    fmt.Println(string(f), ok)
}
