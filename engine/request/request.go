package request

const (
	DefaultDepth = 16
	DefaultLevel = 20
)

type Request struct {
	Fen   string `json:"fen"`
	Depth int    `json:"depth,omitempty"`
	Level int    `json:"level,omitempty"`
}

func New() *Request {
	return &Request{Depth: DefaultDepth, Level: DefaultLevel}
}
