package main

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {

	// p := Person{
	// 	Name: "sujon",
	// 	Age:  20,
	// 	City: "sylhet",
	// }

	// rawjson, err := json.Marshal(p)

	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println(string(rawjson))

	CustomJson()

}
