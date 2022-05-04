package controllers

import "go.uber.org/dig"

type GatewayController struct {
	dig.In

	UserController *UserController
}
