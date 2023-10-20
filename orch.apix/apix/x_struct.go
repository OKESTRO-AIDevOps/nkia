package apix

type ConnectionCheck struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID   string `json:"id"`
		Args string `json:"args"`
	} `json:"query"`
}

type Keygen struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type AddCluster struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Admin_InstallEnvironment struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID      string `json:"id"`
		LocalIP string `json:"localip"`
	} `json:"query"`
}

type Admin_InstallControlPlane struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		LocalIP  string `json:"localip"`
		OSName   string `json:"osnm"`
		CVersion string `json:"cv"`
	} `json:"query"`
}

type Admin_InstallAnotherControlPlaneOnLocal struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
	} `json:"query"`
}

type Admin_InstallAnotherControlPlaneOnRemote struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
		LocalIP  string `json:"localip"`
		OSName   string `json:"osnm"`
		CVersion string `json:"cv"`
		Token    string `json:"token"`
		NodeRole string `json:"nrole"`
		NodeID   string `json:"nid"`
	} `json:"query"`
}

type Admin_InstallWorker struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
		LocalIP  string `json:"localip"`
		OSName   string `json:"osnm"`
		CVersion string `json:"cv"`
		Token    string `json:"token"`
		NodeRole string `json:"nrole"`
		NodeID   string `json:"nid"`
	} `json:"query"`
}

type Admin_InstallVolumeOnRemote struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
		LocalIP  string `json:"localip"`
	} `json:"query"`
}

type Admin_InstallVolumeOnLocal struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		TargetIP string `json:"targetip"`
	} `json:"query"`
}

type Admin_InstallToolkit struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Admin_InstallLog struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		IsLocal  string `json:"islocal"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
	} `json:"query"`
}

type Admin_InstallLockGet struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		IsLocal  string `json:"islocal"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
	} `json:"query"`
}

type Admin_InstallLockSet struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID       string `json:"id"`
		IsLocal  string `json:"islocal"`
		TargetIP string `json:"targetip"`
		TargetID string `json:"targetid"`
		TargetPW string `json:"targetpw"`
	} `json:"query"`
}

type Admin_Init struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Admin_InitLog struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Setting_CreateNamespace struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RepoAddr  string `json:"repoaddr"`
		RegAddr   string `json:"regaddr"`
	} `json:"query"`
}

type Setting_SetRepository struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RepoAddr  string `json:"repoaddr"`
		RepoID    string `json:"repoid"`
		RepoPW    string `json:"repopw"`
	} `json:"query"`
}

type Setting_SetRegistry struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RegAddr   string `json:"regaddr"`
		RegID     string `json:"regid"`
		RegPW     string `json:"regpw"`
	} `json:"query"`
}

type Setting_CreateMonitoring struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Toolkit_Build struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RepoAddr  string `json:"repoaddr"`
		RegAddr   string `json:"regaddr"`
	} `json:"query"`
}

type Toolkit_BuildLog struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_Nodes struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Pods struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_PodLog struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		PodName   string `json:"podnm"`
	} `json:"query"`
}

type Resource_Service struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Deployment struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Event struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Resource struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Namespace struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_Ingress struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_NodePort struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_PodScheduled struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_PodUnscheduled struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerCPU struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerMemory struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerFSRead struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerFSWrite struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerNetReceive struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_ContainerNetTransmit struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Resource_VolumeAvailable struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_VolumeCapacity struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_VolumeUsed struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeTemperature struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeTemperatureChange struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeTemperatureAverage struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeProcesses struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeCores struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeMemory struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeMemoryTotal struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeDiskRead struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeDiskWrite struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeNetworkReceive struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeNetworkTransmit struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Resource_NodeDiskWritten struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Apply_RegisterSecret struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
	} `json:"query"`
}

type Apply_Distribute struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RepoAddr  string `json:"repoaddr"`
		RegAddr   string `json:"regaddr"`
	} `json:"query"`
}

type Apply_CreateOperationSource struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID        string `json:"id"`
		Namespace string `json:"ns"`
		RepoAddr  string `json:"repoaddr"`
		RegAddr   string `json:"regaddr"`
	} `json:"query"`
}

type Apply_Restart struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_Rollback struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_Kill struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_NetworkRefresh struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID string `json:"id"`
	} `json:"query"`
}

type Apply_HorizontalAutoscale struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_HorizontalAutoscaleUndo struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_QoS struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_QoSUndo struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID           string `json:"id"`
		Namespace    string `json:"ns"`
		Resource     string `json:"resource"`
		ResourceName string `json:"resourcenm"`
	} `json:"query"`
}

type Apply_Ingress struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID          string `json:"id"`
		Namespace   string `json:"ns"`
		HostName    string `json:"hostnm"`
		ServiceName string `json:"svcnm"`
	} `json:"query"`
}

type Apply_IngressUndo struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID          string `json:"id"`
		Namespace   string `json:"ns"`
		HostName    string `json:"hostnm"`
		ServiceName string `json:"svcnm"`
	} `json:"query"`
}

type Apply_NodePort struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID          string `json:"id"`
		Namespace   string `json:"ns"`
		ServiceName string `json:"svcnm"`
	} `json:"query"`
}

type Apply_NodePortUndo struct {
	Target string `json:"target"`
	Option string `json:"option"`
	Query  struct {
		ID          string `json:"id"`
		Namespace   string `json:"ns"`
		ServiceName string `json:"svcnm"`
	} `json:"query"`
}
