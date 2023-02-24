package main

import (
	"flag"
	"log"

	"sapgorfc/gorfc"
	"sapgorfc/utils/handler"
	"sapgorfc/utils/service"
	"sapgorfc/utils/config"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

)

func abapSystem(c *config.Config) gorfc.ConnectionParameters {
        return gorfc.ConnectionParameters{
                "user": c.Sapconn.Destination,
                "passwd": c.Sapconn.Password,
                "ashost": c.Sapconn.Ashost,
                "sysnr":  c.Sapconn.Sysnr,
                "client": c.Sapconn.Client,
                "lang": c.Sapconn.Language,
                "dest": c.Sapconn.Destination,
                "saprouter": c.Sapconn.Saprouter,
        }
}

func main() {
	pathJsonFile := flag.String("f", "config.json", "Config json file")
	flag.Parse()

	c := new(config.Config)
	if err := c.Parse(*pathJsonFile); err != nil {
		log.Println(err)
		return
	}

	log.Fatalf("Start Connecting to SAP ...")
	sap, saperr := gorfc.ConnectionFromParams(abapSystem(c))
	if saperr != nil {
        	log.Fatalf("Error SAP connection: %s", saperr)
			//panic(saperr)
        }

	listenAddr := c.Host+":"+c.Port

	r := router.New()
	obj := service.NewService(sap)
	handler.NewHandler(r, obj, c)

	if err := fasthttp.ListenAndServe(listenAddr, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
		panic(err)
	}
	log.Fatalf("Start Server ...")
}
