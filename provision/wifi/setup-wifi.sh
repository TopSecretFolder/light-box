#!/bin/bash
#
# Script to configure Wi-Fi using NetworkManager (nmcli) by trying multiple
# credential files found in /boot/firmware/.
#

set -e # Exit immediately if a command exits with a non-zero status.

echo "Starting NetworkManager Wi-Fi setup script..."

CREDS_PATH_PATTERN="/boot/firmware/wifi_credentials_*.txt"

# 1. Check if any credential files exist
# Use 'ls' to check for files matching the pattern. Redirect output to /dev/null.
# If ls returns a non-zero status, no files were found.
if ! ls $CREDS_PATH_PATTERN >/dev/null 2>&1; then
    echo "Error: No Wi-Fi credential files found matching '$CREDS_PATH_PATTERN'. Aborting."
    exit 1
fi

# ==================== WAIT AND RESCAN SECTION ====================
# Give the Wi-Fi adapter a few seconds to power on and initialize.
echo "Waiting 5 seconds for the Wi-Fi hardware to initialize..."
sleep 5

# Explicitly tell NetworkManager to perform a new Wi-Fi scan.
echo "Forcing a new Wi-Fi scan..."
nmcli device wifi rescan
echo "Scan complete."
# ======================================================================

# 2. To ensure a clean slate, delete ALL existing Wi-Fi profiles.
echo "Searching for and deleting all existing Wi-Fi connection profiles..."
nmcli -g NAME,TYPE c show | grep ':802-11-wireless$' | cut -d':' -f1 | while read -r conn_name; do
    if [ -n "$conn_name" ]; then
        echo "Deleting profile: '$conn_name'"
        nmcli connection delete "$conn_name"
    fi
done

# 3. Loop through all found credential files and attempt to connect.
echo "Searching for networks to join..."
for creds_file in $CREDS_PATH_PATTERN; do
    echo "--- Processing credentials from: $creds_file ---"

    # Read credentials from the file, removing any carriage returns (^M)
    SSID=$(grep -E '^SSID=' "$creds_file" | cut -d'=' -f2- | tr -d '\r')
    PSK=$(grep -E '^PSK=' "$creds_file" | cut -d'=' -f2- | tr -d '\r')

    # Validate that the variables were read correctly
    if [ -z "$SSID" ] || [ -z "$PSK" ]; then
        echo "Warning: Could not read SSID or PSK from $creds_file. Skipping."
        continue # Move to the next file
    fi

    echo "Attempting to connect to network '$SSID'..."

    # Use 'set +e' to temporarily disable exit-on-error for the connection attempt
    set +e
    nmcli device wifi connect "$SSID" password "$PSK" ifname wlan0

    # Check the exit code of the last command. 0 means success.
    if [ $? -eq 0 ]; then
        # Re-enable exit-on-error
        set -e
        echo "✅ Successfully connected to '$SSID' using credentials from $creds_file."
        echo "Wi-Fi setup script finished."
        exit 0 # Exit the script successfully
    fi
    # Re-enable exit-on-error
    set -e

    echo "Connection to '$SSID' failed. Trying next available credentials..."
done

echo "❌ Could not connect to any network after trying all available credential files. Aborting."
exit 1 # Exit the script with an error status
