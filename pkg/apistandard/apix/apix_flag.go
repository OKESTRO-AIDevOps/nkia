package apix

var APIX_FLAGS = "" +
	"to                : cluster id for orch id" + "\n" +
	"as                : options for orch, it goes like --as a1=a,a2=b,a2=c " + "\n" +
	"clusterid         : target cluster name to be registered" + "\n" +
	"updatetoken       : update token for connection" + "\n" +
	"localip           : local ip adress" + "\n" +
	"osnm              : OS distro version" + "\n" +
	"cv                : Kubernetes version at the same time container runtime version" + "\n" +
	"targetip          : ssh server ip" + "\n" +
	"targetid          : ssh user id" + "\n" +
	"targetpw          : ssh user password" + "\n" +
	"token             : join token (the whole 'print token' output)" + "\n" +
	"nrole             : node role" + "\n" +
	"nid               : node id (given node name)" + "\n" +
	"islocal           : indicates wheter or not we need remote ssh connection [ true | false ]" + "\n" +
	"ns                : namespace" + "\n" +
	"repoaddr          : repository address " + "\n" +
	"regaddr           : registry address" + "\n" +
	"repoid            : repository id" + "\n" +
	"repopw            : repository password" + "\n" +
	"regid             : registry id" + "\n" +
	"regpw             : registry password" + "\n" +
	"varnm             : variable name (for pipe usually)" + "\n" +
	"varval            : variable value (for pipe usually)" + "\n" +
	"podnm             : pod name" + "\n" +
	"resource          : resource type name such as [ deployment | service ]" + "\n" +
	"resourcenm        : resource name that has a specific resource type" + "\n" +
	"hostnm            : host name" + "\n" +
	"svcnm             : service name" + "\n" +
	""
