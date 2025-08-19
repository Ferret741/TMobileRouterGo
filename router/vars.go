package router

import (
    "time"
)

const (
    // Define fixed strings here
    ROUTER_URL_LOGIN            = "TMI/v1/auth/login"
    ROUTER_URL_TELEMETRY        = "TMI/v1/network/telemetry?get=clients"
    ROUTER_PORT_LOGIN           = 8080
    ROUTER_CREDS_USER           = "admin"

    REQUEST_USER_AGENT          = "jRome741/2.0+"
    REQUEST_TIMEOUT             = 3*time.Second

    AUTH_EXPIRATION_THRESHOLD   = 100

    CONN_TEST_TIMEOUT           = 500*time.Millisecond

    ARG_REGEX_VALIDATE_ADDR     = `^(\d{1,3}\.){3}\d{1,3}$`
    ARG_REGEX_VALIDATE_PORT     = `^\d{2,5}$`
    ARG_REGEX_VALIDATE_PATH     = `^(\.?(/\w+\S+){0,})\w+\S+$`
    ARG_REGEX_VALIDATE_DEVICE   = `^\w[\w\-]+$`



    // Define error strings here
    ERR_NO_CONFIG_FILE          = "Could not access config file: %s\n\tError: %s\n"
    ERR_ACCESS_CONFIG_FILE      = "Error encountered while reading config file: %s\n\tError: %s\n"
    ERR_ROUTER_UNREACHABLE      = "Router is unreachable at %s:%s\n\tError: %s\n"
    ERR_AUTH_CAN_NOT_WRITE      = "Error encountered while writing to auth file: %s\n\tError: %s\n"
    ERR_JSON_MARSHAL            = "Error marshaling JSON content\n\tError: %s\n"
    ERR_JSON_UNMARSHAL          = "Error unmarshaling JSON content\n\tError: %s\n"
    ERR_HTTP_REQUEST_DO         = "Error during HTTP request\n\tError: %s\n"
    ERR_HTTP_REQUEST_GENERATE   = "Error generating HTTP request\n\tError: %s\n"
    ERR_HTTP_READ_BODY          = "Error reading HTTP body\n\tError: %s\n"
    ERR_CONN_TEST_FAILED        = "Could not establish a connection to %s\n\tError: %s\n"
    ERR_ARGS_INSUFFICIENT       = "Could not find argument for '%s'\n"
    ERR_ARGS_REGEX_INVALID      = "Could not compile regular expression\n\tError:%s'\n"
    ERR_ARGS_INVALID_SYNTAX     = "Invalid format for argument '%s' (Given value: '%s')\n"


    HELP_BODY = "\x1B[1;31m\n" +
    "Usage %s [options]\n" +
    "\n" +
    "Options:\n" +
    "\n" +
    "(--config-file|--config|-c) <filepath>\n" +
    "    Path to a configuration file. Configuration file should have the following lines\n" +
    "        - router_username\n" +
    "        - router_password\n" +
    "        - router_address\n" +
    "        - router_port\n" +
    "        - tokenfilepath\n" +
    "\n" +
    "(--device-name|-n) <device_name>\n" +
    "    Device name for which to find IP address\n" +
    "\n" +
    "(--token-file|-t) <token_file_path>\n" +
    "    Path to file in which token string and expiration date are stored\n" +
    "\n" +
    "(--router-address|--router|-A) <router_ip_address>\n" +
    "    IP address of the router \n" +
    "\n" +
    "(--router-port|--port|-p) <router_port>\n" +
    "    RESTAPI port for the router\n" +
    "\n" +
    "(--show-ip-address|--ip-address|-i)\n" +
    "    Search for the IP address of device given by <device_name> (See above)\n" +
    "\n" +
    "(--print-table|--all|-t|-a)\n" +
    "    Print all devices along with full gateway device telemetry\n" +
    "\n" +
    "(--help|-h)\n" +
    "    Show this help page and exit\n" +
    "\n" +
    "\n" +
    "Examples\n" +
    "%s --device-name optiplex --show-ip-address --config ./config  --device optiplex\n" +
    "\n" +
    "\x1B[0m\n"
)
