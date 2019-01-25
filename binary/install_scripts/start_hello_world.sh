#!/bin/bash
cp /home/ubuntu/install_scripts/systemd_unit_file/hello_world.service /etc/systemd/system/
systemctl daemon-reload
systemctl start hello_world.service
systemctl status hello_world.service

