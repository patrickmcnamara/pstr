package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const pstrURL = "https://paster.xyz/"

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("pstr: pstr [ID] prints paste of that ID to stdout")
			fmt.Println("      pstr posts paste from stdin to paster.xyz and prints it's URL")
		} else {
			// get paste by id
			r, err := http.Get(pstrURL + os.Args[1])
			chk(err)
			// print paste to stdout
			io.Copy(os.Stdout, r.Body)
		}
	} else {
		// create form
		vals := make(url.Values)
		// read paste value from stdin
		buf, err := ioutil.ReadAll(os.Stdin)
		chk(err)
		// add values to form
		vals.Add("Value", string(buf))
		vals.Add("List", "list")
		vals.Add("Expiry", "")
		// post form
		r, err := http.PostForm(pstrURL, vals)
		chk(err)
		// print url of paste
		fmt.Println(r.Request.URL.String())
	}
}

func chk(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Errorf("pstr: %w", err))
		os.Exit(1)
	}
}
