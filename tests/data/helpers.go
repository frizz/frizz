package data

import (
	"testing"

	"github.com/davelondon/ktest/require"
	"kego.io/json"
	"kego.io/process/parser"
	"kego.io/system/node"
	"kego.io/tests"
)

func Setup(t *testing.T) (*tests.ContextBuilder, *node.Node) {

	simple := false
	if !simple {
		cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)
		m := `"type": "multi",
			"js": "js1",
			"ss": "ss1",
			"sr": "sr1",
			"jn": 1.1,
			"sn": 1.2,
			"jb": true,
			"sb": false,
			"alajs": ["alajs0", "alajs1", "alajs2"],
			"alas": [{"type": "simple", "js": "alas0"}, {"type": "simple", "js": "alas1"}],
			"alass": ["alass0", "alass1", "alass2"],
			"aljb": true,
			"aljn": 2.0,
			"aljs": "aljs",
			"almjs": {"a": "almjsa", "b": "almjsb"},
			"alms": {"a": {"type": "simple", "js": "almsa"}, "b": {"type": "simple", "js": "almsb"}},
			"almss": {"a": "almssa", "b": "almssb"},
			"alss": "alss",
			"als": {"type": "als", "js": "als"},
			"i": { "type": "facea", "a": "ia" },
			"ajs": [ "ajs0", "ajs1", "ajs2", "ajs3" ],
			"ass": [ "ass0", "ass1", "ass2", "ass3" ],
			"ajn": [ 2.1, 2.2, 2.3, 2.4 ],
			"asn": [ 3.1, 3.2, 3.3, 3.4 ],
			"ajb": [ true, false, false, false ],
			"asb": [ false, true, false, false ],
			"mjs": { "a": "mjsa", "b": "mjsb" },
			"mss": { "a": "mssa", "b": "mssb" },
			"mjn": { "a": 4.1, "b": 4.2 },
			"msn": { "a": 5.1, "b": 5.2 },
			"mjb": { "a": true, "b": false },
			"msb": { "a": false, "b": true },
			"anri": [ "anri0", { "type": "facea", "a": "anri1" }, { "type": "multi", "ss": "anri2" } ],
			"mnri": { "a": "mnria", "b": { "type": "facea", "a": "mnrib" }, "c": { "type": "multi", "ss": "mnric" } }
						`
		mm := m + `,
			"m": { "type": "multi" },
			"am": [ { "type": "multi", "js": "amjs0" }, { "type": "multi", "js": "amjs1" } ],
			"mm": { "a": { "type": "multi", "js": "mmjsa" }, "b": { "type": "multi", "js": "mmjsb" } }`

		s := `{
				` + m + `,
				"m": {` + mm + `},
				"am": [ {` + mm + `}, {` + mm + `} ],
				"mm": { "a": {` + mm + `}, "b": {` + mm + `} }
			}`
		n := node.NewNode()
		require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
		return cb, n

	} else {
		cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

		m := `"type": "multi",
		"alms": {"a": {"type": "simple", "js": "almsa"}, "b": {"type": "simple", "js": "almsb"}}
			`
		mm := m + `,
			"m": { "type": "multi" },
			"am": [ { "type": "multi", "js": "amjs0" }, { "type": "multi", "js": "amjs1" } ],
			"mm": { "a": { "type": "multi", "js": "mmjsa" }, "b": { "type": "multi", "js": "mmjsb" } }`

		s := `{
		` + m + `,
		"m": {` + mm + `},
		"am": [ {` + mm + `}, {` + mm + `} ],
		"mm": { "a": {` + mm + `}, "b": {` + mm + `} }
	}`
		n := node.NewNode()
		require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
		return cb, n
	}

}

func Empty(t *testing.T) (*tests.ContextBuilder, *node.Node) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	s := `{
	"type": "multi",
	"m": { "type": "multi" },
	"am": [ { "type": "multi" }, { "type": "multi" } ],
	"mm": { "a": { "type": "multi" }, "b": { "type": "multi" } }
}`
	n := node.NewNode()
	require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
	return cb, n
}

func Run(t *testing.T, n *node.Node, m *Multi, test func(t *testing.T, n *node.Node, m *Multi)) {
	test(t, n.Map["mm"].Map["b"], m.Mm["b"])
	test(t, n.Map["mm"].Map["a"], m.Mm["a"])
	test(t, n.Map["am"].Array[1], m.Am[1])
	test(t, n.Map["am"].Array[0], m.Am[0])
	test(t, n.Map["m"], m.M)
	test(t, n, m)
}
