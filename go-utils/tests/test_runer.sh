#!/bin/sh

Help()
{
   echo "It allows you to see the results and summary of all the tests in your project."
   echo
   echo "options:"
   echo "all      Print the test output verbose mode."
   echo "sum      Print the test summury mode."
   echo "help     Print this Help."
   echo
}

if [[ $# -eq 0 ]] ; then
    echo 'some message'
    exit 1
elif [[ $1 == "sum" ]] ; then
all=$(go test ./...  -v)

SUB_STRING1="PASS"
SUB_STRING2="FAIL"
SUB_STRING3="RUN"
s1=${all//"$SUB_STRING1"}
successful="$(((${#all} - ${#s1}) / ${#SUB_STRING1}))"
s2=${all//"$SUB_STRING2"}
fail="$(((${#all} - ${#s2}) / ${#SUB_STRING2}))"
s3=${all//"$SUB_STRING3"}
total="$(((${#all} - ${#s3}) / ${#SUB_STRING3}))"


echo "Test Count:" \
$'\n\t'"Success Test: $successful" \
$'\n\t'"Fail    Test: $fail" \
$'\n\t'"Total   Test: $total" \ | sed \
-e "s/Test Count:/"$'\e[1;35m'"&"$'\e[m'"/" \
-e "s/Success Test: $successful/"$'\e[1;32m'"&"$'\e[m'"/" \
-e "s/Fail    Test: $fail/"$'\e[1;31m'"&"$'\e[m'"/" \
-e "s/Total   Test: $total/"$'\e[1;34m'"&"$'\e[m'"/" 
elif [[ $1 == "all" ]] ; then
    go test ./...  -v
elif [[ $1 == "help" ]] ; then
    Help
else
    echo "Wrong argument"
fi
