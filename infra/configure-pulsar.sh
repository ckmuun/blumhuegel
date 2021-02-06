#!/bin/bash
echo "simple setup script for a given pulsar instance. This does not create the pulsar deployment."

echo "checking if pulsarctl already installed"

if ! command -v pulsarctl &>/dev/null; then
  echo "Pulsarctl could not be found"
  PULSARCTL_TMP="pulsarctl-tmp"
  PULSARCTL_VERSION="0.4.1"
  echo "setup 0 installing pulsarctl v0.4.1"
  mkdir $PULSARCTL_TMP
  cd $PULSARCTL_TMP || exit

  echo "using pulsarctl version $PULSARCTL_VERSION"
  curl -L "https://github.com/streamnative/pulsarctl/releases/download/v$PULSARCTL_VERSION/pulsarctl-386-linux.tar.gz" >>pulsarctl-lix.tar.gz
  tar -xvf pulsarctl-lix.tar.gz

  sudo cp pulsarctl-lix/pulsarctl /usr/bin/pulsarctl
  cd ..
  cd ..
  rmdir --ingore-fail-on-non-empty $PULSARCTL_TMP

  pulsarctl help
  echo "installed pulsarctl"
fi

# configure the cluster for blumhuegel usage

