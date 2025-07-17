package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`
	DB1    DB     `yaml:"db1"`
}
