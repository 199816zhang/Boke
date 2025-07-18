package conf

import "fmt"

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	Debug    bool   `yaml:"debug"`
	Source   string `yaml:"source"`
}

func (d DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.DB)
}
func (d DB) Empty() bool {
	return d.User == "" && d.Password == "" && d.Host == "" && d.Port == 0
}
