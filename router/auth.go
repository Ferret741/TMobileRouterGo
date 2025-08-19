package router

import (
    "encoding/json"
    "io"
    "os"
    "fmt"
    "strings"
    "strconv"
    "time"
)


// The "class" for the Authentication struct. For the time being it
// only contains the authentication token and token expiration time
type Auth struct {
    AuthToken           string
    AuthTokenExpiration int
}


// Send an HTTP request to the route to obtain a new authentication token
func (a *Auth) RequestNewToken(url, username, password string) {

    // {
    //   "auth": {
    //     "expiration": 1743032529,
    //     "refreshCountLeft": 4,
    //     "refreshCountMax": 4,
    //     "token": "<long_string>"
    //   }
    // }
    // Define the JSON response template
    type AuthResponse struct {
        Auth struct {
            Expiration       int    `json:"expiration"`
            RefreshCountLeft int    `json:"refreshCountLeft"`
            RefreshCountMax  int    `json:"refreshCountMax"`
            Token            string `json:"token"`
        } `json:"auth"`
    }


    // Define the json template for auth request
    type authRequestTemplate struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // Populate a JSON representation
    authRequestJson := authRequestTemplate{
        Username: username,
        Password: password,
    }

    // Marshal JSON representation -> json
    jsonbytes, err := json.Marshal(authRequestJson)

    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_JSON_MARSHAL, err)
        os.Exit(1)
    }

    // Define request headers
    headers := make([]Header,0)
    headers = append(headers, Header{Header_key: "Content-Type",    Header_value: "application/json; charset=UTF-8"})
    headers = append(headers, Header{Header_key: "Accept",          Header_value: "application/json"})
    headers = append(headers, Header{Header_key: "Accept-Encoding", Header_value: "gzip"})
    headers = append(headers, Header{Header_key: "Connection",      Header_value: "Keep-Alive"})
    headers = append(headers, Header{Header_key: "User-Agent",      Header_value: REQUEST_USER_AGENT})

    // Define the router request object
    router_req := &RouterRequest{
        URL:        url,
        Method:     "POST",
        PostBody:   jsonbytes,
        Headers:    headers,
    }

    // Dispatch the POST request and receive the request
    auth_resp, err := router_req.Send()
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_HTTP_REQUEST_DO, err)
        os.Exit(1)
    }

    // Pull the response body from the response
    response_body, err := io.ReadAll(auth_resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_HTTP_READ_BODY, err)
        os.Exit(1)
    }
    defer auth_resp.Body.Close()

    // Unmarshal the response JSON
    responseJsonTemplate := &AuthResponse{}
    err = json.Unmarshal(response_body, responseJsonTemplate)
    if err != nil {
        fmt.Fprintf(os.Stderr, ERR_JSON_UNMARSHAL, err)
        os.Exit(1)
    }

    //Assign new token attributes
    a.AuthToken           = responseJsonTemplate.Auth.Token
    a.AuthTokenExpiration = responseJsonTemplate.Auth.Expiration

}


// Attempt to read in token information from the token file. Failure to do so
// does no constitute a failure since we can request a new token from the
// router
func (a *Auth) LoadAuthTokenFile(tokenFile string) {
    // Read file file
    if bytes, err := os.ReadFile(tokenFile); err == nil {

        // Pull the first line for the file
        first_line := strings.SplitN(string(bytes), "\n", 2)[0]

        // Ensure first line is non-zero in length and contains a space
        if len(first_line) > 0 && strings.Contains(first_line, " ") {

            // Split the fields an assign first field to expiration, and second
            // field to token content
            fields := strings.SplitN(first_line, " ", 2)

            // Convert first field to interger
            if newint, err := strconv.Atoi(strings.Trim(fields[0], " "));
               err == nil {
                a.AuthTokenExpiration = newint
            }
            a.AuthToken = strings.Trim(fields[1], " ")
        }
    }
}



// Write out token information to the token file
func (a *Auth) WriteAuthTokenFile(tokenFile string) {
    // Compose the new string
    auth_string := fmt.Sprintf("%d %s\n",a.AuthTokenExpiration, a.AuthToken)

    // Write the string and capture any errors. Do not consider it a critical
    // error if the write fails, since the token can always be re-requested
    if err := os.WriteFile(tokenFile, []byte(auth_string), 0644); err != nil {
        fmt.Fprintf(os.Stderr, ERR_AUTH_CAN_NOT_WRITE, tokenFile, err)
    }
}


// Return a bool indicating whether the current token is expired. This should be
// called prior to any router request, excluding those requesting a new token
func (a *Auth) IsTokenExpired() (bool, int) {
    // Assume token is expired and prove otherwise
    isExpired := true

    // Calculate the delta between token expiration and now. If it exceeds the
    // threshold (defined in vars.go) then consider the token valid
    expiryDelta := a.AuthTokenExpiration - int(time.Now().Unix())
    if expiryDelta > AUTH_EXPIRATION_THRESHOLD {
        isExpired = false
    }

    // Return both the boolean indicating is expired, and the time remaining
    // before expiration
    return isExpired, expiryDelta
}
