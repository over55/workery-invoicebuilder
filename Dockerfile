# The base go-image
FROM golang:1.18-alpine

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as described:
# go build will build an executable file in the current directory
RUN go build -o workery-invoicebuilder .

EXPOSE 8000

# Run the executable
CMD [ "/app/workery-invoicebuilder", "serve" ]


## BUILD:
##   $ docker build -f Dockerfile -t workery-invoicebuilder:latest .
##
## EXECUTE:
##   $ docker run -it --rm -p 1337:80 workery-invoicebuilder:latest
