package middle_ware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shawon1fb/go_api/util"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//gin
func LoginUserFilterMiddleware(ck *util.CuckooFilter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req UserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
			return
		}



		checkUserExist := ck.LookupItem(req.Username)

		if checkUserExist == true {
			ctx.Set("user", req)
			ctx.Next()
			return
		} else {
			err2 := errors.New("user not found")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err2))
			return
		}

	}
}
