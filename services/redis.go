package services

import (
	"log"
	"net/http"

	"github.com/brenoandrade/estrategia/model"
	"github.com/brenoandrade/estrategia/utils"
	pkgredis "github.com/go-redis/redis"
)

var redis *pkgredis.Client

// InitRedis as
func InitRedis(url string) {
	redis = pkgredis.NewClient(&pkgredis.Options{
		Addr: url,
	})

	log.Println("[REDIS] connected...")
}

// SetRepo asdsa
func SetRepo(repo *model.Repo) *model.Error {
	if err := redis.Set(repo.Key(), utils.ToJSON(repo), 0).Err(); err != nil {
		return model.NewError("services.redis.set_repo", err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// SetRepos a
func SetRepos(repos []*model.Repo) *model.Error {
	pipe := redis.Pipeline()
	for _, repo := range repos {
		pipe.Set(repo.Key(), utils.ToJSON(repo), 0)
	}

	if _, err := pipe.Exec(); err != nil {
		return model.NewError("services.redis.set_repos", err.Error(), http.StatusInternalServerError)
	}

	return nil
}

// SearchRepos dasda
func SearchRepos(key string) ([]*model.Repo, *model.Error) {
	keys, err := redis.Keys(key).Result()
	if err != nil {
		return nil, model.NewError("services.redis.search_repos", err.Error(), http.StatusInternalServerError)
	}

	data := make([]*model.Repo, len(keys))
	for _, key := range keys {
		repo, err := GetRepo(key)
		if err == nil {
			data = append(data, repo)
		}
	}

	return nil, nil
}

// GetRepo sadsa
func GetRepo(key string) (*model.Repo, *model.Error) {
	bt, err := redis.Get(key).Bytes()
	if err != nil {
		return nil, model.NewError("services.redis.get_repo", err.Error(), http.StatusInternalServerError)
	}

	data := &model.Repo{}
	utils.ByteFromJSON(bt, &data)

	return data, nil
}
