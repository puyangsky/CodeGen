package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/puyangsky/codegen"
	"log"
)

func main()  {

	hint := `Usage: main <command>

Commands:
  help     Get help on a command.
  gen      Generate code.
  version  Display version information.

Global Options:
  -t  	   Specify config type, yaml or json
  -f	   Specify the config file
  -o	   Specify the output directory

For more help, you can use 'main help [command]' for the detailed information
or you can check the docs: http://codegen.io/docs/`

	if len(os.Args) == 1 {
		fmt.Println(hint)
		os.Exit(0)
	}

	var filename string
	var genType string
	var output string
	flag.StringVar(&filename,"f", "config.yaml", "config file")
	flag.StringVar(&genType, "t", "yaml", "yaml or json")
	flag.StringVar(&output, "o", ".", "output dir")

	flag.Parse()

	exist, _ := codegen.PathExists(filename)
	if exist {
		fmt.Println("[INFO] FileName:", filename)
	}else {
		log.Fatalln("Invalid file path!")
		os.Exit(-1)
	}

	exist, _ = codegen.PathExists(output)
	if exist {
		fmt.Println("[INFO] Output:", output)
	}else {
		log.Fatalln("Invalid output path!")
		os.Exit(-1)
	}

	if genType != "json" && genType != "yaml" {
		log.Fatalln("Invalid generator type!")
		os.Exit(-1)
	}
	fmt.Println("[INFO] GenType:",genType)

	err := codegen.GenCode(filename, genType, output)
	if err != nil {
		log.Fatalln("Generate code failed!")
		os.Exit(-1)
	}
}
