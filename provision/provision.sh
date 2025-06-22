#!/bin/bash
#
# Provisioning script for the Raspberry Pi
# This script downloads and installs the necessary components from GitHub to 
# run the light-box services properly
#

GITHUB_USER="TopSecretFolder"
GITHUB_REPO="light-box"
BRANCH="main"

# Derived URL for raw file access
GITHUB_RAW_URL="https://raw.githubusercontent.com/$GITHUB_USER/$GITHUB_REPO/refs/heads/$BRANCH"

# installer scripts
INSTALL_REMOTE_SCRIPT_URL="$GITHUB_RAW_URL/provision/install/install-remote-script.sh"
INSTALL_REMOTE_SERVICE_URL="$GITHUB_RAW_URL/provision/install/install-remote-service.sh"

CB=$((RANDOM % 1000 + 1))
# Wifi service scripts
WIFI_REMOTE_SCRIPT_URL="$GITHUB_RAW_URL/provision/wifi/setup-wifi.sh?cb=$CB"
WIFI_REMOTE_SERVICE_URL="$GITHUB_RAW_URL/provision/wifi/setup-wifi.service?cb=$CB"

# install the wifi script and service
curl -sSL $INSTALL_REMOTE_SCRIPT_URL | sudo bash -s $WIFI_REMOTE_SCRIPT_URL
curl -sSL $INSTALL_REMOTE_SERVICE_URL | sudo bash -s $WIFI_REMOTE_SERVICE_URL

echo ""
echo "--- Installation Complete! ---"
echo ""
echo "ACTION REQUIRED:"
echo "1. Create and edit the credentials file:"
echo "   sudo nano /boot/wifi_credentials.txt"
echo ""
echo "   Add the following content (replace with your details):"
echo "   COUNTRY=US"
echo "   SSID=YourNetworkName"
echo "   PSK=YourNetworkPassword"
echo ""
echo "2. Reboot the Raspberry Pi to apply all changes:"
echo "   sudo reboot"
echo ""

# install docker
# Add Docker's official GPG key:
echo "update deps"
sudo apt-get update
echo "install some stuff"
sudo apt-get install -y ca-certificates curl
echo "add docker keyring"
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo "add docker to apt sources"
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

echo "apt get update"
sudo apt-get update

echo "install docker packages"
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

echo "docker install done"
