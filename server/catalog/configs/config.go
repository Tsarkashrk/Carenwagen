package configs

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        Host     string
        Port     string
        User     string
        Password string
        Name     string
    }
}

func LoadConfig() *Config {
    cfg := &Config{}
    cfg.Server.Port = "8080"
    cfg.Database.Host = "localhost"
    cfg.Database.Port = "5432"
    cfg.Database.User = "postgres"
    cfg.Database.Password = "postgres"
    cfg.Database.Name = "catalog"
    return cfg
}
