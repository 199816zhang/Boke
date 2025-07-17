package conf

type Jwt struct {
	Expire int64  `yaml:"expire"`
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}
