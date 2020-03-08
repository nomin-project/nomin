package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
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

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/graphics/icon.icns",
			AppIconDefaultPath: "resources/graphics/icon-512.png",
		},
		Debug: *debug,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astikit.StrPtr("Nomin"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astikit.StrPtr("About"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("https://github.com/nomin-project/nomin#about")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
					{
						Label: astikit.StrPtr("Contribute"),
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
				Label: astikit.StrPtr("Help"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astikit.StrPtr("Report Bug"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							err := browser.OpenURL("http://www.github.com/nomin-project/nomin/issues")
							if err != nil {
								fmt.Println(err)
							}
							return
						},
					},
					{
						Label: astikit.StrPtr("Contact Developer"),
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
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			// Store global variables
			window = ws[0]

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
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(850),
				Width:           astikit.IntPtr(1100),
			},
		}},
	}); err != nil {
		log.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
