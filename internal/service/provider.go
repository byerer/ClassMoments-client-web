package service

import (
	"ClassMoments-client-web/internal/service/content"
	"ClassMoments-client-web/internal/service/uploader"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"github.com/google/wire"
)

var ProviderSetService = wire.NewSet(
	usercommon.NewUserCommon,
	content.NewUserService,
	uploader.NewUploaderService,
)
