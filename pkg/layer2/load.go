package layer2

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func loadYaml(sourcePath string, target interface{}) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	return decoder.Decode(target)
}

func loadJson(sourcePath string, target interface{}) error {
	return fmt.Errorf("loadJson not implemented [%s, %s]", sourcePath, target)
}

func (c *Catalog) loadControlFamilyFiles(sourcePaths []string) error {
	for _, sourcePath := range sourcePaths {
		family := &ControlFamily{}
		if strings.Contains(sourcePath, ".yaml") {
			err := loadYaml(sourcePath, family)
			if err != nil {
				return err
			}
		} else if strings.Contains(sourcePath, ".json") {
			err := loadJson(sourcePath, family)
			if err != nil {
				return fmt.Errorf("error loading json: %w", err)
			}
		} else {
			return fmt.Errorf("unsupported file type")
		}
		c.ControlFamilies = append(c.ControlFamilies, *family)
	}
	return nil
}

// LoadControlFamily loads a single control family from a YAML
// file at the provided path. JSON support is pending development.
func (c *Catalog) LoadControlFamily(sourcePath string) error {
	return c.loadControlFamilyFiles([]string{sourcePath})
}

// LoadControlFamilyFiles loads multiple control families from mulitple
// YAML files at the provided paths. JSON support is pending development.
func (c *Catalog) LoadControlFamilyFiles(sourcePaths []string) error {
	return c.loadControlFamilyFiles(sourcePaths)
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

func (c *Catalog) LoadThreat(sourcePath string) error {
	return fmt.Errorf("not implemented")
}

func (c *Catalog) LoadCapability(sourcePath string) error {
	return fmt.Errorf("not implemented")
}
