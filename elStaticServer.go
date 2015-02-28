package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

//Configuration Application
type ElStaticServerConfig struct {
	//url to call
	bindingAddress string
	//directory with file to read
	baseDirectory string
}

//configuration to use
var elStaticServerConfig = new(ElStaticServerConfig)

//parse command in configuration
func init() {
	const (
		defaultBindingAddress = "localhost:4000"
		elCurrent             = "elCurrent"
		defaultBaseDirectory  = elCurrent //elCurrent if default is current directory

	)
	flag.StringVar(&elStaticServerConfig.bindingAddress, "bindingAddress", defaultBindingAddress, "The binding address")
	flag.StringVar(&elStaticServerConfig.baseDirectory, "baseDirectory", defaultBaseDirectory, "directory with file to read (elCurrent to use)")

	//use home'user
	switch elStaticServerConfig.baseDirectory {
	case "":
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		elStaticServerConfig.baseDirectory = usr.HomeDir
		break

	case elCurrent:
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		elStaticServerConfig.baseDirectory = dir
		break
	}
	flag.Parse()

}

func main() {
	panic(http.ListenAndServe(elStaticServerConfig.bindingAddress, http.FileServer(http.Dir(elStaticServerConfig.baseDirectory))))
}
