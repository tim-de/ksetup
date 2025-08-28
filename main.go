package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"strings"
	"ksetup/throbber"
)

func main() {
	if !fileExists("_ag.exe") {
		log.Println("Program must be run in Kohan Ahriman's Gift installation directory")
		os.Exit(1)
	}

	kg_target, err := getDownloadInfo("Kohan-Citadel/kohangold-KG-")
	if err != nil {
		log.Fatalf("Couldn't get release information for KohanGold: %s\n", err.Error())
	}
	ddraw_target, err := getDownloadInfo("FunkyFr3sh/cnc-ddraw")
	if err != nil {
		log.Fatalf("Couldn't get release information for cnc-ddraw: %s\n", err.Error())
	}
	khaldun_target, err := getDownloadInfo("Kohan-Citadel/khaldun.net-client")
	if err != nil {
		log.Fatalf("Couldn't get release information for khaldun.net client: %s\n", err.Error())
	}

	kg_fix := Fix{
		description: "Kohan Gold",
		check: []string{kg_target.filename},
		download: []TargetInfo{kg_target},
	}

	launcher_fix := Fix{
		description: "Kohan Launcher",
		check: []string{"KohanLauncher.exe"},
		remove: nil,
		download: []TargetInfo{{
			downloadurl: "https://github.com/Kohan-Citadel/kohangold-KG-/releases/download/v0.9.6/KohanLauncher.exe",
			filename: "KohanLauncher.exe",
		}},
	}

	khaldun_client_fix := Fix{
		description: "khaldun.net client",
		check: []string{"dinput.dll"},
		download: []TargetInfo{khaldun_target},
		unpack: []string{khaldun_target.filename},
		cleanup: []string{khaldun_target.filename},
	}

	append_to_cnc_ini := func() {
		throbber.Delay()
		log.Println("cnc-ddraw - Adding Kohan settings to ddraw.ini")
		to_append := strings.Join([]string{
			"; Kohan: Ahriman's Gift",
			"[_ag]",
			"windowed=false",
			"fullscreen=true",
			"maintas=true",
		}, "\r\n")
		filehandle, err := os.OpenFile("ddraw.ini", os.O_APPEND | os.O_WRONLY, 0644)
		if err != nil {
			throbber.Delay()
			log.Printf("cnc-ddraw - Error: Could not open 'ddraw.ini' for writing: %s\n",
				err.Error())
			return
		}
		defer filehandle.Close()

		if _, err := filehandle.WriteString(to_append); err != nil {
			throbber.Delay()
			log.Printf("cnc-ddraw - Error: could not append to 'ddraw.ini': %s",
				err.Error())
		}
	}

	ddraw_fix := Fix{
		description: "cnc-ddraw",
		check: []string{
			"ddraw.dll",
			"ddraw.ini",
			"cnc-ddraw config.exe",
			"Shaders/",
		},
		download: []TargetInfo{ddraw_target},
		remove: []string{
			"D3D8.dll",
			"D3D9.dll",
			"D3DImm.dll",
			"DDraw.dll",
			"dgVoodoo.conf",
			"dgVoodooCpl.exe",
			"libwine.dll",
			"wine3d.dll",
		},
		unpack: []string{ddraw_target.filename},
		cleanup: []string{ddraw_target.filename},
		extra: append_to_cnc_ini,
	}

	fixes := []Fix{
		kg_fix,
		launcher_fix,
		khaldun_client_fix,
		ddraw_fix,
	}

	var waitgroup sync.WaitGroup
	throbber.Start()
	for _, fix := range fixes {
		apply_fix(fix, &waitgroup)
	}
	waitgroup.Wait()
	throbber.Stop()
	log.Println("Done! Press Enter to quit")
	fmt.Scanln()
}
