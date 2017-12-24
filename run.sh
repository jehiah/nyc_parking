#!/bin/bash

[ ! -e /usr/local/bin/brew ] && echo "brew required" && exit 1
if ! brew list postgresql >/dev/null; then 
    echo "brew install postgresql"
    brew install postgresql || exit 1
else
    echo "postgresql installed"
fi

if brew services list postgresql | grep -q stopped; then
    echo "brew services start postgresql"
    brew services start postgresql || exit 1
else
    brew services list postgresql
fi

if [ ! -e data/locations.csv ]; then
    echo "downloading locations.csv"
    curl --silent -o data/locations.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/locations.CSV" || exit 1
fi
if [ ! -e data/signs.csv ]; then
    echo "downloading signs.csv"
    curl --silent -o data/signs.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/signs.CSV" || exit 1
fi

