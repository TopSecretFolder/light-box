#!/bin/bash
#
# Generic script to download and install a remote systemd service file.
# Usage: sudo bash install-remote-service.sh <URL>
#

set -e

# --- VALIDATION ---
# Check if the script is run as root
if [ "$(id -u)" -ne 0 ]; then
  echo "This script must be run with root privileges. Please use sudo." >&2
  exit 1
fi

# Check if a URL argument is provided
if [ -z "$1" ]; then
  echo "Error: No URL provided." >&2
  echo "Usage: $0 <URL_to_service_file>" >&2
  exit 1
fi

# --- SCRIPT LOGIC ---
REMOTE_URL="$1"
SERVICE_NAME=$(basename "$REMOTE_URL")
DEST_PATH="/etc/systemd/system/$SERVICE_NAME"

echo "--- Installing Remote Service ---"
echo "Source URL: $REMOTE_URL"
echo "Destination: $DEST_PATH"

# Download the service file
echo "Downloading..."
if curl -sSL -f -o "$DEST_PATH" "$REMOTE_URL"; then
  echo "Download successful."
else
  echo "Error: Failed to download from $REMOTE_URL." >&2
  exit 1
fi

# Reload systemd and enable the new service
echo "Reloading systemd daemon..."
systemctl daemon-reload

echo "Enabling service '$SERVICE_NAME'..."
systemctl enable "$SERVICE_NAME"

echo "--- Service '$SERVICE_NAME' installed and enabled successfully! ---"
exit 0
