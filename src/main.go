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
	// users API
	users := p.r.Group("/users")
	{
		users.GET("/list", func(c *gin.Context) {
			go p.usersList(c)
		})
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
