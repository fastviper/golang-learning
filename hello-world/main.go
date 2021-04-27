package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"rsc.io/quote"
)

// Greetings of several types
func main() {
	// plain
	fmt.Println("Just print " + "something")

	// quote
	fmt.Println(quote.Hello())

	// rand number -- cheat sheet for formatting https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#printf
	rand.Seed(time.Now().Unix()) // seed the generator with current time
	fmt.Printf("The random number of this unique moment is %.3g\n", float64(rand.Intn(10))*rand.Float64()*math.E)

	// looped print in function
	fmt.Println(repeat(3, "hello\n"))

	// You were the choosen one!
	fmt.Println("Imperium strikes back:")
	fmt.Println(matrix_repeat(5, 3, "|-o-|"))
	fmt.Printf("          \t|-@@-|\n\n")
	fmt.Println(matrix_repeat(5, 3, "|-o-|"))

	// go also has multiple results returned, niice
	looking_for := "spark"
	host, port := zookeeper(looking_for)
	if port > 0 {
		fmt.Println(host, port)
	} else {
		fmt.Println("service not found")
	}

}

func repeat(n int, hello string) string {
	var ret string

	for i := 0; i < n; i++ {
		ret += hello
	}

	return ret
}

func matrix_repeat(x, y int, txt string) string {
	var ret string

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ret += txt + "\t"
		}
		ret += "\n"
	}

	return ret
}

func zookeeper(service string) (string, int) {
	var services = make(map[string]string)
	services["kafka"] = "kafka.dev.endor:8082"
	services["hadoop"] = "hdfs.dev.endor:9864"
	services["spark"] = "spark.dev.endor:9000"

	svc, found := services[service]
	if found {
		parsed := strings.Split(svc, ":")
		host := parsed[0]
		port, _ := strconv.ParseInt(parsed[1], 10, 16)
		return host, int(port)
	}

	return "", -1
}
