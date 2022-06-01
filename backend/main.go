package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Body struct {
	Consul         Consul         `json:"Consul"`
	Deltas         []Deltas       `json:"Deltas"`
	Invalid        interface{}    `json:"Invalid"`
	Kubernetes     Kubernetes     `json:"Kubernetes"`
	AmbassadorMeta AmbassadorMeta `json:"AmbassadorMeta"`
}

type Deltas struct {
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	DeltaType  string   `json:"deltaType"`
	APIVersion string   `json:"apiVersion"`
}

type Consul struct {
}

type MatchLabels struct {
	Hostname string `json:"hostname,omitempty"`
}

type TLSSecret struct {
	Name string `json:"name,omitempty"`
}

type PreviewURL struct {
	Enabled bool `json:"enabled"`
}

type PrivateKeySecret struct {
	Name string `json:"name,omitempty"`
}

type AcmeProvider struct {
	Email            string           `json:"email,omitempty"`
	Authority        string           `json:"authority,omitempty"`
	Registration     string           `json:"registration,omitempty"`
	PrivateKeySecret PrivateKeySecret `json:"privateKeySecret,omitempty"`
}

type Spec struct {
	Hostname     									string       				`json:"hostname,omitempty"`
	Selector     									Selector     				`json:"selector,omitempty,omitempty"`
	TLSSecret    									TLSSecret    				`json:"tlsSecret,omitempty"`
	PreviewURL   									PreviewURL   				`json:"previewUrl,omitempty"`
	AcmeProvider 									AcmeProvider 				`json:"acmeProvider,omitempty"`
	AmbassadorID 									[]string     				`json:"ambassador_id,omitempty"`
	Type            							string   						`json:"type,omitempty"`
	Ports           							[]Ports  						`json:"ports,omitempty"`
	ClusterIP       							string   						`json:"clusterIP,omitempty"`
	SessionAffinity 							string   						`json:"sessionAffinity"`	
	ExternalTrafficPolicy 				string   						`json:"externalTrafficPolicy,omitempty"`
	Proto       									string 							`json:"proto,omitempty"`
	AuthService 									string 							`json:"auth_service,omitempty"`
	Service 											string 							`json:"service,omitempty"`
	Volumes                       []Volumes       		`json:"volumes,omitempty"`
	NodeName                      string          		`json:"nodeName,omitempty"`
	Priority                      int             		`json:"priority,omitempty"`
	DNSPolicy                     string          		`json:"dnsPolicy,omitempty"`
	Containers                    []Containers    		`json:"containers,omitempty"`
	Tolerations                   []Tolerations   		`json:"tolerations,omitempty"`
	NodeSelector                  NodeSelector    		`json:"nodeSelector,omitempty"`
	RestartPolicy                 string          		`json:"restartPolicy,omitempty"`
	SchedulerName                 string          		`json:"schedulerName,omitempty"`
	ServiceAccount                string          		`json:"serviceAccount,omitempty"`
	SecurityContext               SecurityContext 		`json:"securityContext,omitempty"`
	EnableServiceLinks            bool            		`json:"enableServiceLinks,omitempty"`
	ServiceAccountName            string          		`json:"serviceAccountName,omitempty"`
	TerminationGracePeriodSeconds int             		`json:"terminationGracePeriodSeconds,omitempty"`
	Affinity                      Affinity        		`json:"affinity, omitempty"`
	Config 												Config 							`json:"config,omitempty"`
	Rewrite 											string 							`json:"rewrite,omitempty"`
	Prefix             						string             	`json:"prefix,omitempty"`
	PrefixRegex        						bool               	`json:"prefix_regex,omitempty"`
	RegexRewrite       						RegexRewrite       	`json:"regex_rewrite,omitempty"`
	AddResponseHeaders 						AddResponseHeaders 	`json:"add_response_headers,omitempty"`
}

type Status struct {
	State                	string 								`json:"state,omitempty"`
	TLSCertificateSource 	string 								`json:"tlsCertificateSource,omitempty"`
	Phase             		string              	`json:"phase,omitempty"`
	PodIP             		string              	`json:"podIP,omitempty"`
	HostIP            		string              	`json:"hostIP,omitempty"`
	PodIPs            		[]PodIPs            	`json:"podIPs,omitempty"`
	QosClass          		string              	`json:"qosClass,omitempty"`
	StartTime         		time.Time           	`json:"startTime,omitempty"`
	Conditions        		[]Conditions        	`json:"conditions,omitempty"`
	ContainerStatuses 		[]ContainerStatuses 	`json:"containerStatuses,omitempty"`
	LoadBalancer 					LoadBalancer 					`json:"loadBalancer,omitempty"`
}

type Annotations struct {
	KubectlKubernetesIoLastAppliedConfiguration 	string 			`json:"kubectl.kubernetes.io/last-applied-configuration,omitempty"`
	MetaHelmShReleaseName      										string 			`json:"meta.helm.sh/release-name,omitempty"`
	MetaHelmShReleaseNamespace 										string 			`json:"meta.helm.sh/release-namespace,omitempty"`
	A8RIoIgnore                										string 			`json:"a8r.io/ignore,omitempty"`
	AppGetambassadorIoDescription               	string 			`json:"app.getambassador.io/description,omitempty"`
	AppGetambassadorIoOwner                     	string 			`json:"app.getambassador.io/owner,omitempty"`
	EndpointsKubernetesIoLastChangeTriggerTime 		time.Time 	`json:"endpoints.kubernetes.io/last-change-trigger-time,omitempty"`
	SidecarIstioIoInject            							string 			`json:"sidecar.istio.io/inject,omitempty"`
	ConsulHashicorpComConnectInject 							string 			`json:"consul.hashicorp.com/connect-inject,omitempty"`
}

type Host struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
	Metadata   Metadata `json:"metadata"`
	APIVersion string   `json:"apiVersion"`
}

type Items struct {
	Key  			string 		`json:"key"`
	Path 			string 		`json:"path"`
	FieldRef 	FieldRef 	`json:"fieldRef"`
}

type ConfigMap struct {
	Name        string  `json:"name"`
	Items       []Items `json:"items"`
	DefaultMode int     `json:"defaultMode"`
}

type Secret struct {
	SecretName  	string 		`json:"secretName"`
	DefaultMode 	int    		`json:"defaultMode"`
	Data       		Data     	`json:"data,omitempty"`
	Kind       		string   	`json:"kind"`
	Type       		string   	`json:"type"`
	Metadata   		Metadata 	`json:"metadata"`
	APIVersion 		string   	`json:"apiVersion"`
}

type Volumes struct {
	Name      		string    		`json:"name"`
	ConfigMap 		ConfigMap 		`json:"configMap,omitempty"`
	Secret    		Secret    		`json:"secret,omitempty"`
	EmptyDir 			EmptyDir 			`json:"emptyDir,omitempty"`
	DownwardAPI 	DownwardAPI 	`json:"downwardAPI,omitempty"`
}

type Ports struct {
	Name          	string 		`json:"name,omitempty"`
	Port       			int    		`json:"port,omitempty"`
	Protocol      	string 		`json:"protocol,omitempty"`
	ContainerPort 	int    		`json:"containerPort,omitempty"`
	TargetPort 			string 		`json:"targetPort,omitempty"`
	NodePort   			int    		`json:"nodePort,omitempty"`
}

type Limits struct {
	Memory 	string 	`json:"memory"`
	CPU    	string 	`json:"cpu"`
}

type Requests struct {
	CPU    	string 	`json:"cpu"`
	Memory 	string 	`json:"memory"`
}

type Resources struct {
	Limits   Limits   `json:"limits"`
	Requests Requests `json:"requests"`
}

type VolumeMounts struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readOnly"`
	MountPath string `json:"mountPath"`
}

type HTTPGet struct {
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type LivenessProbe struct {
	HTTPGet             HTTPGet `json:"httpGet"`
	PeriodSeconds       int     `json:"periodSeconds"`
	TimeoutSeconds      int     `json:"timeoutSeconds"`
	FailureThreshold    int     `json:"failureThreshold"`
	SuccessThreshold    int     `json:"successThreshold"`
	InitialDelaySeconds int     `json:"initialDelaySeconds"`
}

type ReadinessProbe struct {
	HTTPGet             HTTPGet `json:"httpGet"`
	PeriodSeconds       int     `json:"periodSeconds"`
	TimeoutSeconds      int     `json:"timeoutSeconds"`
	FailureThreshold    int     `json:"failureThreshold"`
	SuccessThreshold    int     `json:"successThreshold"`
	InitialDelaySeconds int     `json:"initialDelaySeconds"`
}

type Capabilities struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}

type SecurityContext struct {
	Capabilities             	Capabilities 	`json:"capabilities"`
	ReadOnlyRootFilesystem   	bool         	`json:"readOnlyRootFilesystem"`
	AllowPrivilegeEscalation 	bool         	`json:"allowPrivilegeEscalation"`
	RunAsUser 								int 					`json:"runAsUser"`
}

type Containers struct {
	Args                     []string        `json:"args"`
	Name                     string          `json:"name"`
	Image                    string          `json:"image"`
	Ports                    []Ports         `json:"ports"`
	Resources                Resources       `json:"resources"`
	VolumeMounts             []VolumeMounts  `json:"volumeMounts"`
	LivenessProbe            LivenessProbe   `json:"livenessProbe"`
	ReadinessProbe           ReadinessProbe  `json:"readinessProbe"`
	ImagePullPolicy          string          `json:"imagePullPolicy"`
	SecurityContext          SecurityContext `json:"securityContext"`
	TerminationMessagePath   string          `json:"terminationMessagePath"`
	TerminationMessagePolicy string          `json:"terminationMessagePolicy"`
	Env                      []Env           `json:"env"`
}

type Tolerations struct {
	Key               string `json:"key"`
	Operator          string `json:"operator"`
	Effect            string `json:"effect,omitempty"`
	TolerationSeconds int    `json:"tolerationSeconds,omitempty"`
}

type NodeSelector struct {
	BetaKubernetesIoOs string `json:"beta.kubernetes.io/os,omitempty"`
}

type PodIPs struct {
	IP string `json:"ip"`
}

type Conditions struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastProbeTime      interface{} `json:"lastProbeTime"`
	LastTransitionTime time.Time   `json:"lastTransitionTime"`
}

type Running struct {
	StartedAt time.Time `json:"startedAt"`
}

type State struct {
	Running Running `json:"running"`
}

type LastState struct {
}

type ContainerStatuses struct {
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Ready        bool      `json:"ready"`
	State        State     `json:"state"`
	ImageID      string    `json:"imageID"`
	Started      bool      `json:"started"`
	LastState    LastState `json:"lastState"`
	ContainerID  string    `json:"containerID"`
	RestartCount int       `json:"restartCount"`
}

type Labels struct {
	K8SApp          							string 		`json:"k8s-app,omitempty"`
	PodTemplateHash 							string 		`json:"pod-template-hash,omitempty"`
	Product                  			string 		`json:"product,omitempty"`
	Service                  			string 		`json:"service,omitempty"`
	HelmShChart              			string 		`json:"helm.sh/chart,omitempty"`
	AppKubernetesIoName      			string 		`json:"app.kubernetes.io/name,omitempty"`
	AppKubernetesIoPartOf    			string 		`json:"app.kubernetes.io/part-of,omitempty"`
	AppKubernetesIoInstance  			string 		`json:"app.kubernetes.io/instance,omitempty"`
	AppKubernetesIoManagedBy 			string 		`json:"app.kubernetes.io/managed-by,omitempty"`
	ServiceKubernetesIoHeadless 	string 		`json:"service.kubernetes.io/headless,omitempty"`
	AppKubernetesIoComponent 			string 		`json:"app.kubernetes.io/component,omitempty"`
	Provider  										string 		`json:"provider,omitempty"`
	Component 										string 		`json:"component,omitempty"`
}

type OwnerReferences struct {
	UID                	string 	`json:"uid"`
	Kind               	string 	`json:"kind"`
	Name               	string 	`json:"name"`
	APIVersion         	string 	`json:"apiVersion"`
	Controller         	bool   	`json:"controller"`
	BlockOwnerDeletion 	bool   	`json:"blockOwnerDeletion"`
}

type EmptyDir struct {
}

type FieldRef struct {
	FieldPath  string `json:"fieldPath"`
	APIVersion string `json:"apiVersion"`
}

type DownwardAPI struct {
	Items       []Items `json:"items"`
	DefaultMode int     `json:"defaultMode"`
}

type LabelSelector struct {
	MatchLabels MatchLabels `json:"matchLabels,omitempty"`
}

type PodAffinityTerm struct {
	TopologyKey   string        `json:"topologyKey"`
	LabelSelector LabelSelector `json:"labelSelector,omitempty"`
}

type PreferredDuringSchedulingIgnoredDuringExecution struct {
	Weight          int             `json:"weight"`
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm,omitempty"`
}

type PodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredDuringSchedulingIgnoredDuringExecution `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

type Affinity struct {
	PodAntiAffinity PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

type ValueFrom struct {
	FieldRef FieldRef `json:"fieldRef"`
}

type Env struct {
	Name      string    `json:"name"`
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
	Value     string    `json:"value,omitempty"`
}

type Metadata struct {
	UID               string            `json:"uid,omitempty"`
	Name              string            `json:"name,omitempty"`
	Labels            Labels            `json:"labels,omitempty"`
	SelfLink          string            `json:"selfLink,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Annotations       Annotations       `json:"annotations,omitempty"`
	GenerateName      string            `json:"generateName,omitempty"`
	OwnerReferences   []OwnerReferences `json:"ownerReferences,omitempty"`
	ResourceVersion   string            `json:"resourceVersion,omitempty"`
	CreationTimestamp time.Time         `json:"creationTimestamp,omitempty"`
}

type Pods struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status"`
	Metadata   Metadata `json:"metadata,omitempty"`
	APIVersion string   `json:"apiVersion"`
}

type Cepc struct {
	Label     string `json:"label,omitempty"`
	Address   string `json:"address,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	AccountID string `json:"account_id,omitempty"`
}

type Config struct {
	Cepc Cepc `json:"cepc,omitempty"`
}

type Module struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec"`
	Metadata   Metadata `json:"metadata"`
	APIVersion string   `json:"apiVersion"`
}

type Data struct {
	TLSCrt 												string 	`json:"tls.crt"`
	TLSKey 												string 	`json:"tls.key"`
	LicenseKey 										string 	`json:"license-key"`
	UserKey 											string 	`json:"user.key"`
	ApplicationInstanceLabelKey 	string 	`json:"application.instanceLabelKey"`
	URL                         	string 	`json:"url"`
}

type RegexRewrite struct {
	Pattern      string `json:"pattern,omitempty"`
	Substitution string `json:"substitution,omitempty"`
}

type CacheControl struct {
	Value  string  `json:"value,omitempty"`
	Append bool    `json:"append,omitempty"`
}

type AddResponseHeaders struct {
	CacheControl CacheControl `json:"cache-control,omitempty"`
}

type Mapping struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec,omitempty"`
	Metadata   Metadata `json:"metadata"`
	APIVersion string   `json:"apiVersion"`
}

type Selector struct {
	AppKubernetesIoName     	string			 `json:"app.kubernetes.io/name,omitempty"`
	AppKubernetesIoInstance 	string			 `json:"app.kubernetes.io/instance,omitempty"`
	MatchLabels 							MatchLabels	 `json:"matchLabels,omitempty"`
}

type LoadBalancer struct {
}

type Service struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status"`
	Metadata   Metadata `json:"metadata,omitempty"`
	APIVersion string   `json:"apiVersion"`
}

type ConfigMaps struct {
	APIVersion string   `json:"apiVersion"`
	Data       Data     `json:"data"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
}

type TargetRef struct {
	UID             string `json:"uid"`
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	ResourceVersion string `json:"resourceVersion"`
}
type Addresses struct {
	IP        string    `json:"ip"`
	NodeName  string    `json:"nodeName"`
	TargetRef TargetRef `json:"targetRef"`
}
type Subsets struct {
	Ports     []Ports     `json:"ports"`
	Addresses []Addresses `json:"addresses"`
}

type Endpoints struct {
	Kind       string    	`json:"kind"`
	Subsets    []Subsets 	`json:"subsets"`
	APIVersion string    	`json:"apiVersion"`
	Metadata   Metadata  	`json:"metadata,omitempty"`
}

type AuthService struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec"`
	Metadata   Metadata `json:"metadata"`
	APIVersion string   `json:"apiVersion"`
}

type RateLimitService struct {
	Kind       string   `json:"kind"`
	Spec       Spec     `json:"spec"`
	Metadata   Metadata `json:"metadata"`
	APIVersion string   `json:"apiVersion"`
}
type Kubernetes struct {
	Host                       []Host             `json:"Host"`
	Pods                       []Pods             `json:"Pods"`
	Module                     []Module           `json:"Module"`
	Secret                     []Secret           `json:"secret"`
	Mapping                    []Mapping          `json:"Mapping"`
	Service                    []Service          `json:"service"`
	ConfigMaps                 []ConfigMaps       `json:"ConfigMaps"`
	DevPortal                  interface{}        `json:"DevPortal"`
	Endpoints                  []Endpoints        `json:"Endpoints"`
	Ingresses                  interface{}        `json:"ingresses"`
	LogService                 interface{}        `json:"LogService"`
	TCPMapping                 interface{}        `json:"TCPMapping"`
	TLSContext                 interface{}        `json:"TLSContext"`
	AuthService                []AuthService      `json:"AuthService"`
	ConsulResolver             interface{}        `json:"ConsulResolver"`
	TracingService             interface{}        `json:"TracingService"`
	Ingressclasses             interface{}        `json:"ingressclasses"`
	RateLimitService           []RateLimitService `json:"RateLimitService"`
	KubernetesServiceResolver  interface{}        `json:"KubernetesServiceResolver"`
	KubernetesEndpointResolver interface{}        `json:"KubernetesEndpointResolver"`
}
type AmbassadorMeta struct {
	Sidecar           interface{} `json:"sidecar"`
	ClusterID         string      `json:"cluster_id"`
	KubeVersion       string      `json:"kube_version"`
	AmbassadorID      string      `json:"ambassador_id"`
	AmbassadorVersion string      `json:"ambassador_version"`
}

func enableCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	// OPEN FILE
	jsonFile, err := os.Open("raw_snapshot.json")
	if err != nil {
			fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// DECODE JSON
	var body Body
	json.Unmarshal(byteValue, &body)
	var newService []Service
	var newJson []byte
	var jsonLen	int = 0;
	// var parts []string

	// CREATE SERVICE
	var simplifiedService Service
	for _, value := range body.Kubernetes.Service {
		jsonLen++
		simplifiedService.Kind = value.Kind
		simplifiedService.Spec = value.Spec
		simplifiedService.Status = value.Status
		simplifiedService.Metadata = value.Metadata
		simplifiedService.APIVersion = value.APIVersion
		
		// WRITE TO STRUCT
		newService = append(newService, simplifiedService)
		newJson, _ = json.MarshalIndent(newService, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
	}


	http.HandleFunc("/len", func(w http.ResponseWriter, _ *http.Request) {
		enableCors(w)
		fmt.Fprintf(w, "%d", jsonLen)
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		w.Write(newJson)
	})

	http.HandleFunc("/part", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			w.Write([]byte("nothing to see here\n"))
			return
		}
		part := keys[0]
		intPart, _ := strconv.Atoi(part)
		var index int
		for _, value := range body.Kubernetes.Service {
			if (value.Kind == "Service" && index == intPart) {
				returnedJson, err := json.MarshalIndent(value, "", "    ")
				if err != nil {
					fmt.Println(err)
				}		
				w.Write(returnedJson)}
			index++
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

