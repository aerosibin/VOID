package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerCfg struct {
	Port string `yaml:"port"`
}
type DBCfg struct {
	URL string `yaml:"url"`
}
type RedisCfg struct {
	Addr string `yaml:"addr"`
}
type JWTCfg struct {
	Secret string `yaml:"secret"`
}
type PricingCfg struct {
	BasePrice      float64 `yaml:"base_price"`
	DistanceWeight float64 `yaml:"distance_weight"`
	TrafficWeight  float64 `yaml:"traffic_weight"`
	DemandWeight   float64 `yaml:"demand_weight"`
}

type Config struct {
	Server   ServerCfg  `yaml:"server"`
	Database DBCfg      `yaml:"database"`
	Redis    RedisCfg   `yaml:"redis"`
	JWT      JWTCfg     `yaml:"jwt"`
	Pricing  PricingCfg `yaml:"pricing"`
}

func LoadConfig() *Config {
	f, err := os.Open("config/default.yaml")
	if err != nil {
		log.Fatalf("cannot open config file: %v", err)
	}
	defer f.Close()

	cfg := &Config{}
	dec := yaml.NewDecoder(f)
	if err := dec.Decode(cfg); err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	// env overrides
	if v := os.Getenv("PORT"); v != "" {
		cfg.Server.Port = v
	}
	if v := os.Getenv("DATABASE_URL"); v != "" {
		cfg.Database.URL = v
	}
	if v := os.Getenv("REDIS_ADDR"); v != "" {
		cfg.Redis.Addr = v
	}
	if v := os.Getenv("JWT_SECRET"); v != "" {
		cfg.JWT.Secret = v
	}
	return cfg
}
