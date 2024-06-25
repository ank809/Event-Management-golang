package controllers

import (
	"context"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func RegisterEvent(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var user models.RegisterUser = models.RegisterUser{
		ID:          primitive.NewObjectID(),
		Name:        "Jane Doe",
		Email:       "jane.doe@example.com",
		PhoneNumber: "+1234567890",
	}

	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)
	var res models.Event
	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, "NO documents found")
			return
		} else {
			c.JSON(http.StatusNotFound, err)
			return
		}
	} else {
		res.TotalAttendees += 1
		res.Attendees = append(res.Attendees, user)
		update := bson.M{
			"$set": bson.M{
				"totalattendees": res.TotalAttendees,
				"attendees":      res.Attendees,
			},
		}
		_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, "Registered successfully")

}
