package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/trantho123/warehouse-management/db/sqlc"
	"github.com/trantho123/warehouse-management/utils"
)

type registerUserRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,passwd"`
	Email    string `json:"email" binding:"required,email"`
}

type registerUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResp(user db.User) registerUserResponse {
	return registerUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
	}
}

func (s *Server) registerUser(ctx *gin.Context) {
	var req registerUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok && len(errs) > 0 {
			ctx.JSON(http.StatusBadRequest, errorResponse(validationErrorMessage(errs[0])))
			return
		}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		RoleID: pgtype.Int4{
			Int32: 1,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}

	newUser, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		switch db.ErrorCode(err) {
		case db.UniqueViolation:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "username/email exited"})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "system error, please try again later"})
			return
		}

	}
	userResp := newUserResp(newUser)
	ctx.JSON(http.StatusOK, userResp)
}
