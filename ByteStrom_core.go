package main

import (
	"errors"

	"github.com/snicklin99/flow"
)

func (d *data) getRelatives() ([]data, []data, error) {

	parents, errParents := d.getParents()

	if errParents != nil {
		return nil, nil, errors.New("error getting parents")
	}

	children, errChildren := d.getChildren()

	if errChildren != nil {
		return nil, nil, errors.New("error getting children")
	}

	if len(parents) == 0 && len(children) == 0 {
		return nil, nil, errors.New("no parents or children")
	}

	return parents, children, nil
}

func (d *data) core() error { // This is where the program will do its work (the main loop)

	keepLooping := true

	parents, children, errRelatives := d.getRelatives()

	if errRelatives != nil {
		return errRelatives
	}

	for keepLooping { // Main loop.

		// Ask the parents if they have anything for me to install

		// Listen to the children to see if they have anything for me to install

		flow.SleepMilliseconds(60000) // Wait a minute!
	}

	return nil
}

// Get the parent nodes
func (d *data) getParents() ([]data, error) {

	return nil, nil
}

// Get the children nodes
func (d *data) getChildren() ([]data, error) {

	return nil, nil
}
