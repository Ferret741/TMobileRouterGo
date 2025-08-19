package router

import (
    "fmt"
    "os"
    "io"
    "encoding/json"
)


type Telemetry struct {
    Name       string
    Connected  bool
    Signal     int
    Mac        string
    Ipv4       string
}

// {
//   "clients": {
//     "2.4ghz": [
//       {
//         "connected": true,
//         "ipv4": "192.168.12.139",
//         "ipv6": [],
//         "mac": "00:1C:C2:66:80:5B",
//         "name": "wlan0",
//         "signal": -37
//       },
//
//       ...
//     "ethernet": [
//       {
//         "connected": true,
//         "ipv4": "192.168.12.237",
//         "ipv6": [
//           "fe80::5ef9:ddff:fe75:fafe",
//           "2607:fb90:8d18:8aae:5ef9:ddff:fe75:fafe"
//         ],
//         "mac": "5C:F9:DD:75:FA:FE",
//         "name": "optiplex"
//       }
//     ]
//   }
// }
type TelemetryJSONResponse struct {
    Clients struct{
        Two4Ghz []struct{
            Connected  bool      `json:"connected"`
            Ipv4       string    `json:"ipv4"`
            Ipv6       []string  `json:"ipv6"`
            Mac        string    `json:"mac"`
            Name       string    `json:"name"`
            Signal     int       `json:"signal"`
        } `json:"2.4ghz"`

        Five0Ghz []struct{
            Connected  bool      `json:"connected"`
            Ipv4       string    `json:"ipv4"`
            Ipv6       []string  `json:"ipv6"`
            Mac        string    `json:"mac"`
            Name       string    `json:"name"`
            Signal     int       `json:"signal"`
        } `json:"5.0ghz"`

        Ethernet []struct{
            Connected  bool      `json:"connected"`
            Ipv4       string    `json:"ipv4"`
            Ipv6       []string  `json:"ipv6"`
            Mac        string    `json:"mac"`
            Name       string    `json:"name"`
        } `json:"ethernet"`
    } `json:"clients"`
}


func (t *Telemetry) GetData(url, token string) (*TelemetryJSONResponse) {

    // Build the headers
    headers := make([]Header, 0)
    headers = append(headers, Header{Header_key:"Authorization",    Header_value: fmt.Sprintf("Bearer %s",token)})
    headers = append(headers, Header{Header_key:"Content-type",     Header_value: "application/json"})
    headers = append(headers, Header{Header_key:"Accept",           Header_value: "application/json"})
    headers = append(headers, Header{Header_key:"Connection",       Header_value: "Keep-Alive"})
    headers = append(headers, Header{Header_key:"Accept-Encoding",  Header_value: "gzip"})
    headers = append(headers, Header{Header_key:"User-Agent",       Header_value: REQUEST_USER_AGENT})

    // Create a router request object
    router_request := &RouterRequest{
        URL: url,
        Method: "GET",
        Headers: headers,
    }


    // Send the request
    router_response, err := router_request.Send()
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_HTTP_READ_BODY, err)
        os.Exit(1)
    }
    router_response_body, err := io.ReadAll(router_response.Body)
    defer router_response.Body.Close()

    // Mashal the data into JSON
    var telemetry_json_holder TelemetryJSONResponse
    err = json.Unmarshal(router_response_body, &telemetry_json_holder)
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_JSON_UNMARSHAL, err)
        os.Exit(1)
    }

    return &telemetry_json_holder
}
