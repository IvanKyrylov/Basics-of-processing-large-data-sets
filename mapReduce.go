package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
)

type Count struct {
	Key   string
	Count int
}

var tag int = -1

func MapReduce(path string, index int) <-chan []Count {

	tag = index

	// t := time.Now()
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	lines, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	lists := make(chan []Count)

	finalValue := make(chan []Count)

	var wg sync.WaitGroup

	wg.Add(len(lines) - 1)

	for _, line := range lines[1:] {
		go func(auto []string) {
			defer wg.Done()
			lists <- Map(auto)
		}(line)
	}

	go Reducer(lists, finalValue)

	wg.Wait()
	close(lists)
	// for _, v := range <-finalValue {
	// 	if v.key == "ЧОРНИЙ" {
	// 		fmt.Println(v)
	// 	}
	// }
	return finalValue
}

func Map(auto []string) []Count {
	list := []Count{}
	if tag == -1 {
		for _, v := range auto {
			temp := Count{
				Key:   v,
				Count: 1,
			}
			list = append(list, temp)
		}
	}
	temp := Count{
		Key:   auto[tag],
		Count: 1,
	}
	list = append(list, temp)
	return list
}

func Reducer(mapList chan []Count, sendFinalValue chan []Count) {
	mapTemp := make(map[string]int)
	for list := range mapList {
		for _, v := range list {
			mapTemp[v.Key] += v.Count
		}
	}

	final := []Count{}
	for k, v := range mapTemp {
		temp := Count{
			Key:   k,
			Count: v,
		}
		final = append(final, temp)
	}

	sendFinalValue <- final
}
