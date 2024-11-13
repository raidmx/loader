package config

import (
	"encoding/json"
	"os"
)

// Load loads the config from the specified path and loads it into the provided object type
// and returns an error if found.
func Load(path string, data any, defaultBytes []byte) error {
	content, err := os.ReadFile(path)
	if err != nil {
		content = defaultBytes
	}

	if err := json.Unmarshal(content, data); err != nil {
		return err
	}

	return nil
}

// Save saves the config from the specified data structure and saves it by marshalling it into a
// JSON file and returns an error if found.
func Save(path string, data any) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	return nil
}
