package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	f.Write([]byte("# test\n"))
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
	log.Printf("Successfully edited.")

}
