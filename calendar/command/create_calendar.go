package command

import (
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/remote"
)

func (c *Command) createCalendar(parameters ...string) (string, bool, error) {
	if len(parameters) != 1 {
		return "Please provide the name of one calendar to create", false, nil
	}

	calIn := &remote.Calendar{
		Name: parameters[0],
	}

	_, err := c.Engine.CreateCalendar(c.user(), calIn)
	if err != nil {
		return "", false, err
	}
	return "", false, nil
}
