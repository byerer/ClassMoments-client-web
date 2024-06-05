package repo

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/repo/user"
	"github.com/google/wire"
)

var ProviderSetRepo = wire.NewSet(
	data.NewData,
	data.NewDB,
	user.NewUserRepo,
)
