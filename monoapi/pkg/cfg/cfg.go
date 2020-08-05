package cfg

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Validatable interface {
	Validate() error
}

func Init(name, file string, config Validatable) error {
	viper.SetConfigName(name) // name of config file (without extension)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Wrap(err, "failed to get current dir")
	}

	viper.AddConfigPath(dir)
	viper.AutomaticEnv()

	if file != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(file)
	}

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "viper: failed to read config")
	}

	if err = viper.Unmarshal(config); err != nil {
		return errors.Wrap(err, "failed to unmarshal config to obj")
	}

	return config.Validate()
}
