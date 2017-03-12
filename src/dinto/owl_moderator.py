import subprocess
import sys

stdoutdata = subprocess.getoutput("python3 /go/src/app/dinto/ontology.py " + sys.argv[1])

if(stdoutdata.split()[0][:7] == "WARNING"):
	print("\nError: OWL ontology not correctly structured" + stdoutdata + '\n')
	#sys.stdout = open('errlog.txt', 'a')
	#print(stdoutdata)
else:
	print('\n' + stdoutdata + '\n')
	#sys.stdout = open('log.txt', 'a')
	#print(stdoutdata)

