package model

import "strconv"

// Repo model
type Repo struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Language    string   `json:"language"`
	URL         string   `json:"html_url"`
	Tags        []string `json:"tags"`
}

// Key Generate key for db
func (repo *Repo) Key() string {
	key := strconv.Itoa(repo.ID)
	for _, tag := range repo.Tags {
		key += "|" + tag
	}
	return key
}

// duplicatedTag a
func (repo *Repo) duplicatedTag(tag string) bool {
	for _, item := range repo.Tags {
		if item == tag {
			return true
		}
	}

	return false
}

// AddTag a
func (repo *Repo) AddTag(tag string) bool {
	if repo.duplicatedTag(tag) {
		return false
	}

	repo.Tags = append(repo.Tags, tag)

	return true
}

// DelTag a
func (repo *Repo) DelTag(tag string) bool {
	if !repo.duplicatedTag(tag) {
		return false
	}

	data := make([]string, len(repo.Tags)-1)
	for _, item := range repo.Tags {
		if item != tag {
			data = append(data, item)
		}
	}

	repo.Tags = data

	return true
}
