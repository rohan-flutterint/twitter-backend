package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
)

func (u *UserHandler) Unfollow(c *gin.Context) {
	follow := model.User_followers{}
	follow.Follower = c.Value("id").(string)
	follow.Followee = c.Query("follow")

	err := u.userUseCase.UnfollowUser(follow)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User Unfollowed"})
}
