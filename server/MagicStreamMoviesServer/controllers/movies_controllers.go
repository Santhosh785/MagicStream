package controllers

import (
	"context"
	"net/http"
	"time"

	database "github.com/Santhosh785/MagicStream/server/MagicStreamMoviesServer/database"
	models "github.com/Santhosh785/MagicStream/server/MagicStreamMoviesServer/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies", database.Client)

func GetMovies() gin.HandlerFunc { //handles https endpoints
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		movies := make([]models.Movie, 0)

		cursor, err := movieCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})

		}

		c.JSON(http.StatusOK, movies)
	}
}
