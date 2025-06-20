#!/bin/bash
#
# Script to configure wpa_supplicant from a file in /boot.
#

set -e # Exit immediately if a command exits with a non-zero status.

echo "Starting Wi-Fi setup script..."

CREDS_FILE="/boot/wifi_credentials.txt"
WPA_CONFIG_FILE="/etc/wpa_supplicant/wpa_supplicant.conf"

# 1. Check if the credentials file exists
if [ ! -f "$CREDS_FILE" ]; then
    echo "Error: Credentials file not found at $CREDS_FILE. Aborting."
    exit 1
fi

# 2. Read credentials from the file, removing any carriage returns (^M)
# Using 'source' is not safe here, so we parse it manually.
COUNTRY=$(grep -E '^COUNTRY=' "$CREDS_FILE" | cut -d'=' -f2- | tr -d '\r')
SSID=$(grep -E '^SSID=' "$CREDS_FILE" | cut -d'=' -f2- | tr -d '\r')
PSK=$(grep -E '^PSK=' "$CREDS_FILE" | cut -d'=' -f2- | tr -d '\r')

# 3. Validate that the variables were read correctly
if [ -z "$COUNTRY" ] || [ -z "$SSID" ] || [ -z "$PSK" ]; then
    echo "Error: Could not read COUNTRY, SSID, or PSK from $CREDS_FILE. Please check the file format."
    exit 1
fi

echo "Credentials read successfully. SSID: $SSID"

# 4. Generate the wpa_supplicant configuration block
# wpa_passphrase generates a secured, hashed PSK from the plain-text password.
NETWORK_BLOCK=$(wpa_passphrase "$SSID" "$PSK")

# 5. Check if wpa_passphrase succeeded
if [ $? -ne 0 ]; then
    echo "Error: wpa_passphrase command failed. Could not generate network block."
    exit 1
fi

echo "wpa_passphrase block generated."

# 6. Write the new configuration file
# This overwrites the existing file to ensure a clean state on every boot.
# The header is required for the Pi's networking to work correctly.
cat > "$WPA_CONFIG_FILE" << EOF
country=$COUNTRY
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1

$NETWORK_BLOCK
EOF

echo "Successfully wrote new configuration to $WPA_CONFIG_FILE"

# 7. Restart the wlan0 interface using the modern ip command.
echo "Restarting wlan0 interface to apply new settings..."
ip link set wlan0 down
sleep 1
ip link set wlan0 up

echo "Wi-Fi setup script finished."

exit 0
