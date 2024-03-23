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
	v := NewViper(&c,
		ViperReadConfig("yaml", cbyte),
		ViperSetEnvPrefix("TEST"),
	)
	fmt.Println(c.Name)
	fmt.Println(c.Debug)
	fmt.Println(v.Get("ENV"))
	// Output:
	// Example
	// true
	// env
}
