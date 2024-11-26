package common

import "time"

type User struct {
	Username string `json:"username"`
}

type GameWithEngine struct {
	Username    string `json:"username"`
	Fen         string `json:"fen,omitempty"`
	EngineName  string `json:"engine_name,omitempty"`
	Level       int    `json:"level"`
	TimeControl string `json:"time_control,omitempty"`
	Color       string `json:"color,omitempty"`
}

type Multiplayer struct {
	WhitePlayerName string `json:"white_player_name"`
	BlackPlayerName string `json:"black_player_name"`
	Fen             string `json:"fen,omitempty"`
	TimeControl     string `json:"time_control,omitempty"`
}

type Move struct {
	GameId int    `json:"game_id"`
	Move   string `json:"move"`
}

type Stat struct {
	Username string `json:"username"`
}

type Top struct {
	Amount int `json:"amount"`
}

type StatResponse struct {
	Games int `json:"games"`
	Wins  int `json:"wins"`
	Loses int `json:"loses"`
	Draws int `json:"draws"`
}

type StartGameResponse struct {
	GameId int `json:"game_id"`
}

type MoveResponse struct {
	Fen         string `json:"fen"`
	IsValidMove bool   `json:"is_valid_move"`
	Result      string `json:"result"`
}

type MultiplayerData struct {
	Color         string    `json:"color" db:"color"`
	TimeControl   string    `json:"time_control" db:"time_control"`
	Result        string    `json:"result" db:"result"`
	OpponentId    string    `json:"opponent_name" db:"opponent_name"`
	CreatedTime   time.Time `json:"created_time" db:"created_time"`
	StartPosition string    `json:"start_position" db:"start_position"`
	Moves         string    `json:"moves" db:"moves"`
}

type SingleplayerData struct {
	Color         string    `json:"color" db:"color"`
	TimeControl   string    `json:"time_control" db:"time_control"`
	Result        string    `json:"result" db:"result"`
	EngineName    string    `json:"engine_name" db:"engine_name"`
	EngineLevel   int       `json:"engine_level" db:"level"`
	CreatedTime   time.Time `json:"created_time" db:"created_time"`
	StartPosition string    `json:"start_position" db:"start_position"`
	Moves         string    `json:"moves" db:"moves"`
}

type PuzzleData struct {
	PuzId        int       `json:"puz_id" db:"puz_id"`
	StartPos     int       `json:"start_pos" db:"start_pos"`
	IsCorrect    bool      `json:"is_correct" db:"is_correct"`
	UserSolution string    `json:"user_solution" db:"user_solution"`
	Solution     string    `json:"solution" db:"solution"`
	TimeTaken    time.Time `json:"time_taken" db:"time_taken"`
}

type PuzzleStat struct {
	PuzId     int `db:"puz_id"`
	AuthorId  int `db:"author_id"`
	Solutions int `db:"solutions"`
	Attempts  int `db:"attempts"`
}

type UserStat struct {
	Rating      int `json:"rating" db:"rating"`
	TopPlace    int `json:"top_place" db:"top_place"`
	AmountGames int `json:"amount_games" db:"amount_games"`
	AmountWins  int `json:"amount_wins" db:"amount_wins"`
	AmountLoses int `json:"amount_loses" db:"amount_loses"`
	AmountDraws int `json:"amount_draws" db:"amount_draws"`
}

type TopStat struct {
	Username string `json:"username" db:"username"`
	Rating   int    `json:"rating" db:"rating"`
	TopPlace int    `json:"top_place" db:"top_place"`
}

type TimeControlsStat struct {
	TimeControl string `json:"time_control" db:"time_control"`
	AmountGames int    `json:"amount_games" db:"amount_games"`
	AmountWins  int    `json:"amount_wins" db:"amount_wins"`
	AmountLoses int    `json:"amount_loses" db:"amount_loses"`
	AmountDraws int    `json:"amount_draws" db:"amount_draws"`
}
