// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package engine

import (
	"github.com/pkg/errors"
)

type EventResponder interface {
	AcceptEvent(user *User, eventID string) error
	DeclineEvent(user *User, eventID string) error
	TentativelyAcceptEvent(user *User, eventID string) error
	RespondToEvent(user *User, eventID, response string) error
}

// REVIEW: See if this is still necessary. I believe RespondToEvent handles all cases already
func (m *mscalendar) AcceptEvent(user *User, eventID string) error {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return err
	}

	return m.client.AcceptEvent(user.Remote.ID, eventID)
}

func (m *mscalendar) DeclineEvent(user *User, eventID string) error {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return err
	}

	return m.client.DeclineEvent(user.Remote.ID, eventID)
}

func (m *mscalendar) TentativelyAcceptEvent(user *User, eventID string) error {
	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return err
	}

	return m.client.TentativelyAcceptEvent(user.Remote.ID, eventID)
}

func (m *mscalendar) RespondToEvent(user *User, eventID, response string) error {
	if response == OptionNotResponded {
		return errors.New("not responded is not a valid response")
	}

	err := m.Filter(
		withClient,
		withUserExpanded(user),
	)
	if err != nil {
		return err
	}

	switch response {
	case OptionYes:
		return m.client.AcceptEvent(user.Remote.ID, eventID)
	case OptionNo:
		return m.client.DeclineEvent(user.Remote.ID, eventID)
	case OptionMaybe:
		return m.client.TentativelyAcceptEvent(user.Remote.ID, eventID)
	default:
		return errors.New(response + " is not a valid response")
	}
}
