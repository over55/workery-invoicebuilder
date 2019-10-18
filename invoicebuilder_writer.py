from PyPDF2 import PdfFileWriter, PdfFileReader
import io
from reportlab.pdfgen import canvas
from reportlab.lib.pagesizes import letter

WORKERY_INVOICE_FILENAME = "Skills_At_Work_Invoice_Template_-_Workery.pdf"


class WorkeryInvoicePDFWriter:
    """
    Class used to generate standard workery invoice PDF.
    """
    filename = ""
    data = dict()

    def __init__(self, filename=None):
        """
        Default constructor.
        """
        self.filename = filename

    def set_data(self, data):
        self.data = data

    def set_filename(self, filename):
        self.filename = filename

    def render(self, packet):
        """
        Utility function which will take the data and display it on the page
        using the `Workery Invoice` template structure.
        """
        can = canvas.Canvas(packet, pagesize=letter)
        can.drawString(465, 734, str(self.data['invoice_id']))
        can.drawString(349, 717, str(self.data['invoice_date']))
        can.drawString(400, 696, str(self.data['associate_name']))
        can.drawString(400, 676, str(self.data['associate_telephone']))
        can.drawString(140, 656, str(self.data['client_name']))
        can.drawString(110, 636, str(self.data['client_address']))
        can.drawString(109, 614, str(self.data['client_telephone']))
        can.drawString(270, 614, str(self.data['client_email']))
        can.drawString(80, 579, str(self.data['line_01_qty']))
        can.drawString(130, 579, str(self.data['line_01_desc']))
        can.drawString(393, 579, str(self.data['line_01_price']))
        can.drawString(485, 579, str(self.data['line_01_amount']))
        can.drawString(80, 563, str(self.data['line_02_qty']))
        can.drawString(130, 563, str(self.data['line_02_desc']))
        can.drawString(393, 563, str(self.data['line_02_price']))
        can.drawString(485, 563, str(self.data['line_02_amount']))
        can.drawString(80, 548, str(self.data['line_03_qty']))
        can.drawString(130, 548, str(self.data['line_03_desc']))
        can.drawString(393, 548, str(self.data['line_03_price']))
        can.drawString(485, 548, str(self.data['line_03_amount']))
        can.drawString(80, 533, str(self.data['line_04_qty']))
        can.drawString(130, 533, str(self.data['line_04_desc']))
        can.drawString(393, 533, str(self.data['line_04_price']))
        can.drawString(485, 533, str(self.data['line_04_amount']))
        can.drawString(80, 518, str(self.data['line_05_qty']))
        can.drawString(130, 518, str(self.data['line_05_desc']))
        can.drawString(393, 518, str(self.data['line_05_price']))
        can.drawString(485, 518, str(self.data['line_05_amount']))
        can.drawString(80, 503, str(self.data['line_06_qty']))
        can.drawString(130, 503, str(self.data['line_06_desc']))
        can.drawString(393, 503, str(self.data['line_06_price']))
        can.drawString(485, 503, str(self.data['line_06_amount']))
        can.drawString(80, 488, str(self.data['line_07_qty']))
        can.drawString(130, 488, str(self.data['line_07_desc']))
        can.drawString(393, 488, str(self.data['line_07_price']))
        can.drawString(485, 488, str(self.data['line_07_amount']))
        can.drawString(80, 473, str(self.data['line_08_qty']))
        can.drawString(130, 473, str(self.data['line_08_desc']))
        can.drawString(393, 473, str(self.data['line_08_price']))
        can.drawString(485, 473, str(self.data['line_08_amount']))
        can.drawString(80, 458, str(self.data['line_09_qty']))
        can.drawString(130, 458, str(self.data['line_09_desc']))
        can.drawString(393, 458, str(self.data['line_09_price']))
        can.drawString(485, 458, str(self.data['line_09_amount']))
        can.drawString(80, 443, str(self.data['line_10_qty']))
        can.drawString(130, 443, str(self.data['line_10_desc']))
        can.drawString(393, 443, str(self.data['line_10_price']))
        can.drawString(485, 443, str(self.data['line_10_amount']))
        can.drawString(80, 428, str(self.data['line_11_qty']))
        can.drawString(130, 428, str(self.data['line_11_desc']))
        can.drawString(393, 428, str(self.data['line_11_price']))
        can.drawString(485, 428, str(self.data['line_11_amount']))
        can.drawString(80, 413, str(self.data['line_12_qty']))
        can.drawString(130, 413, str(self.data['line_12_desc']))
        can.drawString(393, 413, str(self.data['line_12_price']))
        can.drawString(485, 413, str(self.data['line_12_amount']))
        can.drawString(80, 398, str(self.data['line_13_qty']))
        can.drawString(130, 398, str(self.data['line_13_desc']))
        can.drawString(393, 398, str(self.data['line_13_price']))
        can.drawString(485, 398, str(self.data['line_13_amount']))
        can.drawString(80, 383, str(self.data['line_14_qty']))
        can.drawString(130, 383, str(self.data['line_14_desc']))
        can.drawString(393, 383, str(self.data['line_14_price']))
        can.drawString(485, 383, str(self.data['line_14_amount']))
        can.drawString(80, 369, str(self.data['line_15_qty']))
        can.drawString(130, 369, str(self.data['line_15_desc']))
        can.drawString(393, 369, str(self.data['line_15_price']))
        can.drawString(485, 369, str(self.data['line_15_amount']))
        can.drawString(185, 339, str(self.data['invoice_quote_days']))
        can.drawString(220, 293, str(self.data['invoice_associate_tax']))
        can.drawString(215, 267, str(self.data['invoice_quote_date']))
        can.drawString(180, 243, str(self.data['invoice_customers_approval']))
        can.drawString(58, 188, str(self.data['line_01_notes']))
        can.drawString(58, 158, str(self.data['line_02_notes']))
        can.drawString(450, 349, str(self.data['total_labour']))
        can.drawString(450, 323, str(self.data['total_materials']))
        can.drawString(450, 303, str(self.data['waste_removal']))
        can.drawString(450, 278, str(self.data['sub_total'])) # SUBTOTAL
        can.drawString(450, 259, str(self.data['tax']))
        can.drawString(448, 229, str(self.data['total']))
        can.drawString(448, 157, str(self.data['deposit']))
        can.drawString(430, 135, str(self.data['amount_due'])) # AMOUNT DUE - self.data['payment_amount']
        can.drawString(180, 135, str(self.data['payment_date']))
        can.drawString(79, 111, str(self.data['cash']))
        can.drawString(180, 111, str(self.data['cheque']))
        can.drawString(284, 111, str(self.data['debit']))
        can.drawString(386, 111, str(self.data['credit']))
        can.drawString(488, 111, str(self.data['other']))
        can.drawString(250, 93, str(self.data['client_signature']))
        can.drawString(105, 67, str(self.data['associate_sign_date']))
        can.drawString(382, 67, str(self.data['associate_signature']))
        can.drawString(478, 28, str(self.data['work_order_id']))
        can.save()

    def generate_pdf(self):
        """
        Function will generate a PDF based on the inputted `data` and return
        the PDF.

        Special thanks: https://stackoverflow.com/a/29266450
        """
        packet = io.BytesIO()

        # create a new PDF with Reportlab
        self.render(packet)

        # move to the beginning of the StringIO buffer
        packet.seek(0)
        new_pdf = PdfFileReader(packet)
        # read your existing PDF
        existing_pdf = PdfFileReader(open(WORKERY_INVOICE_FILENAME, "rb"))
        output = PdfFileWriter()
        # add the "watermark" (which is the new pdf) on the existing page
        page = existing_pdf.getPage(0)
        page2 = new_pdf.getPage(0)
        page.mergePage(page2)
        output.addPage(page)
        # finally, write "output" to a real file
        outputStream = open(self.filename, "wb")
        output.write(outputStream)
        outputStream.close()
