package schedule

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonalAvailability struct {
	ID                    string `json:"id"`
	Email                 string `json:"email"`
	Name                  string `json:"name"`
	TimeSlotsAvailability []bool `json:"timeSlotsAvailability"`
}

var MembersAvailability = []PersonalAvailability{
	{ID: "1", Name: "Tom", Email: "tom@foo.com", TimeSlotsAvailability: []bool{true, true, true, true, true}},
	{ID: "2", Name: "Jerry", Email: "jerry@bar.com", TimeSlotsAvailability: []bool{true, true, true, true, false}},
}

func GetMembersAvailability(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, MembersAvailability)
}

func PostMembersAvailability(c *gin.Context) {
	var newMemberAvailability PersonalAvailability

	if err := c.BindJSON(&newMemberAvailability); err != nil {
		return
	}

	MembersAvailability = append(MembersAvailability, newMemberAvailability)
	c.IndentedJSON(http.StatusCreated, newMemberAvailability)
}
