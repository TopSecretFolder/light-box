package actions

type ActionType string

type Action[T any] struct {
	Type    ActionType `json:"type"`
	Payload T          `json:"payload"`
}
