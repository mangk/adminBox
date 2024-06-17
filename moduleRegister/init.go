package moduleRegister

import "sync"

var modulesInitChannel chan Module
var once sync.Once

type Module interface {
	InitModule()
}

func ModuleAdd(module Module) {
	if modulesInitChannel == nil {
		modulesInitChannel = make(chan Module, 100)
	}
	modulesInitChannel <- module
}

func ModelInit() {
	once.Do(func() {
		if modulesInitChannel != nil {
			close(modulesInitChannel)
			for {
				v, ok := <-modulesInitChannel
				if !ok {
					break
				}
				v.InitModule()
			}
		}
	})
}
