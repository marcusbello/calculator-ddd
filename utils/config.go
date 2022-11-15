package utils

type config struct {
	PostgresConfig
	RedisConfig
	Port string
}

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Db       string
	Port     string
}

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

func GetConfigs() *config {
	// Change to env or viper a json config for production
	//add postgres data here
	pgCred := PostgresConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Db:       "postgres",
		Port:     "5432",
	}
	//add redis config here
	redisCred := RedisConfig{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		Db:       0,
	}
	//add default http router port
	port := ":3330"
	return &config{pgCred, redisCred, port}
}
