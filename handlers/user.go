package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"strings"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(tokenString) // Пропускаем переменную claims
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	collection := client.Database("myapp").Collection("users")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	var users []models.User
	for cur.Next(context.Background()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
