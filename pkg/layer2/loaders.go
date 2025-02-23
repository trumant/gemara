package layer2

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func loadYamlFromURL(sourcePath string, data interface{}) error {
	resp, err := http.Get(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch URL; response status: %v", resp.Status)
	}

	decoder := yaml.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode YAML from URL: %v", err)
	}
	return nil
}

func loadYaml(sourcePath string, data interface{}) error {
	if strings.HasPrefix(sourcePath, "http") {
		return loadYamlFromURL(sourcePath, data)
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	err = decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("error decoding YAML: %w", err)
	}
	return nil
}

func loadJson(sourcePath string, data interface{}) error {
	return fmt.Errorf("loadJson not implemented [%s, %s]", sourcePath, data)
}

func (c *Catalog) LoadControlFamilyFiles(sourcePaths []string) error {
	for _, sourcePath := range sourcePaths {
		catalog := &Catalog{}
		if strings.Contains(sourcePath, ".yaml") {
			err := loadYaml(sourcePath, catalog)
			if err != nil {
				return err
			}
		} else if strings.Contains(sourcePath, ".json") {
			err := loadJson(sourcePath, catalog)
			if err != nil {
				return fmt.Errorf("error loading json: %w", err)
			}
		} else {
			return fmt.Errorf("unsupported file type")
		}
		c.ControlFamilies = append(c.ControlFamilies, catalog.ControlFamilies...)
	}
	return nil
}

// LoadControlFamiliesFile loads multiple control families from a single
// YAML file at the provided path. JSON support is pending development.
func (c *Catalog) LoadControlFamiliesFile(sourcePath string) error {
	if strings.Contains(sourcePath, ".yaml") {
		err := loadYaml(sourcePath, c)
		if err != nil {
			return err
		}
	} else if strings.Contains(sourcePath, ".json") {
		err := loadJson(sourcePath, c)
		if err != nil {
			return fmt.Errorf("error loading json: %w", err)
		}
	} else {
		return fmt.Errorf("unsupported file type")
	}
	return nil
}

// LoadControlFamily loads a single control family from a YAML
// file at the provided path. JSON support is pending development.
func (c *Catalog) LoadControlFamily(sourcePath string) error {
	return c.LoadControlFamilyFiles([]string{sourcePath})
}

func (c *Catalog) LoadThreat(sourcePath string) error {
	return fmt.Errorf("not implemented")
}

func (c *Catalog) LoadCapability(sourcePath string) error {
	return fmt.Errorf("not implemented")
}
