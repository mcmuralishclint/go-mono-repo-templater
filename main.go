package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var dirName string
	var services string
	flag.StringVar(&dirName, "dir", ".", "Directory name for the mono repo (default: current directory)")
	flag.StringVar(&services, "services", "", "Comma-separated names of services")
	flag.Parse()

	if services == "" {
		fmt.Println("Error: At least one service name is required")
		os.Exit(1)
	}

	serviceNames := strings.Split(services, ",")

	err := os.MkdirAll(filepath.Join(dirName, "services"), 0755)
	if err != nil {
		fmt.Println("Error creating services directory:", err)
		os.Exit(1)
	}

	err = os.MkdirAll(filepath.Join(dirName, "pkg"), 0755)
	if err != nil {
		fmt.Println("Error creating pkg directory:", err)
		os.Exit(1)
	}

	makefileContent := `all:
	@echo "Hello, ` + dirName + `!"`
	err = writeFile(filepath.Join(dirName, "Makefile"), makefileContent)
	if err != nil {
		fmt.Println("Error creating Makefile:", err)
		os.Exit(1)
	}

	err = createGoMod(dirName)
	if err != nil {
		fmt.Println("Error creating go.mod file:", err)
		os.Exit(1)
	}

	for _, serviceName := range serviceNames {
		err = createService(filepath.Join(dirName, "services", serviceName))
		if err != nil {
			fmt.Println("Error creating service", serviceName, ":", err)
			os.Exit(1)
		}
	}

	fmt.Println("Mono repo structure created successfully!")
}

func createService(serviceDir string) error {
	cmdDir := filepath.Join(serviceDir, "cmd")
	internalDir := filepath.Join(serviceDir, "internal")
	apiDir := filepath.Join(serviceDir, "api")

	err := os.MkdirAll(cmdDir, 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(internalDir, 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(apiDir, 0755)
	if err != nil {
		return err
	}

	err = writeFile(filepath.Join(cmdDir, "main.go"), "")
	if err != nil {
		return err
	}
	err = writeFile(filepath.Join(internalDir, filepath.Base(serviceDir)+".go"), "")
	if err != nil {
		return err
	}
	err = writeFile(filepath.Join(apiDir, "handler.go"), "")
	if err != nil {
		return err
	}
	err = writeFile(filepath.Join(apiDir, "router.go"), "")
	if err != nil {
		return err
	}
	err = writeFile(filepath.Join(serviceDir, "Dockerfile"), "")
	if err != nil {
		return err
	}

	return nil
}

func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func createGoMod(dirName string) error {
	modName := filepath.Base(dirName)

	err := os.Chdir(dirName)
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	err = exec.Command("go", "mod", "init", modName).Run()
	if err != nil {
		return err
	}
	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		return err
	}
	return nil
}
