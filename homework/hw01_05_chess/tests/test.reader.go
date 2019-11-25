package tests

import (
	"fmt"
	"github.com/xdire/algrefresh/util"
	"io/ioutil"
	"testing"
)

type testExecutor func(t *testing.T, input string, compareTo string) error

type testObject struct {
	name		string
	dataIn  	string
	dataOut		string
	fileIn		util.File
	fileOut		util.File
	err 		error
}

func testReader(t *testing.T, files []util.File, executor testExecutor) {
	pattern, err := util.CreatePrefixSuffixNamePattern([]string{"test."}, []string{".in",".out"})
	if err != nil {
		t.Error(err)
	}
	testMap := make(map[string]*testObject)
	combs := util.CombinePattern(files, pattern)
	for k, v := range combs {
		if _, ok := testMap[k]; !ok {
			testMap[k] = &testObject{
				name:    k,
			}
		}
		testObj := testMap[k]
		for _, file := range v.Files {
			if testObj.err != nil {
				continue
			}
			if file.SuffixSelector == ".in" {
				testObj.fileIn = file.File
				data, err := ioutil.ReadFile(file.Path)
				if err != nil {
					testObj.err = err
				}
				testObj.dataIn = string(data)
			}
			if file.SuffixSelector == ".out" {
				testMap[k].fileOut = file.File
				data, err := ioutil.ReadFile(file.Path)
				if err != nil {
					testObj.err = err
				}
				testObj.dataOut = string(data)
			}
		}
	}

	for _, file := range testMap {
		t.Run(fmt.Sprintf("Test for input %s", file.name), func(t *testing.T) {
			err := executor(t, file.dataIn, file.dataOut)
			if err != nil {
				t.Fatalf("Test failed for input %s", file.name)
			}
		})
	}

}
