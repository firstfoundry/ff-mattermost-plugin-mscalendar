package command

import (
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/utils"
)

func (c *Command) showCalendars(_ ...string) (string, bool, error) {
	resp, err := c.Engine.GetCalendars(c.user())
	if err != nil {
		return "", false, err
	}
	return utils.JSONBlock(resp), false, nil
}
