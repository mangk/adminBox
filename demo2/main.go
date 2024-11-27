package main

import "github.com/mangk/adminBox/log"

func main() {

	//
	//
	//
	log.Info("fljkds")
	l := log.Trace()

	l.Warn("444")
	l.Warn("555")
	log.Info("fljkds")
	log.Info("fljkds")
	log.Info("fljkds")
	log.Info("fljkds")
	x()
}

func x() {

	l := log.Trace()
	l.Warn("2", "f333", 111)
	l.Warn("3")
	log.Info("4")
	log.Info("5")
	log.Info("6")
	log.Info("7")
	log.Info("7", "1111", 2222)
}
