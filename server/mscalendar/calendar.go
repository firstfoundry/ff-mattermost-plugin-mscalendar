// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package mscalendar

import (
	"time"

	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/server/remote"
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
				_, err = m.Poster.DM(mattermostUserID, "You have been invited to an Microsoft Outlook calendar event but have not linked your account.  Feel free to join us by connecting your Microsoft Outlook account using `/mscalendar connect`")
				if err != nil {
					m.Logger.Warnf("CreateEvent error creating DM. err=%v", err)
					continue
				}
			}
		}
	}

	return m.client.CreateEvent(user.Remote.ID, event)
}

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

	return m.client.GetCalendars(user.Remote.ID)
}
