package modules

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// Class interface
type theInterface interface {
	// Get list of source files recursively
	Get() ([]string, error)
}

// Class properties
type theProperties struct {
	// Path of source folder.
	sourceFolder string

	// The build folder contents will be excluded along with the NodeJs module folder contents.
	buildFolder string

	// The NodeJs module folder contents will be excluded along with the build folder contents.
	nodeModulesFolder string
}

// Class module to get list of source files recursively. Property sourceFolderPath should be ReactJs project folder.
// Property buildFolderPath should be ReactJs project build output inside project folder which is will exclude from listing.
// Property nodeModulesFolderPath should be NPM modules folder which is will exclude from listing.
func GetListOfSourceFiles(sourceFolderPath string, buildFolderPath string, nodeModulesFolderPath string) theInterface {
	return theProperties{
		sourceFolder:      sourceFolderPath,
		buildFolder:       buildFolderPath,
		nodeModulesFolder: nodeModulesFolderPath,
	}
}

// Get list of source files recursively.
func (properties theProperties) Get() ([]string, error) {
	// Get target build folder string
	var buildFolderTarget []byte
	for i := 0; i < len(properties.buildFolder); i++ {
		if properties.buildFolder[i] == 47 {
			buildFolderTarget = nil
			continue
		}

		buildFolderTarget = append(buildFolderTarget, properties.buildFolder[i])
	}

	// List folder contents recursively
	return recursiveThroughPath(recursiveThroughPathParameters{
		excludeFolders: []string{string(buildFolderTarget), properties.nodeModulesFolder},
		theProperties:  properties,
	})
}

// Struct for recursiveThroughPath
type recursiveThroughPathParameters struct {
	// Inherit theProperties to get source folder
	theProperties

	// The folders which is excluded from listing
	excludeFolders []string
}

// Check files inside folder and subfolder instead check inside folder listed in excludeFolder parameter.
func recursiveThroughPath(parameters recursiveThroughPathParameters) ([]string, error) {
	var result []string

	// Walk through source folder and subfolder
	errorGetSourceFolderContents := filepath.WalkDir(parameters.sourceFolder,
		func(contents string, fileInfo fs.DirEntry, errorFilePathWalk error) error {
			if errorFilePathWalk != nil {
				return errorFilePathWalk
			}

			isContentsExcludedFolders := stringContainsSlice(stringContainsSliceParameter{
				stringContent:                  contents,
				recursiveThroughPathParameters: parameters,
			})

			if !fileInfo.IsDir() && !isContentsExcludedFolders {
				result = append(result, contents)
			}
			return nil
		},
	)
	if errorGetSourceFolderContents != nil {
		return nil, errorGetSourceFolderContents
	}

	return result, nil
}

// Struct for stringContainsSlice
type stringContainsSliceParameter struct {
	stringContent string
	recursiveThroughPathParameters
}

// Check string content contain any substring listed in excludeFolders.
func stringContainsSlice(parameters stringContainsSliceParameter) bool {
	for i := 0; i < len(parameters.excludeFolders); i++ {
		if strings.Contains(parameters.stringContent, string(parameters.excludeFolders[i])) {
			return true
		}
	}

	return false
}
