// Package util contains utility functions to help in code development
package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mathstylish/initzr/metadata/model"
)

func FindIndexByOptionID(options []model.Option, key string) int {
	index := 0
	for idx, opt := range options {
		if opt.ID == key {
			index = idx
			break
		}
	}
	return index
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, inFile); err != nil {
			return err
		}
	}

	return nil
}

func SaveAndUnzip(fileName string, zipContent io.Reader) error {
	tmpFile := filepath.Join(os.TempDir(), fileName+".zip")

	out, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, zipContent); err != nil {
		return fmt.Errorf("failed to write zip content: %w", err)
	}

	outputDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get output directory: %w", err)
	}

	if err := unzip(tmpFile, outputDir); err != nil {
		return fmt.Errorf("failed to unzip: %w", err)
	}

	return nil
}
