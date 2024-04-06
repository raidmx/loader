package dragonfly

import (
	_ "embed"
	"encoding/json"
	"os"
)

// Operators specify the structure in which the server operators are stored.
type Operators struct {
	List []string `json:"list"`
}

//go:embed operators.json
var defaultOperators []byte

// loadOperators loads all the operators from the operators.json file
func loadOperators() {
	content, err := os.ReadFile("./operators.json")

	if err != nil {
		content = defaultOperators
	}

	if err := json.Unmarshal(content, &operators); err != nil {
		panic(err)
	}
}

// Operators is the list of server Operators which is saved after closing the server
// in the file Operators.json
var operators Operators

// Sets the player with the provided XUID as the operator
func SetOP(xuid string) {
	operators.List = append(operators.List, xuid)
}

// Returns whether the player with the provided XUID is an operator
func IsOP(xuid string) bool {
	for _, id := range operators.List {
		if xuid == id {
			return true
		}
	}

	return false
}

// RemoveOP removes the player with the provided UUID from the operator status
func RemoveOP(xuid string) {
	for index, id := range operators.List {
		if xuid == id {
			operators.List[index] = operators.List[len(operators.List)-1]
			operators.List = operators.List[:len(operators.List)-1]
		}
	}
}
