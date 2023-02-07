package usecases

import (
	"github.com/vincent87720/daymood/app/internal/settings"
)

type Tradings struct {
	Settings *settings.Settings
}

func NewTradings(settings *settings.Settings) *Tradings {
	return &Tradings{
		Settings: settings,
	}
}

func (tradings *Tradings) Read() (*settings.Trading, error) {
	trading, err := tradings.Settings.GetTradingSettings()
	if err != nil {
		return nil, err
	}
	return &trading, nil
}

func (tradings *Tradings) Update(trading settings.Trading) error {

	err := tradings.Settings.SetTradingSettings(trading)
	if err != nil {
		return err
	}
	return nil
}
