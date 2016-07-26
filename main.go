package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func usage() {
	baseName := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Syntax: %s name parameter...\r\n", baseName)
	flag.PrintDefaults()
	os.Exit(1)
}

func gotemplate(params ...string) error {
	gopath := os.Getenv("GOPATH")
	cmd := exec.Command(gopath+"/bin/gotemplate", params...)
	output, err := cmd.CombinedOutput()
	if nil != err {
		log.Printf("%s\r\n", output)
		log.Fatalln("Generate template fail")
		return err
	}
	log.Print(string(output))
	return nil
}

func arrayName(typeName string) string {
	return "Array" + escapeName(typeName)
}

func escapeName(typeName string) string {
	if '*' == typeName[0] {
		return strings.Replace(typeName, "*", "p", 1)
	}
	return typeName
}

func generateArray(args []string) {
	if len(args) != 1 {
		log.Fatal("Syntax: array type")
		os.Exit(2)
	}
	typeOrigin := args[0]

	if err := gotemplate("github.com/RainInFall/js-like/array", fmt.Sprintf("%s(%s)", arrayName(typeOrigin), typeOrigin)); nil != err {
		os.Exit(3)
	}
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
			generateArray(args)
		}
	case "array2":
		{
			if len(args) != 2 {
				log.Fatal("Syntax: array2 typeA typeB")
				os.Exit(2)
			}
			typeA := args[0]
			typeB := args[1]
			generateArray(args[0:1])
			generateArray(args[1:2])
			if err := gotemplate(
				"-r", fmt.Sprintf("D=%s", escapeName(typeB)),
				"github.com/RainInFall/js-like/array2",
				fmt.Sprintf("%s(%s,%s,%s,%s)",
					"array"+escapeName(typeA)+escapeName(typeB), typeA, typeB, arrayName(typeA), arrayName(typeB))); nil != err {
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

			generateArray(args[0:1])
			generateArray(args[1:2])
			if err := gotemplate(
				"github.com/RainInFall/js-like/object",
				fmt.Sprintf("Object%s%s(%s,%s,%s,%s)",
					escapeName(keyType), escapeName(valueType), keyType, valueType, arrayName(keyType), arrayName(valueType))); nil != err {
				os.Exit(3)
			}
		}
	}
}
