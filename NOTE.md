# note

# TODO

- install log status check and attach to session log



# 20231114

keygen from the website


# 20231115 


./nokubectl --as admin orch conncheck


./nokubectl --as admin orch add cl --clusterid test-cs


./nokubectl --as admin orch install cl --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.94 --osnm ubuntu --cv 1.27 --updatetoken 8c5b871631794bc2e5dfe3b8004b924a


./nokubectl --as admin orch install cl log --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu 


./nokubectl --to test-cs admin install worker --targetip 192.168.50.95:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.95 --osnm ubuntu --cv 1.27 --token -


./nokubectl --to test-cs admin install log --targetip 192.168.50.95:22 --targetid ubuntu --targetpw ubuntu


./nokubectl --to test-cs resource nodes --ns -


./nokubectl --to test-cs install volume --localip 192.168.50.94 