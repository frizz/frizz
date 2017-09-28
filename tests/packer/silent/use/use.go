package use

import (
	"time"

	"frizz.io/tests/packer/silent"
)

// TODO: This package highlights why we need the horrible frizz-import kludge. The generated file
// TODO: needs "silent.Package.GetImportedPackages(packages)" but "time.Package.GetImportedPackages(packages)"
// TODO: would fail.
//
// TODO: We read this package using basic AST, so we hve no idea about the content of the silent
// TODO: package vs the time package.
//
// TODO: Maybe we can use go/types and scan all imports. However, this is slow, requires the source
// TODO: of all imports, and we have a problem regenerating when the signature of a type changes
// TODO: and the previously generated source becomes invalid. This will cause go/types to fail.
// TODO: Perhaps there's a way to provide a custom loader that ignores the generated files?

// frizz-import: "frizz.io/tests/packer/silent"

// frizz
type Uses silent.Silent

// frizz
type Time time.Time
