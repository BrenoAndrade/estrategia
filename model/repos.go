package model

import (
	"strconv"
)

// Repo estrutura para um repositório
type Repo struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Language    string   `json:"language"`
	URL         string   `json:"html_url"`
	Tags        []string `json:"tags"`
}

// Key gera uma chave para ser utilizada no redis
func (repo *Repo) Key() string {
	key := strconv.Itoa(repo.ID)
	for _, tag := range repo.Tags {
		key += "|" + tag
	}
	return key
}

func (repo *Repo) duplicatedTag(tag string) bool {
	for _, item := range repo.Tags {
		if item == tag {
			return true
		}
	}

	return false
}

// AddTag adiciona uma tag a um repo
func (repo *Repo) AddTag(tag string) bool {
	if repo.duplicatedTag(tag) {
		return false
	}

	repo.Tags = append(repo.Tags, tag)

	return true
}

// DelTag remove uma tag de um repo
func (repo *Repo) DelTag(tag string) bool {
	if !repo.duplicatedTag(tag) {
		return false
	}

	data := make([]string, 0)
	for _, item := range repo.Tags {
		if item != tag && item != "" {
			data = append(data, item)
		}
	}

	repo.Tags = data

	return true
}
