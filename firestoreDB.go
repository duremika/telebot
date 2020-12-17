package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	projectID      = "protection-equip"
	collectionName = "equips"
)

func getNewClient() (client *firestore.Client) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil
	}
	return
}

func Update(equip *Equip) {
	client := getNewClient()
	defer client.Close()

	it := client.Collection(collectionName).Documents(context.Background())
	var docName string
	for {
		if doc, err := it.Next(); err != nil {
			break
		} else if doc.Data()["Name"] == equip.Name {
			docName = doc.Ref.ID
		}
	}

	client.Collection(collectionName).Doc(docName).Set(
		context.Background(),
		map[string]interface{}{
			"Name": equip.Name,
			"Date": equip.Date.Format(layout)},
		firestore.MergeAll)
}

func FindAll() ([]Equip, error) {
	client := getNewClient()

	defer client.Close()

	var equipsDB []Equip
	it := client.Collection(collectionName).Documents(context.Background())
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		ss := strings.Split(doc.Data()["Date"].(string), " ")
		y, _ := strconv.Atoi(ss[2])
		m, _ := strconv.Atoi(ss[1])
		d, _ := strconv.Atoi(ss[0])
		an := []int{2020 + y, m, d}

		equip := Equip{
			Name: doc.Data()["Name"].(string),
			Date: time.Date(an[0], time.Month(an[1]), an[2], 0, 0, 0, 0, time.Local),
		}
		equipsDB = append(equipsDB, equip)
	}
	sort.SliceStable(equipsDB, func(i, j int) bool {
		return int(equipsDB[i].Date.Unix()) < int(equipsDB[j].Date.Unix())
	})
	return equipsDB, nil
}
