package router

import (
    "fmt"
    "net/http"
    "bytes"
    "os"
)


type Header struct{
    Header_key   string
    Header_value string
}

type RouterRequest struct {
    // The URL string
    URL string

    // The Post body
    PostBody []byte

    // The HTTP method
    Method string

    // A list of structs representing headers
    Headers []Header
}


func (r *RouterRequest) Send() (*http.Response, error) {

    // So that we can handle various request methods later
    // on in this method
    var http_request *http.Request
    var err error

    // Generate a new HTTP client
    http_client := http.Client{Timeout: REQUEST_TIMEOUT}

    // Depending on the request method, perform certain
    // functions
    switch r.Method {

        // POST method
        case "POST":
            http_request, err = http.NewRequest(r.Method, r.URL, bytes.NewReader(r.PostBody))

        case "GET":
            http_request, err = http.NewRequest(r.Method, r.URL, nil)

        default:
    }

    // Collect the error here since its nly generated in the
    // switch statemetn
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_HTTP_REQUEST_GENERATE, err)
        os.Exit(1)
    }

    // Iterate over the headers and add them to the request
    for _, val := range r.Headers {
        http_request.Header.Add(val.Header_key, val.Header_value)
    }

    // Send off the request and return directly the results
    return http_client.Do(http_request)

}
