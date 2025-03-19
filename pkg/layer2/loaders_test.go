package layer2

// This file contains table tests for the following functions:
// - loadYaml
// - LoadControlFamily
// - LoadControlFamilyFiles
// - LoadControlFamiliesFile
// - loadYamlFromURL (placeholder, pending a URL to test against)
// - loadJson (placeholder, pending implementation)
// - LoadThreat (placeholder, pending implementation)
// - LoadCapability (placeholder, pending implementation)

// The test data is pulled from ./test-data.yaml

import (
	"testing"
)

var tests = []struct {
	name       string
	sourcePath string
	wantErr    bool
}{
	{
		name:       "Bad path",
		sourcePath: "./bad-path.yaml",
		wantErr:    true,
	},
	{
		name:       "Bad YAML",
		sourcePath: "./test-data/bad.yaml",
		wantErr:    true,
	},
	{
		name:       "Good YAML — CCC",
		sourcePath: "./test-data/good-ccc.yaml",
		wantErr:    false,
	},
	{
		name:       "Good YAML — OSPS",
		sourcePath: "./test-data/good-osps.yaml",
		wantErr:    false,
	},
}

func Test_loadYaml(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Catalog{}
			if err := loadYaml(tt.sourcePath, data); (err == nil) == tt.wantErr {
				t.Errorf("loadYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_LoadFile(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalog{}
			err := c.LoadFile(tt.sourcePath)
			if (err == nil) == tt.wantErr {
				t.Errorf("Catalog.LoadControlFamily() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(c.ControlFamilies) == 0 {
				t.Errorf("Catalog.LoadControlFamily() did not load any control families")
			}
		})
	}
}

func Test_LoadFiles(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalog{}
			err := c.LoadFiles([]string{tt.sourcePath})
			if (err == nil) == tt.wantErr {
				t.Errorf("Catalog.LoadControlFamilyFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(c.ControlFamilies) == 0 {
				t.Errorf("Catalog.LoadControlFamilyFiles() did not load any control families")
			}
		})
	}
}

func Test_loadYamlFromURL(t *testing.T) {
	// Placeholder test
}

func Test_loadJson(t *testing.T) {
	// Placeholder test
}
