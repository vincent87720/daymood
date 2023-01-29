package settings

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Trading struct {
	Ajeossi      float32 `yaml:"Ajeossi"`
	ShippingFee  float32 `yaml:"ShippingFee"`
	ExchangeRate float32 `yaml:"ExchangeRate"`
	Tariff       float32 `yaml:"Tariff"`
	Markup       float32 `yaml:"Markup"`
	Costs        []struct {
		Key   string  `yaml:"Key"`
		Value float32 `yaml:"Value"`
	} `yaml:"Costs"`
}

type Y struct {
	Trading Trading `yaml:"Trading"`
}

type Settings struct {
	yamlByteXi []byte
	rootPath   string
	yamlPath   string
	mode       string //DEV: development, PROD: production, BIN: binary
	y          Y
	backend    struct {
		host string
		port string
	}
	database struct {
		host     string
		db       string
		username string
		password string
	}
	easypick struct {
		host     string
		port     string
		username string
		password string
	}
	session struct {
		secret string
	}
}

func Init() Settings {
	s := Settings{}
	s.setEnvVar()
	s.setPath()
	s.ReadFile()
	s.UnmarshalSettings()
	return s
}

func (s *Settings) setEnvVar() {
	godotenv.Load()
	s.mode = os.Getenv("APP_MODE")
	s.backend.host = os.Getenv("APP_HOST")
	s.backend.port = os.Getenv("APP_PORT")
	s.database.host = os.Getenv("DB_HOSTNAME")
	s.database.db = os.Getenv("DB_DATABASE")
	s.database.username = os.Getenv("DB_USERNAME")
	s.database.password = os.Getenv("DB_PASSWORD")
	s.easypick.username = os.Getenv("EZSTORE_USERNAME")
	s.easypick.password = os.Getenv("EZSTORE_PASSWORD")
	s.session.secret = os.Getenv("SESSION_SECRET")
}

func (s *Settings) setPath() {
	if s.mode == "DEV" {
		s.rootPath = "."
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		s.rootPath = exPath
	}
	s.yamlPath = s.rootPath + "/settings.yaml"

}

func (s *Settings) ReadFile() {
	ya, err := ioutil.ReadFile(s.yamlPath)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		return
	}
	s.yamlByteXi = ya
}

func (s *Settings) WriteCSV(filepath string, preparecsv [][]string) error {
	// if _, err := os.Stat(s.binaryPath + "/product.csv"); errors.Is(err, os.ErrNotExist) {

	// }
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("\xef\xbb\xbf") //加上BOM頭讓Excel可以正確讀取

	writer := csv.NewWriter(file)

	for _, val := range preparecsv {
		writer.Write(val)
	}
	writer.Flush()
	return nil
}

func (s *Settings) UnmarshalSettings() {

	s.y = Y{}
	err := yaml.Unmarshal(s.yamlByteXi, &s.y)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	// fmt.Printf("settings: %+v\n", s.t)
}

func (s *Settings) Print() {
	fmt.Println(s.y.Trading.Ajeossi)
	fmt.Println(s.y.Trading.ShippingFee)
	fmt.Println(s.y.Trading.ExchangeRate)
	fmt.Println(s.y.Trading.Tariff)
}

func (s *Settings) GetExeFilePath() string {
	return s.rootPath
}

func (s *Settings) GetAppMode() string {
	return s.mode
}

func (s *Settings) GetBackendAddr() string {
	return s.backend.host + ":" + s.backend.port
}

func (s *Settings) GetEasyPickAddr() string {
	return s.easypick.host + ":" + s.easypick.port
}

func (s *Settings) GetDBConnectionString() string {
	return "postgresql://" + s.database.username + ":" + s.database.password + "@" + s.database.host + "/" + s.database.db + "?sslmode=disable"
}

func (s *Settings) GetSessionSecret() string {
	return s.session.secret
}

func (s *Settings) GetTradingSettings() (Trading, error) {
	return s.y.Trading, nil
}

func (s *Settings) SetTradingSettings(trading Trading) error {
	s.y.Trading = trading
	return nil
}

func (s *Settings) WriteFile() error {
	err := ioutil.WriteFile(s.yamlPath, s.yamlByteXi, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) MarshalSettings() error {

	var byteXi []byte
	byteXi, err := yaml.Marshal(s.y)
	if err != nil {
		return err
	}
	s.yamlByteXi = byteXi
	return nil
}
