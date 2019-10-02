# Developer Machine
Please follow the instructions provided to get started setting up the server.

1. To get started, clone a copy of the latest code to your development environment.

        git clone https://github.com/over55/workery-invoicebuilder.git

2. Setup your virtual environment.

        cd workery-invoicebuilder
        virtualenv -p python3.6 env
        source env/bin/activate

3. Install our projects dependencies via ``pip``.

        pip install requirements.txt

4. Confirm the server will start up by running the following in your terminal.

        python invoice_builder_server.py

5. (Optional) If you are running **MacOS**, you will get a message asking if Python should accepting income connection from ``Python3``. Please select accept and this program will work.
