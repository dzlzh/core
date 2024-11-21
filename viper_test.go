package core

import (
	"fmt"
	"os"
)

func ExampleNewViper() {
	type config struct {
		Name  string `mapstructure:"name"`
		Debug bool   `mapstructure:"debug"`
	}

	cbyte := []byte(`
name: Example
debug: true
`)

	os.Setenv("TEST_ENV", "env")
	var c config
	NewViper(&c,
		ViperReadConfig("yaml", cbyte),
		ViperSetEnvPrefix("TEST"),
	)
	fmt.Println(c.Name)
	fmt.Println(c.Debug)
	fmt.Println(G_VIPER.Get("ENV"))
	// Output:
	// Example
	// true
	// env
}
