package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func Gemini(text string) (*genai.GenerateContentResponse, error) {

	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	viper.AutomaticEnv()

	API_KEY := viper.GetString("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// text generation using gemini model
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(text))
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// func printResponse(resp *genai.GenerateContentResponse) {
// 	for _, cand := range resp.Candidates {
// 		if cand.Content != nil {
// 			for _, part := range cand.Content.Parts {
// 				fmt.Println(part)
// 			}
// 		}
// 	}
// 	fmt.Println("---")
// }
