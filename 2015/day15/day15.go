package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ingredient_t struct {
	label      string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type cookie_t struct {
	ingredients []ingredient_t
	quantities  []int
}

func (cookie *cookie_t) Calories() int {
	calories := 0
	for i := range cookie.ingredients {
		calories += cookie.ingredients[i].calories * cookie.quantities[i]
	}
	return calories
}

func (cookie *cookie_t) Score() int {

	capacity_score := 0
	durability_score := 0
	flavor_score := 0
	texture_score := 0

	for i := range cookie.ingredients {
		capacity_score += cookie.ingredients[i].capacity * cookie.quantities[i]
		durability_score += cookie.ingredients[i].durability * cookie.quantities[i]
		flavor_score += cookie.ingredients[i].flavor * cookie.quantities[i]
		texture_score += cookie.ingredients[i].texture * cookie.quantities[i]
	}

	if capacity_score <= 0 || durability_score <= 0 || flavor_score <= 0 || texture_score <= 0 {
		return 0
	}

	return capacity_score * durability_score * flavor_score * texture_score
}

func ParseDay15Input(s string) []ingredient_t {
	re := regexp.MustCompile(`(\w+): capacity (\-?\d+), durability (\-?\d+), flavor (\-?\d+), texture (\-?\d+), calories (\-?\d+)`)

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ingredients []ingredient_t

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		var ingredient ingredient_t
		ingredient.label = r[1]
		ingredient.capacity, _ = strconv.Atoi(r[2])
		ingredient.durability, _ = strconv.Atoi(r[3])
		ingredient.flavor, _ = strconv.Atoi(r[4])
		ingredient.texture, _ = strconv.Atoi(r[5])
		ingredient.calories, _ = strconv.Atoi(r[6])
		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}

func SearchAllCookies(ingredients []ingredient_t) (cookie_t, cookie_t) {
	var best_cookie, best_cookie_500 cookie_t
	best_cookie_score := 0
	best_cookie_500_score := 0

	for i := 0; i < 100; i++ {
		for ii := 0; ii < 100-i; ii++ {
			for iii := 0; iii < 100-i-ii; iii++ {
				iv := 100 - i - ii - iii
				cookie := cookie_t{
					ingredients: make([]ingredient_t, len(ingredients)),
					quantities:  make([]int, len(ingredients)),
				}
				copy(cookie.ingredients, ingredients)
				cookie.quantities[0] = i
				cookie.quantities[1] = ii
				cookie.quantities[2] = iii
				cookie.quantities[3] = iv

				s := cookie.Score()
				c := cookie.Calories()

				if s > best_cookie_score {
					best_cookie_score = s
					best_cookie = cookie
				}

				if c == 500 && s > best_cookie_500_score {
					best_cookie_500_score = s
					best_cookie_500 = cookie
				}
			}
		}
	}
	return best_cookie, best_cookie_500
}

func main() {
	fmt.Println("Day15")
	best_cookie, best_cookie_500 := SearchAllCookies(ParseDay15Input(`input\input15.txt`))
	part1 := best_cookie.Score()
	part2 := best_cookie_500.Score()
	fmt.Printf("part1: %v\n", part1) // 18965440
	fmt.Printf("part2: %v\n", part2) // 15862900
}
