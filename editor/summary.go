package editor

import "kego.io/editor/mdl"

type summary struct {
	*mdl.TableStruct
}

type objectSummary struct {
	*summary
}

type objectSummaryRow struct {
	*mdl.TableRowStruct
}

func newObjectSummaryRow() {

}

func NewObjectSummary(node *Node) *objectSummary {
	s := &summary{TableStruct: mdl.Table()}
	s.Head("name", "origin", "holds", "value", "options")
	return nil
}

type mapSummary struct {
	*summary
}

type arraySummary struct {
	*summary
}
