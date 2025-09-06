package layer2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

// decodeYAMLFromReader decodes YAML from an io.Reader into the provided target.
func decodeYAMLFromReader(reader io.Reader, target interface{}) error {
	decoder := yaml.NewDecoder(reader, yaml.DisallowUnknownField())
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("error decoding YAML: %w", err)
	}
	return nil
}

// decodeYAMLFromURL fetches a URL and decodes YAML into the provided target.
func decodeYAMLFromURL(sourceURL string, target interface{}) error {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch URL; response status: %v", resp.Status)
	}
	return decodeYAMLFromReader(resp.Body, target)
}

// decodeYAMLFromFile opens a file and decodes YAML into the provided target.
func decodeYAMLFromFile(filePath string, target interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	return decodeYAMLFromReader(file, target)
}

// loadYaml opens a provided path to unmarshal its data as YAML. It takes a URL or local path to a file as a sourcePath and a pointer to a Catalog object.
func loadYaml(sourcePath string, data *Catalog) error {
	if strings.HasPrefix(sourcePath, "http") {
		return decodeYAMLFromURL(sourcePath, data)
	}
	return decodeYAMLFromFile(sourcePath, data)
}

// loadJson opens a provided path to unmarshal its data as JSON. It takes a URL or local path to a file as a sourcePath and a pointer to a Catalog object.
func loadJson(sourcePath string, data *Catalog) error {
	return fmt.Errorf("loadJson not implemented [%s, %v]", sourcePath, data)
}

// LoadFiles loads data from any number of YAML files at the provided paths. JSON support is pending development.
// If run multiple times, this method will append new data to previous data.
func (c *Catalog) LoadFiles(sourcePaths []string) error {
	for _, sourcePath := range sourcePaths {
		catalog := &Catalog{}
		err := c.LoadFile(sourcePath)
		if err != nil {
			return err
		}
		c.ControlFamilies = append(c.ControlFamilies, catalog.ControlFamilies...)
		c.Capabilities = append(c.Capabilities, catalog.Capabilities...)
		c.Threats = append(c.Threats, catalog.Threats...)
	}
	return nil
}

// LoadFile loads data from a single YAML file at the provided path. JSON support is pending development.
// If run multiple times for the same data type, this method will override previous data.
func (c *Catalog) LoadFile(sourcePath string) error {
	if strings.Contains(sourcePath, ".yaml") || strings.Contains(sourcePath, ".yml") {
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

// LoadNestedCatalog loads a YAML file containing a nested catalog.
// Only supports a single layer of nesting.
// Accepts file URIs with the 'file:///' prefix.
// Throws an error if the URL is not https.
// TODO: Consider validating/sanitizing inputs to reduce injection risks.
func (c *Catalog) LoadNestedCatalog(sourcePath, fieldName string) error {
	if fieldName == "" {
		return fmt.Errorf("fieldName cannot be empty")
	}

	var yamlData map[string]interface{}
	parsedURL, err := url.Parse(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to parse sourcePath: %w", err)
	}

	switch parsedURL.Scheme {
	case "http":
		return fmt.Errorf("insecure http URLs are not allowed: %s", sourcePath)
	case "https":
		err = decodeYAMLFromURL(sourcePath, &yamlData)
		if err != nil {
			return err
		}
	case "file":
		err = decodeYAMLFromFile(parsedURL.Path, &yamlData)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported sourcePath scheme: %s", parsedURL.Scheme)
	}

	fieldData, exists := yamlData[fieldName]
	if !exists {
		return fmt.Errorf("field '%s' not found in YAML file", fieldName)
	}
	fieldYamlBytes, err := yaml.Marshal(fieldData)
	if err != nil {
		return fmt.Errorf("error marshaling field data to YAML: %w", err)
	}
	return decodeYAMLFromReader(strings.NewReader(string(fieldYamlBytes)), c)
}

// ...existing code...
