package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"redis-caching/model"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


// setup dependency injection
type UserHandler struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserHandler(db *gorm.DB, rdb *redis.Client) *UserHandler {
	return &UserHandler{db: db, rdb: rdb}
}

// user handler
func (h *UserHandler) GetUserWithRedis(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	id := r.URL.Path[len("/redis/user/"):]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("user:%s", id)

	cachedUser, err := h.rdb.Get(r.Context(), cacheKey).Result()

	if err == redis.Nil {
		var user model.User

		result := h.db.First(&user, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		userDataBytes, _ := json.Marshal(user)
		h.rdb.Set(r.Context(), cacheKey, userDataBytes, 10*time.Minute)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else if err != nil {
		http.Error(w, "failed to query Redis", http.StatusInternalServerError)
		return
	} else {
		var user model.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err != nil {
			http.Error(w, "failed to unmarshal user data", http.StatusInternalServerError)
			return
		}

		// Respond with the user data from Redis
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}

	duration := time.Since(startTime)

	fmt.Printf("Request getting user detail with Redis: %s \n", duration)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	id := r.URL.Path[len("/user/"):]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	var user model.User

	result := h.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	duration := time.Since(startTime)

	fmt.Printf("Request getting user detail without Redis: %s\n", duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
