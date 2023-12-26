package data

import "fmt"

var (
	Countries [10]string
	Slice     []int
	Codes     map[int]string
)

func init() {
	Countries[0] = "United States"
	Countries[1] = "Canada"
	Countries[2] = "Germany"
	Countries[3] = "France"
	Countries[4] = "United Kingdom"
	Countries[5] = "Japan"
	Countries[6] = "Italy"
	Countries[7] = "Australia"
	Countries[8] = "Brazil"
	Countries[9] = "China"

	names := []string{"John", "Doe"}
	names = append(names, "Carol")

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	qty := len(Countries)

	fmt.Println(Countries, Slice, Codes, names, wellKnownPorts, qty)
}
