package dragonfly

import (
	"github.com/sirupsen/logrus"
	"github.com/stcraft/dragonfly/server"
)

// Server is a global instance of dragonfly
var Server *server.Server

// Logger is an instance of the Logger used for logging stuff
// to the console.
var Logger *logrus.Logger
