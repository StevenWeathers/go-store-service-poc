package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"strings"
)

func openDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster("db")
	cluster.Keyspace = "storeservices"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

func getall() (Response, error) {
	var r Response
	session, err := openDB()
	if err != nil {
		return r, err
	}

	var storeid, mondayopen, mondayclose, tuesdayopen, tuesdayclose,
		wednesdayopen, wednesdayclose, thursdayopen, thursdayclose,
		fridayopen, fridayclose, saturdayopen, saturdayclose,
		sundayopen, sundayclose int
	var address1, city, country, fax,
		isOnboardedStore, latitude, longitude, mapUrl,
		milesToStore, phone, state, storeName, storeNumber, zip string
	var flags []string

	iter := session.Query(`SELECT storeid, address1,
    city, country, fridayopen, fridayclose,
    mondayopen, mondayclose, saturdayopen, saturdayclose,
    sundayopen, sundayclose, thursdayopen, thursdayclose,
    tuesdayopen, tuesdayclose, wednesdayopen, wednesdayclose,
    fax, flags, isonboardedstore, latitude,
    longitude, mapurl, milestostore, phone,
    state, storename, storenumber, zip FROM stores`).Iter()
	for iter.Scan(&storeid, &address1,
		&city, &country, &fridayopen, &fridayclose,
		&mondayopen, &mondayclose, &saturdayopen, &saturdayclose,
		&sundayopen, &sundayclose, &thursdayopen, &thursdayclose,
		&tuesdayopen, &tuesdayclose, &wednesdayopen, &wednesdayclose,
		&fax, &flags, &isOnboardedStore, &latitude, &longitude, &mapUrl,
		&milesToStore, &phone, &state, &storeName, &storeNumber, &zip) {
		r.StoreLocation = append(r.StoreLocation, Store{
			Address1: address1,
			CanonicalStoreUrl: fmt.Sprintf("/%s-%s/%s", strings.ToUpper(state),
				strings.ToUpper(city), storeNumber),
			City:    city,
			Country: country,
			DailyHours: Hours{
				Friday:    []int{fridayopen, fridayclose},
				Monday:    []int{mondayopen, mondayclose},
				Saturday:  []int{saturdayopen, saturdayclose},
				Sunday:    []int{sundayopen, sundayclose},
				Thursday:  []int{thursdayopen, thursdayclose},
				Tuesday:   []int{tuesdayopen, tuesdayclose},
				Wednesday: []int{wednesdayopen, wednesdayclose},
			},
			Fax:              fax,
			Flags:            flags,
			IsOnboardedStore: isOnboardedStore,
			Latitude:         latitude,
			Longitude:        longitude,
			MapURL:           mapUrl,
			MilesToStore:     milesToStore,
			Phone:            phone,
			State:            state,
			StoreName:        storeName,
			StoreNumber:      storeNumber,
			Zip:              zip,
		})
	}
	if err := iter.Close(); err != nil {
		return r, err
	}

	return r, nil
}

func getstore(sn string) (Response, error) {
	var r Response
	session, err := openDB()
	if err != nil {
		return r, err
	}

	var storeid, mondayopen, mondayclose, tuesdayopen, tuesdayclose,
		wednesdayopen, wednesdayclose, thursdayopen, thursdayclose,
		fridayopen, fridayclose, saturdayopen, saturdayclose,
		sundayopen, sundayclose int
	var address1, city, country, fax,
		isOnboardedStore, latitude, longitude, mapUrl,
		milesToStore, phone, state, storeName, storeNumber, zip string
	var flags []string

	iter := session.Query(`SELECT storeid, address1,
    city, country, fridayopen, fridayclose,
    mondayopen, mondayclose, saturdayopen, saturdayclose,
    sundayopen, sundayclose, thursdayopen, thursdayclose,
    tuesdayopen, tuesdayclose, wednesdayopen, wednesdayclose,
    fax, flags, isonboardedstore, latitude,
    longitude, mapurl, milestostore, phone,
    state, storename, storenumber, zip FROM stores WHERE storenumber = ?`, sn).Iter()
	for iter.Scan(&storeid, &address1, &city, &country,
		&fridayopen, &fridayclose, &mondayopen, &mondayclose,
		&saturdayopen, &saturdayclose, &sundayopen, &sundayclose,
		&thursdayopen, &thursdayclose, &tuesdayopen, &tuesdayclose,
		&wednesdayopen, &wednesdayclose, &fax, &flags, &isOnboardedStore,
		&latitude, &longitude, &mapUrl, &milesToStore, &phone, &state,
		&storeName, &storeNumber, &zip) {
		r.StoreLocation = append(r.StoreLocation, Store{
			Address1: address1,
			CanonicalStoreUrl: fmt.Sprintf("/%s-%s/%s", strings.ToUpper(state),
				strings.ToUpper(city), storeNumber),
			City:    city,
			Country: country,
			DailyHours: Hours{
				Friday:    []int{fridayopen, fridayclose},
				Monday:    []int{mondayopen, mondayclose},
				Saturday:  []int{saturdayopen, saturdayclose},
				Sunday:    []int{sundayopen, sundayclose},
				Thursday:  []int{thursdayopen, thursdayclose},
				Tuesday:   []int{tuesdayopen, tuesdayclose},
				Wednesday: []int{wednesdayopen, wednesdayclose},
			},
			Fax:              fax,
			Flags:            flags,
			IsOnboardedStore: isOnboardedStore,
			Latitude:         latitude,
			Longitude:        longitude,
			MapURL:           mapUrl,
			MilesToStore:     milesToStore,
			Phone:            phone,
			State:            state,
			StoreName:        storeName,
			StoreNumber:      storeNumber,
			Zip:              zip,
		})
	}

	if err := iter.Close(); err != nil {
		return r, err
	}

	return r, nil
}
