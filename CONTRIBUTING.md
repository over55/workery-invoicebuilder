1. Run the following to setup our dependencies.

  ```
  virtualenv -p python3.6 env
  source env/bin/activate
  pip install PyPDF2      # https://github.com/mstamy2/PyPDF2
  pip install reportlab
  pip install grpcio
  pip install grpcio-tools
  pip freeze > requirements.txt
  ```


2. Run the following code to compile our `protos` definition for python code.

  ```
  $ cd workery-invoicebuilder
  $ python -m grpc_tools.protoc -Iprotos --python_out=. --grpc_python_out=. protos/invoice.proto
  ```
