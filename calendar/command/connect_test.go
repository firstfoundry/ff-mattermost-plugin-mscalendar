package command

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/config"
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/engine"
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/engine/mock_engine"
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/remote"
)

func TestConnect(t *testing.T) {
	tcs := []struct {
		name           string
		command        string
		setup          func(m engine.Engine)
		expectedOutput string
		expectedError  string
	}{
		{
			name:    "user already connected",
			command: "connect",
			setup: func(m engine.Engine) {
				mscal := m.(*mock_engine.MockEngine)
				mscal.EXPECT().GetRemoteUser("user_id").Return(&remote.User{Mail: "user@email.com"}, nil).Times(1)
			},
			expectedOutput: fmt.Sprintf(
				"Your Mattermost account is already connected to %s account `user@email.com`. To connect to a different account, first run `/%s disconnect`.",
				config.Provider.DisplayName, config.Provider.CommandTrigger,
			),
			expectedError: "",
		},
		{
			name:    "user not connected",
			command: "connect",
			setup: func(m engine.Engine) {
				mscal := m.(*mock_engine.MockEngine)
				mscal.EXPECT().GetRemoteUser("user_id").Return(nil, errors.New("remote user not found")).Times(1)
				mscal.EXPECT().Welcome("user_id").Return(nil)
			},
			expectedOutput: "",
			expectedError:  "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			conf := &config.Config{
				PluginURL: "http://localhost",
			}

			mscal := mock_engine.NewMockEngine(ctrl)
			command := Command{
				Context: &plugin.Context{},
				Args: &model.CommandArgs{
					Command: fmt.Sprintf("/%s %s", config.Provider.CommandTrigger, tc.command),
					UserId:  "user_id",
				},
				ChannelID: "channel_id",
				Config:    conf,
				Engine:    mscal,
			}

			if tc.setup != nil {
				tc.setup(mscal)
			}

			out, _, err := command.Handle()
			if tc.expectedOutput != "" {
				require.Equal(t, tc.expectedOutput, out)
			}

			if tc.expectedError != "" {
				require.Equal(t, tc.expectedError, err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
