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
	"kego.io/process/generate"
)

type SourceType string

const (
	S_STRUCTS SourceType = "structs"
	S_TYPES   SourceType = "types"
)

func GenerateAll(ctx context.Context, path string, done map[string]bool) error {

	if _, found := done[path]; found {
		return nil
	}

	scache := sysctx.FromContext(ctx)
	pi, ok := scache.Get(path)
	if !ok {
		return kerr.New("XMVXECGDOX", "%s not found in ctx", path)
	}
	if path != "kego.io/system" {
		if err := GenerateAll(ctx, "kego.io/system", done); err != nil {
			return kerr.Wrap("HBKXDVYWUP", err)
		}
	}
	for aliasPath, _ := range pi.Aliases {
		if err := GenerateAll(ctx, aliasPath, done); err != nil {
			return kerr.Wrap("WVXTUBQYVT", err)
		}
	}

	info, found, err := getInfo(ctx, pi.Dir)
	if err != nil {
		return kerr.Wrap("SIMBVNBWOV", err)
	}

	if !found || info.Hash != pi.Hash {
		if err := Generate(ctx, pi.Env); err != nil {
			return kerr.Wrap("TUFKDUPWMD", err)
		}
	}

	done[path] = true

	return nil

}

func getInfo(ctx context.Context, dir string) (info *generate.InfoStruct, found bool, err error) {
	f, err := os.Open(filepath.Join(dir, "generated.go"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, kerr.Wrap("TLFTCRNBKK", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, false, kerr.Wrap("HLDYEAPLEQ", err)
	}

	if !strings.HasPrefix(scanner.Text(), "// info:{") {
		return nil, false, nil
	}

	data := []byte(scanner.Text()[8:])

	var i generate.InfoStruct
	if err := json.UnmarshalPlain(data, &i); err != nil {
		return nil, false, kerr.Wrap("UJXKJVLXHG", err)
	}
	return &i, true, nil
}

// Generate generates the source code for type structs, and writes the generated.go to the
// filesystem.
func Generate(ctx context.Context, env *envctx.Env) error {

	wgctx.Add(ctx, "Generate")
	defer wgctx.Done(ctx, "Generate")

	cmd := cmdctx.FromContext(ctx)

	cmd.Printf("Generating types for %s... ", env.Path)

	outputDir := env.Dir
	filename := "generated.go"
	source, err := generate.Structs(ctx, env)
	if err != nil {
		return kerr.Wrap("XFNESBLBTQ", err)
	}

	// We only backup in the system structs and types files because they are the only
	// generated files we ever need to roll back
	backup := env.Path == "kego.io/system"

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.Wrap("UONJTTSTWW", err)
	} else {
		cmd.Println("OK.")
	}

	return nil
}

func save(dir string, contents []byte, name string, backup bool) error {

	if len(contents) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.Wrap("BPGOUIYPXO", err)
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
		return kerr.Wrap("NWLWHSGJWP", err)
	}
	defer output.Close()

	if _, err := output.Write(contents); err != nil {
		return kerr.Wrap("FBMGPRWQBL", err)
	}

	if err := output.Sync(); err != nil {
		return kerr.Wrap("EGFNTMNKFX", err)
	}

	return nil
}
