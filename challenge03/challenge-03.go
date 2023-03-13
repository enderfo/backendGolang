package main

import (
    "fmt"
    "sort"
    "unicode"
)

func main() {
	kalimat := "selamat malam"
    pecahKatas := []rune(kalimat)
    for i := 0; i < len(pecahKatas); i++ {
        perKata := string(pecahKatas[i])
        fmt.Println(perKata)
    }

    m := map[rune]int{}
    for _, char := range kalimat {
        if unicode.IsSpace(char) {
            continue
        }
        m[char] += 1
    }
    chars := []rune{}
    for char := range m {
        chars = append(chars, char)
    }
    sort.Slice(chars, func(i, j int) bool {
        return m[chars[i]] < m[chars[j]]
    })
	fmt.Print("map[ ")
    for _, char := range chars {
        fmt.Printf(`%c:%d `, char, m[char])
    }
	fmt.Println("]")
}
