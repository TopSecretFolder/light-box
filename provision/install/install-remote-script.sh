#!/bin/bash
#
# Generic script to download and install a remote script to /usr/local/bin.
# Usage: sudo bash install-remote-script.sh <URL>
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
  echo "Usage: $0 <URL_to_script>" >&2
  exit 1
fi

# --- SCRIPT LOGIC ---
REMOTE_URL="$1"
# Extract the filename from the URL (e.g., http://.../setup-wifi.sh -> setup-wifi.sh)
SCRIPT_NAME=${$(basename "$REMOTE_URL")%%\?*}
DEST_PATH="/usr/local/bin/$SCRIPT_NAME"

echo "--- Installing Remote Script ---"
echo "Source URL: $REMOTE_URL"
echo "Destination: $DEST_PATH"

# Download the file using curl
echo "Downloading..."
if curl -sSL -f -o "$DEST_PATH" "$REMOTE_URL"; then
  echo "Download successful."
else
  echo "Error: Failed to download from $REMOTE_URL." >&2
  exit 1
fi

# Make the script executable
echo "Setting executable permissions..."
chmod +x "$DEST_PATH"

echo "--- Script '$SCRIPT_NAME' installed successfully! ---"
exit 0
