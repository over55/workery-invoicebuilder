FROM python:3.6

ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

# Install dependencies:
COPY requirements.txt .
RUN pip install -r requirements.txt

# Run the application:
COPY . .
CMD ["python", "invoicebuilder_server.py"]





## Using the production Dockerfile, build and tag the Docker image:
##   $ docker build -f Dockerfile -t workery-invoicebuilder:latest .
##
## Spin up the container:
##   $ docker run -it --rm -p 1337:80 workery-invoicebuilder:latest
##
## Navigate to http://localhost:1337/ in your browser to view the app.
##
## SPECIAL THANKS TO:
## https://pythonspeed.com/articles/activate-virtualenv-dockerfile/
