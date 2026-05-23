package frontier

import (
	"log"

	"github.com/aditya-sutar-45/search--/crawler/utils"
)

type VisistedSet struct {
	*ReddisConnection
}

func NewVisistedSet(r *ReddisConnection) *VisistedSet {
	return &VisistedSet{
		ReddisConnection: r,
	}
}

func (v *VisistedSet) IsVisited(url string) bool {
	urlHash := utils.HashURL(url)

	exists, err := v.client.SIsMember(v.ctx, VISITED_SET_KEY, urlHash).Result()
	if err != nil {
		return false
	}

	return exists
}

func (v *VisistedSet) MarkVisisted(url string) {
	urlHash := utils.HashURL(url)

	err := v.client.SAdd(v.ctx, VISITED_SET_KEY, urlHash).Err()
	if err != nil {
		log.Printf("ERROR could not add url %s to set: %v\n", url, err)
		return
	}
}
