package adapter

import (
	"entity"
	"fmt"
)

func UpdateUsers(users []*entity.User) {
	fmt.Println(users)
}
