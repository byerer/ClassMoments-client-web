package controller

import (
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentController struct {
	commentService comment.CommentService
}

func NewCommentController(commentService comment.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

func (cc *CommentController) AddComment(ctx *gin.Context) {
	var commentReq schema.CommentReq
	if err := ctx.ShouldBindJSON(&commentReq); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := cc.commentService.AddComment(&commentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
