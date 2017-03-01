import ontospy
import sys

print("Loading OWL ontology...")
print("\n")

model = ontospy.Ontospy(sys.argv[1])

print(model)
print("\n")
print(model.stats())


#print(model.getClass("obo2:OAE_0000218"))
#print("\n")

#model.printClassTree()


