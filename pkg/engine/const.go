// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

const (
	// Mesos is the string constant for MESOS orchestrator type
	Mesos string = "Mesos"
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
)

const (
	// DefaultVNETCIDR is the default CIDR block for the VNET
	DefaultVNETCIDR = "10.0.0.0/8"
	// DefaultInternalLbStaticIPOffset specifies the offset of the internal LoadBalancer's IP
	// address relative to the first consecutive Kubernetes static IP
	DefaultInternalLbStaticIPOffset = 10
	// NetworkPolicyNone is the string expression for the deprecated NetworkPolicy usage pattern "none"
	NetworkPolicyNone = "none"
	// NetworkPolicyCalico is the string expression for calico network policy config option
	NetworkPolicyCalico = "calico"
	// NetworkPolicyCilium is the string expression for cilium network policy config option
	NetworkPolicyCilium = "cilium"
	// NetworkPluginCilium is the string expression for cilium network plugin config option
	NetworkPluginCilium = NetworkPolicyCilium
	// NetworkPolicyAzure is the string expression for Azure CNI network policy manager
	NetworkPolicyAzure = "azure"
	// NetworkPluginAzure is the string expression for Azure CNI plugin
	NetworkPluginAzure = "azure"
	// NetworkPluginKubenet is the string expression for kubenet network plugin
	NetworkPluginKubenet = "kubenet"
	// NetworkPluginFlannel is the string expression for flannel network plugin
	NetworkPluginFlannel = "flannel"
	// DefaultKubeHeapsterDeploymentAddonName is the name of the kube-heapster-deployment addon
	DefaultKubeHeapsterDeploymentAddonName = "kube-heapster-deployment"
	// DefaultKubeDNSDeploymentAddonName is the name of the kube-dns-deployment addon
	DefaultKubeDNSDeploymentAddonName = "kube-dns-deployment"
	// DefaultCoreDNSAddonName is the name of the coredns addon
	DefaultCoreDNSAddonName = "coredns"
	// DefaultDNSAutoscalerAddonName is the name of the coredns addon
	DefaultDNSAutoscalerAddonName = "dns-autoscaler"
	// DefaultKubeProxyAddonName is the name of the kube-proxy config addon
	DefaultKubeProxyAddonName = "kube-proxy-daemonset"
	// DefaultAzureStorageClassesAddonName is the name of the azure storage classes addon
	DefaultAzureStorageClassesAddonName = "azure-storage-classes"
	// DefaultAzureNpmDaemonSetAddonName is the name of the azure npm daemon set addon
	DefaultAzureNpmDaemonSetAddonName = "azure-npm-daemonset"
	// DefaultCalicoDaemonSetAddonName is the name of calico daemonset addon
	DefaultCalicoDaemonSetAddonName = "calico-daemonset"
	// DefaultCiliumDaemonSetAddonName is the name of cilium daemonset addon
	DefaultCiliumDaemonSetAddonName = "cilium-daemonset"
	// DefaultFlannelDaemonSetAddonName is the name of flannel plugin daemonset addon
	DefaultFlannelDaemonSetAddonName = "flannel-daemonset"
	// DefaultAADAdminGroupRBACAddonName is the name of the default admin group RBAC addon
	DefaultAADAdminGroupRBACAddonName = "aad-default-admin-group-rbac"
	// DefaultAzureCloudProviderDeploymentAddonName is the name of the azure cloud provider deployment addon
	DefaultAzureCloudProviderDeploymentAddonName = "azure-cloud-provider-deployment"
	// DefaultAzureCNINetworkMonitorAddonName is the name of the azure cni network monitor addon
	DefaultAzureCNINetworkMonitorAddonName = "azure-cni-networkmonitor"
	// DefaultAuditPolicyAddonName is the name of the audit policy addon
	DefaultAuditPolicyAddonName = "audit-policy"
	// DefaultTillerAddonName is the name of the tiller addon deployment
	DefaultTillerAddonName = "tiller"
	// DefaultAADPodIdentityAddonName is the name of the aad-pod-identity addon deployment
	DefaultAADPodIdentityAddonName = "aad-pod-identity"
	// DefaultACIConnectorAddonName is the name of the aci-connector addon deployment
	DefaultACIConnectorAddonName = "aci-connector"
	// DefaultDashboardAddonName is the name of the kubernetes-dashboard addon deployment
	DefaultDashboardAddonName = "kubernetes-dashboard"
	// DefaultClusterAutoscalerAddonName is the name of the autoscaler addon deployment
	DefaultClusterAutoscalerAddonName = "cluster-autoscaler"
	// DefaultBlobfuseFlexVolumeAddonName is the name of the blobfuse flexvolume addon
	DefaultBlobfuseFlexVolumeAddonName = "blobfuse-flexvolume"
	// DefaultSMBFlexVolumeAddonName is the name of the smb flexvolume addon
	DefaultSMBFlexVolumeAddonName = "smb-flexvolume"
	// DefaultKeyVaultFlexVolumeAddonName is the name of the keyvault flexvolume addon deployment
	DefaultKeyVaultFlexVolumeAddonName = "keyvault-flexvolume"
	// DefaultELBSVCAddonName is the name of the elb service addon deployment
	DefaultELBSVCAddonName = "elb-svc"
	// DefaultGeneratorCode specifies the source generator of the cluster template.
	DefaultGeneratorCode = "aksengine"
	// DefaultReschedulerAddonName is the name of the rescheduler addon deployment
	DefaultReschedulerAddonName = "rescheduler"
	// DefaultHeapsterAddonName is the name of the heapster addon deployment
	DefaultHeapsterAddonName = "heapster"
	// DefaultMetricsServerAddonName is the name of the kubernetes Metrics server addon deployment
	DefaultMetricsServerAddonName = "metrics-server"
	// NVIDIADevicePluginAddonName is the name of the kubernetes NVIDIA Device Plugin daemon set
	NVIDIADevicePluginAddonName = "nvidia-device-plugin"
	// ContainerMonitoringAddonName is the name of the kubernetes Container Monitoring addon deployment
	ContainerMonitoringAddonName = "container-monitoring"
	// AzureCNINetworkMonitoringAddonName is the name of the Azure CNI networkmonitor addon
	AzureCNINetworkMonitoringAddonName = "azure-cni-networkmonitor"
	// AzureNetworkPolicyAddonName is the name of the Azure CNI networkmonitor addon
	AzureNetworkPolicyAddonName = "azure-npm-daemonset"
	// IPMASQAgentAddonName is the name of the ip masq agent addon
	IPMASQAgentAddonName = "ip-masq-agent"
	// DefaultKubernetesKubeletMaxPods is the max pods per kubelet
	DefaultKubernetesKubeletMaxPods = 110
	// DefaultMasterEtcdServerPort is the default etcd server port for Kubernetes master nodes
	DefaultMasterEtcdServerPort = 2380
	// DefaultMasterEtcdClientPort is the default etcd client port for Kubernetes master nodes
	DefaultMasterEtcdClientPort = 2379
	// etcdAccountNameFmt is the name format for a typical etcd account on Cosmos
	etcdAccountNameFmt = "%sk8s"
	// etcdEndpointURIFmt is the name format for a typical etcd account uri
	etcdEndpointURIFmt = "%sk8s.etcd.cosmosdb.azure.com"
)

const (
	//DefaultExtensionsRootURL  Root URL for extensions
	DefaultExtensionsRootURL = "https://raw.githubusercontent.com/Azure/aks-engine/master/"
	// DefaultDockerEngineRepo for grabbing docker engine packages
	DefaultDockerEngineRepo = "https://download.docker.com/linux/ubuntu"
	// DefaultDockerComposeURL for grabbing docker images
	DefaultDockerComposeURL = "https://github.com/docker/compose/releases/download"
)

const (
	//DefaultConfigurationScriptRootURL  Root URL for configuration script (used for script extension on RHEL)
	DefaultConfigurationScriptRootURL = "https://raw.githubusercontent.com/Azure/aks-engine/master/parts/"
)

const (
	// AzureStackSuffix is appended to kubernetes version on Azure Stack instances
	AzureStackSuffix = "-azs"
)

const (
	kubeConfigJSON = "k8s/kubeconfig.json"
	// Windows custom scripts
	kubernetesWindowsAgentCustomDataPS1   = "k8s/kuberneteswindowssetup.ps1"
	kubernetesWindowsAgentFunctionsPS1    = "k8s/kuberneteswindowsfunctions.ps1"
	kubernetesWindowsConfigFunctionsPS1   = "k8s/windowsconfigfunc.ps1"
	kubernetesWindowsKubeletFunctionsPS1  = "k8s/windowskubeletfunc.ps1"
	kubernetesWindowsCniFunctionsPS1      = "k8s/windowscnifunc.ps1"
	kubernetesWindowsAzureCniFunctionsPS1 = "k8s/windowsazurecnifunc.ps1"
	kubernetesWindowsOpenSSHFunctionPS1   = "k8s/windowsinstallopensshfunc.ps1"
)

// cloud-init (i.e. ARM customData) file references
const (
	kubernetesMasterNodeCustomDataYaml = "k8s/cloud-init/masternodecustomdata.yml"
	kubernetesNodeCustomDataYaml       = "k8s/cloud-init/nodecustomdata.yml"
	kubernetesJumpboxCustomDataYaml    = "k8s/cloud-init/jumpboxcustomdata.yml"
	kubernetesCSEMainScript            = "k8s/cloud-init/artifacts/cse_main.sh"
	kubernetesCSEHelpersScript         = "k8s/cloud-init/artifacts/cse_helpers.sh"
	kubernetesCSEInstall               = "k8s/cloud-init/artifacts/cse_install.sh"
	kubernetesCSEConfig                = "k8s/cloud-init/artifacts/cse_config.sh"
	kubernetesCISScript                = "k8s/cloud-init/artifacts/cis.sh"
	kubernetesCSECustomCloud           = "k8s/cloud-init/artifacts/cse_customcloud.sh"
	kubernetesHealthMonitorScript      = "k8s/cloud-init/artifacts/health-monitor.sh"
	// kubernetesKubeletMonitorSystemdTimer     = "k8s/cloud-init/artifacts/kubelet-monitor.timer" // TODO enable
	kubernetesKubeletMonitorSystemdService   = "k8s/cloud-init/artifacts/kubelet-monitor.service"
	kubernetesDockerMonitorSystemdTimer      = "k8s/cloud-init/artifacts/docker-monitor.timer"
	kubernetesDockerMonitorSystemdService    = "k8s/cloud-init/artifacts/docker-monitor.service"
	kubernetesMountEtcd                      = "k8s/cloud-init/artifacts/mountetcd.sh"
	kubernetesMasterGenerateProxyCertsScript = "k8s/cloud-init/artifacts/generateproxycerts.sh"
	kubernetesCustomSearchDomainsScript      = "k8s/cloud-init/artifacts/setup-custom-search-domains.sh"
	sshdConfig                               = "k8s/cloud-init/artifacts/sshd_config"
	sshdConfig1604                           = "k8s/cloud-init/artifacts/sshd_config_1604"
	kubeletSystemdService                    = "k8s/cloud-init/artifacts/kubelet.service"
	kmsSystemdService                        = "k8s/cloud-init/artifacts/kms.service"
	aptPreferences                           = "k8s/cloud-init/artifacts/apt-preferences"
	dockerClearMountPropagationFlags         = "k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf"
	systemdBPFMount                          = "k8s/cloud-init/artifacts/sys-fs-bpf.mount"
	etcdSystemdService                       = "k8s/cloud-init/artifacts/etcd.service"
	etcIssue                                 = "k8s/cloud-init/artifacts/etc-issue"
	etcIssueNet                              = "k8s/cloud-init/artifacts/etc-issue.net"
	cisNetEnforcement                        = "k8s/cloud-init/artifacts/sysctl-d-60-CIS.conf"
	cisLogEnforcement                        = "k8s/cloud-init/artifacts/rsyslog-d-60-CIS.conf"
	modprobeConfCIS                          = "k8s/cloud-init/artifacts/modprobe-CIS.conf"
	pwQuality                                = "k8s/cloud-init/artifacts/pwquality-CIS.conf"
	pamDotDSU                                = "k8s/cloud-init/artifacts/pam-d-su"
	profileDCISSh                            = "k8s/cloud-init/artifacts/profile-d-cis.sh"
	pamDotDCommonAuth                        = "k8s/cloud-init/artifacts/pam-d-common-auth"
	pamDotDCommonPassword                    = "k8s/cloud-init/artifacts/pam-d-common-password"
	auditdRules                              = "k8s/cloud-init/artifacts/auditd-rules"
)

const (
	dcosCustomData188       = "dcos/dcoscustomdata188.t"
	dcosCustomData190       = "dcos/dcoscustomdata190.t"
	dcosCustomData198       = "dcos/dcoscustomdata198.t"
	dcosCustomData110       = "dcos/dcoscustomdata110.t"
	dcosProvision           = "dcos/dcosprovision.sh"
	dcosWindowsProvision    = "dcos/dcosWindowsProvision.ps1"
	dcosProvisionSource     = "dcos/dcosprovisionsource.sh"
	dcos2Provision          = "dcos/bstrap/dcosprovision.sh"
	dcos2BootstrapProvision = "dcos/bstrap/bootstrapprovision.sh"
	dcos2CustomData1110     = "dcos/bstrap/dcos1.11.0.customdata.t"
	dcos2CustomData1112     = "dcos/bstrap/dcos1.11.2.customdata.t"
)

const (
	swarmProvision            = "swarm/configure-swarm-cluster.sh"
	swarmWindowsProvision     = "swarm/Install-ContainerHost-And-Join-Swarm.ps1"
	swarmModeProvision        = "swarm/configure-swarmmode-cluster.sh"
	swarmModeWindowsProvision = "swarm/Join-SwarmMode-cluster.ps1"
)

const (
	agentOutputs                  = "agentoutputs.t"
	agentParams                   = "agentparams.t"
	armParameters                 = "k8s/armparameters.t"
	dcosAgentResourcesVMAS        = "dcos/dcosagentresourcesvmas.t"
	dcosWindowsAgentResourcesVMAS = "dcos/dcosWindowsAgentResourcesVmas.t"
	dcosAgentResourcesVMSS        = "dcos/dcosagentresourcesvmss.t"
	dcosWindowsAgentResourcesVMSS = "dcos/dcosWindowsAgentResourcesVmss.t"
	dcosAgentVars                 = "dcos/dcosagentvars.t"
	dcosBaseFile                  = "dcos/dcosbase.t"
	dcosParams                    = "dcos/dcosparams.t"
	dcosMasterResources           = "dcos/dcosmasterresources.t"
	dcosMasterVars                = "dcos/dcosmastervars.t"
	dcos2BaseFile                 = "dcos/bstrap/dcosbase.t"
	dcos2BootstrapVars            = "dcos/bstrap/bootstrapvars.t"
	dcos2BootstrapParams          = "dcos/bstrap/bootstrapparams.t"
	dcos2BootstrapResources       = "dcos/bstrap/bootstrapresources.t"
	dcos2BootstrapCustomdata      = "dcos/bstrap/bootstrapcustomdata.yml"
	dcos2MasterVars               = "dcos/bstrap/dcosmastervars.t"
	dcos2MasterResources          = "dcos/bstrap/dcosmasterresources.t"
	iaasOutputs                   = "iaasoutputs.t"
	kubernetesParams              = "k8s/kubernetesparams.t"
	masterOutputs                 = "masteroutputs.t"
	masterParams                  = "masterparams.t"
	swarmBaseFile                 = "swarm/swarmbase.t"
	swarmParams                   = "swarm/swarmparams.t"
	swarmAgentResourcesVMAS       = "swarm/swarmagentresourcesvmas.t"
	swarmAgentResourcesVMSS       = "swarm/swarmagentresourcesvmss.t"
	swarmAgentVars                = "swarm/swarmagentvars.t"
	swarmMasterResources          = "swarm/swarmmasterresources.t"
	swarmMasterVars               = "swarm/swarmmastervars.t"
	swarmWinAgentResourcesVMAS    = "swarm/swarmwinagentresourcesvmas.t"
	swarmWinAgentResourcesVMSS    = "swarm/swarmwinagentresourcesvmss.t"
	windowsParams                 = "windowsparams.t"
)
