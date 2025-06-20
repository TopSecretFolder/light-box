#!/bin/bash
#
# Script to configure Wi-Fi using NetworkManager (nmcli) from a file in /boot.
# This is the modern method for recent Raspberry Pi OS versions.
#

set -e # Exit immediately if a command exits with a non-zero status.

echo "Starting NetworkManager Wi-Fi setup script..."

CREDS_FILE="/boot/wifi_credentials.txt"

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

# 4. To ensure a clean slate, delete ALL existing Wi-Fi profiles.
# This is simpler and more robust than trying to match a specific profile.
echo "Searching for and deleting all existing Wi-Fi connection profiles..."
# The output of the nmcli command is redirected to a while loop.
# This is safer for names that might contain spaces.
nmcli -g NAME,TYPE c show | grep ':802-11-wireless$' | cut -d':' -f1 | while read -r conn_name; do
    if [ -n "$conn_name" ]; then
        echo "Deleting profile: '$conn_name'"
        nmcli connection delete "$conn_name"
    fi
done

# 5. Create a new connection and activate it.
# This single command creates the profile, saves it, and connects.
# The 'ifname wlan0' argument makes it specific to the Wi-Fi interface.
echo "Attempting to create and connect to new Wi-Fi network '$SSID'..."
nmcli device wifi connect "$SSID" password "$PSK" ifname wlan0

# nmcli will return a non-zero exit code on failure, which 'set -e' will catch.
echo "Successfully connected to '$SSID'."

echo "Wi-Fi setup script finished."
exit 0
