#!/usr/bin/bash

make_request() {
        curl --silent --request GET "127.0.0.1:8080/api/v1/roll${1}"
}

declare -a DICE
DICE=(4 5 6 8 10 12 20 100)

echo "Default values: "
curl --silent --request GET "127.0.0.1:8080/api/v1/roll"
echo

# Check with the variable, without queries.
echo "Variable for the sides., no queries"
for dice in ${DICE[@]}
do
    make_request "/d${dice}"
    make_request "/D${dice}"
done
echo

# Check with the variable with query for the count.
echo "Variable for the sides with count query."
for dice in ${DICE[@]}
do
    make_request "/d${dice}?count=2"
    make_request "/D${dice}?count=2"
done
echo

# Check with variable for the sides and with variable for the count.
echo "Variable for the sides and for the count."
for dice in ${DICE[@]}
do
    make_request "/d${dice}/2"
    make_request "/D${dice}/2"
done
echo


# Check without the variable with query for the sides.
echo "Query for the sides."
for dice in ${DICE[@]}
do
    make_request "?sides=${dice}"
done
echo

# Check without the variable with query for the count.
echo "Query for the count."
make_request "?count=2"
echo

# Check without the variable with query for the sides and the count.
echo "Query for the sides and the count."
for dice in ${DICE[@]}
do
    make_request "?sides=${dice}&count=2"
done
echo
