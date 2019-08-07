package model

// ConfigFile struct for config file
type ConfigFile struct {
	Port      string
	URLWatson string
	APIKey    string
	URLRedis  string
}

// GetConfig get config file
func GetConfig() ConfigFile {
	return ConfigFile{
		Port:      ":3001",
		URLWatson: "https://gateway-wdc.watsonplatform.net/natural-language-understanding/api",
		APIKey:    "tbi2jvQ9sIfCju9CSony03NVw8GHaxdP2DmmEiHpGrF-",
		URLRedis:  "",
	}
}
