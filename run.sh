#!/bin/bash

# [ ! -e /usr/local/bin/brew ] && echo "brew required" && exit 1
# if ! brew list postgresql >/dev/null; then 
#     echo "brew install postgresql"
#     brew install postgresql || exit 1
# else
#     echo "postgresql installed"
# fi

# if brew services list postgresql | grep -q stopped; then
#     echo "brew services start postgresql"
#     brew services start postgresql || exit 1
# else
#     brew services list postgresql
# fi

# YMD=$(date +%Y-%m-%d)
PREVIOUSYMD=2017-01-13
PREVIOUSYMD=2017-12-24
PREVIOUSYMD=2018-04-08
PREVIOUSYMD=2018-07-29
PREVIOUSYMD=2018-10-18
PREVIOUSYMD=2019-07-10
PREVIOUSYMD=2019-08-03
PREVIOUSYMD=2019-11-29
YMD=${YMD:-2020-01-26}

# through nov 2019 - https://data.cityofnewyork.us/Transportation/Parking-Regulation-Locations-and-Signs/xswq-wnv9/data
# URLBASE="http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg" # files ended in CSV not csv
URLBASE="https://www1.nyc.gov/html/dot/downloads/ParkReg"

if [ ! -e data/locations_$YMD.csv ]; then
    echo "downloading locations_$YMD.csv"
    curl --silent -o data/locations_$YMD.csv "$URLBASE/locations.csv" || exit 1
    # sed -i -e 's/\r/\r\n/g' data/locations_$YMD.csv
fi
if [ ! -e data/signs_$YMD.csv ]; then
    echo "downloading signs_$YMD.csv"
    curl --silent -o data/signs_$YMD.csv "$URLBASE/signs.csv" || exit 1
fi

echo "running nyc_parking analysis"
echo "./nyc_parking --signs=data/signs_$YMD.csv --locations=data/locations_$YMD.csv --sign_summary > data/sign_data_$YMD.json"
./nyc_parking --signs=data/signs_$YMD.csv --locations=data/locations_$YMD.csv --sign_summary > data/sign_data_$YMD.json
if [ "$?" != 0 ]; then
    echo "error running ./nyc_parking"
    exit 1
fi

ARGS="--current-file=data/sign_data_$YMD.json"
if [ -e "data/sign_data_$PREVIOUSYMD.json" ]; then
    ARGS="$ARGS --previous-file=data/sign_data_$PREVIOUSYMD.json"
fi
echo "./generate_report.py $ARGS"
./generate_report.py $ARGS
if [ "$?" != 0 ]; then
    echo "error running ./generate_report.py $ARGS"
    exit 1
fi

# ack 'WEST\s+44 STREET,\s*9' data/locations_2019-07-10.csv 

# 'Order'
# S-087878, S-089254