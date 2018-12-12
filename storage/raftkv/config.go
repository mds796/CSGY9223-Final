package raftkv

import "strconv"

type Config struct {
	NodeID   string
	Host     string
	Port     uint16
	JoinHost string
	JoinPort uint16
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) RaftTarget() string {
	return c.target(c.Host, c.Port+1)
}

func (c *Config) JoinTarget() string {
	return c.target(c.JoinHost, c.JoinPort)
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
