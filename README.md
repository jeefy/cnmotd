# Cloud Native MOTD

A unified message bus for CNCF projects to use for notifying users of changes or critical issues 

## How it works
`cnmotd` will traverse and read all YAML files under the [entries](entries/) folder, including subfolders. Those YAML files have should contain a list object that looks like the following:

```
- projects:
    kubernetes: true
  level: crit
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

This will generate an ATOM Feed that can be queried/customized for consumption across all CNCF projects

## Fields

| Field Name  | Description | Required? |
| ----------- | ----------- | ----------- |
| projects    | map[string]bool, relevant project(s) this message should be displayed for | No |
| level   | The importance level of the message. (crit, warn, info) | Yes |
| startDate | The date the message should be published (YYYY-MM-DD HH:II:SS) | No |
| endDate | The date the message should be published (YYYY-MM-DD HH:II:SS) | Yes |
| item.title | Title of the notice | Yes |
| item.link.href | Link to any relevant post or info | Yes |
| item.description | A more verbose message to end users | Yes |
| item.author.name | Who wrote/published the message | Yes |
| item.author.email | An email address or contact point | No |

## Example Queries

```
<BASE_URL>?projects=kubernetes                 # Display critical messages for Kubernetes
<BASE_URL>?projects=kubernetes&level=warn      # Display warning and critical messages for Kubernetes
<BASE_URL>?projects=kubernetes,etcd&level=info # Display all messages for Kubernetes and etcd
<BASE_URL>?level=info                          # Display all messages for all projects
<BASE_URL>                                     # Display critical messages for all projects
```

## CLI Consumption Example

There is a client/consumption example in [examples/client.go](examples/client.go)

This parses the published feed and outputs color coded and parseable info

```
-- Cloud Native Notices --
CRIT - Update k8s.gcr.io to use registry.k8s.io - https://registry.k8s.io
-- /motd.cncf.io/ --
```

### Consumption Recommendations

- **The Most Important Rule:** 
  
  Never block on the MOTD service. If it's unreachable or returning invalid data, shrug and move on. This is meant to supplement CNCF projects, but we don't want to potentially compromise trust in our projects because the MOTD service isn't available
- Allow users to specify their own `projects` and `level`, but default to `critical` and your project name
- Allow users to specify the URL of the CNMOTD instance, but please default to `https://motd.cncf.io`
- Allow users to skip-or-hide MOTD output. This is especially helpful for CLI-clients that get heavily scripted against