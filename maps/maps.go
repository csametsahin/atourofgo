package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

func main() {
	// m = make(map[string]Vertex)
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	m["Istanbul"] = Vertex{41.0082, 28.9784}
	fmt.Println(m["Istanbul"])
	delete(m, "Istanbul")
	v, ok := m["Istanbul"]
	fmt.Println("The value ", v, "Is it present", ok)
}
