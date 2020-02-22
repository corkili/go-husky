package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-husky/internal/api"
	"go-husky/internal/log"
	"net/http"
)

var logger = log.GetLogger()

type GinServer struct {
	ginEngine *gin.Engine
	port uint16
	requestMappings []RequestMapping
}

type Builder struct {
	port uint16
	requestMappings []RequestMapping
	middlewareList []gin.HandlerFunc
	authHandler func(c *gin.Context) (ok bool, rsp *api.Response)
	enableLog bool
	enableRecovery bool
}

type HttpMethod uint8

const (
	GET HttpMethod = 0
	POST HttpMethod = 1
)

type RequestMapping struct {
	UrlPath string
	Method HttpMethod
	Handler gin.HandlerFunc
	Auth bool
}

func (ginServer *GinServer) Start() (err error) {
	return ginServer.ginEngine.Run(fmt.Sprintf(":%d", ginServer.port))
}

func (builder *Builder) SetPort(port uint16) *Builder {
	builder.port = port
	return builder
}

func (builder *Builder) AddRequestMapping(requestMapping ...RequestMapping) *Builder {
	if builder.requestMappings == nil {
		builder.requestMappings = []RequestMapping{}
	}
	builder.requestMappings = append(builder.requestMappings, requestMapping...)
	return builder
}

func (builder *Builder) AddMiddleware(middleware ...gin.HandlerFunc) *Builder {
	if builder.middlewareList == nil {
		builder.middlewareList = []gin.HandlerFunc{}
	}
	builder.middlewareList = append(builder.middlewareList, middleware...)
	return builder
}

func (builder *Builder) SetAuthHandler(handler func(c *gin.Context) (ok bool, rsp *api.Response)) *Builder {
	builder.authHandler = handler
	return builder
}

func (builder *Builder) EnableLog(enableLog bool) *Builder {
	builder.enableLog = enableLog
	return builder
}

func (builder *Builder) EnableRecovery(enableRecovery bool) *Builder {
	builder.enableRecovery = enableRecovery
	return builder
}

func (builder *Builder) Build() *GinServer {
	ginServer := &GinServer{}
	ginServer.port = builder.port
	ginServer.ginEngine = gin.New()

	if builder.enableLog {
		ginServer.ginEngine.Use(gin.Logger())
	}

	if builder.enableRecovery {
		ginServer.ginEngine.Use(gin.Recovery())
	}

	if builder.middlewareList != nil {
		ginServer.ginEngine.Use(builder.middlewareList...)
	}

	ginServer.requestMappings = []RequestMapping{}

	if builder.requestMappings != nil {
		for _, requestMapping := range builder.requestMappings {
			if ginServer.checkRequestMapping(&requestMapping) {
				ginServer.requestMappings = append(ginServer.requestMappings, requestMapping)
				handler := buildRequestHandler(requestMapping.Handler, requestMapping.Auth, builder.authHandler)
				switch requestMapping.Method {
					case GET:
						ginServer.ginEngine.GET(requestMapping.UrlPath, handler)
					case POST:
						ginServer.ginEngine.POST(requestMapping.UrlPath, handler)
				}
			}
		}
	}

	return ginServer
}

func (ginServer *GinServer) checkRequestMapping(requestMapping *RequestMapping) bool {
	if requestMapping.UrlPath != "" && requestMapping.Method <= POST && requestMapping.Handler != nil {
		for _, rm := range ginServer.requestMappings {
			if rm.UrlPath == requestMapping.UrlPath && rm.Method == requestMapping.Method {
				logger.Error("requestMapping is exist")
				return false
			}
		}
		return true
	} else {
		logger.Error("requestMapping is invalid")
		return false
	}
}

func buildRequestHandler(handlerFunc gin.HandlerFunc, auth bool, authHandler func(c *gin.Context) (ok bool, rsp *api.Response)) gin.HandlerFunc {
	if auth && authHandler != nil {
		return func(c *gin.Context) {
			ok, rsp := authHandler(c)
			if ok {
				handlerFunc(c)
			} else {
				c.JSON(http.StatusOK, rsp)
			}
		}
	} else {
		return handlerFunc
	}
}


