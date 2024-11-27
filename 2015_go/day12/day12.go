package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var ErrRed error = errors.New("found red")

func rSum(i interface{}) (int, error) {
	switch v := i.(type) {

	case float64:

		return int(i.(float64)), nil

	case string:

		if i.(string) == "red" {
			return 0, ErrRed
		} else {
			return 0, nil
		}

	case map[string]interface{}:

		sum := 0
		for key, val := range i.(map[string]interface{}) {
			// look at values
			v, err := rSum(val)
			if err == ErrRed {
				// found a red -> ignore this object
				return 0, nil
			} else {
				sum += v
			}

			// look at keys
			v, _ = rSum(key)
			sum += v
		}
		return sum, nil

	case []interface{}:

		sum := 0
		for k := 0; k < len(i.([]interface{})); k++ {
			v, _ := rSum(i.([]interface{})[k])
			sum += v
		}
		return sum, nil

	default:

		log.Fatalf("Something else (%T): %v\n", v, v)
		return 0, nil

	}
}

func main() {
	fmt.Println("Day 12")

	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var d map[string]interface{}

	re := regexp.MustCompile(`(-{0,1}\d+)`)
	sum := 0
	sum2 := 0

	for scanner.Scan() {
		r := re.FindAllString(scanner.Text(), -1)
		for i := 0; i < len(r); i++ {
			n, err := strconv.Atoi(r[i])
			if err != nil {
				log.Fatal(err, r[i])
			}
			sum += n
		}

		err = json.Unmarshal(scanner.Bytes(), &d)
		if err != nil {
			log.Fatal(err)
		}

		sum2, _ = rSum(d)
	}

	fmt.Printf("sum: %v\n", sum)   // 111754
	fmt.Printf("sum2: %v\n", sum2) // 65402
}
