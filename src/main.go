package main

// Skeleton to build HTTP/JSON APIs

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

// Flag settings
var (
	configFile = flag.String("c", "", "path to the configuration file (e.g, config.json)")
)

func (p *PathwarAPI) run() {
	p.r.Run(p.cfg.ListenOn)
}

func (p *PathwarAPI) init() error {

	// wrapper around handlers that returns an internal error upon failure
	type APIHandler func(*PathwarAPI, *gin.Context) error
	handlerWrapper := func(h APIHandler) func(*gin.Context) {
		return func(c *gin.Context) {
			err := h(p, c)
			if err != nil {
				rsp := Response{RC_INTERNAL_ERROR}
				c.JSON(500, rsp)
			}
		}
	}

	// users API
	users := p.r.Group("/users")
	{
		users.GET("/list", handlerWrapper((*PathwarAPI).usersList))
		users.PUT("/new", handlerWrapper((*PathwarAPI).usersNew))
	}

	// add your own resources

	return nil
}

func NewPathwarAPI(cfg_path string) (*PathwarAPI, error) {
	cfg, err := NewConfig(cfg_path)
	if err != nil {
		return nil, err
	}
	return &PathwarAPI{
		gin.Default(),
		cfg,
	}, nil
}

func main() {
	flag.Parse()
	app, err := NewPathwarAPI(*configFile)
	if err != nil {
		log.Fatalf("failed to init PathwarAPI: %v", err)
	}
	err = app.init()
	if err != nil {
		log.Fatalf("failed to init PathwarAPI: %v", err)
	}
	app.run()
}
