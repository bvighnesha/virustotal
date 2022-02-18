package main

import (
	"encoding/json"
	"fmt"
	"vighnesha.in/virustotal/api"
	"vighnesha.in/virustotal/api/model"
)

func main() {
	jsonData := `{
   "data": {
	"ID":"{object ID}",
  "type": "{object type}",
  "links": {
    "related": "https://www.virustotal.com/api/v3/{collection name}/{object id}/{one-to-one relationship}",
    "self": "https://www.virustotal.com/api/v3/{collection name}/{object id}",
    "next": "https://www.virustotal.com/api/v3/{collection name}/{object id}/relationships/{one-to-many relationship}?cursor=CuABChEKBGRhdGUSCQjA1LC..."
  },
  "attributes": {},
  "relationships": {
    "relationship": {
      "data": {
        "error": {
        "code": "NotFoundError",
        "message": "{item type} with id \"{item id}\" not found"
      },
        "id": "www.google.com",
        "type": "domain"
      },
       "meta": {
        "cursor": "CuABChEKBGRhdGUSCQjA1LC..."
      },
      "links": {
        "related": "https://www.virustotal.com/api/v3/{collection name}/{object id}/{one-to-one relationship}",
        "self": "https://www.virustotal.com/api/v3/{collection name}/{object id}/relationships/{one-to-one relationship}",
        "next": "https://www.virustotal.com/api/v3/{collection name}/{object id}/relationships/{one-to-many relationship}?cursor=CuABChEKBGRhdGUSCQjA1LC..."
      }
    },
	"relationship1": {
      "data": [{
        "error": {
        "code": "NotFoundError",
        "message": "{item type} with id \"{item id}\" not found"
      },
        "id": "www.google.com",
        "type": "domain"
      }],
       "meta": {
        "cursor": "CuABChEKBGRhdGUSCQjA1LC..."
      },
      "links": {
        "related": "https://www.virustotal.com/api/v3/{collection name}/{object id}/{one-to-one relationship}",
        "self": "https://www.virustotal.com/api/v3/{collection name}/{object id}/relationships/{one-to-one relationship}",
        "next": "https://www.virustotal.com/api/v3/{collection name}/{object id}/relationships/{one-to-many relationship}?cursor=CuABChEKBGRhdGUSCQjA1LC..."
      }
    }
	
  }
    }
}`

	var response model.Response
	json.Unmarshal([]byte(jsonData), &response)
	//data, err := json.MarshalIndent(response, "", " ")
	//fmt.Println(string(data), err)
	relationships := response.Data.Relationships["relationship"]
	fmt.Println(relationships.Data[0].ID)
	//fmt.Println(jsonData)

	virusTotal := api.New("0835ea6ebdb5441835e16462e3858728bf6d550fd0a765d3ff6b4a52b238c4fb")
	//virusTotal.Universal().Files().UploadFromPath("/Users/vighnesha/Downloads/Fiddler Everywhere 3.0.1.dmg")
	//virusTotal.Universal().Files().UploadURLForLargerThan32MB()

	virusTotal.Universal().Files().BehaviourSummary("ae339ffd45d5f0870dde1dacf9c572153a302fbf1db0bdd162105aa7049bf2c0")

}
