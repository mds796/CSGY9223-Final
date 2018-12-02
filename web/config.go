package web

import (
	"strconv"
)

// Config is a configuration struct for the web service.
type Config struct {
	StaticPath string
	StaticUrl  string

	Host string
	Port uint16

	UserHost string
	UserPort uint16

	AuthHost string
	AuthPort uint16

	FollowHost string
	FollowPort uint16

	PostHost string
	PostPort uint16

	FeedHost string
	FeedPort uint16
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) UserTarget() string {
	return c.target(c.UserHost, c.UserPort)
}

func (c *Config) AuthTarget() string {
	return c.target(c.AuthHost, c.AuthPort)
}

func (c *Config) PostTarget() string {
	return c.target(c.PostHost, c.PostPort)
}

func (c *Config) FollowTarget() string {
	return c.target(c.FollowHost, c.FollowPort)
}

func (c *Config) FeedTarget() string {
	return c.target(c.FeedHost, c.FeedPort)
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
