package main

import "github.com/snicklin99/flow"

//
// ByteStrom.go
//
// ByteStrom is a product which is used to deploy software across a network in a heirarchical manner.
//

type data struct { // This is where the data will be stored.

	softwareName string
	name         string
	ipAddress    string
	hostname     string
}

func main() {

	d, commenceErr := commence()
	flow.OnErrorQuit(commenceErr, commenceErr.Error())

	coreErr := d.core()
	flow.OnErrorQuit(coreErr, coreErr.Error())

	finishErr := d.finish()
	flow.OnErrorQuit(finishErr, finishErr.Error())

}
