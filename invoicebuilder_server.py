"""
The following code is responsible for generating PDF files based on the
`workery standard invoice template` PDF file and exposing this ability through
a `gRPC` endpoint.
"""
from concurrent import futures
import logging

import grpc

import invoice_pb2
import invoice_pb2_grpc
import invoicebuilder_writer


CHUNK_SIZE = 1024 * 1024  # 1MB


def get_file_chunks(filename):
    """
    Utlity function opens up a binary file for reading and converts it
    to our chunks byte stream. Special thanks to the following link:
    https://github.com/gooooloo/grpc-file-transfer
    """
    with open(filename, 'rb') as f:
        while True:
            piece = f.read(CHUNK_SIZE);
            if len(piece) == 0:
                return
            yield invoice_pb2.Chunk(buffer=piece)


class InvoiceBuilder(invoice_pb2_grpc.InvoiceBuilderServicer):

    def GeneratePDF(self, request, context):
        filename = "bin/"+request.workOrderId+"-"+request.invoiceId+"-invoice.pdf"
        writer = invoicebuilder_writer.WorkeryInvoicePDFWriter(filename)
        writer.set_data({
            'invoice_id': request.invoiceId,
            'invoice_date': request.invoiceDate,
            'associate_name': request.associateName,
            'associate_telephone': request.associateTelephone,
            'client_name': request.clientName,
            'client_address': request.clientAddress,
            'client_telephone': request.clientTelephone,
            'client_email': request.clientEmail,
            'line_01_qty': request.line01Quantity,
            'line_01_desc': request.line01Description,
            'line_01_price': request.line01Price,
            'line_01_amount': request.line01Amount,
            'line_02_qty': request.line02Quantity,
            'line_02_desc': request.line02Description,
            'line_02_price': request.line02Price,
            'line_02_amount': request.line02Amount,
            'line_03_qty': request.line03Quantity,
            'line_03_desc': request.line03Description,
            'line_03_price': request.line03Price,
            'line_03_amount': request.line03Amount,
            'line_04_qty': request.line04Quantity,
            'line_04_desc':request.line04Description,
            'line_04_price': request.line04Price,
            'line_04_amount': request.line04Amount,
            'line_05_qty': request.line05Quantity,
            'line_05_desc': request.line05Description,
            'line_05_price': request.line05Price,
            'line_05_amount': request.line05Amount,
            'line_06_qty': request.line06Quantity,
            'line_06_desc': request.line06Description,
            'line_06_price': request.line06Price,
            'line_06_amount': request.line06Amount,
            'line_07_qty': request.line07Quantity,
            'line_07_desc': request.line07Description,
            'line_07_price': request.line07Price,
            'line_07_amount': request.line07Amount,
            'line_08_qty': request.line08Quantity,
            'line_08_desc': request.line08Description,
            'line_08_price': request.line08Price,
            'line_08_amount': request.line08Amount,
            'line_09_qty': request.line09Quantity,
            'line_09_desc': request.line09Description,
            'line_09_price': request.line09Price,
            'line_09_amount': request.line09Amount,
            'line_10_qty': request.line10Quantity,
            'line_10_desc': request.line10Description,
            'line_10_price': request.line10Price,
            'line_10_amount': request.line10Amount,
            'line_11_qty': request.line11Quantity,
            'line_11_desc': request.line11Description,
            'line_11_price': request.line11Price,
            'line_11_amount': request.line11Amount,
            'line_12_qty': request.line12Quantity,
            'line_12_desc': request.line12Description,
            'line_12_price': request.line12Price,
            'line_12_amount': request.line12Amount,
            'line_13_qty': request.line13Quantity,
            'line_13_desc': request.line13Description,
            'line_13_price': request.line13Price,
            'line_13_amount': request.line13Amount,
            'line_14_qty': request.line14Quantity,
            'line_14_desc': request.line14Description,
            'line_14_price': request.line14Price,
            'line_14_amount': request.line14Amount,
            'line_15_qty': request.line15Quantity,
            'line_15_desc': request.line15Description,
            'line_15_price': request.line15Price,
            'line_15_amount': request.line15Amount,
            'invoice_quote_days': request.invoiceQuoteDays,
            'invoice_associate_tax': request.invoiceAssociateTax,
            'invoice_quote_date': request.invoiceQuoteDate,
            'invoice_customers_approval': request.invoiceCustomersApproval,
            'line_01_notes': request.line01Notes,
            'line_02_notes': request.line02Notes,
            'total_labour': request.totalLabour,
            'total_materials': request.totalMaterials,
            'waste_removal': request.wasteRemoval,
            'sub_total': request.subTotal,
            'tax': request.tax,
            'total': request.total,
            'deposit': request.deposit,
            'payment_amount': request.paymentAmount,
            'payment_date': request.paymentDate,
            'cash': request.cash,
            'cheque': request.cheque,
            'debit': request.debit,
            'credit': request.credit,
            'other': request.other,
            'client_signature': request.clientSignature,
            'associate_sign_date': request.associateSignDate,
            'associate_signature': request.associateSignature,
            'work_order_id': request.workOrderId,
        })
        writer.generate_pdf()
        return get_file_chunks(filename)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    invoice_pb2_grpc.add_InvoiceBuilderServicer_to_server(InvoiceBuilder(), server)
    server.add_insecure_port('[::]:50051')
    server.start()

    # The following code will grant gRPC control of our applications main
    # runtime loop else our application recieves a keyboard interruption
    # which means we will need to stop our gRPC server and stop our
    # application.
    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(0)
    except Exception as e:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()
