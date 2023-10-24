# NKIA API eXpression



## Below are nokubectl specific flags



**nokudectl specific flags cannot be used in combination with each other or other options,** 

**aka mutually exclusive**



**ex) nokubectl {flag}**





```yaml
- init:  

        description:  initiate  (or re-init) a nokubectl config and runtime directory 

- help:  

        description:  given command line argument, prints out relevant information 

- interactive:  

        description:  (recommended) run command in interactive mode 

- apix-md:  

        description:  exports all apix information to a markdown file 

- apix-js:  

        description:  exports all apix information to an importable js file 

- apix-py:  

        description:  exports all apix information to an importable py file 






```

## Below are orch.io request specific options**



**ex) nokubectl {arg1} {arg2} {arg3} {--req_option_name} {req_option_val}**





```yaml
- to:  

        description:  target cluster registered on orch.io server  

- as:  

        description:  request option to apply when making request to orch.io, currently available [ 'admin' ] 


```

## Below are all available queries and corresponding required options



**queries are made of arguments and then joined by the options** 

**ex) nokubectl {arg1} {arg2} {arg3} {--option_name} {option_val}**





```yaml
- conncheck:  

        description:  checking orch.io connection                 

        options: 

                # this is admin query, which is used 

                # with [ --as admin ]  





- keygen:  

        description:  generate and receive private key                 

        options: 

                # this is admin query, which is used 

                # with [ --as admin ]  





- addcluster:  

        description:  generate and receive token for connecting a working cluster             

        options: 

                # this is admin query, which is used 

                # with [ --as admin ]  





- admin install env:  

        description:  set up installation environment          

        options: 

                --localip:  local ip adress 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install mainctrl:  

        description:  set up main control plane          

        options: 

                --localip:  local ip adress 

                --osnm:  OS distro version 

                --cv:  Kubernetes version at the same time container runtime version 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install subctrl prep:  

        description:  prepare sub control plane     

        options: 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install subctrl add:  

        description:  set up sub sontrol plane     

        options: 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                --localip:  local ip adress 

                --osnm:  OS distro version 

                --cv:  Kubernetes version at the same time container runtime version 

                --token:  join token (the whole 'print token' output) 

                --nrole:  node role 

                --nid:  node id (given node name) 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install worker:  

        description:  set up worker            

        options: 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                --localip:  local ip adress 

                --osnm:  OS distro version 

                --cv:  Kubernetes version at the same time container runtime version 

                --token:  join token (the whole 'print token' output) 

                --nrole:  node role 

                --nid:  node id (given node name) 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install volume prep:  

        description:  prepare volume         

        options: 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                --localip:  local ip adress 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install volume add:  

        description:  set up volume        

        options: 

                --targetip:  ssh server ip 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install toolkit:  

        description:  set up toolkit           

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install log:  

        description:  check set up logs         

        options: 

                --islocal:  indicates wheter or not we need remote ssh connection [ true | false ] 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install lock get:  

        description:  get lock info on installation process      

        options: 

                --islocal:  indicates wheter or not we need remote ssh connection [ true | false ] 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin install lock set:  

        description:  set lock info on installation process     

        options: 

                --islocal:  indicates wheter or not we need remote ssh connection [ true | false ] 

                --targetip:  ssh server ip 

                --targetid:  ssh user id 

                --targetpw:  ssh user password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin init:  

        description:  initiate admin function             

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- admin init log:  

        description:  get admin initiation log          

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- setting create namespace:  

        description:  create namespace          

        options: 

                --ns:  namespace 

                --repoaddr:  repository address  

                --regaddr:  registry address 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- setting set repo:  

        description:  set repository per project        

        options: 

                --ns:  namespace 

                --repoaddr:  repository address  

                --repoid:  repository id 

                --repopw:  repository password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- setting set reg:  

        description:  set registry per project         

        options: 

                --ns:  namespace 

                --regaddr:  registry address 

                --regid:  registry id 

                --regpw:  registry password 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- setting create monitoring:  

        description:  create monitoring         

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- toolkit build:  

        description:  build from docker-compose.yaml          

        options: 

                --ns:  namespace 

                --repoaddr:  repository address  

                --regaddr:  registry address 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- toolkit build log:  

        description:  get build log       

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource nodes:  

        description:  get nodes           

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource pods:  

        description:  get pods           

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource pods log:  

        description:  get logs from a specific pod          

        options: 

                --ns:  namespace 

                --podnm:  pod name 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource service:  

        description:  get serivces           

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource deployment:  

        description:  get deployments           

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource event:  

        description:  get events          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource resource:  

        description:  get resorces          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource namespace:  

        description:  get namespace          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource ingress:  

        description:  get ingress         

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource nodeport:  

        description:  get nodeport        

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource pod scheduled:  

        description:  get json info on scheduled pods          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource pod unscheduled:  

        description:  get json info on unscheduled pods        

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container cpu:  

        description:  get json info on cpu usage by containers          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container mem:  

        description:  get json info on memory usage by containers          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container fs read:  

        description:  get json info on filesystem read by containers          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container fs write:  

        description:  get json info on filesystem write by containers          

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container net receive:  

        description:  get json info on network receive by containers         

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource container net transmit:  

        description:  get json info on network transmission by containers         

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource volume available:  

        description:  get json info on volume availability     

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource volume capacity:  

        description:  get json info on volume capacity        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource volume used:  

        description:  get json info on volume usage        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node temperature:  

        description:  get json info on node temperature         

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node temperature change:  

        description:  get json info on node temperature change       

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node temperature average:  

        description:  get json info on node temperature average       

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node process:  

        description:  get json info on node processes        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node cores:  

        description:  get json info on node cores        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node mem:  

        description:  get json info on node memory          

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node mem total:  

        description:  get json info on total node memory       

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node disk read:  

        description:  get json info on node disk read        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node disk write:  

        description:  get json info on node disk write        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node net receive:  

        description:  get json info on network receive        

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node net transmit:  

        description:  get json info on network transmission         

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- resource node disk written:  

        description:  get json info on bytes written per node       

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply reg secret:  

        description:  add registry secret to a cluster           

        options: 

                --ns:  namespace 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply distro:  

        description:  deploy a project onto a cluster             

        options: 

                --ns:  namespace 

                --repoaddr:  repository address  

                --regaddr:  registry address 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply create operation source:  

        description:  create operation source file for various actions        

        options: 

                --ns:  namespace 

                --repoaddr:  repository address  

                --regaddr:  registry address 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply restart:  

        description:  restart (update) deployment          

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply rollback:  

        description:  rollback deployment         

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply kill:  

        description:  delete deployment             

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply net refresh:  

        description:  restart cluster network         

        options: 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply hpa:  

        description:  apply hpa to a deployment              

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply hpa undo:  

        description:  remove hpa from a deployment            

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply qos:  

        description:  apply qos to a deployment              

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply qos undo:  

        description:  remove qos (to default) from a deployment            

        options: 

                --ns:  namespace 

                --resource:  resource type name such as [ deployment | service ] 

                --resourcenm:  resource name that has a specific resource type 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply ingress:  

        description:  apply ingress to a service             

        options: 

                --ns:  namespace 

                --hostnm:  host name 

                --svcnm:  service name 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply ingress undo:  

        description:  remove ingress from a service  

        options: 

                --ns:  namespace 

                --hostnm:  host name 

                --svcnm:  service name 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply nodeport:  

        description:  apply nodeport to a deployment           

        options: 

                --ns:  namespace 

                --svcnm:  service name 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  





- apply nodeport undo:  

        description:  remove nodeport from a deployment         

        options: 

                --ns:  namespace 

                --svcnm:  service name 

                # this is nkia query, which is used 

                # with [ --to $CLUSTER_ID ] usually  






```

