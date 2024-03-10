package storage

import (
	"fmt"
	"mobidev/internal/models"
)

type InMemoryStorage map[string]models.User

func (inMem InMemoryStorage) Save(u models.User) {
	inMem[u.Email] = u
}

func (inMem InMemoryStorage) Load(email string) (models.User, error) {
	user, exists := inMem[email]
	if !exists {
		return user, fmt.Errorf("user %s not exists", email)
	}

	return user, nil

}

var InMemory = make(InMemoryStorage)
