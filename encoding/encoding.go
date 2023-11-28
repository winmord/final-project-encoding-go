package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var jsonFileData []byte
	var err error

	if jsonFileData, err = os.ReadFile(j.FileInput); err != nil {
		return fmt.Errorf("JSON file reading fail: %w", err)
	}

	if json.Unmarshal(jsonFileData, &j.DockerCompose) != nil {
		return fmt.Errorf("JSON decoding fail: %w", err)
	}

	var yamlData []byte
	if yamlData, err = yaml.Marshal(&j.DockerCompose); err != nil {
		return fmt.Errorf("YAML encoding fail: %w", err)
	}

	var yamlFile *os.File
	if yamlFile, err = os.Create(j.FileOutput); err != nil {
		return fmt.Errorf("YAML file creating fail: %w", err)
	}

	if _, err = yamlFile.Write(yamlData); err != nil {
		return fmt.Errorf("YAML file writing fail: %w", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var yamlFileData []byte
	var err error

	if yamlFileData, err = os.ReadFile(y.FileInput); err != nil {
		return fmt.Errorf("YAML file reading fail: %w", err)
	}

	if err = yaml.Unmarshal(yamlFileData, &y.DockerCompose); err != nil {
		return fmt.Errorf("YAML decoding fail: %w", err)
	}

	var jsonData []byte
	if jsonData, err = json.Marshal(&y.DockerCompose); err != nil {
		return fmt.Errorf("JSON encoding fail: %w", err)
	}

	var jsonFile *os.File
	if jsonFile, err = os.Create(y.FileOutput); err != nil {
		return fmt.Errorf("JSON file creating fail: %w", err)
	}

	if _, err = jsonFile.Write(jsonData); err != nil {
		return fmt.Errorf("JSON file writing fail: %w", err)
	}

	return nil
}
