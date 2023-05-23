package main

import (
	"errors"
	"strings"

	"github.com/snicklin99/env"
	"github.com/snicklin99/web"
)

var softwareName string = "ByteStrom" // This is the name of the software.

func commence() (*data, error) { // This is where the program will start.

	d := new(data)
	d.initialise(softwareName)

	return d, nil
}

func (d *data) initialise(software string) error {

	d.softwareName = softwareName
	theHostname, getHostnameErr := web.GetHostname()

	if getHostnameErr != nil {
		return errors.New("error getting hostname")
	}

	d.hostname = theHostname

	theIPAddress, getIPAddressErr := web.GetIp()

	if getIPAddressErr != nil {
		return errors.New("error getting ip address")
	}

	d.ipAddress = theIPAddress

	// Get the name of the system from an environment variable.
	envVarName := strings.ToUpper(softwareName) + "_NAME"
	theEnv, getNameErr := env.Get(envVarName)

	if getNameErr != nil {
		return errors.New("error getting name. Set environment variable: " + envVarName)
	}

	d.name = string(theEnv.GetValue())

	return nil
}
