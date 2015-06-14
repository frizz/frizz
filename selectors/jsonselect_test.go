package jsonselect_test

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"encoding/json"

	"kego.io/assert"
	. "kego.io/jsonselect"
	_ "kego.io/jsonselect/types"
	"kego.io/kego"
	"kego.io/kerr"
	"kego.io/system"
)

// Used for storing the results of the benchmarking tests below
// to avoid compiler optimizations.
var parser *Parser
var values []interface{}

func getTestParser(testDocuments map[string]*Element, testName string) (*Parser, error) {
	jsonDocument := testDocuments[testName[0:strings.Index(testName, "_")]]
	return CreateParser(jsonDocument, "kego.io/jsonselect", map[string]string{})
}

func runTestsInDirectory(t *testing.T, baseDirectory string, path string, imports map[string]string) {
	var testDocuments = make(map[string]*Element)
	var testSelectors = make(map[string]string)
	var testOutput = make(map[string][]string)

	files, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		t.Error("Error encountered while loading conformance tests ", err)
	}

	for _, fileInfo := range files {
		name := fileInfo.Name()
		if strings.HasSuffix(name, ".json") {
			json_document, err := ioutil.ReadFile(baseDirectory + name)
			if err != nil {
				t.Error("Error encountered while reading ", name, ": ", err)
				continue
			}
			var i interface{}
			unknown, err := kego.Unmarshal(json_document, &i, "kego.io/jsonselect", map[string]string{})
			assert.False(t, unknown)
			assert.NoError(t, err)

			typer, ok := i.(system.Typer)
			assert.True(t, ok)

			ty, ok := typer.GetType()

			r := system.NewMinimalRuleHolder(ty, "kego.io/jsonselect", map[string]string{})

			e := &Element{
				Data:  i,
				Value: reflect.ValueOf(i),
				Rule:  r,
			}
			if err != nil {
				t.Error("Error encountered while deserializing ", name, ": ", err)
				continue
			}
			testDocuments[name[0:len(name)-len(".json")]] = e
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
		//if testName != "typed_interface" {
		//	continue
		//}
		var passed bool = true
		//t.Log("Running test ", testName)
		parser, err := getTestParser(testDocuments, testName)
		if err != nil {
			t.Error("Test ", testName, "failed: ", err)
			passed = false
		}
		selectorString := testSelectors[testName]
		expectedOutput := testOutput[testName]

		results, err := parser.GetElements(selectorString)
		if err != nil {
			t.Error("Test ", testName, "failed: ", err)
			passed = false
		}
		if len(results) != len(expectedOutput) {
			t.Error("Test ", testName, " failed due to number of results being mismatched; ", len(results), " != ", len(expectedOutput), ": [Actual] ", results, " != [Expected] ", expectedOutput)
			passed = false
		} else {
			var expected = make([]interface{}, 0, 10)
			var actual = make([]*Element, 0, 10)
			//matchType := "string"

			for idx, result := range results {
				expectedEncoded := expectedOutput[idx]

				var expectedJson interface{}
				err := json.Unmarshal([]byte(expectedEncoded), &expectedJson)
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
					matched, _ = comparison(actualElement, expectedElement, path, imports)
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

func comparison(actual *Element, expected interface{}, path string, imports map[string]string) (bool, error) {
	switch actual.Rule.ParentType.Native.Value {
	case "string":
		ns, ok := actual.Data.(system.NativeString)
		if !ok {
			return false, kerr.New("ENXGPAJVYL", nil, "jsonselect.comparison", "actual.Data %T does not implement system.NativeString", actual.Data)
		}
		actualString, ok := ns.NativeString()
		if !ok {
			// If we're expecting null, return true
			if expected == nil {
				return true, nil
			}
			return false, kerr.New("IKMQTDAKSM", nil, "jsonselect.comparison", "ns.NativeString returned false")
		}
		expectedString, ok := expected.(string)
		if !ok {
			return false, kerr.New("MPXARMUVPH", nil, "jsonselect.comparison", "expected %T is not a string", expected)
		}
		return expectedString == actualString, nil
		break
	case "number":
		nn, ok := actual.Data.(system.NativeNumber)
		if !ok {
			return false, kerr.New("LWVOBVJLAS", nil, "jsonselect.comparison", "actual.Data %T does not implement system.NativeNumber", actual.Data)
		}
		actualNumber, ok := nn.NativeNumber()
		if !ok {
			// If we're expecting null, return true
			if expected == nil {
				return true, nil
			}
			return false, kerr.New("TIESASEDYJ", nil, "jsonselect.comparison", "ns.NativeNumber returned false")
		}
		expectedNumber, ok := expected.(float64)
		if !ok {
			return false, kerr.New("GPOPXFJQSA", nil, "jsonselect.comparison", "expected %T is not a float64", expected)
		}
		return expectedNumber == actualNumber, nil
		break
	case "bool":
		nb, ok := actual.Data.(system.NativeBool)
		if !ok {
			return false, kerr.New("IKIBAMGJSG", nil, "jsonselect.comparison", "actual.Data %T does not implement system.NativeBool", actual.Data)
		}
		actualBool, ok := nb.NativeBool()
		if !ok {
			// If we're expecting null, return true
			if expected == nil {
				return true, nil
			}
			return false, kerr.New("AAWULQRLHA", nil, "jsonselect.comparison", "ns.NativeBool returned false")
		}
		expectedBool, ok := expected.(bool)
		if !ok {
			return false, kerr.New("YEVEBUIQUH", nil, "jsonselect.comparison", "expected %T is not a bool", expected)
		}
		return expectedBool == actualBool, nil
		break
	case "array":
		expectedArray, ok := expected.([]interface{})
		if !ok {
			return false, kerr.New("AIHXLBFOQA", nil, "jsonselect.comparison", "expected %T is not []interface{}", expected)
		}

		length := actual.Value.Len()
		expectedLength := len(expectedArray)
		if length != expectedLength {
			return false, nil
		}

		itemsRule, err := actual.Rule.ItemsRule()
		if err != nil {
			return false, kerr.New("IAOMSWSSGR", err, "jsonselect.comparison", "actual.Rule.ItemsRule (array)")
		}

		for i := 0; i < length; i++ {
			object, _, value, found, _, err := system.GetArrayMember(actual.Value, i)
			if err != nil {
				return false, kerr.New("YLYWUVDEXX", err, "jsonselect.comparison", "system.GetArrayMember")
			}
			if !found {
				return false, nil
			}
			child := &Element{Data: object, Rule: itemsRule, Value: value}
			match, err := comparison(child, expectedArray[i], path, imports)
			if err != nil {
				return false, kerr.New("CTHINNYIRI", err, "jsonselect.comparison", "comparison (array)")
			}
			if !match {
				return false, nil
			}
		}
		break
	case "map":
		expectedMap, ok := expected.(map[string]interface{})
		if !ok {
			return false, kerr.New("CCIRAQFFSR", nil, "jsonselect.comparison", "expected %T is not map[string]interface{} (map)", expected)
		}
		itemsRule, err := actual.Rule.ItemsRule()
		if err != nil {
			return false, kerr.New("XOBKDDGNQR", err, "jsonselect.comparison", "actual.Rule.ItemsRule (map)")
		}
		compareChild := func(key string) (bool, error) {
			object, _, value, found, _, err := system.GetMapMember(actual.Value, key)
			if err != nil {
				return false, kerr.New("LVRSXLXCIJ", err, "jsonselect.comparison", "system.GetMapMember")
			}
			if !found {
				return false, nil
			}
			child := &Element{Data: object, Rule: itemsRule, Value: value}
			match, err := comparison(child, expectedMap[key], path, imports)
			if err != nil {
				return false, kerr.New("QTVTEIETXV", err, "jsonselect.comparison", "getNodes (map)")
			}
			return match, nil
		}
		for _, k := range actual.Value.MapKeys() {
			key, ok := k.Interface().(string)
			if !ok {
				return false, kerr.New("GLUYWFLJTN", nil, "jsonselect.comparison", "Map nodes must be strings, not %T", k.Interface())
			}
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.New("EAQCUKTFBW", err, "jsonselect.comparison", "compareChild map (actual)")
			}
			if !matched {
				return false, nil
			}
		}
		for key, _ := range expectedMap {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.New("YGCQDYMOEA", err, "jsonselect.comparison", "compareChild map (expected)")
			}
			if !matched {
				return false, nil
			}
		}
		break
	case "object":
		expectedMap, ok := expected.(map[string]interface{})
		if !ok {
			return false, kerr.New("OJEQQPYXJP", nil, "jsonselect.comparison", "expected %T is not map[string]interface{} (object)", expected)
		}
		typer, ok := actual.Data.(system.Typer)
		if !ok {
			return false, kerr.New("KQJCVJSTKH", nil, "jsonselect.comparison", "actual %T does not implement system.Typer")
		}
		parentType, ok := typer.GetType()
		if !ok {
			return false, kerr.New("TWVUMUFLST", nil, "jsonselect.comparison", "typer.GetType couldn't find type")
		}
		compareChild := func(key string) (bool, error) {
			object, _, value, found, _, err := system.GetObjectField(actual.Value, system.IdToGoName(key))
			if err != nil {
				return false, kerr.New("KOLTDOJSJY", err, "jsonselect.comparison", "system.GetObjectField")
			}
			if !found {
				return false, nil
			}
			field, ok := parentType.Fields[key]
			if !ok {
				return false, kerr.New("DXRELESKCB", nil, "jsonselect.comparison", "field %s not found in %s", key, actual.Rule.ParentType.Id)
			}
			itemRule, err := system.NewRuleHolder(field, path, imports)
			if err != nil {
				return false, kerr.New("ERPYTUODXO", err, "jsonselect.comparison", "system.NewRuleHolder")
			}
			child := &Element{Data: object, Rule: itemRule, Value: value}
			match, err := comparison(child, expectedMap[key], path, imports)
			if err != nil {
				return false, kerr.New("NNCBWVRAJC", err, "jsonselect.comparison", "comparison (object)")
			}
			return match, nil
		}
		for key, _ := range parentType.Fields {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.New("DIUQUBQAJN", err, "jsonselect.comparison", "compareChild object (actual)")
			}
			if !matched {
				return false, nil
			}
		}
		for key, _ := range expectedMap {
			matched, err := compareChild(key)
			if err != nil {
				return false, kerr.New("GXBHGSNDFO", err, "jsonselect.comparison", "compareChild object (expected)")
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
	runTestsInDirectory(t, "./tests/level_1/", "kego.io/jsonselect", map[string]string{})
}

func TestLevel2(t *testing.T) {
	runTestsInDirectory(t, "./tests/level_2/", "kego.io/jsonselect", map[string]string{})
}

func TestLevel3(t *testing.T) {
	runTestsInDirectory(t, "./tests/level_3/", "kego.io/jsonselect", map[string]string{})
}

func TestKego(t *testing.T) {
	runTestsInDirectory(t, "./tests/kego/", "kego.io/jsonselect", map[string]string{})
}

func TestGallery(t *testing.T) {
	runTestsInDirectory(t, "./tests/gallery/", "kego.io/jsonselect", map[string]string{})
}
