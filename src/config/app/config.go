package app

type Configuration struct {
	App App `map-structure:"app" json:"app" yaml:"app"`
}
