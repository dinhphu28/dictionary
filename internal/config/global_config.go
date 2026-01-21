package config

type GlobalConfig struct {
	Priority []string `json:"priority"`
	DataDir  string   `json:"data_dir"`
}

type Config struct {
	Version  int      `toml:"version"`
	Priority []string `toml:"priority"`

	Paths struct {
		Resources string `toml:"resources"`
	} `toml:"paths"`
}
