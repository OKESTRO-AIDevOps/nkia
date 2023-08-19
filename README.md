# nkia

don't get me wrong. I love Kubernetes. it's just that i have a lot of spare time.

# Offcially: eNabling Kubernetes Integration Architecture

*Unofficially but preferably: No Kube In the A-S*


## nkia api

Here, kubernetes api has been abstracted and auxiliary functions needed to interact with them\
are bundled so that the nkia server or orchestrator agent can easily communicate with\
the abstraction layer


## secure server & orchestrator (http/https/ws/wss)

Here, nkia-api is wrapped within a server that implements \
KubeConfig X509 Data based Challange Protocol (KCXD), which consists of\
Single Terminal Transfer Challenge (STTC) and Multi Terminal Socket Challenge (MTSC)\
in order to provide a practical and secure option for utilizing npia-api \
with or without external reverse proxy managing https or wss 


## secure client (http/https/ws/wss)

Here, client can be either a command line executable that communicates with the\
nkia http/https server or a orchestrator agent & browser js that communicate with\
the nkia ws/wss orchestrator 


# Documentation

Refer to [/docs](https://okestro-aidevops.github.io/nkia/)
