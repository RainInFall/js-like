package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func usage() {
	baseName := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Syntax: %s name parameter...\r\n", baseName)
	flag.PrintDefaults()
	os.Exit(1)
}

func gotemplate(pkg, params string) error {
	gopath := os.Getenv("GOPATH")
	cmd := exec.Command(gopath+"/bin/gotemplate", pkg, params)
	output, err := cmd.CombinedOutput()
	if nil != err {
		log.Printf("%s\r\n", output)
		log.Fatalln("Generate template fail")
		return err
	}
	log.Print(string(output))
	return nil
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("Need at least 2 arguments!")
		os.Exit(2)
	}

	templateName := args[0]
	args = args[1:]
	switch templateName {
	case "array":
		{
			if len(args) != 1 {
				log.Fatal("Syntax: array type")
				os.Exit(2)
			}
			typ := args[0]
			if err := gotemplate("github.com/RainInFall/js-like/array", fmt.Sprintf("Array%s(%s)", typ, typ)); nil != err {
				os.Exit(3)
			}
		}
	case "object":
		{
			if len(args) != 2 {
				log.Fatal("Syntax: object key_type value_type")
				os.Exit(2)
			}

			keyType := args[0]
			valueType := args[1]
			arrayKeyType := fmt.Sprintf("Array%s", keyType)
			arrayValueType := fmt.Sprintf("Array%s", valueType)

			if err := gotemplate("github.com/RainInFall/js-like/array", fmt.Sprintf("%s(%s)", arrayKeyType, keyType)); nil != err {
				os.Exit(3)
			}
			if err := gotemplate("github.com/RainInFall/js-like/array", fmt.Sprintf("%s(%s)", arrayValueType, valueType)); nil != err {
				os.Exit(3)
			}
			if err := gotemplate("github.com/RainInFall/js-like/object",
				fmt.Sprintf("Object%s%s(%s,%s,%s,%s)", keyType, valueType, keyType, valueType, arrayKeyType, arrayValueType)); nil != err {
				os.Exit(3)
			}
		}
	}
}
