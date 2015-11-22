package editor

import "kego.io/editor/mdl"

type summary struct {
	*mdl.TableStruct
	path    string
	aliases map[string]string
}

type summaryRow struct {
	*mdl.TableRowStruct
	node  *Node
	table *summary
}

type summaryCell struct {
	*mdl.TableCellStruct
}

type mapSummary struct {
	*summary
}

type arraySummary struct {
	*summary
}
