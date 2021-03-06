/*
локальные отображения, каналы
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)


type Frequencies struct {
	Key   rune
	Value int
}

func main() {

	lenArgs := len(os.Args)
	if lenArgs == 1 {
		log.Println("No arguments")
		return
	}

	countFiles := lenArgs - 1
	result := make(chan map[rune]int, countFiles)

	for _, filename := range os.Args[1:] {
		go countFrequencies(filename, result)
	}

	totalsFrequency := make(map[rune]int)
	merge(result, totalsFrequency, countFiles)
	showResult(totalsFrequency)
}

func countFrequencies(filename string, result chan map[rune]int) {
	frequencySymbols := make(map[rune]int)
	symbols, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("failed to open the file:", err)
	}
	for _, s := range string(symbols) {
		frequencySymbols[s]++
	}
	result <- frequencySymbols
}

func merge(result chan map[rune]int, totalsFrequency map[rune]int, countFiles int) {
	for i := 0; i < countFiles; i++ {
		frequencyFile := <-result
		for s, f := range frequencyFile {
			totalsFrequency[s] += f
		}
	}
}

func showResult(totalsFrequency map[rune]int) {
	var totalsSort []Frequencies
	for s, f := range totalsFrequency {
		totalsSort = append(totalsSort, Frequencies{s, f})
	}
	sort.Slice(totalsSort, func(i, j int) bool { return totalsSort[i].Value > totalsSort[j].Value })
	for _, Elem := range totalsSort {
		fmt.Printf("%q : %v\n", Elem.Key, Elem.Value)
	}
}
