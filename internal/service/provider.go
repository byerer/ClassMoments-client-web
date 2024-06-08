package service

import (
	"ClassMoments-client-web/internal/service/content"
	"ClassMoments-client-web/internal/service/like"
	"ClassMoments-client-web/internal/service/uploader"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"github.com/google/wire"
)

var ProviderSetService = wire.NewSet(
	usercommon.NewUserCommon,
	like.NewLikeService,
	content.NewUserService,
	uploader.NewUploaderService,
)
