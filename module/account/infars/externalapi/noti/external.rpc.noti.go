package accountnotirpc

type externalNotiService struct {
	apiURL string
}

func NewNotiRPC(apiURL string) *externalNotiService {
	return &externalNotiService{apiURL: apiURL}
}
