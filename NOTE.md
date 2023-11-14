# note


keygen from the website

./nokubectl --as admin orch conncheck

./nokubectl --as admin orch add cl --clusterid test-cs

 - ./nokubectl read

 - 149220e323812e51201998cd90e5e998

./nokubectl --as admin orch install cl --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.94 --osnm ubuntu --cv 1.27 --updatetoken bd7e54e8c578fb4d636d425811670420


./nokubectl --as admin orch install cl log --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu 

./nokubectl --to test-cs admin install worker --targetip 192.168.50.95:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.95 --osnm ubuntu --cv 1.27 --token -

./nokubectl --to test-cs admin install log --targetip 192.168.50.95:22 --targetid ubuntu --targetpw ubuntu