package askai

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/jsonutil"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
)

type askaiHandler struct {
	validate *validator.Validate
	configAI *configs.AI
}

func NewAskAI(router fiber.Router, configAI *configs.AI, validate *validator.Validate) {
	ah := askaiHandler{
		validate: validate,
		configAI: configAI,
	}

	router.Get("/ask-ai", ah.askAI)
}

func (ah *askaiHandler) askAI(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 20*time.Second)
	defer cancel()

	req := AskAiRequest{}

	c.BodyParser(&req)

	err := ah.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	channel := make(chan string)
	go ah.generateAIResponse(ctx, req.Prompt, channel)
	defer close(channel)
	response := <-channel
	return c.Status(http.StatusOK).JSON(jsonutil.SuccessResponse("Successfully retrieved AI response", fiber.Map{
		"response": response,
	}))
}

func (ah *askaiHandler) generateAIResponse(ctx context.Context, prompt string, c chan string) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  ah.configAI.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Models.GenerateContent(
		ctx,
		ah.configAI.Model,
		genai.Text(fmt.Sprintf(
			"Jawab hanya jika pertanyaan berikut berkaitan dengan makanan. Jika tidak, balas dengan: 'Maaf, saya hanya dapat menjawab pertanyaan seputar makanan.' Pertanyaannya: %s", prompt)),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	c <- result.Text()
}
