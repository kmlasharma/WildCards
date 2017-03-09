import test_utils

def test_correct_pml():
    app = test_utils.start_app()
    app.input_pml_file('~/test.pml')
    app.input_owl_file('~/test.owl')
