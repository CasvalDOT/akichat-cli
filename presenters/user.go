package presenters

import (
	"fmt"

	core "github.com/CasvalDOT/akichat-core"
)

type usersPresenter struct {
	users []core.User
}

func (m *usersPresenter) Print() {
	for _, user := range m.users {
		fmt.Printf("%s - %s\n", user.ID, user.Name)
	}
}

func NewUsersPresenter(users []core.User) Presenter {
	return &usersPresenter{
		users,
	}
}
