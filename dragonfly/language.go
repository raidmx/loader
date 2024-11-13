package dragonfly

import (
	_ "embed"
	"fmt"
	"github.com/stcraft/loader/config"

	"github.com/stcraft/dragonfly/server/player"
)

// languageConfig specifies the structure in which translations and
// toasts are saved in the config file.
type languageConfig struct {
	Messages map[string]string       `json:"messages"`
	Toasts   map[string]config.Toast `json:"toasts"`
}

// This is an unexported instance of language translations
var lang languageConfig

//go:embed language.json
var defaultLang []byte

// loadLanguage Parses the language.json file if it exists and creates a new one
// with the default values if it doesn't exist and loads all the content.
func LoadLanguage() {
	if err := config.Load("language.json", &lang, defaultLang); err != nil {
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
