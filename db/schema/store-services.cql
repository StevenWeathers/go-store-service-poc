CREATE KEYSPACE storeservices WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': '1'
};

CREATE TABLE storeservices.stores (
  "storeid" int,
  "address1" varchar,
  "city" varchar,
  "country" varchar,
  "fridayopen" int,
  "fridayclose" int,
  "mondayopen" int,
  "mondayclose" int,
  "saturdayopen" int,
  "saturdayclose" int,
  "sundayopen" int,
  "sundayclose" int,
  "thursdayopen" int,
  "thursdayclose" int,
  "tuesdayopen" int,
  "tuesdayclose" int,
  "wednesdayopen" int,
  "wednesdayclose" int,
  "fax" varchar,
  "flags" list<varchar>,
  "isonboardedstore" varchar,
  "latitude" varchar,
  "longitude" varchar,
  "mapurl" varchar,
  "milestostore" varchar,
  "phone" varchar,
  "state" varchar,
  "storename" varchar,
  "storenumber" varchar,
  "zip" varchar,
  PRIMARY KEY ("storeid")
);

CREATE INDEX ON storeservices.stores("storenumber");
