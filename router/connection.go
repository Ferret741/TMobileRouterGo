package router

import (
    "net"
    "fmt"
    "os"
)


func TestConnection(router_ip_address, router_port string) {

    // Construct the full address string
    router_host_string := fmt.Sprintf("%s:%s",router_ip_address,router_port)

    // Make the connection
    c, err := net.Dial("tcp4", router_host_string)
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_CONN_TEST_FAILED, router_host_string, err)
        os.Exit(5)

    }
    c.Close()
}
