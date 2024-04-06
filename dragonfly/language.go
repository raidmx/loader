package dragonfly

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/STCraft/dragonfly/server/player"
)

// Toast specifies the structure in which toasts are saved in the config file.
type Toast struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// language specifies the structure in which translations and toasts are saved in the config
// file.
type language struct {
	Messages map[string]string `json:"messages"`
	Toasts   map[string]Toast  `json:"toasts"`
}

// This is an unexported instance of language translations
var lang language

//go:embed language.json
var defaultLanguage []byte

// loadLanguage Parses the language.json file if it exists and creates a new one
// with the default values if it doesn't exist and loads all the content.
func loadLanguage() {
	content, err := os.ReadFile("./language.json")

	if err != nil {
		content = defaultLanguage

		if err := os.WriteFile("./language.json", defaultLanguage, 0755); err != nil {
			panic(err)
		}
	}

	if err := json.Unmarshal(content, &lang); err != nil {
		panic(err)
	}
}

// Returns the translation for the provided key from the loaded translations
func Translation(key string, args ...any) string {
	v, ok := lang.Messages[key]

	if !ok {
		panic(fmt.Sprintf("Translation with the key %s does not exist", key))
	}

	return fmt.Sprintln(fmt.Sprintf(v, args...))
}

// SendToast sends the toast for the provided key to the specified player
func SendToast(p *player.Player, key string) {
	t, ok := lang.Toasts[key]

	if !ok {
		panic(fmt.Sprintf("Toast with the key %s does not exist", key))
	}

	p.SendToast(t.Title, t.Content)
}
