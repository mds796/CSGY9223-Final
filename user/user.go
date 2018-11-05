package user

import (
	"github.com/google/uuid"
	"strings"
)

type User struct {
	uuid     uuid.UUID
	username string
}

// create assigns a uuid to a username
func (user *User) create(username string) uuid.UUID {
	uuid := uuid.New()
	users[uuid] = username
	return uuid
}

// view returns a uuid's corresponding username
func (user *User) view(uuid uuid.UUID) (username string) {
	return users[uuid]
}

// search returns a list of uuids that match a given query
func (user *User) search(query string) (uuids []uuid.UUID) {
	for uuid, username := range users {
		if strings.Contains(username, query) {
			uuids = append(uuids, uuid)
		}
	}
	return uuids
}
