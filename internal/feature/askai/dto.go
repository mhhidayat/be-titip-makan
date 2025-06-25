package askai

type AskAiRequest struct {
	Prompt string `json:"prompt" validate:"required"`
}
