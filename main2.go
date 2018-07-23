/*
общее отображение, без каналов
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
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

	totalsFrequency := make(map[rune]int)
	var wait sync.WaitGroup
	var m sync.Mutex

	for _, filename := range os.Args[1:] {
		wait.Add(1)
		go countFrequencies(filename, totalsFrequency, &wait, &m)
	}
	wait.Wait()
	showResult(totalsFrequency)
}

func countFrequencies(filename string, totalsFrequency map[rune]int, wait *sync.WaitGroup, m *sync.Mutex) {
	symbols, err := ioutil.ReadFile(filename)
	m.Lock()
	if err != nil {
		log.Fatal("failed to open the file:", err)
	}
	for _, s := range string(symbols) {
		totalsFrequency[s]++
	}
	m.Unlock()
	wait.Done()
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
