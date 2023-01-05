package launch

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	_ "github.com/lib/pq"
	"github.com/vincent87720/daymood/app/internal/routers"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func Launch() {
	s := settings.Init()

	if s.GetAppMode() == "DEV" || s.GetAppMode() == "PROD" {
		runApp(s)
	} else {
		//production
		// nsAppInit(s)
	}
}

// func nsAppInit(s settings.Settings) {
// 	runtime.LockOSThread()

// 	cocoa.TerminateAfterWindowsClose = false
// 	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
// 		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
// 		obj.Retain()

// 		//go run server
// 		go runApp(s)

// 		//set app icon
// 		imageBoxBytes := loadImage(s.GetExeFilePath() + "/box.pdf")
// 		imageBoxData := core.NSData_WithBytes(imageBoxBytes, uint64(len(imageBoxBytes)))
// 		imageBox := cocoa.NSImage_InitWithData(imageBoxData)
// 		obj.Button().SetImage(imageBox)

// 		//Set "Open" option
// 		itemOpen := cocoa.NSMenuItem_New()
// 		itemOpen.SetTitle("Open")
// 		itemOpen.SetAction(objc.Sel("openClicked:"))
// 		cocoa.DefaultDelegateClass.AddMethod("openClicked:", func(_ objc.Object) {
// 			openbrowser("http://" + s.GetBackendAddr() + "/daymood")
// 			// openClicked <- true
// 		})

// 		//Set "Quit" option
// 		itemQuit := cocoa.NSMenuItem_New()
// 		itemQuit.SetTitle("Quit")
// 		itemQuit.SetAction(objc.Sel("terminate:"))

// 		//new separator
// 		itemSeparator := cocoa.NSMenuItem_Separator()

// 		//set menu title
// 		r := core.NSMakeRect(0, 0, 200, 10)
// 		textView := cocoa.NSTextView_Init(r)
// 		textView.SetString("DAYMOOD MANAGMENT")
// 		textView.SetBackgroundColor(cocoa.NSColor_Clear())
// 		textView.SetAlignment(2)
// 		textView.SetEditable(false)
// 		textView.SetFieldEditor(false)
// 		itemTitle := cocoa.NSMenuItem_New()
// 		itemTitle.SetView(textView)

// 		empty := cocoa.NSView_Init(r)
// 		itemSpace := cocoa.NSMenuItem_New()
// 		itemSpace.SetView(empty)
// 		itemSpace2 := cocoa.NSMenuItem_New()
// 		itemSpace2.SetView(empty)

// 		//create menu
// 		menu := cocoa.NSMenu_New()
// 		menu.SetMinimumWidth_(150)
// 		menu.AddItem(itemOpen)
// 		menu.AddItem(itemQuit)
// 		menu.AddItem(itemSeparator)
// 		menu.AddItem(itemSpace)
// 		menu.AddItem(itemTitle)
// 		menu.AddItem(itemSpace2)
// 		obj.SetMenu(menu)

// 	})
// 	app.Run()
// }

func runApp(s settings.Settings) {
	connStr := s.GetDBConnectionString()
	backendAddr := s.GetBackendAddr()

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router, err := routers.SetupRouters(db, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Run(backendAddr)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func loadImage(filepath string) []byte {
	// imageFile := "Box.pdf"
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes
}
