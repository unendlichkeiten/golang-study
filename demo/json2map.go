package main

import (
	"encoding/json"
	"fmt"
)

const Json = `[
  {
    "caller": "API",
    "permissions": [
      {
        "interface": "interface_1",
        "method": "GET"
      },
      {
        "interface": "interface_2",
        "method": "POST"
      }
    ]
  },
  {
    "caller": "GOV",
    "permissions": [
      {
        "interface": "interface_3",
        "method": "GET"
      },
      {
        "interface": "interface_4",
        "method": "POST"
      }
    ]
  }
]`

type Permission struct {
	Interface string `json:"interface,omitempty"`
	Method    string `json:"method,omitempty"`
}

type Auth struct {
	Caller      string       `json:"caller,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}
func main() {
	AuthInfos := make([]Auth, 0)
	if err := json.Unmarshal([]byte(Json), &AuthInfos); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", AuthInfos)
	}

	authMap := make(map[string]map[string]string)
	permissionMap := make(map[string]string)
	for _, auth := range AuthInfos {
		for _, permission := range auth.Permissions {
			permissionMap[permission.Interface] = permission.Method
		}
		authMap[auth.Caller] = permissionMap
		permissionMap = make(map[string]string)
	}

	fmt.Printf("%#v\n", authMap)
}
