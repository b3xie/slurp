package internal

import (
	"context"
	"crypto/sha256"
	"fmt"

	model "github.com/b3xie/slurp/graph/model"
)

func UserIDMake() uint {

	return 0
}
func CreateUser(ctx context.Context, userInfo model.NewUser) {
	rdb := RedisConnect()
	userLen := func() (int64, map[string]string) {
		len := rdb.HLen(ctx, "users").Val()
		usr := rdb.HGetAll(ctx, "users").Val()
		return len, usr
	}
	_, usr := userLen()
	for i, e := range usr {
		fmt.Printf(e)
		fmt.Printf(i)
	}
	sum := sha256.Sum256([]byte(string(userInfo.Secret)))
	hashString := string(sum[:])
	rdb.HSet(ctx, "users", []string{string(userInfo.Username), hashString})

}
