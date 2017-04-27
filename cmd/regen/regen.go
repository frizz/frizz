package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"frizz.io/context/envctx"
	"frizz.io/process"

	"strings"

	"github.com/dave/gopackages"
)

func main() {
	packages := []string{
		"frizz.io",
		"github.com/dave/images",
		"github.com/dave/uploader",
	}
	done := map[string]bool{}
	for _, path := range packages {
		if err := regenRecursive(path, done); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func regenRecursive(dir string, done map[string]bool) error {
	dir, err := gopackages.GetDirFromPackage(os.Environ(), os.Getenv("GOPATH"), "frizz.io")
	if err != nil {
		return err
	}
	err = filepath.Walk(dir, filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only process dirs (skip files).
		if !info.IsDir() {
			return nil
		}
		// Does the dir name start with a period? If so, ignore.
		_, dirName := filepath.Split(path)
		if strings.HasPrefix(dirName, ".") {
			return filepath.SkipDir
		}
		// Does the dir contain generated.go?
		if _, err := os.Stat(filepath.Join(path, "generated.go")); err != nil {
			if os.IsNotExist(err) {
				return nil
			} else {
				return err
			}
		}
		if err := regenSingle(path, done); err != nil {
			return err
		}
		return nil
	}))
	if err != nil {
		return err
	}
	return nil
}

func regenSingle(dir string, done map[string]bool) error {
	path, err := gopackages.GetPackageFromDir(os.Getenv("GOPATH"), dir)
	if err != nil {
		return err
	}
	if done[path] {
		return nil
	}
	options := process.Options{
		Log:  true,
		Path: path,
	}
	ctx, _, err := process.Initialise(context.Background(), options)
	if err != nil {
		return err
	}
	env := envctx.FromContext(ctx)

	if err := process.GenerateAll(ctx, env.Path, done); err != nil {
		return err
	}

	return nil
}
