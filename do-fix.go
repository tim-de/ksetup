package main

import (
	"log"
	"os"
	"sync"
	
	"ksetup/throbber"
)

type Fix struct {
	description string
	check []string
	remove []string
	download []TargetInfo
	unpack []string
	cleanup []string
	rename []struct{from string; to string}
	extra func()
}

func apply_fix(fix Fix, waitgroup *sync.WaitGroup) {
	waitgroup.Add(1)
	go func(fix Fix) {
		defer waitgroup.Done()
		for _, check_target := range fix.check {
			if !fileExists(check_target) {
				break
			}
			throbber.Delay()
			log.Printf("%s - Files already present, skipping\n", fix.description)
			return
		}
		for _, rm_target := range fix.remove {
			if fileExists(rm_target) {
				os.Remove(rm_target)
			}
		}

		for _, dl_target := range fix.download {
			throbber.Delay()
			log.Printf("%s - Starting download of '%s' from '%s'\n",
				fix.description,
				dl_target.filename,
				dl_target.downloadurl)
			downloadFile(dl_target.downloadurl, dl_target.filename)
			throbber.Delay()
			log.Printf("%s - Finished downloading '%s'\n",
				fix.description,
				dl_target.filename)
		}

		for _, unpack_target := range fix.unpack {
			throbber.Delay()
			log.Printf("%s - Unpacking %s\n", fix.description, unpack_target)
			unzip_file(unpack_target)
		}

		for _, rename_target := range fix.rename {
			if err := os.Rename(rename_target.from, rename_target.to); err != nil {
				throbber.Delay()
				log.Printf("%s - Error: Failed to rename '%s' to '%s': %s\n",
					fix.description,
					rename_target.from,
					rename_target.to,
					err.Error())
			}
		}

		if fix.cleanup != nil {
			throbber.Delay()
			log.Printf("%s - cleaning up intermediate files\n", fix.description)
		}

		for _, cleanup_target := range fix.cleanup {
			_ = os.Remove(cleanup_target)
		}

		if fix.extra != nil {
			fix.extra()
		}
	} (fix)
}
