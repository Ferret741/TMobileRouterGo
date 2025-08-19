package router

import (
    "fmt"
    "os"
    "strings"
)

type Config struct {
    // Exportable
    ConfigFilepath      string
    RouterAddr          string
    RouterPort          string
    TokenFilepath       string
    DeviceName          string
    Output              string

    // Non-exportable
    Password    string
    Username    string
}



// Attempt to load the configuration file if it exists. Since
// the configuration is paramount to proper execution of the
// program, exit with an error
func (c *Config) loadConfigFile() []byte {
   // Attempt to read the file
   bytes, err := os.ReadFile(c.ConfigFilepath)

   // Handle errors
   if err != nil {
       if os.IsNotExist(err) {
           fmt.Fprintf(os.Stderr, ERR_NO_CONFIG_FILE, c.ConfigFilepath, err)
           os.Exit(1)
       }
       fmt.Fprintf(os.Stderr, ERR_ACCESS_CONFIG_FILE, c.ConfigFilepath, err)
       os.Exit(1)
   }

   return bytes
}


// Populate the configuration object from the configuration file.
func (c *Config) PopulateConfigFromFile() {
    content := string(c.loadConfigFile())
    for _, value := range strings.Split(content,"\n") {

        if ! strings.Contains(value, ":") {
            continue
        }

        key := strings.Trim(strings.SplitN(value, ":", 2)[0], " ")
        val := strings.Trim(strings.SplitN(value, ":", 2)[1], " ")

        // Only set the value if they have not been
        // overriden by command line arguments
        switch strings.ToLower(key) {
            case "router_username":
                if c.Username == "" {
                    c.Username = val
                }

            case "router_password":
                c.Password = val

            case "router_port":
                if c.RouterPort == "" {
                    c.RouterPort = val
                }

            case "router_address":
                if c.RouterAddr == "" {
                    c.RouterAddr = val
                }

            case "tokenfilepath":
                if c.TokenFilepath == "" {
                    c.TokenFilepath = val
                }
        }
    }
}
