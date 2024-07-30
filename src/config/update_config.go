package config

import (
	"encoding/json"
)

type UpdateConfig struct {
	Destination     string `json:"destination"`
	Source          string `json:"source"`
	InstructionSets []Set  `json:"instruction_sets"`
	ScriptPath      string `json:"script_path"`
}

type Set struct {
	Description  string        `json:"description"`
	TargetTable  string        `json:"target_table"`
	Instructions []Instruction `json:"instructions"`
}

type Instruction struct {
	Description string `json:"description"`
	Method      string `json:"method"`
	Execute     string `json:"execute"`
	Rollback    string `json:"rollback"`
}

type Replacement struct {
	Key   string
	Value string
}

var Updcfg *UpdateConfig

// Simple function to handle unmarshaling the Application Config
func HandleUpdateConfig(body []byte) (*UpdateConfig, error) {
	cfg := new(UpdateConfig)
	err := json.Unmarshal(body, &cfg)
	//store for use later
	Updcfg = cfg
	return cfg, err

}

func GetUpdateConfig() (config *UpdateConfig) {

	return Updcfg
}
