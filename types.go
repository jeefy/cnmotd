package main

import (
	"log"
	"time"

	"github.com/gorilla/feeds"
)

type MOTDLevel string

const (
	LevelInfo     MOTDLevel = "info"
	LevelWarning  MOTDLevel = "warn"
	LevelCritical MOTDLevel = "crit"
)

func (l MOTDLevel) Int() int {
	switch l {
	case LevelCritical:
		return 0
	case LevelWarning:
		return 1
	case LevelInfo:
		return 2
	default:
		log.Printf("%v is an invalid MOTDLevel", l)
		return 0
	}
}

type MOTDItem struct {
	Projects  map[string]bool `json:"projects,omitempty" yaml:"projects,omitempty"`
	Level     MOTDLevel       `json:"level,omitempty" yaml:"level"`
	StartDate time.Time       `json:"start_date,omitempty" yaml:"startDate,omitempty"`
	EndDate   time.Time       `json:"end_date,omitempty" yaml:"endDate"`
	Item      feeds.Item      `json:"item" yaml:"item"`
}

type MOTDItems []MOTDItem

func (m MOTDItems) Filter(projects []string, level MOTDLevel) MOTDItems {
	var filtered MOTDItems

	for _, item := range m {
		if !item.StartDate.IsZero() && time.Now().Before(item.StartDate) {
			if args.debug {
				log.Printf("Skipping item %s because it is not yet at the start date", item.Item.Title)
			}
			continue
		}
		if item.EndDate.IsZero() {
			if args.debug {
				log.Printf("No end date is set for %s, skipping", item.Item.Title)
			}
			continue
		} else if time.Now().After(item.EndDate) {
			if args.debug {
				log.Printf("Skipping item %s because it is past the end date", item.Item.Title)
			}
			continue
		}
		if item.Level.Int() > level.Int() {
			if args.debug {
				log.Printf("Skipping item %s because it (%d) is below the level threshold (%d)", item.Item.Title, item.Level.Int(), level.Int())
			}
			continue
		}

		for _, project := range projects {

			if _, ok := item.Projects[project]; ok || project == "" {
				filtered = append(filtered, item)
				if args.debug {
					log.Printf("Adding item %s because it matches project %s", item.Item.Title, project)
				}
			} else {
				if args.debug {
					log.Printf("Skipping item %s because it does not match project %s", item.Item.Title, project)
				}
			}
		}
	}
	log.Printf("Returning %d items", len(filtered))

	return filtered
}
