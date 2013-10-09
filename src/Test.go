package main
 
import (
    "net"
    "fmt"
)
 
func main() {
    addrs, _ := net.InterfaceAddrs()
    fmt.Printf("IP [%v]", addrs)
}