package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	env  string
	port string
)

// Load loads variable from config/env.yml
// return true if load successfully
// return false if can not load env.yml
func Load(filename string) bool {
	if filename == "" {
		panic("filename can not be blank.")
	}

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("can not found config file.")
		return false
	}

	var config = make(map[string]string)
	yamlErr := yaml.Unmarshal(file, &config)

	if yamlErr != nil {
		panic(yamlErr)
	}

	for k, v := range config {
		os.Setenv(k, v)
	}

	flag.StringVar(&env, "env", os.Getenv("env"), "set up env variable")
	flag.StringVar(&port, "port", os.Getenv("port"), "set up server port")

	flag.Parse()

	if env != "" {
		os.Setenv("env", env)
	}

	if port != "" {
		os.Setenv("port", port)
	}

	return true
}

// Get returns value and exist
func Get(key string) (result string, exist bool) {
	value := os.Getenv(key)
	if env == "" {
		return "", false
	}

	return value, true
}

// MustGet return env and return error if not found key
func MustGet(key string) (result string) {
	env := os.Getenv(key)
	if env == "" {
		msg := fmt.Sprintf(
			"can not find %s in `ENV`, please checkout your env file.",
			key,
		)

		panic(msg)
	}

	return env
}
