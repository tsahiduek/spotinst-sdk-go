package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// AwsGroupService is an interface for interfacing with the AwsGroup
// endpoints of the Spotinst API.
type AwsGroupService interface {
	List(*ListAwsGroupInput) (*ListAwsGroupOutput, error)
	Create(*CreateAwsGroupInput) (*CreateAwsGroupOutput, error)
	Read(*ReadAwsGroupInput) (*ReadAwsGroupOutput, error)
	Update(*UpdateAwsGroupInput) (*UpdateAwsGroupOutput, error)
	Delete(*DeleteAwsGroupInput) (*DeleteAwsGroupOutput, error)
}

// AwsGroupServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type AwsGroupServiceOp struct {
	client *Client
}

var _ AwsGroupService = &AwsGroupServiceOp{}

type AwsGroup struct {
	ID          *string              `json:"id,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Description *string              `json:"description,omitempty"`
	Capacity    *AwsGroupCapacity    `json:"capacity,omitempty"`
	Compute     *AwsGroupCompute     `json:"compute,omitempty"`
	Strategy    *AwsGroupStrategy    `json:"strategy,omitempty"`
	Scaling     *AwsGroupScaling     `json:"scaling"`
	Scheduling  *AwsGroupScheduling  `json:"scheduling"`
	Integration *AwsGroupIntegration `json:"thirdPartiesIntegration"`
	Multai      *AwsGroupMultai      `json:"multai"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string `json:"-"`

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string `json:"-"`
}

type AwsGroupIntegration struct {
	EC2ContainerService *AwsGroupEC2ContainerServiceIntegration `json:"ecs"`
	ElasticBeanstalk    *AwsGroupElasticBeanstalkIntegration    `json:"elasticBeanstalk"`
	Rancher             *AwsGroupRancherIntegration             `json:"rancher"`
	Kubernetes          *AwsGroupKubernetesIntegration          `json:"kubernetes"`
	Mesosphere          *AwsGroupMesosphereIntegration          `json:"mesosphere"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMultai struct {
	Token     *string                   `json:"token"`
	Balancers []*AwsGroupMultaiBalancer `json:"balancers"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMultaiBalancer struct {
	ProjectID   *string `json:"projectId"`
	BalancerID  *string `json:"balancerId"`
	TargetSetID *string `json:"targetSetId"`
	AzAwareness *bool   `json:"azAwareness"`
	AutoWeight  *bool   `json:"autoWeight"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupRancherIntegration struct {
	MasterHost *string `json:"masterHost,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty"`
	SecretKey  *string `json:"secretKey,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupElasticBeanstalkIntegration struct {
	EnvironmentID *string `json:"environmentId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupEC2ContainerServiceIntegration struct {
	ClusterName *string `json:"clusterName,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupKubernetesIntegration struct {
	Server *string `json:"apiServer,omitempty"`
	Token  *string `json:"token,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupMesosphereIntegration struct {
	Server *string `json:"apiServer,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScheduling struct {
	Tasks []*AwsGroupScheduledTask `json:"tasks"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScheduledTask struct {
	Frequency           *string `json:"frequency,omitempty"`
	CronExpression      *string `json:"cronExpression,omitempty"`
	TaskType            *string `json:"taskType,omitempty"`
	ScaleTargetCapacity *int    `json:"scaleTargetCapacity,omitempty"`
	ScaleMinCapacity    *int    `json:"scaleMinCapacity,omitempty"`
	ScaleMaxCapacity    *int    `json:"scaleMaxCapacity,omitempty"`
	BatchSizePercentage *int    `json:"batchSizePercentage,omitempty"`
	GracePeriod         *int    `json:"gracePeriod,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScaling struct {
	Up   []*AwsGroupScalingPolicy `json:"up"`
	Down []*AwsGroupScalingPolicy `json:"down"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScalingPolicy struct {
	PolicyName        *string                           `json:"policyName,omitempty"`
	MetricName        *string                           `json:"metricName,omitempty"`
	Statistic         *string                           `json:"statistic,omitempty"`
	Unit              *string                           `json:"unit,omitempty"`
	Threshold         *float64                          `json:"threshold,omitempty"`
	Adjustment        *int                              `json:"adjustment,omitempty"`
	MinTargetCapacity *int                              `json:"minTargetCapacity,omitempty"`
	MaxTargetCapacity *int                              `json:"maxTargetCapacity,omitempty"`
	Namespace         *string                           `json:"namespace,omitempty"`
	EvaluationPeriods *int                              `json:"evaluationPeriods,omitempty"`
	Period            *int                              `json:"period,omitempty"`
	Cooldown          *int                              `json:"cooldown,omitempty"`
	Operator          *string                           `json:"operator,omitempty"`
	Dimensions        []*AwsGroupScalingPolicyDimension `json:"dimensions,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupScalingPolicyDimension struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupStrategy struct {
	Risk                     *float64                  `json:"risk,omitempty"`
	OnDemandCount            *int                      `json:"onDemandCount,omitempty"`
	DrainingTimeout          *int                      `json:"drainingTimeout,omitempty"`
	AvailabilityVsCost       *string                   `json:"availabilityVsCost,omitempty"`
	UtilizeReservedInstances *bool                     `json:"utilizeReservedInstances,omitempty"`
	FallbackToOnDemand       *bool                     `json:"fallbackToOd,omitempty"`
	SpinUpTime               *int                      `json:"spinUpTime,omitempty"`
	Signals                  []*AwsGroupStrategySignal `json:"signals"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupStrategySignal struct {
	Name *string `json:"name"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupCapacity struct {
	Minimum *int    `json:"minimum,omitempty"`
	Maximum *int    `json:"maximum,omitempty"`
	Target  *int    `json:"target,omitempty"`
	Unit    *string `json:"unit,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupCompute struct {
	Product             *string                             `json:"product,omitempty"`
	InstanceTypes       *AwsGroupComputeInstanceType        `json:"instanceTypes,omitempty"`
	LaunchSpecification *AwsGroupComputeLaunchSpecification `json:"launchSpecification,omitempty"`
	AvailabilityZones   []*AwsGroupComputeAvailabilityZone  `json:"availabilityZones,omitempty"`
	ElasticIPs          []string                            `json:"elasticIps,omitempty"`
	EBSVolumePool       []*AwsGroupComputeEBSVolume         `json:"ebsVolumePool"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeEBSVolume struct {
	DeviceName *string  `json:"deviceName,omitempty"`
	VolumeIDs  []string `json:"volumeIds,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeInstanceType struct {
	OnDemand *string                              `json:"ondemand,omitempty"`
	Spot     []string                             `json:"spot,omitempty"`
	Weights  []*AwsGroupComputeInstanceTypeWeight `json:"weights,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeInstanceTypeWeight struct {
	InstanceType *string `json:"instanceType,omitempty"`
	Weight       *int    `json:"weightedCapacity,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeAvailabilityZone struct {
	Name     *string `json:"name,omitempty"`
	SubnetID *string `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLaunchSpecification struct {
	LoadBalancerNames      []string                            `json:"loadBalancerNames"`
	LoadBalancersConfig    *AwsGroupComputeLoadBalancersConfig `json:"loadBalancersConfig,omitempty"`
	SecurityGroupIDs       []string                            `json:"securityGroupIds,omitempty"`
	HealthCheckType        *string                             `json:"healthCheckType,omitempty"`
	HealthCheckGracePeriod *int                                `json:"healthCheckGracePeriod,omitempty"`
	ImageID                *string                             `json:"imageId,omitempty"`
	KeyPair                *string                             `json:"keyPair,omitempty"`
	UserData               *string                             `json:"userData,omitempty"`
	Tenancy                *string                             `json:"tenancy,omitempty"`
	Monitoring             *bool                               `json:"monitoring,omitempty"`
	EBSOptimized           *bool                               `json:"ebsOptimized,omitempty"`
	IamInstanceProfile     *AwsGroupComputeIamInstanceProfile  `json:"iamRole,omitempty"`
	BlockDevices           []*AwsGroupComputeBlockDevice       `json:"blockDeviceMappings"`
	NetworkInterfaces      []*AwsGroupComputeNetworkInterface  `json:"networkInterfaces"`
	Tags                   []*AwsGroupComputeTag               `json:"tags"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLoadBalancersConfig struct {
	LoadBalancers []*AwsGroupComputeLoadBalancer `json:"loadBalancers"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeLoadBalancer struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`
	Type *string `json:"type,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeNetworkInterface struct {
	ID                             *string  `json:"networkInterfaceId,omitempty"`
	Description                    *string  `json:"description,omitempty"`
	DeviceIndex                    *int     `json:"deviceIndex,omitempty"`
	SecondaryPrivateIPAddressCount *int     `json:"secondaryPrivateIpAddressCount,omitempty"`
	AssociatePublicIPAddress       *bool    `json:"associatePublicIpAddress,omitempty"`
	DeleteOnTermination            *bool    `json:"deleteOnTermination,omitempty"`
	SecurityGroupsIDs              []string `json:"groups,omitempty"`
	PrivateIPAddress               *string  `json:"privateIpAddress,omitempty"`
	SubnetID                       *string  `json:"subnetId,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeBlockDevice struct {
	DeviceName  *string             `json:"deviceName,omitempty"`
	VirtualName *string             `json:"virtualName,omitempty"`
	EBS         *AwsGroupComputeEBS `json:"ebs,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeEBS struct {
	DeleteOnTermination *bool   `json:"deleteOnTermination,omitempty"`
	Encrypted           *bool   `json:"encrypted,omitempty"`
	SnapshotID          *string `json:"snapshotId,omitempty"`
	VolumeType          *string `json:"volumeType,omitempty"`
	VolumeSize          *int    `json:"volumeSize,omitempty"`
	IOPS                *int    `json:"iops,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeIamInstanceProfile struct {
	Name *string `json:"name,omitempty"`
	Arn  *string `json:"arn,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type AwsGroupComputeTag struct {
	Key   *string `json:"tagKey,omitempty"`
	Value *string `json:"tagValue,omitempty"`

	forceSendFields []string `json:"-"`
	nullFields      []string `json:"-"`
}

type ListAwsGroupInput struct{}

type ListAwsGroupOutput struct {
	Groups []*AwsGroup `json:"groups,omitempty"`
}

type CreateAwsGroupInput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type CreateAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type ReadAwsGroupInput struct {
	ID *string `json:"groupId,omitempty"`
}

type ReadAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type UpdateAwsGroupInput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type UpdateAwsGroupOutput struct {
	Group *AwsGroup `json:"group,omitempty"`
}

type DeleteAwsGroupInput struct {
	ID *string `json:"groupId,omitempty"`
}

type DeleteAwsGroupOutput struct{}

func awsGroupFromJSON(in []byte) (*AwsGroup, error) {
	b := new(AwsGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func awsGroupsFromJSON(in []byte) ([]*AwsGroup, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AwsGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsGroupsFromHttpResponse(resp *http.Response) ([]*AwsGroup, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsGroupsFromJSON(body)
}

func (s *AwsGroupServiceOp) List(input *ListAwsGroupInput) (*ListAwsGroupOutput, error) {
	r := s.client.newRequest("GET", "/aws/ec2/group")

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListAwsGroupOutput{Groups: gs}, nil
}

func (s *AwsGroupServiceOp) Create(input *CreateAwsGroupInput) (*CreateAwsGroupOutput, error) {
	r := s.client.newRequest("POST", "/aws/ec2/group")
	r.obj = input
	r.forceSendFields = input.Group.forceSendFields
	r.nullFields = input.Group.nullFields

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Read(input *ReadAwsGroupInput) (*ReadAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest("GET", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Update(input *UpdateAwsGroupInput) (*UpdateAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.Group.ID),
	})
	if err != nil {
		return nil, err
	}

	// We do not need the ID anymore so let's drop it.
	input.Group.ID = nil

	r := s.client.newRequest("PUT", path)
	r.obj = input
	r.forceSendFields = input.Group.forceSendFields
	r.nullFields = input.Group.nullFields

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateAwsGroupOutput)
	if len(gs) > 0 {
		output.Group = gs[0]
	}

	return output, nil
}

func (s *AwsGroupServiceOp) Delete(input *DeleteAwsGroupInput) (*DeleteAwsGroupOutput, error) {
	path, err := uritemplates.Expand("/aws/ec2/group/{groupId}", map[string]string{
		"groupId": StringValue(input.ID),
	})
	if err != nil {
		return nil, err
	}

	r := s.client.newRequest("DELETE", path)
	r.obj = input

	_, resp, err := requireOK(s.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteAwsGroupOutput{}, nil
}
