package services

import (
	"log"
	"strings"

	"github.com/brenoandrade/estrategia/utils"
	nlu "github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

var (
	urlWatson string
	apiKey    string
)

var watson *nlu.NaturalLanguageUnderstandingV1

// InitWatson inicia o singleton do watson
func InitWatson(url, key string) {
	watson, _ = nlu.NewNaturalLanguageUnderstandingV1(&nlu.NaturalLanguageUnderstandingV1Options{
		URL:       url,
		Version:   "2019-07-12",
		IAMApiKey: key,
	})
}

// GetWatsonKeywords o watson analisa o repo e sugere algumas categorias em que ele se enquadra
func GetWatsonKeywords(url string) (data []string) {
	log.Println("[WATSON] analyze url:", url)

	var limit int64 = 1
	resp, _ := watson.Analyze(
		&nlu.AnalyzeOptions{
			URL: &url,
			Features: &nlu.Features{
				Categories: &nlu.CategoriesOptions{
					Limit: &limit,
				},
			},
		},
	)

	result := &nlu.AnalysisResults{}
	utils.ByteFromJSON(utils.ToJSON(resp.GetResult()), &result)

	if len(result.Categories) > 0 {
		label := *result.Categories[0].Label
		data = strings.Split(label, "/")
		if len(data) > 0 {
			data = data[1:]
		}
	}

	return data
}
