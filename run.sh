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

YMD=$(date +%Y-%m-%d)
PREVIOUSYMD=2017-01-13
PREVIOUSYMD=2017-12-24
#PREVIOUSYMD=2018-04-08
#PREVIOUSYMD=2018-07-29
PREVIOUSYMD=2018-10-18
YMD=2019-07-10

if [ ! -e data/locations_$YMD.csv ]; then
    echo "downloading locations_$YMD.csv"
    curl --silent -o data/locations_$YMD.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/locations.CSV" || exit 1
fi
if [ ! -e data/signs_$YMD.csv ]; then
    echo "downloading signs_$YMD.csv"
    curl --silent -o data/signs_$YMD.csv "http://a841-dotweb01.nyc.gov/datafeeds/ParkingReg/signs.CSV" || exit 1
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