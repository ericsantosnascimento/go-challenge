package main

import (
	"flag"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

var maxExecutionTime = int64(500000000)

func main() {
	listenAddr := flag.String("http.addr", ":8080", "http listen address")
	flag.Parse()

	http.HandleFunc("/numbers", handleRequest)

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	urls, ok := r.URL.Query()["u"]

	result := make([]int, 0)

	//validating u parameter presence
	if !ok || len(urls[0]) < 1 {
		log.Println("not url informed")
		parseResult(w, result)
		return
	}

	//processing which url
	for i := 0; i < len(urls); i++ {

		data, e := callApi(urls[i])// I wish i could have done some make some parallel process here
		numbers := extractNumbersFromJson(e, data)
		result = append(result, numbers...)

	}

	result = removeDuplicated(result) //removing possible duplicated from array
	sort.Ints(result)                 //sorting the smaller subset

	//logging and checking if took to long to fetch the items
	duration := timed(start, "function handleRequest")
	if duration.Nanoseconds() > maxExecutionTime {
		parseResult(w, make([]int, 0))
		return
	}

	parseResult(w, result)
}

func extractNumbersFromJson(e error, data []byte) []int {
	if e == nil {
		var m Message
		err := json.Unmarshal(data, &m)
		if err == nil {
			return m.Numbers
		}
	}
	return make([]int, 0)
}

func parseResult(w http.ResponseWriter, result []int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"numbers": result})
}

func callApi(endpoint string) ([]byte, error) {

	response, err := http.Get(endpoint)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return data, nil
	}

}

func removeDuplicated(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func timed(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	return elapsed
}

type Message struct {
	Numbers []int
}


