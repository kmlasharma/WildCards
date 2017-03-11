import signal
import time
from os import O_NONBLOCK, R_OK, read, access
from fcntl import fcntl, F_GETFL, F_SETFL
from subprocess import Popen, PIPE, STDOUT

class Application:

    def __init__(self):
        self.log = ''
        self.app = Popen(['app'], stdout=PIPE, stdin=PIPE, stderr=PIPE, shell=False)
        flags = fcntl(self.app.stdout, F_GETFL)
        fcntl(self.app.stdout, F_SETFL, flags | O_NONBLOCK) # set nonblocking flag so we can access output on the fly

    # wait until the given request for input is printed
    def wait_for_line(self, expected, timeout=10):
        signal.signal(signal.SIGALRM, handle_timeout)
        signal.alarm(timeout)
        output = ''
        try:
            while not output.endswith(expected):
                try:
                    line = read(self.app.stdout.fileno(), 1024).decode('utf-8')
                    output += line
                    self.log += line
                except OSError as e:
                    time.sleep(0.5) # output is not readable yet
        except Exception as e:
            print('Timed out waiting for ' + expected  + ', actual: ' + output)
            assert False
        finally:
            signal.alarm(0)

    def command(self, cmd):
        self.app.stdin.write(bytearray(cmd + '\n', 'ascii'))
        self.app.stdin.flush()

    def get_process_start(self):
        signal.signal(signal.SIGALRM, handle_timeout)
        signal.alarm(5)
        line = self.app.stdout.readline()
        while 'Drugs in Process' not in line.decode('utf-8'):
            line = self.app.stdout.readline()

    def get_process_drugs(self):
        self.get_process_start()
        drugs = []
        line = self.app.stdout.readline()
        while 'Analysing OWL file' not in line.decode('utf-8'):
            print(line.decode('utf-8'))
            drugs.insert(len(drugs), line.decode('utf-8').strip())
            line = self.app.stdout.readline()
        return drugs

    def input_pml_file(self, name):
        self.wait_for_line('Enter path to PML File: [default is test.pml] ') 
        self.command(name)

    def input_owl_file(self, name):
        self.wait_for_line('Enter path to OWL File: [default is test.owl] ')
        self.command(name)

def start_app():
    return Application()

def handle_timeout(signum, frame):
    assert False

