package storage

import (
	"fmt"
	"net/http"
	"time"

	"mobidev/internal/models"
)

type InMemoryStorage map[string]models.User

var InMemory = make(InMemoryStorage)

func (inMem InMemoryStorage) Save(u models.User) {
	inMem[u.Email] = u
}

// TODO: change func signature.
// return `bool` instead `error`.
func (inMem InMemoryStorage) Load(email string) (models.User, error) {
	user, exists := inMem[email]
	if !exists {
		return user, fmt.Errorf("user %s not exists", email)
	}

	return user, nil
}

func (inMem InMemoryStorage) ValidCookies(reqCookie *http.Cookie) bool {
	for _, user := range inMem {
		if user.Cookie == reqCookie.Value && time.Now().Unix()-user.CookieExpires<0{
			return true
		}
	}
	return false
}
