package settings

type Config struct {
	Addr string
}

type fnConfig func(*Config)

func defaultConfig() *Config {
	return &Config{}
}

func NewConfig(fn ...fnConfig) *Config {
	deflt := defaultConfig()
	for _, fn := range fn {
		fn(deflt)
	}
	return deflt
}

// config functions
func ConfigAddr(addr string) fnConfig {
	return func(cfg *Config) {
		cfg.Addr = addr
	}
}
