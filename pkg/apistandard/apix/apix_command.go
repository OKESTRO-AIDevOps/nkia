package apix

var APIX_COMMAND = "" +
	/*a00*/ "init                                              : initiate nokubectl                " + "\n" +
	/*a01*/ "set                                               : set nokubectl cluster id                " + "\n" +
	/*a02*/ "set-opts                                          : set nokubectl options for orch io                " + "\n" +
	/*b00*/ "orch-conncheck                                    : checking orch.io connection                " + "\n" +
	/*b01*/ "orch-keygen                                       : generate and receive private key                " + "\n" +
	/*b02*/ "orch-get-cl                                       : receive cluster names that are available for communication" + "\n" +
	/*b03*/ "orch-add-cl                                       : generate and receive token for connecting a working cluster  " + "\n" +
	/*b04*/ "orch-install-cl                                   : install main control plane through orch.io       " + "\n" +
	/*b05*/ "orch-install-cl-log                               : get installation log for the main control plane through orch.io     " + "\n" +
	/*c00*/ "io-connect-update                                 : start nokubelt with update token         " + "\n" +
	/*c01*/ "io-connect                                        : start nokubelt with token         " + "\n" +
	/*d00*/ "install-mainctrl                                  : install control plane using nokubeadm        " + "\n" +
	/*d01*/ "install-worker                                    : install worker using nokubeadm        " + "\n" +
	/*f00*/ "admin-install-worker                              : remote hook for installing worker        " + "\n" +
	/*d02*/ "install-volume                                    : install volume using nokubeadm       " + "\n" +
	/*f01*/ "admin-install-volume                              : remote hook for installing worker       " + "\n" +
	/*d03*/ "install-toolkit                                   : install build toolkit using nokubeadm        " + "\n" +
	/*f02*/ "admin-install-toolkit                             : remote hook for installing toolkit       " + "\n" +
	/*d04*/ "install-log                                       : get install log using nokubeadm       " + "\n" +
	/*f03*/ "admin-install-log                                 : remote hook for getting log       " + "\n" +
	/*f04*/ "admin-init                                        : initiate admin function            " + "\n" +
	/*f05*/ "admin-init-log                                    : get admin initiation log         " + "\n" +
	/*g00*/ "setting-create-namespace                          : create namespace         " + "\n" +
	/*g01*/ "setting-set-repo                                  : set repository per project       " + "\n" +
	/*g02*/ "setting-set-reg                                   : set registry per project        " + "\n" +
	/*g03*/ "setting-create-volume                             : set volume provisioner       " + "\n" +
	/*g04*/ "setting-create-monitoring                         : set monitoring api without persistent data        " + "\n" +
	/*g05*/ "setting-create-monitoring-persist                 : set monitoring api with persistent data    " + "\n" +
	/*h00*/ "toolkit-build                                     : build from docker-compose.yaml         " + "\n" +
	/*h01*/ "toolkit-build-log                                 : get build log      " + "\n" +
	/*h02*/ "toolkit-pipe                                      : start pipe from .npia/build.yaml          " + "\n" +
	/*h03*/ "toolkit-pipe-log                                  : get pipe log       " + "\n" +
	/*h04*/ "toolkit-pipe-set-var                              : set pipe variable    " + "\n" +
	/*h05*/ "toolkit-pipe-get-var                              : get variables set for pipe    " + "\n" +
	/*i00*/ "resource-nodes                                    : get nodes          " + "\n" +
	/*i01*/ "resource-pods                                     : get pods          " + "\n" +
	/*i02*/ "resource-pods-log                                 : get logs from a specific pod         " + "\n" +
	/*i03*/ "resource-service                                  : get serivces          " + "\n" +
	/*i04*/ "resource-deployment                               : get deployments          " + "\n" +
	/*i05*/ "resource-event                                    : get events         " + "\n" +
	/*i06*/ "resource-resource                                 : get resorces         " + "\n" +
	/*i07*/ "resource-namespace                                : get namespace         " + "\n" +
	/*i08*/ "resource-ingress                                  : get ingress        " + "\n" +
	/*i09*/ "resource-nodeport                                 : get nodeport       " + "\n" +
	/*i10*/ "resource-pod-scheduled                            : get json info on scheduled pods         " + "\n" +
	/*i11*/ "resource-pod-unscheduled                          : get json info on unscheduled pods       " + "\n" +
	/*i12*/ "resource-container-cpu                            : get json info on cpu usage by containers         " + "\n" +
	/*i13*/ "resource-container-mem                            : get json info on memory usage by containers         " + "\n" +
	/*i14*/ "resource-container-fs-read                        : get json info on filesystem read by containers         " + "\n" +
	/*i15*/ "resource-container-fs-write                       : get json info on filesystem write by containers         " + "\n" +
	/*i16*/ "resource-container-net-receive                    : get json info on network receive by containers        " + "\n" +
	/*i17*/ "resource-container-net-transmit                   : get json info on network transmission by containers        " + "\n" +
	/*i18*/ "resource-volume-available                         : get json info on volume availability    " + "\n" +
	/*i19*/ "resource-volume-capacity                          : get json info on volume capacity       " + "\n" +
	/*i20*/ "resource-volume-used                              : get json info on volume usage       " + "\n" +
	/*i21*/ "resource-node-temperature                         : get json info on node temperature        " + "\n" +
	/*i22*/ "resource-node-temperature-change                  : get json info on node temperature change      " + "\n" +
	/*i23*/ "resource-node-temperature-average                 : get json info on node temperature average      " + "\n" +
	/*i24*/ "resource-node-process                             : get json info on node processes       " + "\n" +
	/*i25*/ "resource-node-cores                               : get json info on node cores       " + "\n" +
	/*i26*/ "resource-node-mem                                 : get json info on node memory         " + "\n" +
	/*i27*/ "resource-node-mem-total                           : get json info on total node memory      " + "\n" +
	/*i28*/ "resource-node-disk-read                           : get json info on node disk read       " + "\n" +
	/*i29*/ "resource-node-disk-write                          : get json info on node disk write       " + "\n" +
	/*i30*/ "resource-node-net-receive                         : get json info on network receive       " + "\n" +
	/*i31*/ "resource-node-net-transmit                        : get json info on network transmission        " + "\n" +
	/*i32*/ "resource-node-disk-written                        : get json info on bytes written per node      " + "\n" +
	/*j00*/ "apply-reg-secret                                  : add registry secret to a cluster          " + "\n" +
	/*j01*/ "apply-distro                                      : deploy a project onto a cluster            " + "\n" +
	/*j02*/ "apply-create-operation-source                     : create operation source file for various actions       " + "\n" +
	/*j03*/ "apply-restart                                     : restart (update) deployment         " + "\n" +
	/*j04*/ "apply-rollback                                    : rollback deployment        " + "\n" +
	/*j05*/ "apply-kill                                        : delete deployment            " + "\n" +
	/*j06*/ "apply-net-refresh                                 : restart cluster network        " + "\n" +
	/*j07*/ "apply-hpa                                         : apply hpa to a deployment             " + "\n" +
	/*j08*/ "apply-hpa-undo                                    : remove hpa from a deployment           " + "\n" +
	/*j09*/ "apply-qos                                         : apply qos to a deployment             " + "\n" +
	/*j10*/ "apply-qos-undo                                    : remove qos (to default) from a deployment           " + "\n" +
	/*j11*/ "apply-ingress                                     : apply ingress to a service            " + "\n" +
	/*j12*/ "apply-ingress-undo                                : remove ingress from a service " + "\n" +
	/*j13*/ "apply-nodeport                                    : apply nodeport to a deployment          " + "\n" +
	/*j14*/ "apply-nodeport-undo                               : remove nodeport from a deployment        " + "\n" +
	""
