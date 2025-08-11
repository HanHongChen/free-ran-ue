package backend

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
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

type console struct {
	router *gin.Engine
	routes util.Routes

	server *http.Server

	username string
	password string

	port int

	jwt

	*logger.ConsoleLogger
}

func NewConsole(config *model.ConsoleConfig, logger *logger.ConsoleLogger) *console {
	c := &console{
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

	c.routes = c.initRoutes()
	c.router = util.NewGinRouter(util.CONSOLE_API_PREFIX, c.routes)

	c.router.NoRoute(c.returnPages())
	return c
}

func (c *console) Start() {
	c.ConsoleLog.Infoln("Starting console")

	c.server = &http.Server{
		Addr:    ":" + strconv.Itoa(c.port),
		Handler: c.router,
	}

	go func() {
		if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.ConsoleLog.Errorf("Failed to start console: %v", err)
		}
	}()
	time.Sleep(1 * time.Second)

	c.ConsoleLog.Infoln("Console started")
}

func (c *console) Stop() {
	c.ConsoleLog.Infoln("Stopping console")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := c.server.Shutdown(shutdownCtx); err != nil {
		c.ConsoleLog.Errorf("Failed to stop console: %v", err)
	} else {
		c.ConsoleLog.Infoln("Console stopped successfully")
	}
}

func (c *console) returnPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodGet {

			destPath := filepath.Join("build/console", c.Request.URL.Path)
			if _, err := os.Stat(destPath); err == nil {
				c.File(filepath.Clean(destPath))
				return
			}

			c.File(filepath.Clean("build/console/index.html"))
		} else {
			c.Next()
		}
	}
}

func (c *console) initRoutes() util.Routes {
	return util.Routes{
		{
			Name:        "Console Login",
			Method:      http.MethodPost,
			Pattern:     "/login",
			HandlerFunc: c.handleConsoleLogin,
		},
		{
			Name:        "Console Logout",
			Method:      http.MethodDelete,
			Pattern:     "/logout",
			HandlerFunc: c.handleConsoleLogout,
		},
	}
}
