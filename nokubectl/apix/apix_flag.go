package apix

var APIX_FLAGS = "" +
	"to                : target cluster registered on orch.io server " + "\n" +
	"as                : request option to apply when making request to orch.io, currently available [ 'admin' ]" + "\n" +
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
	"podnm             : pod name" + "\n" +
	"resource          : resource type name such as [ deployment | service ]" + "\n" +
	"resourcenm        : resource name that has a specific resource type" + "\n" +
	"hostnm            : host name" + "\n" +
	"svcnm             : service name" + "\n" +
	""
