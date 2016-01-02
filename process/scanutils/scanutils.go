package scanutils // import "kego.io/process/scanutils"

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"

	"golang.org/x/net/context"
	"kego.io/kerr"
)

type File struct {
	File string
	Err  error
}

func ScanDirToFiles(ctx context.Context, dir string, recursive bool) chan File {

	out := make(chan File)

	process := func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("XAAMWOOVFJ", err, "walker (%s)", file)
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		select {
		case out <- File{file, nil}:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	}

	go func() {

		defer close(out)

		if recursive {
			if err := filepath.Walk(dir, process); err != nil {
				out <- File{"", kerr.New("RNHDOWOSSA", err, "filepath.Walk")}
			}
			return
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			out <- File{"", kerr.New("UFYVXCPCTI", err, "ioutil.ReadDir")}
			return
		}
		for _, f := range files {
			if err := process(filepath.Join(dir, f.Name()), f, nil); err != nil {
				out <- File{"", kerr.New("AWXBSBBVWN", err, "walker")}
				return
			}
		}

	}()

	return out
}

type Content struct {
	File  string
	Bytes []byte
	Err   error
}

// ScanFiles takes a chanel of files
func ScanFilesToBytes(ctx context.Context, in chan File) chan Content {

	out := make(chan Content)

	go func() {

		defer close(out)

		for {
			select {
			case value, open := <-in:
				if !open {
					return
				}
				if value.Err != nil {
					out <- Content{"", nil, kerr.New("PQUCOUYLJE", value.Err, "Received error")}
					return
				}
				bytes, err := ProcessFile(value.File)
				// process returns Bytes == nil for non json files, so we should skip them
				if bytes != nil || err != nil {
					out <- Content{value.File, bytes, err}
				}
			case <-ctx.Done():
				out <- Content{"", nil, kerr.New("AFBJCTFOKX", ctx.Err(), "Context done")}
				return
			}
		}

	}()

	return out

}

func ProcessFile(file string) ([]byte, error) {

	isYaml := strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml")
	isJson := strings.HasSuffix(file, ".json")

	if !isYaml && !isJson {
		return nil, nil
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, kerr.New("NMWROTKPLJ", err, "ioutil.ReadFile (%s)", file)
	}

	if isYaml {
		j, err := yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, kerr.New("FAFJCYESRH", err, "yaml.YAMLToJSON")
		}
		bytes = j
	}

	return bytes, nil
}

/*
type NodeError struct {
	Node *node.Node
	Err  error
}

func ScanBytesToNodes(ctx context.Context, in chan ContentError) chan NodeError {

	out := make(chan NodeError)

	process := func(bytes []byte) NodeError {

	}

	go func() {

		defer close(out)

		for {
			select {
			case value, open := <-in:
				if !open {
					return
				}
				if value.Err != nil {
					out <- NodeError{nil, kerr.New("OHEUBTKFQU", value.Err, "Received error")}
					return
				}
				out <- process(value.Bytes)
			case <-ctx.Done():
				out <- ContentError{nil, kerr.New("CEEPFUFHLA", ctx.Err(), "Context done")}
				return
			}
		}

	}()

	return out

}
*/
