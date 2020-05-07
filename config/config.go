package config

type Configurations struct {
	Server ServerConfigurations
	Cache  CacheConfigurations
}

type ServerConfigurations struct {
	Port int
}

type CacheConfigurations struct {
	TTLMinutes int
}
