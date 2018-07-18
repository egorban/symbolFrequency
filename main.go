package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"log"
)

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
	if err !=nil{
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
	invertFrequency := invertMap(totalsFrequency)
	frequencies := make([]int, 0, len(invertFrequency))
	for frequency := range invertFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	for i := len(frequencies) - 1; i >= 0; i-- {
		for _, symb := range invertFrequency[frequencies[i]] {
			fmt.Println(string(symb), ":", frequencies[i])
		}
	}
}

func invertMap(frequencySymbol map[rune]int) map[int][]rune {
	invertFreqSymb := make(map[int][]rune, len(frequencySymbol))
	for key, value := range frequencySymbol {
		invertFreqSymb[value] = append(invertFreqSymb[value], key)
	}
	return invertFreqSymb
}


