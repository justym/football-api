#!/bin/sh

#create user information and db
mongo admin </mongo_seed/init.js

#import datasets(.json)
#(1)clubs#
#mongoimport --host mongodb --db footballDB --collection clubs --type json --file /mongo_seed/datasets/20*/en.1.clubs.json --jsonArray
mongoimport --host mongodb --db footballDB --collection clubs --type json --file /mongo_seed/datasets/2010-11/en.1.clubs.json
