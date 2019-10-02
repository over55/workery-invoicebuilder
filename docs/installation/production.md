# CentOS 7 - Installation and Configuration
## Overview
The purpose of this section is to provide details on how to setup the ``workery-invoicebuilder`` project to work for **CentOS** operating system. This project assumes the following:

* You have followed the instructions on setting up ``workery-backend`` on a **CentOS 7** server.
* You created a service account in **CentOS 7** for the name ``django``.
* Your server supports ``Python3``.

## Installation
1. While being logged in as the ``techops`` or ``root`` user account, run the following code. This will switch your user account to be the service user account.

        su - django

2. Clone our project.

        cd ~/
        git clone https://github.com/over55/workery-invoicebuilder.git

3. Setup our virtual environment.

        cd workery-invoicebuilder
        virtualenv -p python3.6 env
        source env/bin/activate

4. Install our projects dependencies via ``pip``.

        pip install requirements.txt

5. Confirm the server will start up by running the following in your terminal.

        python invoicebuilder_server.py

6. You should see the application run successfully. Now terminate it by holding the following keys on your keyboard ``CTRL + X``.

7. Type the following to switch back to your ``techops`` or ``root`` user.

        exit

8. We will now modify ``systemd`` so our server will start on every computer boot / restart

        sudo vi /etc/systemd/system/invoicebuilder.service

9. Copy and paste the following code into your terminal and save it.

        [Unit]
        Description=Workery InvoiceBuilder Daemon
        After=network.target

        [Service]
        User=django
        Group=nginx
        WorkingDirectory=/opt/django/workery-invoicebuilder
        ExecStart=/opt/django/workery-invoicebuilder/scripts/startup.sh

        [Install]
        WantedBy=multi-user.target

10. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    sudo systemctl start invoicebuilder
    sudo systemctl enable invoicebuilder
    ```

11. Confirm our service is running.

    ```
    systemctl status invoicebuilder.service
    ```

12. If the service is working correctly you should see something like this at the bottom:

        raspberrypi systemd[1]: Started Workery InvoiceBuilder Daemon.

13. Congradulations, you have setup the microservice!

14. If you see any problems, run the following service to see what is wrong. More information can be found in [this article](https://unix.stackexchange.com/a/225407).

        sudo journalctl -u invoicebuilder

8. To reload the latest modifications to systemctl file.

        systemctl daemon-reload
