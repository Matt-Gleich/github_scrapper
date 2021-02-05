package main

import (
	"time"

	"github.com/Matt-Gleich/github_scraper/pkg/account"
	"github.com/Matt-Gleich/github_scraper/pkg/db"
	"github.com/Matt-Gleich/github_scraper/pkg/gh_api"
)

func main() {
	// Connecting to db and resetting tables
	db.Connect()
	db.ResetTable(account.CreateQuery, account.TableName)

	client := gh_api.GenClient()

	for {
		// Getting account information
		rawAccountData := account.GetData(client)
		formattedAccountData := account.CleanData(rawAccountData)
		account.Insert(formattedAccountData)

		time.Sleep(5 * time.Minute)
	}
}
