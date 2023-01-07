package conf

var RedisConfig = &redis{}

var DatabaseConfig = &database{}

var SystemConfig = &system{}

const defaultConf = `[System]
	Debug = true
	Listen = ":10080"
	Secret = "{Secret}"
	HashIDSalt = "{HashIDSalt}"
`
