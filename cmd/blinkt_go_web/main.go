/*
reads a config.yml of such format:

function:
  - path: "/red"
    r: 255
    g: 0
    b: 0
    pixel: 1
  - path: "/blue"
    r: 0
    g: 0
    b: 255
    pixel: 1
*/

package main

import (
	"fmt"
	"net/http"
	. "github.com/alexellis/blinkt_go"
)

var blinkt *Blinkt

type MyHandler struct {
	session struct {
		Path  string `yaml:"path"`
		R     int    `yaml:"r"`
		G     int    `yaml:"g"`
		B     int    `yaml:"b"`
		Pixel int    `yaml:"pixel"`
	}
}

func (f *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if all?
	//pixels should be a list
	//blinking support?

	blinkt.SetPixel(f.session.Pixel,f.session.R,f.session.G,f.session.B)
	blinkt.Show()
	fmt.Printf("%+v", f) //Handle GPIO

}

func main() {
	var cfg Config
        brightness := 0.5
        newblinkt := NewBlinkt(brightness)
        newblinkt.SetClearOnExit(true)
        newblinkt.Setup()
        blinkt = &newblinkt

	readFile(&cfg)
	mux := http.NewServeMux()
	for _, element := range cfg.Function {
		fmt.Printf("%+v", element.Path)
		mux.Handle(element.Path, &MyHandler{element})
	}
	http.ListenAndServe(":3000", mux)

}
