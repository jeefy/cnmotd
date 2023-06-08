package main

import (
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func LoadFeed() {
	log.Printf("Loading feed from %s", args.entryDir)
	err := filepath.Walk(args.entryDir,
		func(path string, info os.FileInfo, err error) error {
			c := &MOTDItems{}
			if err != nil {
				return err
			}
			if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
				return nil
			}
			log.Printf("Reading entries from %s", path)
			yamlFile, err := os.ReadFile(path)
			if err != nil {
				log.Printf("yamlFile.Get err   #%v ", err)
			}
			err = yaml.Unmarshal(yamlFile, c)
			if err != nil {
				log.Fatalf("Unmarshal: %v", err)
			}

			fullFeed = append(fullFeed, *c...)

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
