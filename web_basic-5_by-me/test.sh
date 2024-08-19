#!/bin/bash
# Set IP and port variables
IP="127.0.0.1"
PORT="8080"
USER_NAME="John"
USER_DATA='{"name": "'"$USER_NAME"'", "email": "sample@gmail", "password": "123456"}'
UPDATE_USER_DATA='{"name": "", "age": 40}'

# Health check
printf "\nHealth check...\n"
curl -X GET $IP:$PORT



# Test POST request
printf "\n\n\nTesting POST request...\n"
curl -X POST -H "Content-Type: application/json" -d "$USER_DATA" $IP:$PORT/user

# Test GET request
printf "\n\n\nTesting GET request...\n"
curl -X GET $IP:$PORT/user/$USER_NAME

# # Test PUT request
# printf "\n\n\nTesting PUT request...\n"
# curl -X PUT -H "Content-Type: application/json" -d "$USER_DATA" $IP:$PORT/user

# # Test DELETE request
printf "\n\n\nTesting DELETE request...\n"
curl -X DELETE $IP:$PORT/user/$USER_NAME