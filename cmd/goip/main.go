package main

import "fmt"
import doip "github.com/pstvm/goip/pkg/doip"

func main() {
	c := doip.DoIPClient {
		ECUAddress: "a",
	}
	fmt.Println(c.GetECUAddress())
	c.ConnectToECU()
}
