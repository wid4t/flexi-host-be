//go:build ignore

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func run() error {
	dir, err := os.MkdirTemp("", "flexi-host-build-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	binary := filepath.Join(dir, "app")
	cmd := exec.Command("go", "build", "-o", binary, ".")
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=arm64", "CGO_ENABLED=0")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build Linux ARM64: %w", err)
	}

	input, err := os.Open(binary)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create("app.zip")
	if err != nil {
		return err
	}
	writer := zip.NewWriter(output)
	header := &zip.FileHeader{Name: "app", Method: zip.Deflate}
	header.SetMode(0755)
	entry, err := writer.CreateHeader(header)
	if err == nil {
		_, err = io.Copy(entry, input)
	}
	if closeErr := writer.Close(); err == nil {
		err = closeErr
	}
	if closeErr := output.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		os.Remove("app.zip")
		return fmt.Errorf("create app.zip: %w", err)
	}
	fmt.Println("Berhasil membuat app.zip (Linux ARM64)")
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
