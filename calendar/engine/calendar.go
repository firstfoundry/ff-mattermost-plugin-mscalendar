// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package engine

import (
	"time"

	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/remote"
)

type Calendar interface {
	CreateCalendar(user *User, calendar *remote.Calendar) (*remote.Calendar, error)
	CreateEvent(user *User, event *remote.Event, mattermostUserIDs []string) (*remote.Event, error)
	DeleteCalendar(user *User, calendarID string) error
	FindMeetingTimes(user *User, meetingParams *remote.FindMeetingTimesParameters) (*remote.MeetingTimeSuggestionResults, error)
	GetCalendars(user *User) ([]*remote.Calendar, error)
	ViewCalendar(user *User, from, to time.Time) ([]*remote.Event, error)
}

func (m *mscalendar) ViewCalendar(user *User, from, to time.Time) ([]*remote.Event, error) {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return nil, err
	}
	return m.client.GetDefaultCalendarView(user.Remote.ID, from, to)
}

func (m *mscalendar) getTodayCalendarEvents(user *User, now time.Time, timezone string) ([]*remote.Event, error) {
	err := m.Filter(
		withClient,
	)
	if err != nil {
		return nil, err
	}

	err = m.ExpandRemoteUser(user)
	if err != nil {
		return nil, err
	}

	from, to := getTodayHoursForTimezone(now, timezone)
	return m.client.GetDefaultCalendarView(user.Remote.ID, from, to)
}

func (m *mscalendar) excludeDeclinedEvents(events []*remote.Event) (result []*remote.Event) {
	for ix, evt := range events {
		if evt.ResponseStatus == nil || evt.ResponseStatus.Response != remote.EventResponseStatusDeclined {
			result = append(result, events[ix])
		}
	}

	return
}

func (m *mscalendar) CreateCalendar(user *User, calendar *remote.Calendar) (*remote.Calendar, error) {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return nil, err
	}
	return m.client.CreateCalendar(user.Remote.ID, calendar)
}

func (m *mscalendar) CreateEvent(user *User, event *remote.Event, mattermostUserIDs []string) (*remote.Event, error) {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return nil, err
	}

	// invite non-mapped Mattermost
	for id := range mattermostUserIDs {
		mattermostUserID := mattermostUserIDs[id]
		_, err := m.Store.LoadUser(mattermostUserID)
		if err != nil {
			if err.Error() == "not found" {
				_, err = m.Poster.DM(mattermostUserID, "You have been invited to a %s event but have not linked your account.  Feel free to join us by connecting your %s account using `/%s connect`", m.Provider.DisplayName, m.Provider.DisplayName, m.Provider.CommandTrigger)
				if err != nil {
					m.Logger.Warnf("CreateEvent error creating DM. err=%v", err)
					continue
				}
			}
		}
	}

	return m.client.CreateEvent(user.Remote.ID, event)
}

// REVIEW: remove all delete calendar references. dead code/feature
func (m *mscalendar) DeleteCalendar(user *User, calendarID string) error {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return err
	}

	return m.client.DeleteCalendar(user.Remote.ID, calendarID)
}

func (m *mscalendar) FindMeetingTimes(user *User, meetingParams *remote.FindMeetingTimesParameters) (*remote.MeetingTimeSuggestionResults, error) {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return nil, err
	}

	// REVIEW: need to figure out exact usages of this return value, and shorten the scope into a smaller struct, to make it so there are no missed expectations of present data between providers
	return m.client.FindMeetingTimes(user.Remote.ID, meetingParams)
}

func (m *mscalendar) GetCalendars(user *User) ([]*remote.Calendar, error) {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return nil, err
	}

	// REVIEW: same here. need to figure out exact usages of this return value, and shorten the scope into a smaller struct, to make it so there are no missed expectations of present data between providers
	return m.client.GetCalendars(user.Remote.ID)
}
