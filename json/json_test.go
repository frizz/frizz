package json

import (
	"encoding/json"
	"strconv"
	"testing"

	"kego.io/assert"
)

func TestSimplejson(t *testing.T) {
	var ok bool
	var err error

	js, err := NewJson([]byte(`{ 
		"test": { 
			"string_array": ["asdf", "ghjk", "zxcv"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "json",
			"bool": true 
		}
	}`))

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	_, ok = js.CheckGet("test")
	assert.Equal(t, true, ok)

	_, ok = js.CheckGet("missing_key")
	assert.Equal(t, false, ok)

	aws := js.Get("test").Get("arraywithsubs")
	assert.NotEqual(t, nil, aws)
	var awsval int
	awsval, _ = aws.GetIndex(0).Get("subkeyone").Int()
	assert.Equal(t, 1, awsval)
	awsval, _ = aws.GetIndex(1).Get("subkeytwo").Int()
	assert.Equal(t, 2, awsval)
	awsval, _ = aws.GetIndex(1).Get("subkeythree").Int()
	assert.Equal(t, 3, awsval)

	i, _ := js.Get("test").Get("int").Int()
	assert.Equal(t, 10, i)

	f, _ := js.Get("test").Get("float").Float64()
	assert.Equal(t, 5.150, f)

	s, _ := js.Get("test").Get("string").String()
	assert.Equal(t, "json", s)

	b, _ := js.Get("test").Get("bool").Bool()
	assert.Equal(t, true, b)

	mi := js.Get("test").Get("int").MustInt()
	assert.Equal(t, 10, mi)

	mi2 := js.Get("test").Get("missing_int").MustInt(5150)
	assert.Equal(t, 5150, mi2)

	ms := js.Get("test").Get("string").MustString()
	assert.Equal(t, "json", ms)

	ms2 := js.Get("test").Get("missing_string").MustString("fyea")
	assert.Equal(t, "fyea", ms2)

	ma2 := js.Get("test").Get("missing_array").MustArray([]interface{}{"1", 2, "3"})
	assert.Equal(t, ma2, []interface{}{"1", 2, "3"})

	mm2 := js.Get("test").Get("missing_map").MustMap(map[string]interface{}{"found": false})
	assert.Equal(t, mm2, map[string]interface{}{"found": false})

	strs, err := js.Get("test").Get("string_array").StringArray()
	assert.Equal(t, err, nil)
	assert.Equal(t, strs[0], "asdf")
	assert.Equal(t, strs[1], "ghjk")
	assert.Equal(t, strs[2], "zxcv")

	gp, _ := js.GetPath("test", "string").String()
	assert.Equal(t, "json", gp)

	gp2, _ := js.GetPath("test", "int").Int()
	assert.Equal(t, 10, gp2)

	assert.Equal(t, js.Get("test").Get("bool").MustBool(), true)

	js.Set("float2", 300.0)
	assert.Equal(t, js.Get("float2").MustFloat64(), 300.0)

	js.Set("test2", "setTest")
	assert.Equal(t, "setTest", js.Get("test2").MustString())
}

func TestStdlibInterfaces(t *testing.T) {
	val := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})
	val2 := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})

	raw := `{"name":"myobject","params":{"string":"json"}}`

	assert.Equal(t, nil, json.Unmarshal([]byte(raw), val))

	assert.Equal(t, "myobject", val.Name)
	assert.NotEqual(t, nil, val.Params.data)
	s, _ := val.Params.Get("string").String()
	assert.Equal(t, "json", s)

	p, err := json.Marshal(val)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, json.Unmarshal(p, val2))
	assert.Equal(t, val, val2) // stable
}

func TestSimplejsonGo11(t *testing.T) {
	js, err := NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807
		}
	}`))

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	arr, _ := js.Get("test").Get("array").Array()
	assert.NotEqual(t, nil, arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			assert.Equal(t, nil, err)
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		assert.Equal(t, i+1, iv)
	}

	ma := js.Get("test").Get("array").MustArray()
	assert.Equal(t, ma, []interface{}{json.Number("1"), "2", json.Number("3")})

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	assert.Equal(t, mm, map[string]interface{}{"subkeyone": json.Number("1")})

	assert.Equal(t, js.Get("test").Get("bignum").MustInt64(), int64(9223372036854775807))
}
