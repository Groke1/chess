package config

type SingleGameConfig struct {
	Username             string
	EngineName           string
	EngineLevel          int
	IsPlayerWhite        bool
	Result               string
	PositionFen          string
	TimeControlMinutes   int
	TimeControlIncrement int
	Moves                string
}
