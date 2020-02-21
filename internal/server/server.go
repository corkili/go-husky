package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-husky/internal/log"
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
				switch requestMapping.Method {
					case GET:
						ginServer.ginEngine.GET(requestMapping.UrlPath, requestMapping.Handler)
					case POST:
						ginServer.ginEngine.POST(requestMapping.UrlPath, requestMapping.Handler)
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


