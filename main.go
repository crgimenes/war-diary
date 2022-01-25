package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	addText := flag.String("a", "", "add log")

	flag.Parse()

	fmt.Println("add:", *addText)

	f, err := os.CreateTemp("", "*.md")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	fmt.Println("file:", f.Name())

	date := time.Now().Format("2006-01-02 15:04:05")

	header := fmt.Sprintf("# %s\n\n", date)

	f.Write([]byte(header))
	f.Close()

	editor, ok := os.LookupEnv("VISUAL")
	if !ok {
		editor, ok = os.LookupEnv("EDITOR")
		if !ok {
			fmt.Println("Warning EDITOR environment variable not set, thing to use vi")
			editor = "vi"
		}
	}
	cmd := exec.Command(editor, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("Error while editing. Error: %v\n", err)
		os.Exit(1)
	}

	note, err := os.ReadFile(f.Name())

	log.Println(string(note))

}
