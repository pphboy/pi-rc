package pirc

type ServiceManager interface {
	StartAll()
	RestartAll()
	ShutdownAll()
}

type Service interface {
	Start()
	Restart()
	Shutdown()
}

type service struct {
}

func (s *service) Start() {

}

func (s *service) Restart() {

}

func (s *service) Shutdown() {

}
