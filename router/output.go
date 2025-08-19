package router

import (
    "fmt"
    "strings"
)

func IterateDevices(telemetry * TelemetryJSONResponse, output, deviceName string) {
    if output == "table" {
        fmt.Println()
    }
    iterate2point4GClient(telemetry, output, deviceName)
    iterate5naughtGClient(telemetry, output, deviceName)
    iterateEthernetClient(telemetry, output, deviceName)
}

func iterate2point4GClient(telemetry * TelemetryJSONResponse, output, deviceName string) {
    networkType     := "2.4Ghz"
    stringSpecifier := "%-10s %-21s %-15s %-5t %5d %-50s\n"
    for _, content := range telemetry.Clients.Two4Ghz {
        switch output {
            case "table":
                fmt.Printf(stringSpecifier, networkType,
                                            content.Mac,
                                            content.Ipv4,
                                            content.Connected,
                                            content.Signal,
                                            content.Name)
            case "ip-address":
                if strings.Trim(content.Name, " ") == deviceName {
                    fmt.Print(content.Ipv4)
                    break
                }
        }
    }
}

func iterate5naughtGClient(telemetry * TelemetryJSONResponse, output, deviceName string) {
    networkType     := "5.0Ghz"
    stringSpecifier := "%-10s %-21s %-15s %-5t %5d %-50s\n"
    for _, content := range telemetry.Clients.Five0Ghz {
        switch output {
            case "table":
                fmt.Printf(stringSpecifier, networkType,
                                            content.Mac,
                                            content.Ipv4,
                                            content.Connected,
                                            content.Signal,
                                            content.Name)
            case "ip-address":
                if strings.Trim(content.Name, " ") == deviceName {
                    fmt.Print(content.Ipv4)
                    break
                }
        }
    }
}

func iterateEthernetClient(telemetry * TelemetryJSONResponse, output, deviceName string) {
    networkType     := "Ethernet"
    stringSpecifier := "%-10s %-21s %-15s %-5t %5d %-50s\n"
    for _, content := range telemetry.Clients.Ethernet {
        switch output {
            case "table":
                fmt.Printf(stringSpecifier, networkType,
                                            content.Mac,
                                            content.Ipv4,
                                            content.Connected,
                                            0,
                                            content.Name)
            case "ip-address":
                if strings.Trim(content.Name, " ") == deviceName {
                    fmt.Print(content.Ipv4)
                    break
                }
        }
    }
}


func printHeader(){
}

func printFooter(){
}

// func FindDevice(telemetry *TelemetryJSONResponse) ([]struct) {
// }
//
// func GetIPAddress(deviceList []struct) ([]string) {
// }

func out() {
    fmt.Println()
}



// &{
//     Clients:{
//         Two4Ghz:[
//             {Connected:true Ipv4:192.168.12.139 Ipv6:[] Mac:00:1C:C2:66:80:5B Name:wlan0 Signal:-44}
//             {Connected:true Ipv4:192.168.12.143 Ipv6:[] Mac:60:74:F4:60:E5:FE Name: Signal:-48}
//             {Connected:true Ipv4:192.168.12.199 Ipv6:[] Mac:60:74:F4:5E:05:60 Name: Signal:-42}
//             {Connected:true Ipv4:192.168.12.249 Ipv6:[] Mac:D4:AD:FC:FD:42:8F Name:Shenzhen Intellirocks Tech co.,ltd Signal:-46}
//             {Connected:true Ipv4:192.168.12.245 Ipv6:[] Mac:D4:AD:FC:F1:7D:08 Name:Shenzhen Intellirocks Tech co.,ltd Signal:-50}
//             {Connected:true Ipv4:192.168.12.109 Ipv6:[] Mac:60:74:F4:5D:F6:98 Name: Signal:-44}
//             {Connected:true Ipv4:192.168.12.227 Ipv6:[] Mac:D4:AD:FC:F1:F3:3A Name:Shenzhen Intellirocks Tech co.,ltd Signal:-40}
//         ]
//
//         Five0Ghz:[
//             {Connected:true Ipv4:192.168.12.247 Ipv6:[fe80::36c6:7911:fca9:b7c] Mac:14:C1:4E:88:8C:78 Name:Dormitorio Signal:-57}
//             {Connected:true Ipv4:192.168.12.208 Ipv6:[fe80::18ee:c8ff:fe0d:49b 2607:fb90:8d18:8aae:18ee:c8ff:fe0d:49b] Mac:1A:EE:C8:0D:04:9B Name:Keesha Signal:-54}
//             {Connected:true Ipv4:192.168.12.138 Ipv6:[fe80::48d6:40ff:fef7:f447] Mac:4A:D6:40:F7:F4:47 Name:Gadget-jRome741 Signal:-57}
//             {Connected:true Ipv4:192.168.12.177 Ipv6:[] Mac:82:1F:AC:C4:6B:53 Name:Bianca-jRome741 Signal:-58}
//             {Connected:true Ipv4:192.168.12.108 Ipv6:[] Mac:44:BB:3B:48:D4:9A Name:09AF01AJ49230832 Signal:-55}
//             {Connected:true Ipv4:192.168.12.206 Ipv6:[fe80::4816:18ff:fefb:38a1 2607:fb90:8d18:8aae:cd30:cc4a:ca00:69fb] Mac:4A:16:18:FB:38:A1 Name:Hurones-II Signal:-44}
//             {Connected:false Ipv4:192.168.12.175 Ipv6:[] Mac:44:EA:30:D9:26:B8 Name:Samsung Electronics Co.,Ltd Signal:-59}
//         ]
//
//         Ethernet:[
//             {Connected:true Ipv4:192.168.12.237 Ipv6:[fe80::5ef9:ddff:fe75:fafe 2607:fb90:8d18:8aae:5ef9:ddff:fe75:fafe] Mac:5C:F9:DD:75:FA:FE Name:optiplex}
//         ]
//     }
// }
