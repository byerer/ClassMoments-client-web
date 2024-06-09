package controller

import "github.com/google/wire"

var ProviderSetController = wire.NewSet(
	NewUserController,
	NewUploadController,
	NewLikeController,
	NewMomentController,
	NewCommentController,
)
