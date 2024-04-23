package dragonfly

import (
	"github.com/STCraft/dragonfly/server"
	"github.com/sirupsen/logrus"
)

// Server is a global instance of dragonfly
var Server *server.Server

// Logger is an instance of the Logger used for logging stuff
// to the console.
var Logger *logrus.Logger
