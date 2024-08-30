package main

import (
	"archive/zip"
	"fmt"
	"internal/web"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	_ = build_stockfish()
	web.Run()
}

const (
	stockfishURL   = "https://github.com/official-stockfish/Stockfish/archive/refs/heads/master.zip"
	zipFile        = "stockfish.zip"
	buildDir       = "Stockfish-master/src"
	executableName = "stockfish" // The name of the Stockfish executable
)

func build_stockfish() string {
	path, err := ensureStockfish()
	if err != nil {
		fmt.Println("Error ensuring Stockfish:", err)
		return ""
	}
	fmt.Printf("Stockfish is available at: %s\n", path)
	return path
}

func ensureStockfish() (string, error) {
	path, exists := executableExists()
	if exists {
		fmt.Println("Stockfish executable already exists.")
		return path, nil
	}

	if err := downloadStockfish(); err != nil {
		return "", fmt.Errorf("error downloading Stockfish: %w", err)
	}

	if err := unzipStockfish(); err != nil {
		return "", fmt.Errorf("error unzipping Stockfish: %w", err)
	}

	if err := buildStockfish(); err != nil {
		return "", fmt.Errorf("error building Stockfish: %w", err)
	}

	path = filepath.Join(buildDir, executableName)
	if err := os.Chmod(path, 0755); err != nil {
		return "", fmt.Errorf("error making Stockfish executable: %w", err)
	}

	fmt.Println("Stockfish has been successfully downloaded, unzipped, built, and made executable.")
	return path, nil
}

func executableExists() (string, bool) {
	// Check in the current directory
	if _, err := os.Stat(executableName); err == nil {
		absPath, _ := filepath.Abs(executableName)
		return absPath, true
	}

	// Check in the build directory
	buildPath := filepath.Join(buildDir, executableName)
	if _, err := os.Stat(buildPath); err == nil {
		absPath, _ := filepath.Abs(buildPath)
		return absPath, true
	}

	// Check in the PATH
	path, err := exec.LookPath(executableName)
	if err == nil {
		return path, true
	}

	return "", false
}

func downloadStockfish() error {
	resp, err := http.Get(stockfishURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzipStockfish() error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		path := filepath.Join(".", file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func buildStockfish() error {
	cmd := exec.Command("make", "build", "ARCH=x86-64-modern")
	cmd.Dir = buildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
