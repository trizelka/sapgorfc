package handler

import (
	"encoding/json"
	"log"
	"time"
	"fmt"
	"sapgorfc/utils/service"
	"sapgorfc/utils/models"
	"sapgorfc/utils/config"
	"sapgorfc/async"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type APIHandler struct {
	APIService service.APIService
	Config	*config.Config
}

func NewHandler(r *router.Router, s service.APIService, c *config.Config) {
	handler := &APIHandler{
		APIService: s,
		Config: c,
	}
	r.POST("/"+c.Path, handler.Method)
	//r.GET("/", Index)
	//r.GET("/hello/{name}", Hello)
}

func (b *APIHandler) Method(ctx *fasthttp.RequestCtx) {
	var m map[string]interface{}
	if err := json.Unmarshal(ctx.PostBody(), &m); err != nil {
		doJSONWrite(ctx, 400, models.ResponseError{Message: err.Error()})
		return
	}

	fcname, fcnameok := m["fcname"].(string)
	if !fcnameok {
                log.Println("Invalid RFC Name")
                doJSONWrite(ctx, 400, models.ResponseError{Message: fmt.Sprintf("Invalid RFC Name")})
                return
        }

	rfcallow, rfcallowok := m["rfc"].([]map[string]interface{})
        if !rfcallowok {
                log.Println("Invalid RFC Allow List")
                doJSONWrite(ctx, 400, models.ResponseError{Message: fmt.Sprintf("Invalid RFC Allow List")})
                return
        }

	allow := false
	for _, rfcValue := range rfcallow {
		if rfcValue["name"].(string) == fcname {
			allow = true
			break
		}
        }
	if allow == false {
                log.Println("RFC not Allowed")
                doJSONWrite(ctx, 400, models.ResponseError{Message: fmt.Sprintf("RFC not Allowed")})
                return
        }

	params, paramserr := m["params"].(interface{})
        if !paramserr {
                log.Println("Invalid RFC Params")
                doJSONWrite(ctx, 400, models.ResponseError{Message: fmt.Sprintf("Invalid RFC Params")})
                return
        }

	var future async.Future
	future = async.Exec(func() (interface{}, error) {
		data, err := b.APIService.SapCall(fcname, params)
		return data, err
	})
	val, err := future.Await()
	if err != nil {
		log.Println("service hit failed: %s\n", err)
		doJSONWrite(ctx, 400, models.ResponseError{Message: err.Error()})
		return
	}

        d, saperr := json.Marshal(val)
	if saperr != nil {
                log.Println("service Marshall failed: %s\n", saperr)
                doJSONWrite(ctx, 400, models.ResponseError{Message: saperr.Error()})
                return
        }

	doJSONWrite(ctx, 200, string(d))
}

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(code)
	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		elapsed := time.Since(start)
		log.Println(elapsed)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
