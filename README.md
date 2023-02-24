# SAP NW RFC HTTP Connector 

**Written in Golang**

**this is only a fork, but it seems that [origin](https://github.com/SAP/gorfc) is very *static***  

**all credit for gorfc goes to [@bsrdjan](https://github.com/bsrdjan)**  

The **sapgorfc** package provides bindings for **SAP NW RFC Library**, for an easy way of interacting with SAP systems

The goal of this project is to make deployment easier using Docker

Using JSON format for HTTP request and response

## Platforms and Prerequisites

The SAP NW RFC Library is a prerequsite for using the GO RFC connector and must be installed on a same system. It is available on platforms supported by GO, except OSX.

Docker, you can much easier to deploy together with library SDK SAP NW RFC

This Platform can be implemented for Microservices architecture by adding new services more easily

## Install SAPGORFC
### Download or Clone via git
Download from github: https://github.com/trizelka/sapgorfc.git or
git clone https://github.com/trizelka/sapgorfc.git

### Run Docker-Compose
Install docker-compose package

```bash
docker-compose up
```

## Getting Started
### Credential and Access to SAP NW RFC
Edit Parameters SAP Access in config.json

```bash
    destination: IXX
    client: 800
    user: demo
    password: password
    language: EN
    ashost: 11.111.11.111
    sysnr: 00
    saprouter: /H/111.22.333.22/S/2222/W/xxxxx/H/222.22.222.222/H/
```

Edit Allow RFC Name in config.json

```bash
"rfc": [
		{"name": "STFC_STRUCTURE"},
		{"name": "RFC_READ_TABLE"}
    ],
```

### Sample request
url: http://localhost:9090/call

method: POST

Body:

```bash
{
	"fcname":"RFC_READ_TABLE",
	"params":{
		"QUERY_TABLE":"USR01",
		"DELIMITER":";",
		"NO_DATA":"",
		"ROWSKIPS":0,
		"ROWCOUNT":0
	}
}
```

SAP ABAP Code:

```bash
DATA T_DATA TYPE STANDARD TABLE OF TAB512.

CALL FUNCTION 'RFC_READ_TABLE' destination 'destination'
  EXPORTING
     QUERY_TABLE = 'KNB5'
*   DELIMITER                  = ' '
*   NO_DATA                    = ' '
*   ROWSKIPS                   = 0
*   ROWCOUNT                   = 0
  TABLES
*     OPTIONS = ???
*     FIELDS = ???
     DATA = T_DATA
  EXCEPTIONS
     TABLE_NOT_AVAILABLE = 1
     TABLE_WITHOUT_DATA = 2
     OPTION_NOT_VALID = 3
     FIELD_NOT_VALID = 4
     NOT_AUTHORIZED = 5
     DATA_BUFFER_EXCEEDED = 6
    OTHERS = 7.
```

## To Do
* Improve the documentation
* Integration to SSO Credential with Keycloak
* Advance Security with SAP Secure Network Communication (SNC)

## References
The SAP NW Download is [here](https://launchpad.support.sap.com/#/softwarecenter/template/products/%20_APP=00200682500000001943&_EVENT=DISPHIER&HEADER=Y&FUNCTIONBAR=N&EVENT=TREE&NE=NAVIGATE&ENR=01200314690200010197&V=MAINT&TA=ACTUAL&PAGE=SEARCH), but the SAP page is the worst, maybe it's better to search for a torrent or ask a friend at SAP.

If you are SAP employee please check SAP OSS note [1037575 - Software download authorizations for SAP employees](http://service.sap.com/sap/support/notes/1037575).

Docker Image Howto [here]https://devopscube.com/build-docker-image/

SAP with SNC [here]https://help.sap.com/docs/SAP_NETWEAVER_750/e73bba71770e4c0ca5fb2a3c17e8e229/2870ca68118047389852ec53f075f76d.html?version=7.5.21&locale=en-US

## Credits
if you have any question, please don't hesitate to contact me at:

trizelka@gmail.com

or

https://id.linkedin.com/in/trisia-juniarto-5abba0a2