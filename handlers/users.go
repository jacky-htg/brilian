package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jacky-htg/brilian/dto"
	"github.com/jacky-htg/brilian/models"
	"github.com/jacky-htg/brilian/pkg/token"
	"github.com/jacky-htg/brilian/repositories"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Db  *sql.DB
	Log *log.Logger
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	switch ctx.Err() {
	case context.Canceled:
		u.Log.Println("Context Cancelled")
		w.WriteHeader(http.StatusInternalServerError)
		return
	case context.DeadlineExceeded:
		u.Log.Println("Deadline Exeeded")
		w.WriteHeader(http.StatusInternalServerError)
		return
	default:
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userRepo := repositories.UserRepository{Db: u.Db, Log: u.Log}
	err = userRepo.Find(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := dto.GetUserResponse{}
	response.FromEntity(userRepo.UserEntity)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	request := dto.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Log.Printf("Error decode request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repo := repositories.UserRepository{Db: u.Db, Log: u.Log, UserEntity: request.ToEntity()}
	pass, err := bcrypt.GenerateFromPassword([]byte(repo.UserEntity.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		http.Error(w, "error generate password", http.StatusInternalServerError)
		return
	}

	repo.UserEntity.Password = string(pass)

	err = repo.Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := dto.GetUserResponse{}
	response.FromEntity(repo.UserEntity)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request := dto.UpdateUserRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Log.Printf("Error decode request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userRepo := repositories.UserRepository{Db: u.Db, Log: u.Log, UserEntity: request.ToEntity()}
	userRepo.UserEntity.Id = id
	err = userRepo.Update(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = userRepo.Find(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := dto.GetUserResponse{}
	response.FromEntity(userRepo.UserEntity)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userRepo := repositories.UserRepository{Db: u.Db, Log: u.Log, UserEntity: models.User{Id: id}}
	err = userRepo.Delete(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	search := r.URL.Query().Get("search")

	userRepo := repositories.UserRepository{Db: u.Db, Log: u.Log}
	list, err := userRepo.List(ctx, search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := []dto.GetUserResponse{}
	for _, v := range list {
		res := dto.GetUserResponse{}
		res.FromEntity(v)
		response = append(response, res)
	}

	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	request := dto.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		u.Log.Printf("Error decode request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userRepo := repositories.UserRepository{Db: u.Db, Log: u.Log, UserEntity: models.User{Email: request.Email}}
	err = userRepo.GetByEmail(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRepo.UserEntity.Password), []byte(request.Password))
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := token.ClaimToken(userRepo.UserEntity.Email)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := dto.LoginResponse{Token: token}
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}
