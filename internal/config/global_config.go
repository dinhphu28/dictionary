package config

const (
	ConfigFile = "config.toml"
)

type Config struct {
	Version int      `toml:"version"`
	Order   []string `toml:"order"`

	Paths struct {
		Resources string `toml:"resources"`
	} `toml:"paths"`
}
