package scanner // import "kego.io/process/scanner"

// ke: {"package": {"complete": true}}

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"

	"context"

	"github.com/dave/kerr"
)

type File struct {
	File string
	Err  error
}

func ScanDirToFiles(ctx context.Context, dir string, recursive bool) chan File {

	out := make(chan File)

	process := func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return kerr.Wrap("XAAMWOOVFJ", err)
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		select {
		case out <- File{file, nil}:
		case <-ctx.Done():
			return kerr.Wrap("IGULDFRIFY", ctx.Err())
		}
		return nil
	}

	go func() {

		defer close(out)

		if recursive {
			if err := filepath.Walk(dir, process); err != nil {
				out <- File{"", kerr.Wrap("RNHDOWOSSA", err)}
			}
			return
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			out <- File{"", kerr.Wrap("UFYVXCPCTI", err)}
			return
		}
		for _, f := range files {
			if err := process(filepath.Join(dir, f.Name()), f, nil); err != nil {
				out <- File{"", kerr.Wrap("AWXBSBBVWN", err)}
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
					out <- Content{"", nil, kerr.Wrap("PQUCOUYLJE", value.Err)}
					return
				}
				bytes, err := ProcessFile(value.File)
				// process returns Bytes == nil for non json files, so we should skip them
				if bytes != nil || err != nil {
					out <- Content{value.File, bytes, err}
				}
			case <-ctx.Done():
				out <- Content{"", nil, kerr.Wrap("AFBJCTFOKX", ctx.Err())}
				return
			}
		}

	}()

	return out

}

func ProcessFile(file string) ([]byte, error) {

	ext := filepath.Ext(file)
	isYaml := ext == ".yaml" || ext == ".yml"
	isJson := ext == ".json"

	if !isYaml && !isJson {
		return nil, nil
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, kerr.Wrap("NMWROTKPLJ", err)
	}

	if isYaml {
		j, err := yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, kerr.Wrap("FAFJCYESRH", err)
		}
		bytes = j
	}

	return bytes, nil
}
