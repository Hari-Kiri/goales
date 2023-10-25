package main

import (
	"flag"

	"github.com/Hari-Kiri/goales/modules"
	"github.com/Hari-Kiri/temboLog"
)

func main() {
	var arguments []*string
	arguments = append(arguments, flag.String("source-folder", "", "Specify the source folder where You hold Your react code."))
	arguments = append(arguments, flag.String("build-folder", "", "Specify Your build folder to hold react code build version."))
	flag.Parse()

	sourceFolderContents := modules.GetListOfSourceFiles(*arguments[0], *arguments[1], "node_modules")
	getSourceFolderContents, errorGetSourceFolderContents := sourceFolderContents.Get()
	if errorGetSourceFolderContents != nil {
		temboLog.FatalLogging("can't get source folder contents:", errorGetSourceFolderContents)
	}

	temboLog.InfoLogging(getSourceFolderContents)
	temboLog.InfoLogging(*arguments[1])

	// result := api.Build(api.BuildOptions{
	// 	EntryPoints:       []string{"test-ui/source/app.jsx"},
	// 	Bundle:            true,
	// 	MinifySyntax:      true,
	// 	MinifyWhitespace:  true,
	// 	MinifyIdentifiers: true,
	// 	Outfile:           "test-ui/build/index.js",
	// 	Write:             true,
	// })

	// for i := 0; i < len(result.OutputFiles); i++ {
	// 	temboLog.InfoLogging(string(result.OutputFiles[i].Contents))
	// 	temboLog.InfoLogging(result.OutputFiles[i].Path)
	// }

	// if len(result.Errors) != 0 {
	// 	temboLog.ErrorLogging("error: ", result.Errors)
	// 	os.Exit(1)
	// }
}
