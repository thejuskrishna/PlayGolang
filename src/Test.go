package main
 
import (
    "net"
    "fmt"
)
 
func main() {

	//List all the IPs
    addrs, _ := net.InterfaceAddrs()
    fmt.Printf("IP [%v]", addrs)
    
}