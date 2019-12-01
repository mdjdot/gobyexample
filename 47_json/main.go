package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Response1 .
type Response1 struct {
	Page   int
	Fruits []string
}

// Response2 .
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB)) // true

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB)
	fmt.Println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)
	fmt.Println(string(strB)) // "gopher"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	mapD := map[string]int{"apple": 2, "peach": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB)
	fmt.Println(string(mapB)) // {"apple":2,"peach":7}

	res1D := Response1{Page: 2, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)
	fmt.Println(string(res1B)) // {"Page":2,"Fruits":["apple","peach","pear"]}

	res2D := &Response2{
		Page:   3,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)
	fmt.Println(string(res2B)) // {"page":3,"fruits":["apple","peach","pear"]}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil { // map[num:6.13 strs:[a b]]
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	num := dat["num"].(float64)
	fmt.Println(num) // 6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1) // a

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           // {1 [apple peach]}
	fmt.Println(res.Fruits[0]) // apple

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "peach": 7}
	enc.Encode(d) // {"apple":5,"peach":7}

}
