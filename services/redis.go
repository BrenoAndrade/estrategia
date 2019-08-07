package services

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/utils"
	pkgredis "github.com/go-redis/redis"
)

var redis *pkgredis.Client

// InitRedis inicia o singleton do redis
func InitRedis(url string) {
	redis = pkgredis.NewClient(&pkgredis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
}

// SetRepo cria um repo no redis
func SetRepo(repo *model.Repo) *model.Error {
	keys, err := redis.Keys(strconv.Itoa(repo.ID) + "*").Result()
	if err != nil {
		log.Println("[REDIS-ERROR] SetRepo:", err.Error())
		return model.NewError("services.redis.set_repo", err.Error(), http.StatusInternalServerError)
	}

	if len(keys) > 0 {
		data, _ := GetRepo(keys[0])
		repo.Tags = data.Tags
	}

	if err := redis.Set(repo.Key(), string(utils.ToJSON(repo)), 0).Err(); err != nil {
		log.Println("[REDIS-ERROR] SetRepo:", err.Error())
		return model.NewError("services.redis.set_repo", err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// DelRepo deleta um repo do redis
func DelRepo(key string) {
	redis.Del(key)
}

// SearchRepos procura por repos e os retorna
func SearchRepos(key string) ([]*model.Repo, *model.Error) {
	keys, err := redis.Keys(key).Result()
	if err != nil {
		log.Println("[REDIS-ERROR] SearchRepos:", err.Error())
		return nil, model.NewError("services.redis.search_repos", err.Error(), http.StatusInternalServerError)
	}

	data := make([]*model.Repo, 0)
	for _, key := range keys {
		repo, err := GetRepo(key)
		if err == nil {
			data = append(data, repo)
		}
	}

	return data, nil
}

// GetRepo retorna um determinado repo
func GetRepo(key string) (*model.Repo, *model.Error) {
	str, err := redis.Get(key).Result()
	if err != nil {
		log.Println("[REDIS-ERROR] GetRepo:", err.Error())
		return nil, model.NewError("services.redis.get_repo", err.Error(), http.StatusInternalServerError)
	}

	data := &model.Repo{}
	utils.ReaderFromJSON(strings.NewReader(str), &data)

	return data, nil
}
