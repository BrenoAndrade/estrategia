package model

// ConfigFile estrutura de configuração
type ConfigFile struct {
	Port      string
	URLWatson string
	APIKey    string
	URLRedis  string
}

// GetConfig leria o config.json se tivesse um kk :D
func GetConfig() ConfigFile {
	return ConfigFile{
		Port:      ":3001",
		URLWatson: "https://gateway-wdc.watsonplatform.net/natural-language-understanding/api",
		APIKey:    "tbi2jvQ9sIfCju9CSony03NVw8GHaxdP2DmmEiHpGrF-",
		URLRedis:  "localhost:6379",
	}
}
