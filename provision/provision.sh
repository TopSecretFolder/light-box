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
echo "docker install script..."
curl -fsSL https://get.docker.com | sudo bash

# add docker group stuff
echo "add docker group"
sudo groupadd docker
sudo usermod -aG docker $USER

# run on boot
echo "start docker and containerd service on boot"
sudo systemctl enable docker.service
sudo systemctl enable containerd.service
