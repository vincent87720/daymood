package settings

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Y struct {
	Backend struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Easypick struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Db       string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Trading struct {
		Ajeossi      float32 `yaml:"ajeossi"`
		ShippingFee  float32 `yaml:"shippingFee"`
		ExchangeRate float32 `yaml:"exchangeRate"`
		Tariff       float32 `yaml:"tariff"`
		Markup       float32 `yaml:"markup"`
	}
}

type Settings struct {
	yamlByteXi []byte
	binaryPath string
	yamlPath   string
	debugMode  bool
	y          Y
}

func Init(debug bool) Settings {
	s := Settings{}
	s.debugMode = debug
	s.SetPath()
	s.ReadFile()
	s.UnmarshalSettings()
	return s
}

func (s *Settings) SetPath() {
	if s.debugMode == true {
		s.binaryPath = "."
		s.yamlPath = s.binaryPath + "/settings.development.yaml"
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		s.binaryPath = exPath
		s.yamlPath = s.binaryPath + "/settings.production.yaml"
	}

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
	return s.binaryPath
}

func (s *Settings) GetDebugMode() bool {
	return s.debugMode
}

func (s *Settings) GetBackendAddr() string {
	return s.y.Backend.Host + ":" + s.y.Backend.Port
}

func (s *Settings) GetEasyPickAddr() string {
	return s.y.Easypick.Host + ":" + s.y.Easypick.Port
}

func (s *Settings) GetDBConnectionString() string {
	return "postgresql://" + s.y.Database.Username + ":" + s.y.Database.Password + "@" + s.y.Database.Host + "/" + s.y.Database.Db + "?sslmode=disable"
}

func (s *Settings) GetTradingSettings() (Y, error) {

	return s.y, nil
}

func (s *Settings) SetTradingSettings(y Y) error {
	s.y = y
	return nil
}

func (s *Settings) WriteFile() error {
	s.Print()
	fmt.Println(string(s.yamlByteXi))
	fmt.Println(s.yamlPath)

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
