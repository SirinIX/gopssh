package template

const (
	defaultYAMLConfigTemplate = `global:
  port: 22
  username: root
  password: cm9vdAo=
  labels:
    all: all
  groups:
  - ips:
    - 192.168.8.8
    - 192.168.8.9
    port: 23
    username: mysql
    password: bXlzcWwK
    labels:
      mysql: master
      # all: all
  - ips:
    - 192.168.8.10
    - 192.168.8.11
    # port: 22
    # username: root
    # password: cm9vdAo=
    labels:
      mysql: slave
      # all: all`
      
	defaultJSONConfigTemplate = `{
  "global": {
    "port": 22,
    "username": "root",
    "password": "cm9vdAo=",
    "labels": {
      "all": "all"
	}
  },
  "groups": [
    {
      "ips": [
        "192.168.8.8",
        "192.168.8.9"
      ],
      "port": 23,
      "username": "mysql",
      "password": "bXlzcWwK",
      "labels": {
        "mysql": "master"
      }
    },
    {
      "ips": [
        "192.168.8.10",
        "192.168.8.11"
      ],
      "labels": {
        "mysql": "slave"
      }
    }
  ]
}`
)

func GetYAMLConfigTemplate() string {
	return defaultYAMLConfigTemplate
}

func GetJSONConfigTemplate() string {
	return defaultJSONConfigTemplate
}
