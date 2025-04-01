package gomarkdoc_test

import (
	"errors"
	"go/build"
	"os"
	"strings"
	"testing"

	"github.com/frgrisk/gomarkdoc"
	"github.com/frgrisk/gomarkdoc/format/formatcore"
	"github.com/frgrisk/gomarkdoc/lang"
	"github.com/frgrisk/gomarkdoc/logger"
	"github.com/matryer/is"
)

func TestWithTemplateFunc(t *testing.T) {
	is := is.New(t)

	fn, err := loadFunc("./testData/docs", "Func")
	is.NoErr(err)

	r, err := gomarkdoc.NewRenderer()
	is.NoErr(err)

	r2, err := gomarkdoc.NewRenderer(
		gomarkdoc.WithTemplateFunc("escape", func(text string) string {
			return formatcore.Escape(strings.ToUpper(text))
		}),
	)
	is.NoErr(err)

	f, err := r.Func(fn)
	is.NoErr(err)

	f2, err := r2.Func(fn)
	is.NoErr(err)

	is.True(strings.Contains(f, "Func is present in this file."))
	is.True(strings.Contains(f2, "FUNC IS PRESENT IN THIS FILE."))
}

func getBuildPackage(path string) (*build.Package, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return build.Import(path, wd, build.ImportComment)
}

func loadFunc(dir, name string) (*lang.Func, error) {
	buildPkg, err := getBuildPackage(dir)
	if err != nil {
		return nil, err
	}

	log := logger.New(logger.ErrorLevel)
	pkg, err := lang.NewPackageFromBuild(log, buildPkg)
	if err != nil {
		return nil, err
	}

	for _, f := range pkg.Funcs() {
		if f.Name() == name {
			return f, nil
		}
	}

	for _, t := range pkg.Types() {
		for _, f := range t.Funcs() {
			if f.Name() == name {
				return f, nil
			}
		}

		for _, f := range t.Methods() {
			if f.Name() == name {
				return f, nil
			}
		}
	}

	return nil, errors.New("func not found")
}
