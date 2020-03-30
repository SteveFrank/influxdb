package config

// MockConfigService mocks the ConfigService.
type MockConfigService struct {
	WriteConfigsFn              func(pp Configs) error
	ParseConfigsFn              func() (Configs, error)
	ParsePreviousActiveConfigFn func() (Config, error)
}

// WriteConfigs returns the write fn.
func (s *MockConfigService) WriteConfigs(pp Configs) error {
	return s.WriteConfigsFn(pp)
}

// ParseConfigs returns the parse fn.
func (s *MockConfigService) ParseConfigs() (Configs, error) {
	return s.ParseConfigsFn()
}

// ParsePreviousActiveConfig returns the parse fn.
func (s *MockConfigService) ParsePreviousActiveConfig() (Config, error) {
	return s.ParsePreviousActiveConfigFn()
}
