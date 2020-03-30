#!/bin/bash

# Instantiate an AWX server using the AWX Installer:
# - This installer is powered by docker-composer
# - AWX will be available as http://localhost:80

rc=$(docker 2&>/dev/null; echo $?)
if [[ $rc -ne 0 ]]; then
  echo "You MUST have docker installed"
fi

git clone https://github.com/ansible/awx.git

ansible-playbook -i awx/installer/inventory awx/installer/install.yml
