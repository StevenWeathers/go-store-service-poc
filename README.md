# Store Hours Service POC

This is an example RESTful API in [Go](https://golang.org/)
using [Cassandra](http://cassandra.apache.org/)
and [Docker](https://www.docker.com/)
for [Store Hours](https://www.lowes.com/bin/stores.csv)

## Installation

This API is best served with [Docker](https://www.docker.com/)
```bash
docker build -t store-services .
docker build -t store-services-db ./db
docker run -d --name=store-services-db store-services-db
docker run -d -p 8080:8080 --network container:store-services-db --name=store-services store-services
```
What about [Docker Compose](https://docs.docker.com/compose/)?
```bash
docker-compose up -d --build
```

## Open Source Projects
Project | License
--- | ---
[gocql](https://github.com/gocql/gocql) | [BSD-3-Clause](https://github.com/gocql/gocql/blob/master/LICENSE)
[Docker](https://github.com/docker/docker) | [Apache-2.0](https://github.com/docker/docker/blob/master/LICENSE)
[Cassandra](https://github.com/apache/cassandra) | [Apache-2.0](https://github.com/apache/cassandra/blob/trunk/LICENSE.txt)
[Go](https://github.com/golang/go) | [BSD-3-Clause](https://github.com/golang/go/blob/master/LICENSE)

## Example Service Response
```json
{
    "storeLocation": [
        {
            "address1": "1041 Charlotte Highway",
            "canonicalStoreUrl": "/NC-TROUTMAN/2750",
            "city": "Troutman",
            "country": "US",
            "dailyHours": {
                "friday": [
                    25200,
                    75600
                ],
                "monday": [
                    25200,
                    75600
                ],
                "saturday": [
                    25200,
                    75600
                ],
                "sunday": [
                    28800,
                    68400
                ],
                "thursday": [
                    25200,
                    75600
                ],
                "tuesday": [
                    25200,
                    75600
                ],
                "wednesday": [
                    25200,
                    75600
                ]
            },
            "fax": "(704) 980-5003",
            "flags": [
                "inStockPageView",
                "metaProducts",
                "productAisleNumber",
                "productBayNumber",
                "productLocations",
                "storeMapView"
            ],
            "isOnboardedStore": "true",
            "latitude": "35.67098105",
            "longitude": "-80.85143012",
            "mapUrl": "http://mobilepti.lowes.com/staticmap?api_key=977b2334e34711e18edc12313d23aa6e&devid=%7BdevId%7D&storeId=2750&size=1024x1024",
            "milesToStore": "7.29",
            "phone": "(704) 980-5000",
            "state": "NC",
            "storeName": "LOWE'S OF TROUTMAN, NC",
            "storeNumber": "2750",
            "zip": "28166"
        },
        {
            "address1": "16830 Statesville Rd",
            "canonicalStoreUrl": "/NC-HUNTERSVILLE/489",
            "city": "Huntersville",
            "country": "US",
            "dailyHours": {
                "friday": [
                    21600,
                    79200
                ],
                "monday": [
                    21600,
                    79200
                ],
                "saturday": [
                    21600,
                    79200
                ],
                "sunday": [
                    28800,
                    72000
                ],
                "thursday": [
                    21600,
                    79200
                ],
                "tuesday": [
                    21600,
                    79200
                ],
                "wednesday": [
                    21600,
                    79200
                ]
            },
            "fax": "(704) 892-6746",
            "flags": [
                "inStockPageView",
                "metaProducts",
                "productAisleNumber",
                "productLocations",
                "storeMapView"
            ],
            "isOnboardedStore": "false",
            "latitude": "35.44581072",
            "longitude": "-80.86484249",
            "mapUrl": "http://mobilepti.lowes.com/staticmap?api_key=977b2334e34711e18edc12313d23aa6e&devid=%7BdevId%7D&storeId=489&size=1024x1024",
            "milesToStore": "8.8",
            "phone": "(704) 892-9449",
            "state": "NC",
            "storeName": "LOWE'S OF HUNTERSVILLE, NC",
            "storeNumber": "489",
            "zip": "28078"
        },
        {
            "address1": "7144 Highway 73",
            "canonicalStoreUrl": "/NC-DENVER/2636",
            "city": "Denver",
            "country": "US",
            "dailyHours": {
                "friday": [
                    21600,
                    79200
                ],
                "monday": [
                    21600,
                    79200
                ],
                "saturday": [
                    21600,
                    79200
                ],
                "sunday": [
                    28800,
                    72000
                ],
                "thursday": [
                    21600,
                    79200
                ],
                "tuesday": [
                    21600,
                    79200
                ],
                "wednesday": [
                    21600,
                    79200
                ]
            },
            "fax": "(704) 398-3704",
            "flags": [
                "inStockPageView",
                "metaProducts",
                "productAisleNumber",
                "productBayNumber",
                "productLocations",
                "storeMapView"
            ],
            "isOnboardedStore": "false",
            "latitude": "35.44830916",
            "longitude": "-81.00445884",
            "mapUrl": "http://mobilepti.lowes.com/staticmap?api_key=977b2334e34711e18edc12313d23aa6e&devid=%7BdevId%7D&storeId=2636&size=1024x1024",
            "milesToStore": "10.69",
            "phone": "(704) 398-3701",
            "state": "NC",
            "storeName": "LOWE'S OF E. LINCOLN COUNTY, NC",
            "storeNumber": "2636",
            "zip": "28037"
        }
    ]
}
```
