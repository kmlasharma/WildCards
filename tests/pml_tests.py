import test_utils

from os import environ
from unittest import TestCase

class Test(TestCase):

    # start app before each test case
    def setUp(self):
        self.res_dir = environ['RES_DIR']
        self.test_pml_file = self.res_dir + '/test.pml'

        self.app = test_utils.start_app()

    # kill app process after each test case
    def tearDown(self):
        self.app.app.kill()

    # test for PML file selection & PML file loading
    def test_loading_pml(self):
        self.app.input_pml_file(self.test_pml_file)
        self.app.wait_for_line('Enter path to OWL File: [default is test.pml] ')

    # test for running PML analyis
    def test_pml_analysis(self):
        self.app.input_pml_file(self.test_pml_file)
        res = self.app.get_process_drugs()
        #print(res)
        assert res == ['Plavix', 'Lipitor', 'Nexium']
        
