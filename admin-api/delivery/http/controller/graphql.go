package controller

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gitlab.com/konstellation/kre/admin-api/adapter/gql"

	"gitlab.com/konstellation/kre/admin-api/adapter/config"
	"gitlab.com/konstellation/kre/admin-api/domain/usecase"
	"gitlab.com/konstellation/kre/admin-api/domain/usecase/logging"
)

const UserIDContextKey = "userID"

type GraphQLController struct {
	cfg                    *config.Config
	logger                 logging.Logger
	runtimeInteractor      *usecase.RuntimeInteractor
	userInteractor         *usecase.UserInteractor
	settingInteractor      *usecase.SettingInteractor
	userActivityInteractor *usecase.UserActivityInteractor
	versionInteractor      *usecase.VersionInteractor
	metricsInteractor      *usecase.MetricsInteractor
}

func NewGraphQLController(
	cfg *config.Config,
	logger logging.Logger,
	runtimeInteractor *usecase.RuntimeInteractor,
	userInteractor *usecase.UserInteractor,
	settingInteractor *usecase.SettingInteractor,
	userActivityInteractor *usecase.UserActivityInteractor,
	versionInteractor *usecase.VersionInteractor,
	metricsInteractor *usecase.MetricsInteractor,
) *GraphQLController {
	return &GraphQLController{
		cfg,
		logger,
		runtimeInteractor,
		userInteractor,
		settingInteractor,
		userActivityInteractor,
		versionInteractor,
		metricsInteractor,
	}
}

func (g *GraphQLController) GraphQLHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	g.logger.Info("Request from user " + userID)

	h := gql.NewHttpHandler(
		g.logger,
		g.runtimeInteractor,
		g.userInteractor,
		g.settingInteractor,
		g.userActivityInteractor,
		g.versionInteractor,
		g.metricsInteractor,
	)

	r := c.Request()
	ctx := context.WithValue(r.Context(), UserIDContextKey, userID)

	h.ServeHTTP(c.Response(), r.WithContext(ctx))

	return nil
}

func (g *GraphQLController) PlaygroundHandler(c echo.Context) error {
	h := playground.Handler("GraphQL playground", "/graphql")
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}
