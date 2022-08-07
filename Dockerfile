#-------------#
# Build Stage #
#-------------#

# First pull Golang image
FROM golang:1.18-alpine as build-env

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as described:
# go build will build an executable file in the current directory
RUN go build -o workery-invoicebuilder .


# Budild Linux 64bit application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/workery-invoicebuilder /app/main.go

#-----------#
# Run Stage #
#-----------#

FROM alpine:3.14


# Copy only required data into this image
COPY --from=build-env /app/workery-invoicebuilder .

# Expose application port
EXPOSE 8001

# Start app
CMD ["./workery-invoicebuilder", "serve"]

# SPECIAL THANKS:
# https://www.bacancytechnology.com/blog/dockerize-golang-application

#-------------------------------------------------------------------------------------------------------------
# HOWTO: BUILD AN IMAGE.
# docker build -t bmika/workery-invoicebuilder:1.0 .

# HOWTO: RUN A CONTAINER.
# docker run -d -p 8001:8001 --name=workery-invoicebuilder -e WORKERY_INVOICEBUILDER_IP="0.0.0.0" WORKERY_INVOICEBUILDER_PORT="8001" WORKERY_INVOICEBUILDER_PDF_TEMPLATE_FILE_PATH="/app/Template_Live_v2.pdf" bmika/workery-invoicebuilder:1.0
#--------------------------------------------------------------------------------------------------------------
