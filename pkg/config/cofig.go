package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

// to store env variables
type Config struct {
	StripSecretKey      string `mapstructure:"STRIPE_SECRET"`
	StripPublishKey     string `mapstructure:"STRIPE_PUBLISH_KEY"`
	StripeWebhookSecret string `mapstructure:"STRIPE_WEBHOOK"`
}

// to hold all names of env variables
var envsNames = []string{
	"STRIPE_SECRET", "STRIPE_PUBLISH_KEY", "STRIPE_WEBHOOK", // stripe
}

var config Config // create an instance of Config
// func to get env variable and store it on struct Config and retuen it with error as nil or error
func LoadConfig() (Config, error) {

	// set-up viper
	viper.AddConfigPath("./")   // add the config path
	viper.SetConfigFile(".env") // set up the file name to viper
	viper.ReadInConfig()        // read the env file

	// range through through the envNames and take each envName and bind that env variable to viper
	for _, env := range envsNames {
		if err := viper.BindEnv(env); err != nil {
			return config, err // error when binding the env to viper
		}
	}

	// then unmarshel the viper into config variable
	if err := viper.Unmarshal(&config); err != nil {
		return config, err // error when unmarsheling the viper to env
	}

	// atlast validate the config file using validator pakage
	// create instance and direct validte
	if err := validator.New().Struct(config); err != nil {
		return config, err // error when validating struct
	}

	//successfully loaded the env values into struct config
	return config, nil
}

func GetCofig() Config {
	return config
}
