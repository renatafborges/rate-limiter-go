#! /bin/bash
status=$?
 
## 1. Run the date command ##
cmd="./rate-limiter-tests"
$cmd
 
## 2. Get exist status  and store into '$status' var ##
status=$?
 
## 3. Now take some decision based upon '$status' ## 
[ $status -eq 0 ] && echo "Automated test executed with sucess" || echo "Failed to execute automated tests"