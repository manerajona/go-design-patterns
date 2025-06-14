package flyweight

import (
	"strings"
)

var nameRepo []string

type User struct {
	names []uint8
}

func NewUser(fullName string) (user *User) {
	getOrAdd := func(name string) uint8 {
		for id := range nameRepo {
			if nameRepo[id] == name {
				return uint8(id)
			}
		}
		nameRepo = append(nameRepo, name)
		return uint8(len(nameRepo) - 1)
	}

	user = &User{}
	names := strings.Split(fullName, " ")
	for _, name := range names {
		user.names = append(user.names, getOrAdd(name))
	}
	return user
}

func (u *User) FullName() string {
	var names []string
	for _, id := range u.names {
		names = append(names, nameRepo[id])
	}
	return strings.Join(names, " ")
}
