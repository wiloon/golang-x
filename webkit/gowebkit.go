package main

import (
	"fmt"
	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
	"github.com/sqs/gojs"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	webView.Connect("load-changed", func(ctx *glib.CallbackContext) {
		loadEvent := webkit2.LoadEvent(ctx.Arg(0).Int())
		switch loadEvent {
		case webkit2.LoadFinished:
			fmt.Println("Load finished.")
			fmt.Printf("Title: %q\n", webView.Title())
			fmt.Printf("URI: %s\n", webView.URI())
			webView.RunJavaScript("window.location.hostname", func(val *gojs.Value, err error) {
				if err != nil {
					fmt.Println("JavaScript error.")
				} else {
					fmt.Printf("Hostname (from JavaScript): %q\n", val)
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI("https://www.baidu.com/")
		return false
	})

	gtk.Main()

	// output:
	// Load finished.
	// Title: "Google"
	// URI: https://www.google.com/
	// Hostname (from JavaScript): "www.google.com"
}