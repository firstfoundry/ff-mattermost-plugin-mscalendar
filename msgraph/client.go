// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package msgraph

import (
	"context"
	"net/http"

	msgraph "github.com/yaegashi/msgraph.go/v1.0"

	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/config"
	"github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/utils/bot"
)

type client struct {
	// caching the context here since it's a "single-use" client, usually used
	// within a single API request
	ctx context.Context

	httpClient *http.Client
	rbuilder   *msgraph.GraphServiceRequestBuilder

	conf *config.Config
	bot.Logger
}
