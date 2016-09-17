package selectors_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/ktest/require"
	"kego.io/json"
	"kego.io/process/parser"
	. "kego.io/process/validate/selectors"
	_ "kego.io/process/validate/selectors/tests"
	"kego.io/system/node"
	"kego.io/tests"
)

// Used for storing the results of the benchmarking tests below
// to avoid compiler optimizations.
var selectorParser *Parser
var values []interface{}

func getTestParser(ctx context.Context, testDocuments map[string]*node.Node, testName string) (*Parser, error) {
	jsonDocument := testDocuments[testName[0:strings.Index(testName, "_")]]
	return CreateParser(ctx, jsonDocument)
}

func runTestsInDirectory(t *testing.T, baseDirectory string) {
	var testDocuments = make(map[string]*node.Node)
	var testSelectors = make(map[string]string)
	var testOutput = make(map[string][]string)

	files, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		t.Error("Error encountered while loading conformance tests ", err)
	}

	cb := tests.Context("kego.io/process/validate/selectors/tests").Jauto().Sauto(parser.Parse)

	for _, fileInfo := range files {
		name := fileInfo.Name()
		if strings.HasSuffix(name, ".json") {
			json_document, err := ioutil.ReadFile(baseDirectory + name)
			if err != nil {
				t.Error("Error encountered while reading ", name, ": ", err)
				continue
			}
			n, err := node.Unmarshal(cb.Ctx(), json_document)
			require.NoError(t, err, name)
			require.NotNil(t, n)

			testDocuments[name[0:len(name)-len(".json")]] = n
		} else if strings.HasSuffix(name, ".output") {
			output_document, err := ioutil.ReadFile(baseDirectory + name)
			if err != nil {
				t.Error("Error encountered while reading ", name, ": ", err)
				continue
			}
			if strings.HasPrefix(string(output_document), "Error") {
				// We won't be handling errors in the same way.
				continue
			}
			var actualOutput []string
			var stringTemporary string
			for _, str := range strings.Split(string(output_document), "\n") {
				stringTemporary = stringTemporary + str
				// Try to parse -- if it works, we have the whole object
				if strings.Index(stringTemporary, "{") == 0 {
					if strings.Count(stringTemporary, "{") != strings.Count(stringTemporary, "}") {
						continue
					}
					actualOutput = append(actualOutput, stringTemporary)
					stringTemporary = ""
				} else if strings.Index(stringTemporary, "[") == 0 {
					if strings.Count(stringTemporary, "[") != strings.Count(stringTemporary, "]") {
						continue
					}
					actualOutput = append(actualOutput, stringTemporary)
					stringTemporary = ""
				} else if len(stringTemporary) > 0 {
					actualOutput = append(actualOutput, stringTemporary)
					stringTemporary = ""
				}
			}
			testOutput[name[0:len(name)-len(".output")]] = actualOutput
		} else if strings.HasSuffix(name, ".selector") {
			selector_document, err := ioutil.ReadFile(baseDirectory + name)
			if err != nil {
				t.Error("Error encountered while reading ", name, ": ", err)
				continue
			}
			testSelectors[name[0:len(name)-len(".selector")]] = string(selector_document)
		}
	}

	for testName := range testOutput {
		single := ""
		if single != "" && testName != single {
			continue
		}
		var passed bool = true
		//t.Log("Running test ", testName)
		parser, err := getTestParser(cb.Ctx(), testDocuments, testName)
		if err != nil {
			t.Error("Test ", testName, "failed: ", err)
			passed = false
		}
		selectorString := testSelectors[testName]
		expectedOutput := testOutput[testName]

		results, err := parser.GetNodes(selectorString)
		if err != nil {
			t.Error("Test ", testName, "failed: ", err)
			passed = false
		}

		//fmt.Println("Output")
		//for i, n := range results {
		//	fmt.Println(i, n.Type.Id.Value(), n.ValueString, n.Key, n.Null)
		//}

		if len(results) != len(expectedOutput) {
			t.Error("Test ", testName, " failed due to number of results being mismatched; ", len(results), " != ", len(expectedOutput), ": [Actual] ", results, " != [Expected] ", expectedOutput)
			passed = false
		} else {
			var expected = make([]interface{}, 0, 10)
			var actual = make([]*node.Node, 0, 10)
			//matchType := "string"

			for idx, result := range results {
				expectedEncoded := expectedOutput[idx]

				var expectedJson interface{}
				err := json.UnmarshalPlain([]byte(expectedEncoded), &expectedJson)
				if err != nil {
					t.Error(
						"Test ", testName, " failed due to a JSON decoding error while decoding expectation: ", err,
					)
					passed = false
				}
				expected = append(expected, expectedJson)
				actual = append(actual, result)
			}

			// Iterate over each of the actual elements; if:
			// * We find a match, remove the match from the expected.
			// * We do not find a match, report an error
			for _, actualElement := range actual {
				var matched bool = false
				for expectedIdx, expectedElement := range expected {
					// TODO: Should we ignore the error here? I guess so...
					matched, _ = comparison(actualElement, expectedElement)
					if matched {
						expected = append(
							expected[:expectedIdx],
							expected[expectedIdx+1:]...,
						)
						break
					}
				}
				if !matched {
					t.Error("Actual element", actualElement, "not found in expected.")
					passed = false
					break
				}
			}
			if len(expected) > 0 {
				for _, value := range expected {
					t.Error("Expected element", value, "not found in actual.")
				}
				passed = false
			}
		}
		if passed {
			//t.Log("Test ", testName, " PASSED")
		} else {
			t.Error("Test ", testName, " FAILED")
		}
	}
}

func comparison(actual *node.Node, expected interface{}) (bool, error) {
	switch actual.JsonType {
	case json.J_NULL:
		// If we're expecting null, return true
		if expected == nil {
			return true, nil
		}
		return false, kerr.New("IKMQTDAKSM", "ns.NativeString returned false")
	case json.J_STRING:
		actualString := actual.ValueString
		expectedString, ok := expected.(string)
		if !ok {
			return false, kerr.New("MPXARMUVPH", "expected %T is not a string", expected)
		}
		return expectedString == actualString, nil
		break
	case json.J_NUMBER:
		actualNumber := actual.ValueNumber
		expectedNumber, ok := expected.(float64)
		if !ok {
			return false, kerr.New("GPOPXFJQSA", "expected %T is not a float64", expected)
		}
		return expectedNumber == actualNumber, nil
		break
	case json.J_BOOL:
		actualBool := actual.ValueBool
		expectedBool, ok := expected.(bool)
		if !ok {
			return false, kerr.New("YEVEBUIQUH", "expected %T is not a bool", expected)
		}
		return expectedBool == actualBool, nil
		break
	case json.J_ARRAY:
		expectedArray, ok := expected.([]interface{})
		if !ok {
			return false, kerr.New("AIHXLBFOQA", "expected %T is not []interface{}", expected)
		}

		length := len(actual.Array)
		expectedLength := len(expectedArray)
		if length != expectedLength {
			return false, nil
		}

		for i, child := range actual.Array {
			match, err := comparison(child, expectedArray[i])
			if err != nil {
				return false, kerr.Wrap("CTHINNYIRI", err)
			}
			if !match {
				return false, nil
			}
		}
		break
	case json.J_MAP:
		expectedMap, ok := expected.(map[string]interface{})
		if !ok {
			return false, kerr.New("CCIRAQFFSR", "expected %T is not map[string]interface{} (map)", expected)
		}
		compareChild := func(key string) (bool, error) {
			child, ok := actual.Map[key]
			if !ok {
				return false, nil
			}
			match, err := comparison(child, expectedMap[key])
			if err != nil {
				return false, kerr.Wrap("QTVTEIETXV", err)
			}
			return match, nil
		}
		for key, _ := range actual.Map {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.Wrap("EAQCUKTFBW", err)
			}
			if !matched {
				return false, nil
			}
		}
		for key, _ := range expectedMap {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.Wrap("YGCQDYMOEA", err)
			}
			if !matched {
				return false, nil
			}
		}
		break
	case "object":
		expectedMap, ok := expected.(map[string]interface{})
		if !ok {
			return false, kerr.New("OJEQQPYXJP", "expected %T is not map[string]interface{} (object)", expected)
		}
		compareChild := func(key string) (bool, error) {
			child, ok := actual.Map[key]
			if !ok {
				return false, nil
			}
			match, err := comparison(child, expectedMap[key])
			if err != nil {
				return false, kerr.Wrap("NNCBWVRAJC", err)
			}
			return match, nil
		}
		for key, field := range actual.Map {
			if field.Missing {
				continue
			}
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.Wrap("DIUQUBQAJN", err)
			}
			if !matched {
				return false, nil
			}
		}
		for key, _ := range expectedMap {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.Wrap("GXBHGSNDFO", err)
			}
			if !matched {
				return false, nil
			}
		}
		break
	}

	return true, nil
}

func TestLevel1(t *testing.T) {
	runTestsInDirectory(t, "./tests/level_1/")
}

func TestLevel2(t *testing.T) {
	runTestsInDirectory(t, "./tests/level_2/")
}

func TestLevel3(t *testing.T) {
	runTestsInDirectory(t, "./tests/level_3/")
}

func TestKego(t *testing.T) {
	runTestsInDirectory(t, "./tests/kego/")
}

func TestGallery(t *testing.T) {
	runTestsInDirectory(t, "./tests/gallery/")
}

func TestExtra(t *testing.T) {
	runTestsInDirectory(t, "./tests/extra/")
}
