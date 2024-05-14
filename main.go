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
	var modPath string
	var services string
	flag.StringVar(&dirName, "dir", "", "Directory name for the mono repo")
	flag.StringVar(&modPath, "mod", "", "Path to the go.mod file")
	flag.StringVar(&services, "services", "", "Comma-separated names of services")
	flag.Parse()

	if dirName == "" {
		fmt.Println("Error: Directory name is required")
		os.Exit(1)
	}
	if modPath == "" {
		fmt.Println("Error: go.mod file path is required")
		os.Exit(1)
	}
	if services == "" {
		fmt.Println("Error: At least one service name is required")
		os.Exit(1)
	}

	serviceNames := strings.Split(services, ",")

	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	err = os.Chdir(dirName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir("services", 0755)
	if err != nil {
		fmt.Println("Error creating services directory:", err)
		os.Exit(1)
	}

	err = os.Mkdir("pkg", 0755)
	if err != nil {
		fmt.Println("Error creating pkg directory:", err)
		os.Exit(1)
	}

	makefileContent := `all:
	@echo "Hello, ` + dirName + `!"`
	err = writeFile("Makefile", makefileContent)
	if err != nil {
		fmt.Println("Error creating Makefile:", err)
		os.Exit(1)
	}

	_, modFileName := filepath.Split(modPath)
	err = exec.Command("cp", modPath, modFileName).Run()
	if err != nil {
		fmt.Println("Error copying go.mod:", err)
		os.Exit(1)
	}

	for _, serviceName := range serviceNames {
		err = createService(serviceName)
		if err != nil {
			fmt.Println("Error creating service", serviceName, ":", err)
			os.Exit(1)
		}
	}

	fmt.Println("Mono repo structure created successfully!")
}

func createService(serviceName string) error {
	serviceDir := filepath.Join("services", serviceName)
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
	err = writeFile(filepath.Join(internalDir, serviceName+".go"), "")
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
