package editor

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/mdl"
)

type summary struct {
	*mdl.TableStruct
	ctx context.Context
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
