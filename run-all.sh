#!/bin/bash

BASE_DIR=$(pwd)
PIDS=()

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

cleanup() {
  echo -e "${RED}Caught interrupt signal. Cleaning up...${NC}"
  pkill main || true
  pkill rust-rest || true
  pkill rest-exe || true
  pkill node || true
  exit 1
}

trap cleanup SIGINT

# Declare two indexed arrays for directories and their respective commands
dirs=(
  "$BASE_DIR/rest/go"
  "$BASE_DIR/rest/rust"
  "$BASE_DIR/rest/haskell"
  "$BASE_DIR/rest/javascript"
  #... add more as needed
)

commands=(
  'go build main.go && ./main'
  'cargo build --release && ./target/release/rust-rest'
  'stack build && stack exec rest-exe'
  '&& node index.js'
  #... add corresponding commands for the directories above
)

pkill main || true
pkill rust-rest || true
pkill rest-exe || true
pkill node || true

# Loop through each directory
for index in "${!dirs[@]}"; do
  dir="${dirs[$index]}"
  cmd="${commands[$index]}"
  
  # Extract the name from the directory path (e.g., "go", "rust", etc.)
  name=$(basename "$dir")
  # Generate the log filename
  timestamp=$(date +"%Y%m%d%H%M%S")
  LOGFILE="$BASE_DIR/logs/${name}_$timestamp.log"

  echo -e "${YELLOW}Processing directory: $dir${NC}"
  cd "$dir" || exit 1  # Exit the script if cd fails
  
  # Split the combined command into build and run
  CMD[0]=$(echo $cmd | awk -F'&&' '{print $1}')
  CMD[1]=$(echo $cmd | awk -F'&&' '{print $2}')

  # Build the app
  echo -e "${GREEN}Building... ${CMD[0]}${NC}"
  eval "${CMD[0]}"
  
  # Run the app in the background (&)
  echo -e "${GREEN}Starting the app... ${CMD[1]}${NC}"
  eval "${CMD[1]}" &

  # Get the PID of the last background process
  APP_PID=$!
  PIDS+=($APP_PID)

  # Give the application some time to start (adjust as needed)
  sleep 2

  # Run the benchmark and log results
  echo -e "${YELLOW}Benchmarking...${NC}"
  "$BASE_DIR/wrk/./wrk" -t12 -c400 -d15s http://127.0.0.1:8080/compute >> "$LOGFILE"

  # Kill the app
  echo -e "${RED}Stopping the app...${NC}"
  pkill main || echo "Failed to kill main"
  pkill rust-rest || echo "Failed to kill rust-rest"
  pkill rest-exe || echo "Failed to kill rest-exe"
  pkill node || echo "Failed to kill node"
  
  # Give some time before moving to the next directory
  sleep 5
done

echo -e "${GREEN}All benchmarks completed.${NC}"
