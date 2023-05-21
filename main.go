package main

import "github.com/samber/do"
import "go.uber.org/zap"

func main() {
	container := do.New()
	defer container.Shutdown()

	do.Provide(container, loggerService)
	do.Provide(container, newHelloService)
	run := do.MustInvoke[*helloService](container)

	run.sayHello(container)
}

func loggerService(i *do.Injector) (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	} else {
		return logger, nil
	}
}

type helloService struct{}

func newHelloService(i *do.Injector) (*helloService, error) {
	return &helloService{}, nil
}

func (hs *helloService) sayHello(i *do.Injector) error {
	log := do.MustInvoke[*zap.Logger](i)
	log.Info("hello world", zap.Int("count", 43))
	return nil
}
