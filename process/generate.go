package process

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/sysctx"
	"kego.io/context/wgctx"
	"kego.io/json"
	"kego.io/kerr"
)

type SourceType string

const (
	S_STRUCTS SourceType = "structs"
	S_TYPES              = "types"
)

func GenerateAll(ctx context.Context, path string, done map[string]bool) error {

	if _, found := done[path]; found {
		return nil
	}

	scache := sysctx.FromContext(ctx)
	pi, ok := scache.Get(path)
	if !ok {
		return kerr.New("XMVXECGDOX", nil, "%s not founc in ctx", path)
	}
	if path != "kego.io/system" {
		if err := GenerateAll(ctx, "kego.io/system", done); err != nil {
			return kerr.New("WVXTUBQYVT", err, "GenerateAll (kego.io/system)")
		}
	}
	for aliasPath, _ := range pi.Environment.Aliases {
		if err := GenerateAll(ctx, aliasPath, done); err != nil {
			return kerr.New("WVXTUBQYVT", err, "GenerateAll (%s)", aliasPath)
		}
	}

	info, found, err := getInfo(ctx, pi.Environment.Dir)
	if err != nil {
		return kerr.New("SIMBVNBWOV", err, "getInfo")
	}

	if !found || info.Hash != pi.Environment.Hash {
		if err := Generate(ctx, pi.Environment); err != nil {
			return kerr.New("TUFKDUPWMD", err, "Generate (%s)", path)
		}
	}

	done[path] = true

	return nil

}

func getInfo(ctx context.Context, dir string) (info *InfoStruct, found bool, err error) {
	f, err := os.Open(filepath.Join(dir, "generated.go"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, kerr.New("TLFTCRNBKK", err, "os.OpenFile")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, false, kerr.New("HLDYEAPLEQ", err, "Scan")
	}

	if !strings.HasPrefix(scanner.Text(), "// info:{") {
		return nil, false, nil
	}

	data := []byte(scanner.Text()[8:])

	var i InfoStruct
	if err := json.UnmarshalPlain(data, &i); err != nil {
		return nil, false, kerr.New("UJXKJVLXHG", err, "DecodeUntyped")
	}
	return &i, true, nil
}

// Generate generates the source code for type structs, and writes the generated.go to the
// filesystem.
func Generate(ctx context.Context, env *envctx.Env) error {

	wgctx.FromContext(ctx).Add(1)
	defer wgctx.FromContext(ctx).Done()

	cmd := cmdctx.FromContext(ctx)

	if cmd.Log {
		fmt.Printf("Generating types for %s... ", env.Path)
	}

	outputDir := env.Dir
	filename := "generated.go"
	source, err := Structs(ctx, env)
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "parse.Structs")
	}

	// We only backup in the system structs and types files because they are the only
	// generated files we ever need to roll back
	backup := env.Path == "kego.io/system"

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "save")
	} else {
		if cmd.Log {
			fmt.Println("OK.")
		}
	}

	return nil
}

func save(dir string, contents []byte, name string, backup bool) error {

	if len(contents) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.New("BPGOUIYPXO", err, "vos.MkdirAll")
		}
	}

	file := filepath.Join(dir, name)

	if backup {
		backupPath := filepath.Join(dir, fmt.Sprintf("%s.backup", name))
		if _, err := os.Stat(backupPath); err == nil {
			os.Remove(backupPath)
		}
		if _, err := os.Stat(file); err == nil {
			os.Rename(file, backupPath)
		}
	} else {
		os.Remove(file)
	}

	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return kerr.New("NWLWHSGJWP", err, "os.OpenFile (could not open output file)")
	}
	defer output.Close()

	if _, err := output.Write(contents); err != nil {
		return kerr.New("FBMGPRWQBL", err, "output.Write")
	}

	if err := output.Sync(); err != nil {
		return kerr.New("EGFNTMNKFX", err, "output.Sync")
	}

	return nil
}
