package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	services := os.Args[1:]

	for _, service := range services {
		log.Println("Generating service of " + service + " from interface")
		if err := OpenFile("/source/"+service+"/proto.pb.go", "/source/"+service+"/service.go", service, false); err != nil {
			log.Fatal(err)
		}

		log.Println("Generating mock of gateway of " + service + " from interface")
		if err := OpenFile("/source/"+service+"/proto.pb.go", "/source/"+service+"/proto.pb.mock.go", service, true); err != nil {
			log.Fatal(err)
		}

	}
}

func OpenFile(inputPath, outputPath, serviceName string, gateway bool) error {
	// load content of service.go file
	if _, err := ioutil.ReadFile(outputPath); err != nil {
		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		file.WriteString(CreateHeader(serviceName))
	}

	data, err := ParseInterfaces(inputPath, outputPath, gateway)
	if err != nil {
		return err
	}

	// open file to append interfaces
	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}
	return nil
}

func CreateHeader(serviceName string) string {
	return "package " + serviceName + "\n\n" +
		"import (\n" +
		"\t\"golang.org/x/net/context\"\n" +
		"\tgoogle_protobuf \"github.com/golang/protobuf/ptypes/empty\"\n" +
		"\t\"google.golang.org/grpc\"\n" +
		")\n\n" +
		"type " + serviceName + "ServiceClientMock struct {\n" +
		"}\n\n" +
		"func New" + serviceName + "ServiceClientMock() " + serviceName + "ServiceClient {\n" +
		"\treturn &" + serviceName + "ServiceClientMock{}\n" +
		"}\n"
}

var regInterfaces = regexp.MustCompile(`type .[^\s]*Client interface {(\s*.[^\n]*\s*[^\}]*)}`)

func ParseInterfaces(inputPath, outputPath string, gateway bool) (string, error) {
	data := ""

	// load content of proto.pb.go file
	protoContent, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	// load content of service.go file
	serviceContent, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return "", err
	}

	// this regexp find all interfaces
	for _, match := range regInterfaces.FindAllStringSubmatch(string(protoContent), -1) {

		// parse line by line
		for _, lines := range strings.Split(match[1], "\n") {
			result := strings.FieldsFunc(lines, func(r rune) bool {
				return r == '(' || r == ')' || r == ',' || r == '\b' || r == ' ' || r == '\t'
			})

			if len(result) < 9 {
				continue
			}

			// check if is interface already at service file
			conditional := `\s*` + regexp.QuoteMeta(result[0]) +
				`\s*\(\s*` +
				regexp.QuoteMeta(result[1]) +
				`\s*` +
				regexp.QuoteMeta(result[2]) +
				`\s*\,\s*\S*\s*` +
				regexp.QuoteMeta(result[4])

			if gateway {
				conditional += `,\s*opts\s*...grpc.CallOption`
			}
			conditional += `\s*\)\s*\(\s*` +
				regexp.QuoteMeta(result[7]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[8])

			if !regexp.MustCompile(conditional).
				MatchString(string(serviceContent)) {

				log.Println("Adding " + result[0] + "interface to service")

				// added to file new interfaces
				data += "\nfunc (s *Service) " +
					result[0] +
					"(" +
					result[1] +
					" " +
					result[2] +
					", " +
					result[3] +
					" " +
					result[4]

				if gateway {
					data += ", opts ...grpc.CallOption"
				}
				data += ") (" +
					result[7] +
					", " +
					result[8] +
					") {\n" +
					"\n" +
					"\treturn &" + result[7][1:] + "{}, nil\n" +
					"}\n"
			}
		}
	}

	return data, nil
}
