package configurator

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Initialize() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Cannot read in configs: %v\n", err)
		return
	}
	log.Printf("Read in configs successfully\n")
}
