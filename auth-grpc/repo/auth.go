package repo

import (
	"context"
	"errors"
	"log"

	"github.com/ilmsg/studious-barnacle/auth-grpc/model"
	"github.com/ilmsg/studious-barnacle/auth-grpc/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoAuth struct {
	collection *mongo.Collection
}

var (
	errInvalidPassword = errors.New("invalid password")
	errCreateToken     = errors.New("something wrong on create jwt token")
	errExistsEmail     = errors.New("exists email")
)

func NewRepoAuth(collection *mongo.Collection) *RepoAuth {
	return &RepoAuth{collection}
}

func (r *RepoAuth) Register(email, password string) (string, error) {
	var authLogin model.AuthLogin
	filter := bson.D{{"email", email}}
	if err := r.collection.FindOne(context.Background(), filter).Decode(&authLogin); err != nil {
		return "", errExistsEmail
	}

	

	return "", nil
}

func (r *RepoAuth) Login(email, password string) (string, error) {
	var authLogin model.AuthLogin
	filter := bson.D{{"email", email}}
	if err := r.collection.FindOne(context.Background(), filter).Decode(&authLogin); err != nil {
		log.Fatal(err)
	}

	if !utils.CheckPasswordHash(password, authLogin.Password) {
		return "", errInvalidPassword
	}

	token, err := utils.CreateToken(authLogin.Id)
	if err != nil {
		return "", errCreateToken
	}

	return token, nil
}
