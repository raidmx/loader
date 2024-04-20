package dragonfly

import (
	_ "embed"
	"encoding/json"
	"os"
	"sync"
)

//go:embed operators.json
var defaultOperators []byte

// loadOperators loads all the operators from the operators.json file
func LoadOperators() {
	content, err := os.ReadFile("./operators.json")

	if err != nil {
		content = defaultOperators
	}

	if err := json.Unmarshal(content, &operators); err != nil {
		panic(err)
	}

	operators.mu = &sync.RWMutex{}
}

// Saves all the operators to the disk.
func SaveOperators() {
	defer operators.mu.RUnlock()
	operators.mu.RLock()

	content, err := json.MarshalIndent(operators, "", "  ")

	if err != nil {
		panic(err)
	}

	os.RemoveAll("./operators.json")

	if err := os.WriteFile("./operators.json", content, 0755); err != nil {
		panic(err)
	}
}

// OperatorRegistry is the registry of all the operators of the server
type OperatorRegistry struct {
	mu   *sync.RWMutex
	List []string `json:"list"`
}

// operators is the instance of OperatorRegistry containing all the operators
// of the server
var operators OperatorRegistry

// Sets the player with the provided XUID as the operator
func SetOP(xuid string) {
	defer operators.mu.Unlock()
	operators.mu.Lock()

	operators.List = append(operators.List, xuid)
}

// Returns whether the player with the provided XUID is an operator
func IsOP(xuid string) bool {
	defer operators.mu.RUnlock()
	operators.mu.RLock()

	for _, id := range operators.List {
		if xuid == id {
			return true
		}
	}

	return false
}

// RemoveOP removes the player with the provided UUID from the operator status
func RemoveOP(xuid string) {
	defer operators.mu.Unlock()
	operators.mu.Lock()

	for index, id := range operators.List {
		if xuid == id {
			operators.List[index] = operators.List[len(operators.List)-1]
			operators.List = operators.List[:len(operators.List)-1]
		}
	}
}
