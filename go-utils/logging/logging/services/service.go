package services

type controllersClient interface {
	CreateTransfer(code int) error
	CanceledTransfer(code int) error
}

type dummyService struct {
	controllersClient controllersClient
}

func NewDummyService(controllersClient controllersClient) *dummyService {
	return &dummyService{
		controllersClient: controllersClient,
	}
}

func (s *dummyService) CreateTransfer(code int) error {
	return s.controllersClient.CreateTransfer(code)
}

func (s *dummyService) CanceledTransfer(code int) error {
	return s.controllersClient.CanceledTransfer(code)
}
