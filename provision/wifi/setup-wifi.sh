#!/bin/bash
#
# Script to configure Wi-Fi using NetworkManager (nmcli) from a file in /boot.
# Modified for robustness when run as a systemd service.
#

set -e # Exit immediately if a command exits with a non-zero status.

echo "Starting NetworkManager Wi-Fi setup script..."

CREDS_FILE="/boot/firmware/wifi_credentials.txt"

# 1. Check if the credentials file exists
if [ ! -f "$CREDS_FILE" ]; then
    echo "Error: Credentials file not found at $CREDS_FILE. Aborting."
    exit 1
fi

# 2. Read credentials from the file, removing any carriage returns (^M)
SSID=$(grep -E '^SSID=' "$CREDS_FILE" | cut -d'=' -f2- | tr -d '\r')
PSK=$(grep -E '^PSK=' "$CREDS_FILE" | cut -d'=' -f2- | tr -d '\r')

# 3. Validate that the variables were read correctly
if [ -z "$SSID" ] || [ -z "$PSK" ]; then
    echo "Error: Could not read SSID or PSK from $CREDS_FILE. Please check the file format."
    exit 1
fi

echo "Credentials read successfully for SSID: $SSID"

# ==================== NEW: WAIT AND RESCAN SECTION ====================
# Give the Wi-Fi adapter a few seconds to power on and initialize.
echo "Waiting 5 seconds for the Wi-Fi hardware to initialize..."
sleep 5

# Explicitly tell NetworkManager to perform a new Wi-Fi scan.
echo "Forcing a new Wi-Fi scan..."
# The 'nmcli device wifi rescan' command tells NetworkManager to look for networks right now.
nmcli device wifi rescan
echo "Scan complete."
# ======================================================================

# 4. To ensure a clean slate, delete ALL existing Wi-Fi profiles.
echo "Searching for and deleting all existing Wi-Fi connection profiles..."
nmcli -g NAME,TYPE c show | grep ':802-11-wireless$' | cut -d':' -f1 | while read -r conn_name; do
    if [ -n "$conn_name" ]; then
        echo "Deleting profile: '$conn_name'"
        nmcli connection delete "$conn_name"
    fi
done

# 5. Create a new connection and activate it.
# This loop will now retry the connection up to 4 times.
echo "Attempting to create and connect to new Wi-Fi network '$SSID'..."
for i in {1..4}; do
    # Use 'set +e' to temporarily disable exit-on-error for the connection attempt
    set +e
    nmcli device wifi connect "$SSID" password "$PSK" ifname wlan0
    
    # Check the exit code of the last command. 0 means success.
    if [ $? -eq 0 ]; then
        # Re-enable exit-on-error
        set -e
        echo "Successfully connected to '$SSID'."
        echo "Wi-Fi setup script finished."
        exit 0 # Exit the script successfully
    fi
    # Re-enable exit-on-error
    set -e

    # If it failed, wait before retrying
    if [ $i -lt 4 ]; then
      echo "Connection failed (Attempt $i). Retrying in 5 seconds..."
      sleep 5
      # Also rescan before the next attempt
      nmcli device wifi rescan
    fi
done

echo "Could not connect to the network after several attempts. Aborting."
exit 1 # Exit the script with an error status
