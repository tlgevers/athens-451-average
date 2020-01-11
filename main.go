package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func readData() (data []byte, err error) {
    read, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        return
    }
    data = read
    return
}

func createArgs(d []byte) (data []string, err error) {
    args := strings.Fields(string(d))
    data = args
    fmt.Println(len(args))
    return
}

func convInt(a []string) (v []int32, err error) {
    var data []int32
    for _, num := range a {
        d, _ := strconv.Atoi(num)
        data = append(data, int32(d))
        v = data
    }
    return
}

func sumArgs(args []int32) (sum int32, err error) {
    sum = 0
    for _, num := range args {
        sum += num
    }
    return
}

func main() {
    fmt.Println("World Starter")
    data, err := readData()
    if err != nil {
        panic(err)
    }
    values, err := createArgs(data)
    if err != nil {
        panic(err)
    }
    args, err := convInt(values)
    if err != nil {
        panic(err)
    }
    sum, err := sumArgs(args)
    if err != nil {
        panic(err)
    }
    fmt.Printf("The sum is: %d", sum)
}
