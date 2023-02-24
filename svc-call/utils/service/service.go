package service

import (
	"fmt"
	"log"
	"sapgorfc/gorfc"
)

type APIService interface {
	SapCall(n string, m interface{}) (interface{}, error)
}

type apiService struct {
	sapconn		*gorfc.Connection
}

func NewService(SAPconnection *gorfc.Connection) APIService {
	return &apiService{
		sapconn: SAPconnection,
	}
}

func (g *apiService) SapCall(fcname string, req interface{}) (interface{}, error) {
	params, ok := req.(map[string]interface{})
	if !ok {
		log.Println("want type map[string]interface{};  got %T", req)
    		return nil, fmt.Errorf("want type map[string]interface{};  got %T", req)
	}
	data, err := g.sapconn.Call(fcname, params)
	if err != nil {
		log.Println("SapCall error: %s", err)
                return nil, err
        }

	return data, nil
}
