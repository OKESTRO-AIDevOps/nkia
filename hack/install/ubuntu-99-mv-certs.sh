USER=<user-on-host> 
mkdir -p /etc/kubernetes/pki/etcd 

mv /home/${USER}/ca.crt /etc/kubernetes/pki/ 

mv /home/${USER}/ca.key /etc/kubernetes/pki/ 

mv /home/${USER}/sa.pub /etc/kubernetes/pki/ 

mv /home/${USER}/sa.key /etc/kubernetes/pki/ 

mv /home/${USER}/front-proxy-ca.crt /etc/kubernetes/pki/ 

mv /home/${USER}/front-proxy-ca.key /etc/kubernetes/pki/ 

mv /home/${USER}/etcd-ca.crt /etc/kubernetes/pki/etcd/ca.crt 

mv /home/${USER}/etcd-ca.key /etc/kubernetes/pki/etcd/ca.key 

mv /home/${USER}/admin.conf /etc/kubernetes/admin.conf 