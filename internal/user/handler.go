package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-go/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (handler *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, handler.GetList)
	router.POST(usersURL, handler.CreateUser)
	router.GET(userURL, handler.GetUserByUUID)
	router.PUT(userURL, handler.UpdateUser)
	router.PATCH(userURL, handler.PatchUser)
	router.DELETE(userURL, handler.DeleteUser)
}

func (handler *handler) GetList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(200)
	writer.Write([]byte("this is list of users"))
}

func (handler *handler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(201)
	writer.Write([]byte("this is create user"))
}

func (handler *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(200)
	writer.Write([]byte("this is user"))
}

func (handler *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(200)
	writer.Write([]byte("this is fully updated user"))
}

func (handler *handler) PatchUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(200)
	writer.Write([]byte("this is patched user"))
}

func (handler *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.WriteHeader(204)
	writer.Write([]byte("this is delete user"))
}
