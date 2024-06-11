package repo

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/repo/comment"
	"ClassMoments-client-web/internal/repo/like"
	"ClassMoments-client-web/internal/repo/moment"
	"ClassMoments-client-web/internal/repo/notification"
	"ClassMoments-client-web/internal/repo/user"
	"github.com/google/wire"
)

var ProviderSetRepo = wire.NewSet(
	data.NewData,
	data.NewDB,
	user.NewUserRepo,
	like.NewLikeRepo,
	moment.NewMomentRepo,
	comment.NewCommentRepo,
	notification.NewNotificationRepo,
)
