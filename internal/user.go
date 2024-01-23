package internal

import (
	"context"
	"crypto/sha256"
	"errors"

	model "github.com/b3xie/slurp/graph/model"
)

func CreateUser(ctx context.Context, userInfo model.NewUser) (*model.Message, error) {
	rdb := RedisConnect()
	exists := rdb.HExists(ctx, "users", userInfo.Username)
	if exists.Val() {
		return nil, errors.New("user already exists")
	}
	sum := sha256.Sum256([]byte(string(userInfo.Secret)))
	hashString := string(sum[:])
	rdb.HSet(ctx, "users", []string{string(userInfo.Username), hashString})
	var reply model.Message
	reply.ID = "200 - ok"
	reply.Text = "User created"
	var res *model.Message = &reply
	return res, nil
}
