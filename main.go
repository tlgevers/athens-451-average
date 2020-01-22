package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "os"
    "errors"
)


func readData() (data []byte, err error) {
    read, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        return
    }
    data = read
    return
}

func createArgs(d []byte) (args []string, err error) {
    args = strings.Fields(string(d))
    if len(args) == 0 {
        err = errors.New("NO DATA WAS FOUND")
    } 
    return
}

func convInt(a []string) (v []int, err error) {
    var data []int
    for _, num := range a {
        d, _ := strconv.Atoi(num)
        data = append(data, int(d))
        v = data
    }
    return
}
 
func sumArgs(args []int) (sum int, err error) {
    sum = 0
    for _, num := range args {
        var n interface{} = num
        _, ok := n.(int)
        if ok != true  {
            err = errors.New("CANNOT SUM NON INTEGERS") 
            return
        }
        sum += num
    }
    return
}

func average(args []int) (avg float32, err error) {
    sum, err := sumArgs(args) 
    if err != nil {
        panic(err)
    }
    val := float32(sum)
    avg = (val / float32(len(args)))
    return
}

func writeData(avg string) {
    f, err := os.Create("./output.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    f.WriteString("The average is: " + avg)
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
    avg, err := average(args)
    if err != nil {
        panic(err)
    }
    roundedAvg := fmt.Sprintf("%.2f", avg)
    writeData(roundedAvg)
}
