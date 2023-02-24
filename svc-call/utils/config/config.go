package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

const (
	TITLE		= "TITLE"
	HOST		= "HOST"
	PORT		= "PORT"
	PATH		= "PATH"
	MAXNUMBER	= "MAXNUMBER"
)

type SapConn struct {
	Destination	string	`json:"destination"`
	Client		string  `json:"client"`
	User		string  `json:"user"`
	Password	string  `json:"password"`
	Language	string  `json:"language"`
	Ashost		string  `json:"ashost"`
	Sysnr		string  `json:"sysnr"`
	Saprouter	string  `json:"saprouter"`
}

type RFC struct {
	Name            string  `json:"name"`
}

type Config struct {
	Title		string	`json:"title"`
	Host		string	`json:"host"`
	Port		string  `json:"port"`
	Path		string  `json:"path"`
	Sapconn 	SapConn	`json:"sapcon"`
	Rfc		[]RFC	`json:"rfc"`
	MaxNumber	int64	`json:"max_number"`
}

func (c *Config) Parse(pathJsonFile string) error {
	err := godotenv.Load(".env")
        if err != nil {
                log.Fatalf("Error loading .env file")
        }

	// log.Println("start parsing")
	jsonData, err := ioutil.ReadFile(pathJsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &c); err != nil {
		log.Println(err)
		return err
	}

	title := os.Getenv(TITLE)
	if title != "" {
		c.Title = title
	}

	host := os.Getenv(HOST)
	if host != "" {
		c.Host = host
	}

	port := os.Getenv(PORT)
	if port != "" {
		c.Port = port
	}

	path := os.Getenv(PATH)
        if path != "" {
                c.Path = path
        }

	maxNumber := os.Getenv(MAXNUMBER)
	if maxNumber != "" {

		number, err := strconv.ParseInt(maxNumber, 10, 64)
		if err != nil {
			return err
		}
		c.MaxNumber = number
	}

	log.Println("from env ", *c)
	return nil
}

// ToDo test this func
func getFromEnv(defaultValue, constname string) string {
	value := os.Getenv(constname)
	if value != "" {
		return value
	}
	return defaultValue
}

