package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var listen = flag.String("l", "localhost:8888", "address to listen on")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage: %s [-l] [<editor> [<editor-args> ...]]\n\n",
			os.Args[0])
		fmt.Fprintln(os.Stderr, "  <editor>=\"gvim\": binary that will be "+
			"started to edit textarea contents")
		fmt.Fprintln(os.Stderr, "  <editor-args>: any editor args that "+
			"should be given prior to filename")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
	}

	flag.Parse()
	http.HandleFunc("/", editHandler)
	log.Println("starting server on", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request from", r.RemoteAddr, r.RequestURI)

	f, err := ioutil.TempFile(os.TempDir(), "chrome")
	if err != nil {
		log.Fatal("can't create temporary file", err)
	}

	defer func() {
		f.Close()
		err := os.Remove(f.Name())
		if err != nil {
			log.Fatal("can't remove temp file", err)
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("can't read body of request", err)
	}

	f.Write(body)

	editor := "gvim"
	args := []string{}
	if flag.NArg() > 0 {
		editor = flag.Arg(0)
	}
	if flag.NArg() > 1 {
		args = flag.Args()[1:]
	}

	cmd := exec.Command(editor, append(args, f.Name())...)
	log.Printf("launching editor as %s %#v", cmd.Path, cmd.Args)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("can't launch editor, command says:")
		log.Fatalf("%s", out)
	}

	f.Seek(0, os.SEEK_SET)
	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("error while reading result text from temp file", err)
	}

	w.Write(data)
}
