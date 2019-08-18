package bolthold

import (
	"time"

	"github.com/solutionroute/hoser"
	bh "github.com/timshannon/bolthold"
)

// Client implements a concrete bolthold (bolt db) backed Client interface
type Client struct {
	Path string
	// App         *hoser.App
	Now         func() time.Time
	userService UserService
	db          *bh.Store
}

// NewClient returns a new hoser bolthold client
func NewClient() *Client {
	client := &Client{
		Now: time.Now,
	}
	client.userService.client = client
	return client
}

// Open the database for path
func (c *Client) Open() error {
	db, err := bh.Open(c.Path, 0600, nil)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

// Close ... Closes a boltdb connection
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// UserService returns the client's concrete UserService
func (c *Client) UserService() hoser.UserService {
	return &c.userService
}
