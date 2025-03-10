package msgraph

import (
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/config"
)

const (
	ProviderMSCalendar            = Kind
	ProviderMSCalendarDisplayName = "Microsoft Calendar"
	ProviderMSCalendarRepository  = "mattermost-plugin-mscalendar"
)

func GetMSCalendarProviderConfig() config.ProviderConfig {
	return config.ProviderConfig{
		Name:        ProviderMSCalendar,
		DisplayName: ProviderMSCalendarDisplayName,
		Repository:  ProviderMSCalendarRepository,

		CommandTrigger: ProviderMSCalendar,

		TelemetryShortName: ProviderMSCalendar,

		BotUsername:    ProviderMSCalendar,
		BotDisplayName: ProviderMSCalendarDisplayName,

		Features: config.ProviderFeatures{
			EncryptedStore:     false,
			EventNotifications: true,
		},
	}
}
