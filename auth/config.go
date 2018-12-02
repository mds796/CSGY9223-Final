package auth

import "strconv"

// Config is a configuration struct for the auth service.
type Config struct {
	Host string
	Port uint16

	UserHost string
	UserPort uint16
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) UserTarget() string {
	return c.target(c.UserHost, c.UserPort)
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
