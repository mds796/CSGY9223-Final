package post

import "strconv"

// Config is a configuration struct for the post service.
type Config struct {
	Host string
	Port uint16
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
