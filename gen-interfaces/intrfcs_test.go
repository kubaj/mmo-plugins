package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"io/ioutil"
	"github.com/sergi/go-diff/diffmatchpatch"
	"fmt"
)

type IntrfcsAccessSuite struct {
	suite.Suite
}

func (s *IntrfcsAccessSuite) SetupSuite() {
}

func (s *IntrfcsAccessSuite) SetupTest() {

}

// Test
type testStruct struct {
	Input       string
	Output      string
	Expected    string
	ServiceName string
}

func (s *IntrfcsAccessSuite) TestService() {
	candidates := []testStruct{
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service1_go",
			Expected: "fixtures/service_orig_go",
		},
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service2_go",
			Expected: "fixtures/service_orig_go",
		},
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service3_go",
			Expected: "fixtures/service_orig_go",
		},
	}

	for _, candidate := range candidates {
		data, err := ParseInterfaces("Horse", candidate.Input, candidate.Output, false)
		s.Nil(err)

		serviceContent, err := ioutil.ReadFile(candidate.Output)
		s.Nil(err)

		serviceContentString := string(serviceContent)

		serviceContentString += data

		serviceContentOriginal, err := ioutil.ReadFile(candidate.Expected)
		s.Nil(err)

		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(serviceContentString, string(serviceContentOriginal), true)
		fmt.Println(len(diffs))
		fmt.Println(dmp.DiffPrettyText(diffs))

		s.Equal(serviceContentString, string(serviceContentOriginal))
	}
}

func (s *IntrfcsAccessSuite) TestGatewayMock() {
	candidates := []testStruct{
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service.proto.bp.mock_go",
			Expected: "fixtures/service.proto.bp.mock_orig_go",
		},
		{
			Input:    "fixtures/proto2_go",
			Output:   "fixtures/service2.proto.bp.mock_go",
			Expected: "fixtures/service2.proto.bp.mock_orig_go",
		},
		{
			Input:    "fixtures/proto2_go",
			Output:   "fixtures/service2.proto.bp.mock_go2",
			Expected: "fixtures/service2.proto.bp.mock_orig_go",
		},
	}

	for _, candidate := range candidates {
		data, err := ParseInterfaces("Horse", candidate.Input, candidate.Output, true)
		s.Nil(err)
		serviceContent, err := ioutil.ReadFile(candidate.Output)
		s.Nil(err)

		serviceContentString := string(serviceContent)

		serviceContentString += data

		serviceContentOriginal, err := ioutil.ReadFile(candidate.Expected)
		s.Nil(err)

		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(string(serviceContentOriginal), serviceContentString, false)
		fmt.Println(len(diffs))
		fmt.Println(dmp.DiffPrettyText(diffs))

		s.Equal(string(serviceContentOriginal), (serviceContentString))
	}
}

func TestIntrfcsAccessSuite(t *testing.T) {
	suite.Run(t, new(IntrfcsAccessSuite))
}

func (s *IntrfcsAccessSuite) TearDownSuite() {

}
