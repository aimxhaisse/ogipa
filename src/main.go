package main

// Skeleton to build HTTP/JSON APIs

import (
	"flag"
	"github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"log"
)

// Flag settings
var (
	configFile = flag.String("c", "", "path to the configuration file (e.g, config.json)")
)

// Main structure of the API
type PathwarAPI struct {
	www *gin.Engine        // Gin's engine
	db  *gorethink.Session // Rethinkdb's engine
	cfg *Config            // App's config
}

func (p *PathwarAPI) run() {
	p.www.Run(p.cfg.ListenOn)
	log.Fatalf("error: can't listen on %s", p.cfg.ListenOn)
}

func (p *PathwarAPI) init() error {
	p.www = gin.Default()

	var err error
	p.db, err = gorethink.Connect(gorethink.ConnectOpts{
		Address:  p.cfg.Rethink.Address,
		Database: p.cfg.Rethink.Database,
	})
	if err != nil {
		return err
	}

	// wrapper around handlers that returns an internal error upon failure
	type APIHandler func(*PathwarAPI, *gin.Context) error
	handlerWrapper := func(h APIHandler) func(*gin.Context) {
		return func(c *gin.Context) {
			err := h(p, c)
			if err != nil {
				rsp := Response{RC_INTERNAL_ERROR}
				c.JSON(500, rsp)
				log.Printf("internal error: %v", err)
			}
		}
	}

	// users API
	users := p.www.Group("/users")
	{
		users.GET("/list", handlerWrapper((*PathwarAPI).usersList))
		users.PUT("/new", handlerWrapper((*PathwarAPI).usersNew))
	}

	return nil
}

func NewPathwarAPI(cfg_path string) (*PathwarAPI, error) {
	cfg, err := NewConfig(cfg_path)
	if err != nil {
		return nil, err
	}
	return &PathwarAPI{
		nil,
		nil,
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
