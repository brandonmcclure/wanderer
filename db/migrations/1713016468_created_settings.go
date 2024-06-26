package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "uavt73rsqcn1n13",
			"created": "2024-04-13 13:54:26.023Z",
			"updated": "2024-04-13 13:54:26.023Z",
			"name": "settings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "0sepzvkh",
					"name": "language",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"en",
							"de",
							"fr",
							"hu",
							"nl",
							"pl",
							"pt",
							"zh"
						]
					}
				},
				{
					"system": false,
					"id": "zwg1jl0d",
					"name": "unit",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"metric",
							"imperial"
						]
					}
				},
				{
					"system": false,
					"id": "jo1zcsbu",
					"name": "mapFocus",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"trails",
							"location"
						]
					}
				},
				{
					"system": false,
					"id": "ufhepjxo",
					"name": "location",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				},
				{
					"system": false,
					"id": "5uip7a4p",
					"name": "user",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("uavt73rsqcn1n13")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
