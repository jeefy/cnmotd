# Cloud Native MOTD

A unified message bus for CNCF projects to use for notifying users of changes or critical issues 

## How it works
`cnmotd` will traverse and read all YAML files under the [entries](entries/) folder, including subfolders. Those YAML files have should contain a list object that looks like the following:

```
- projects:
    kubernetes: true
  level: 0
  endDate: 2023-09-01 00:00:00
  item:
    title: Update k8s.gcr.io to use registry.k8s.io
    link:
      href: https://registry.k8s.io
    description: |
      The k8s.gcr.io container registry is being deprecated in favor of registry.k8s.io. This is a breaking change for all Kubernetes users.
    author:
      name: The Kubernetes Community
      email: community@kubernetes.io
```

This will generate an ATOM Feed that can be queried/customized for consumption across all CNCF projects.

## Fields

| Field Name  | Description | Required? |
| ----------- | ----------- | ----------- |
| projects    | map[string]bool, relevant project(s) this message should be displayed for | No |
| level   | The importance level of the message. 0:critical, 1:important, 2:info | Yes |
| startDate | The date the message should be published (YYYY-MM-DD HH:II:SS) | No |
| endDate | The date the message should be published (YYYY-MM-DD HH:II:SS) | Yes |
| item.title | Title of the notice | Yes |
| item.link.href | Link to any relevant post or info | Yes |
| item.description | A more verbose message to end users | Yes |
| item.author.name | Who wrote/published the message | Yes |
| item.author.email | An email address or contact point | No |

## Example Queries

```
<BASE_URL>?projects=kubernetes              # Display critical messages for Kubernetes
<BASE_URL>?projects=kubernetes&level=1      # Display warning and critical messages for Kubernetes
<BASE_URL>?projects=kubernetes,etcd&level=2 # Display all messages for Kubernetes and etcd
<BASE_URL>?level=2                          # Display all messages for all projects
<BASE_URL>                                  # Display critical messages for all projects
```

## CLI Consumption Example

```
package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	fp.UserAgent = fmt.Sprintf("%s/%s", "kubectl", "v1.28.0")
	feed, err := fp.ParseURL("http://localhost:8080?projects=kubernetes")
	if err != nil {
		log.Printf("Error parsing feed: %s", err)
	} else {
		if len(feed.Items) > 0 {
			color.Blue("-- Cloud Native Notices --")
			for _, item := range feed.Items {
				fmt.Println("- ", color.YellowString(item.Title), " - ", color.GreenString(item.Link))
			}
			color.Blue("-- /motd.cncf.io/ --")
		}
	}
}

```