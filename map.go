package main

import "fmt"
import "strings"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
	var ret = [] string{} 
	for _, value := range p {
		ret = append(ret, fmt.Sprintf("%d", value))
	}
	return fmt.Sprintf("%s", strings.Join(ret, ","))
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	var mn = map[string]Vertex{
    	"Bell Labs": Vertex{
    		40.68433, -74.39967,
    	},
    	"Google": Vertex{
    		37.42202, -122.08408,
    	},
    }
    fmt.Println(mn)
    
    hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
    
    
}