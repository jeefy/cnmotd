package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func LoadFeed() []error {
	errors := []error{}

	log.Printf("Loading feed from %s", args.entryDir)
	err := filepath.Walk(args.entryDir,
		func(path string, info os.FileInfo, err error) error {
			c := &MOTDItems{}
			if err != nil {
				errors = append(errors, err)
				return err
			}
			if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
				return nil
			}
			log.Printf("Reading entries from %s", path)
			yamlFile, err := os.ReadFile(path)
			if err != nil {
				errors = append(errors, err)
				log.Printf("yamlFile.Get err   #%v ", err)
			}
			err = yaml.UnmarshalStrict(yamlFile, c)
			if err != nil {
				errors = append(errors, err)
				log.Fatalf("Unmarshal: %v", err)
			}

			for _, item := range *c {
				if item.EndDate.IsZero() {
					errors = append(errors, fmt.Errorf("no end date is set for `%s`, skipping", item.Item.Title))
				}
				if item.Level == "" {
					errors = append(errors, fmt.Errorf("no level is set for `%s`, skipping", item.Item.Title))
				}
				if item.Item.Link.Href == "" {
					errors = append(errors, fmt.Errorf("no link is set for `%s`, skipping", item.Item.Title))
				}
				if item.Item.Title == "" {
					errors = append(errors, fmt.Errorf("missing title for an entry"))
				}
				if item.Item.Description == "" {
					errors = append(errors, fmt.Errorf("missing description for `%s`", item.Item.Title))
				}
				if item.Item.Author.Name == "" {
					errors = append(errors, fmt.Errorf("missing author for `%s`", item.Item.Title))
				}
			}

			fullFeed = append(fullFeed, *c...)

			return nil
		})
	if err != nil {
		errors = append(errors, err)
		log.Println(err)
	}
	return errors
}
