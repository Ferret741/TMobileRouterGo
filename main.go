package main

import (
    "fmt"
    "os"
    "localhost/tmobile/router"
    "regexp"
)


func validateArgument(currentIndex int, currentArg, regex string){

    // Make certain that there are sufficient arguments in the
    // provided os.Args list
    if currentIndex + 1 >= len(os.Args) {
        fmt.Fprintf(os.Stderr, router.ERR_ARGS_INSUFFICIENT, currentArg)
        os.Exit(13)
    }

    // Perform a regex check to make sure that we are receiving what
    // we are expecting (e.g. that a port is comprised of all numbers,
    // or that a file path is properly formatted)
    if compiledRegex, err := regexp.Compile(regex); err == nil {
        if ! compiledRegex.Match([]byte(os.Args[currentIndex+1])) {
            fmt.Fprintf(os.Stderr, router.ERR_ARGS_INVALID_SYNTAX, currentArg, os.Args[currentIndex+1])
            os.Exit(19)
        }

    // The provided regex string could not be compiled
    } else {
        fmt.Fprintf(os.Stderr, router.ERR_ARGS_REGEX_INVALID, err)
        os.Exit(17)
    }
}


func getCommandLineArgs(conf *router.Config){
    for index, value := range os.Args {

        switch string(value) {
            case "--config-file","--config","-c":
                validateArgument(index, value, router.ARG_REGEX_VALIDATE_PATH)
                conf.ConfigFilepath = os.Args[index+1]

            case "--device-name","-n":
                validateArgument(index, value, router.ARG_REGEX_VALIDATE_DEVICE)
                conf.DeviceName = os.Args[index+1]

            case "--token-file","-t":
                validateArgument(index, value, router.ARG_REGEX_VALIDATE_PATH)
                conf.TokenFilepath = os.Args[index+1]

            case "--router-address","--host","-A":
                validateArgument(index, value, router.ARG_REGEX_VALIDATE_ADDR)
                conf.RouterAddr = os.Args[index+1]

            case "--router-port","--port","-p":
                validateArgument(index, value, router.ARG_REGEX_VALIDATE_PORT)
                conf.RouterPort = os.Args[index+1]

            case "--show-ip-address","--ip-address","-i":
                conf.Output = "ip-address"

            case "--print-table","-T","--all","-a":
                conf.Output = "table"

            case "--help","-h":
                fmt.Printf(router.HELP_BODY, os.Args[0], os.Args[0])
                os.Exit(0)

            default:
        }
    }
}



func main() {

    var routerConfig    router.Config
    var routerAuth      router.Auth
    var routerTelemetry router.Telemetry

    // Override configuration from command line arguments
    getCommandLineArgs(&routerConfig)

    // Populate configuration object from file
    routerConfig.PopulateConfigFromFile()

    // Load the Auth token file
    routerAuth.LoadAuthTokenFile(routerConfig.TokenFilepath)

    // Test the connection to the router
    router.TestConnection(routerConfig.RouterAddr, routerConfig.RouterPort)

    // Determine if the Auth Token is expired
    if isExpired, _ := routerAuth.IsTokenExpired(); isExpired {

        // Generate a request URL
        request_url := fmt.Sprintf("http://%s:%s/%s", routerConfig.RouterAddr,
                                                      routerConfig.RouterPort,
                                                      router.ROUTER_URL_LOGIN)

        // Request a new Auth Token
        routerAuth.RequestNewToken(request_url,
                                   routerConfig.Username,
                                   routerConfig.Password)

        // Write out new data
        routerAuth.WriteAuthTokenFile(routerConfig.TokenFilepath)
    }

    request_telemetry_url := fmt.Sprintf("http://%s:%s/%s", routerConfig.RouterAddr,
                                                            routerConfig.RouterPort,
                                                            router.ROUTER_URL_TELEMETRY)

    // Retreive the data
    telemetryData := routerTelemetry.GetData(request_telemetry_url, routerAuth.AuthToken)

    // Take steps for appropriate output here
    router.IterateDevices(telemetryData, routerConfig.Output, routerConfig.DeviceName)




}
