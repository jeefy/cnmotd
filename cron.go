package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
)

func findOldFiles(dir string) (files []os.FileInfo, err error) {
	tmpfiles, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {
		if file.Type().IsRegular() {
			info, _ := file.Info()
			if time.Since(info.ModTime()) > time.Duration(args.maxCacheAge)*time.Hour {
				if args.debug {
					log.Printf("Found old file %s with mod time of %d", file.Name(), info.ModTime().Unix())
				}
				files = append(files, info)
			}
		}
	}
	return
}

func deleteFiles(path string, files []os.FileInfo) {
	log.Printf("Deleting %d old files\n", len(files))
	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		if args.debug {
			log.Printf("Deleting %s", fullPath)
		}
		err := os.Remove(fullPath)
		if err != nil {
			log.Printf("Error deleting %s: %s\n", fullPath, err)
		}
	}
}

func pruneFiles() {
	log.Printf("Pruning old files in %s\n", args.cacheDir)
	files, err := findOldFiles(args.cacheDir)
	if err != nil {
		return
	}

	deleteFiles(args.cacheDir, files)
}

func StartCron() {
	if args.maxCacheAge > 0 {
		log.Println("Starting CRON pruner")
		cron := gocron.NewScheduler(time.UTC)
		_, err := cron.Cron(args.cronSchedule).Do(pruneFiles)
		if err != nil {
			log.Fatalf("Error creating prune cronjob: %s", err.Error())
		}
		cron.StartAsync()
	} else {
		log.Println("Cache age set to 0. Skipping CRON pruner")
	}
}
