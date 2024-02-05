package config

type httpConfig struct {
	Port string `yaml:"port"`
}

type dbConfig struct {
	Url string `yaml:"url"`
}

type Config struct {
	Db   dbConfig   `yaml:"db"`
	Http httpConfig `yaml:"http"`
}
