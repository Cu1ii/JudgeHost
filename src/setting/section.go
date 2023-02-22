package setting

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

type JudgeEnvironment struct {
	SubmissionPath string `mapstructure:"submission-path" yaml:"submission-path:"`
	ResolutionPath string `mapstructure:"resolution-path" yaml:"resolution-path:"`
}

type JudgeHostExceptions struct {
	ExceptionCodes map[string]string `mapstructure:"codes" json:"codes" yaml:"codes"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
