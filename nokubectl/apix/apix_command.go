package apix

var APIX_COMMAND = "" +
	// front admin option
	"conncheck                                         : checking orch.io connection                " + "\n" +
	"keygen                                            : generate and receive private key                " + "\n" +
	"addcluster                                        : generate and receive token for connecting a working cluster            " + "\n" +
	// query
	"admin-install-env                                 : set up installation environment         " + "\n" +
	"admin-install-mainctrl                            : set up main control plane         " + "\n" +
	"admin-install-subctrl-prep                        : prepare sub control plane    " + "\n" +
	"admin-install-subctrl-add                         : set up sub sontrol plane    " + "\n" +
	"admin-install-worker                              : set up worker           " + "\n" +
	"admin-install-volume-prep                         : prepare volume        " + "\n" +
	"admin-install-volume-add                          : set up volume       " + "\n" +
	"admin-install-toolkit                             : set up toolkit          " + "\n" +
	"admin-install-log                                 : check set up logs        " + "\n" +
	"admin-install-lock-get                            : get lock info on installation process     " + "\n" +
	"admin-install-lock-set                            : set lock info on installation process    " + "\n" +
	"admin-init                                        : initiate admin function            " + "\n" +
	"admin-init-log                                    : get admin initiation log         " + "\n" +
	"setting-create-namespace                          : create namespace         " + "\n" +
	"setting-set-repo                                  : set repository per project       " + "\n" +
	"setting-set-reg                                   : set registry per project        " + "\n" +
	"setting-create-monitoring                         : create monitoring        " + "\n" +
	"toolkit-build                                     : build from docker-compose.yaml         " + "\n" +
	"toolkit-build-log                                 : get build log      " + "\n" +
	"resource-nodes                                    : get nodes          " + "\n" +
	"resource-pods                                     : get pods          " + "\n" +
	"resource-pods-log                                 : get logs from a specific pod         " + "\n" +
	"resource-service                                  : get serivces          " + "\n" +
	"resource-deployment                               : get deployments          " + "\n" +
	"resource-event                                    : get events         " + "\n" +
	"resource-resource                                 : get resorces         " + "\n" +
	"resource-namespace                                : get namespace         " + "\n" +
	"resource-ingress                                  : get ingress        " + "\n" +
	"resource-nodeport                                 : get nodeport       " + "\n" +
	"resource-pod-scheduled                            : get json info on scheduled pods         " + "\n" +
	"resource-pod-unscheduled                          : get json info on unscheduled pods       " + "\n" +
	"resource-container-cpu                            : get json info on cpu usage by containers         " + "\n" +
	"resource-container-mem                            : get json info on memory usage by containers         " + "\n" +
	"resource-container-fs-read                        : get json info on filesystem read by containers         " + "\n" +
	"resource-container-fs-write                       : get json info on filesystem write by containers         " + "\n" +
	"resource-container-net-receive                    : get json info on network receive by containers        " + "\n" +
	"resource-container-net-transmit                   : get json info on network transmission by containers        " + "\n" +
	"resource-volume-available                         : get json info on volume availability    " + "\n" +
	"resource-volume-capacity                          : get json info on volume capacity       " + "\n" +
	"resource-volume-used                              : get json info on volume usage       " + "\n" +
	"resource-node-temperature                         : get json info on node temperature        " + "\n" +
	"resource-node-temperature-change                  : get json info on node temperature change      " + "\n" +
	"resource-node-temperature-average                 : get json info on node temperature average      " + "\n" +
	"resource-node-process                             : get json info on node processes       " + "\n" +
	"resource-node-cores                               : get json info on node cores       " + "\n" +
	"resource-node-mem                                 : get json info on node memory         " + "\n" +
	"resource-node-mem-total                           : get json info on total node memory      " + "\n" +
	"resource-node-disk-read                           : get json info on node disk read       " + "\n" +
	"resource-node-disk-write                          : get json info on node disk write       " + "\n" +
	"resource-node-net-receive                         : get json info on network receive       " + "\n" +
	"resource-node-net-transmit                        : get json info on network transmission        " + "\n" +
	"resource-node-disk-written                        : get json info on bytes written per node      " + "\n" +
	"apply-reg-secret                                  : add registry secret to a cluster          " + "\n" +
	"apply-distro                                      : deploy a project onto a cluster            " + "\n" +
	"apply-create-operation-source                     : create operation source file for various actions       " + "\n" +
	"apply-restart                                     : restart (update) deployment         " + "\n" +
	"apply-rollback                                    : rollback deployment        " + "\n" +
	"apply-kill                                        : delete deployment            " + "\n" +
	"apply-net-refresh                                 : restart cluster network        " + "\n" +
	"apply-hpa                                         : apply hpa to a deployment             " + "\n" +
	"apply-hpa-undo                                    : remove hpa from a deployment           " + "\n" +
	"apply-qos                                         : apply qos to a deployment             " + "\n" +
	"apply-qos-undo                                    : remove qos (to default) from a deployment           " + "\n" +
	"apply-ingress                                     : apply ingress to a service            " + "\n" +
	"apply-ingress-undo                                : remove ingress from a service " + "\n" +
	"apply-nodeport                                    : apply nodeport to a deployment          " + "\n" +
	"apply-nodeport-undo                               : remove nodeport from a deployment        " + "\n" +
	""
