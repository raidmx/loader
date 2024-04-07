package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Object represents a map of all loaded keys and value pairs.
type Object map[string]any

// New returns the Config at the specified path. If no config is found it creates one
// and also returns the default config.
func New(dir string, file string, defaultConfig []byte) Object {
	if err := os.MkdirAll(fmt.Sprintf("./%s", dir), 0755); err != nil {
		panic(err)
	}

	path := fmt.Sprintf("./%s/%s", dir, file)

	content, err := os.ReadFile(path)
	if err != nil {
		if err := os.WriteFile(path, defaultConfig, 0755); err != nil {
			panic(err)
		}

		content = defaultConfig
	}

	var config Object

	if err := json.Unmarshal(content, &config); err != nil {
		panic(err)
	}

	return config
}

// Get ...
func (c Object) Get(key string) any {
	value, ok := c[key]
	if !ok {
		panic("unable to find key")
	}

	return value
}

// GetObject ...
func (c Object) GetObject(key string) Object {
	val, ok := c.Get(key).(map[string]any)
	if !ok {
		panic("value is not a subconfig")
	}

	return val
}

// GetString ...
func (c Object) GetString(key string) string {
	val, ok := c.Get(key).(string)
	if !ok {
		panic("value is not a string")
	}

	return val
}

// GetInt ...
func (c Object) GetInt(key string) int {
	val, ok := c.Get(key).(int)
	if !ok {
		panic("value is not an int")
	}

	return val
}

// GetFloat ...
func (c Object) GetFloat(key string) float64 {
	val, ok := c.Get(key).(float64)
	if !ok {
		panic("value is not a float64")
	}

	return val
}

// GetBool ...
func (c Object) GetBool(key string) bool {
	val, ok := c.Get(key).(bool)
	if !ok {
		panic("value is not a bool")
	}

	return val
}

// GetMessage ...
func (c Object) GetMessage(key string, args ...any) string {
	formatter := c.GetString(key)
	return fmt.Sprintf(formatter, args...)
}

// GetToast ...
func (c Object) GetToast(key string) (string, string) {
	toast := c.GetObject(key)
	return toast.GetString("title"), toast.GetString("content")
}
