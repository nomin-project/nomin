package main

import (
	"flag"
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "if yes, the app is in debug mode")
	window  *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/graphics/icon.icns",
			AppIconDefaultPath: "resources/graphics/icon-512.png",
		},
		Debug:    *debug,
		Homepage: "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr("Nomin"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astilectron.PtrStr("About"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("https://github.com/nomin-project/nomin#about")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
					{
						Label: astilectron.PtrStr("Contribute"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("https://github.com/nomin-project/nomin/blob/master/docs/contribute.adoc")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
					{
						Type: astilectron.MenuItemTypeSeparator,
					},
					{
						Role: astilectron.MenuItemRoleMinimize,
					},
					{
						Role: astilectron.MenuItemRoleClose,
					},
				},
			},
			{
				Label: astilectron.PtrStr("Help"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astilectron.PtrStr("Report Bug"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("http://www.github.com/nomin-project/nomin/issues")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
					{
						Label: astilectron.PtrStr("Contact Developer"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("http://www.github.com/nomin-project/nomin#contact-us")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
				},
			},
		},
		MessageHandler: handleMessages,
		OnWait: func(_ *astilectron.Astilectron, w *astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			// Store global variables
			window = w

			// Add listeners on tray
			//t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) { astilog.Info("Tray has been clicked!"); return })
			return nil
		},
		RestoreAssets: RestoreAssets,
		// Commented out due to #33
		//TrayOptions: &astilectron.TrayOptions{
		//	Image:   astilectron.PtrStr("resources/graphics/icon-20.png"),
			//Tooltip: astilectron.PtrStr("Wow, what a beautiful tray!"),
		//},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(850),
			Width:           astilectron.PtrInt(1200),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
