package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	services := os.Args[1:]

	for _, service := range services {
		protoContent, err := ioutil.ReadFile("/source/" + service + "/proto.pb.go")
		if err != nil {
			log.Fatal(err)
		}

		var re = regexp.MustCompile("`json:\"-\"`")
		s := re.ReplaceAllString(string(protoContent), "`json:\"-\" sql:\"-\"`")

		ioutil.WriteFile("/source/"+service+"/proto.pb.go", []byte(s), 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
}
