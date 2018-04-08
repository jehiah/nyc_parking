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

YMD=$(date +%Y-%m-%d)
YMD=2017-12-24
YMD=2017-01-13
YMD=2018-04-08

if [ ! -e data/locations_$YMD.csv ]; then
    echo "downloading locations_$YMD.csv"
    curl --silent -o data/locations_$YMD.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/locations.CSV" || exit 1
fi
if [ ! -e data/signs_$YMD.csv ]; then
    echo "downloading signs_$YMD.csv"
    curl --silent -o data/signs_$YMD.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/signs.CSV" || exit 1
fi

echo "running nyc_parking analysis"
./nyc_parking --signs=data/signs_$YMD.csv --locations=data/locations_$YMD.csv --sign_summary > data/sign_data_$YMD.json
if [ "$?" != 0 ]; then
    echo "error running ./nyc_parking"
    exit 1
fi
echo "generating report"
./generate_report.py --current-file=data/sign_data_$YMD.json
if [ "$?" != 0 ]; then
    echo "error running ./generate_report.py"
    exit 1
fi
