/*Package main show how to use viper.

Start the app in the root folder with:

	SERVER_HOST=testing go run cmd/main.go

*/
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	fmt.Println(viper.Get("server.host"))
}
