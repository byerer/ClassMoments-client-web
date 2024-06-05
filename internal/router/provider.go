package router

import "github.com/google/wire"

var ProviderSetRouter = wire.NewSet(
	NewClassMomentsAPIRouter,
	NewStaticRouter,
)
