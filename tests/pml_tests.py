import test_utils
import unittest

class Test(unittest.TestCase):

    # start app before each test case
    def setUp(self):
        self.app = test_utils.start_app()

    # kill app process after each test case
    def tearDown(self):
        self.app.app.kill()

    # test for PML file selection & PML file loading
    def test_loading_pml(self):
        self.app.input_pml_file('../test.pml')
        self.app.wait_for_line('Enter path to OWL File: [default is test.owl] ')

    # test for running PML analyis
    def test_pml_analysis(self):
        self.app.input_pml_file('../test.pml')
        self.app.input_owl_file('../test.owl')
        res = self.app.get_process_drugs()
        print(res)
        assert res == ['Plavix', 'Lipitor', 'Nexium']
