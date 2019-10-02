#!/bin/bash
# This is a startup script used by the "workery-invoicebuilder" microservice
# to startup as a seperate processes. The following assumptions must be met
# before running this script:

# (1) This project must be located in the following directory:
#     "/opt/django/workery-invoicebuilder"
# (2) The "env" variable has already been initialized with the projects
#     requirements.txt file.
# (3) The operating system is "CentOS 7".
# (4) There is a "django" user account created on the "CentOS 7" system.
#

cd /opt/django/workery-invoicebuilder
source env/bin/activate
exec python grpc_server.py
