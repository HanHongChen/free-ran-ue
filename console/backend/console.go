package backend

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Alonza0314/free-ran-ue/logger"
	"github.com/Alonza0314/free-ran-ue/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

type jwt struct {
	secret    string
	expiresIn time.Duration
}

type Console struct {
	router *gin.Engine
	routes util.Routes

	server *http.Server

	username string
	password string

	port int

	jwt

	*logger.ConsoleLogger
}

func NewConsole(config *model.ConsoleConfig, logger *logger.ConsoleLogger) *Console {
	r := &Console{
		router: nil,
		routes: nil,

		username: config.Console.Username,
		password: config.Console.Password,

		port: config.Console.Port,

		jwt: jwt{
			secret:    config.Console.JWT.Secret,
			expiresIn: config.Console.JWT.ExpiresIn,
		},

		ConsoleLogger: logger,
	}

	r.routes = r.initRoutes()
	r.router = util.NewGinRouter(util.CONSOLE_API_PREFIX, r.routes)
	return r
}

func (r *Console) Start() {
	r.ConsoleLog.Infoln("Starting console")

	r.server = &http.Server{
		Addr:    ":" + strconv.Itoa(r.port),
		Handler: r.router,
	}

	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.ConsoleLog.Errorf("Failed to start console: %v", err)
		}
	}()
	time.Sleep(1 * time.Second)

	r.ConsoleLog.Infoln("Console started")
}

func (r *Console) Stop() {
	r.ConsoleLog.Infoln("Stopping console")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := r.server.Shutdown(shutdownCtx); err != nil {
		r.ConsoleLog.Errorf("Failed to stop console: %v", err)
	} else {
		r.ConsoleLog.Infoln("Console stopped successfully")
	}
}

func (r *Console) initRoutes() util.Routes {
	return util.Routes{
		{
			Name:        "Console Login",
			Method:      http.MethodPost,
			Pattern:     "/login",
			HandlerFunc: r.handleConsoleLogin,
		},
		{
			Name:        "Console Logout",
			Method:      http.MethodDelete,
			Pattern:     "/logout",
			HandlerFunc: r.handleConsoleLogout,
		},
	}
}
