package client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	types2 "github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/knockbox/matchbox/internal/platform"
	"github.com/knockbox/matchbox/pkg/accessors"
	"github.com/knockbox/matchbox/pkg/enums/ecs_cluster"
	"github.com/knockbox/matchbox/pkg/enums/vpc_instance"
	"github.com/knockbox/matchbox/pkg/models"
)

type Amazon struct {
	ec2Client *ec2.Client
	ecsClient *ecs.Client
	efsClient *efs.Client

	vpci    accessors.VPCInstanceAccessor
	efsi    accessors.EFSInstanceAccessor
	cluster accessors.ECSClusterAccessor

	l hclog.Logger
}

func NewAmazon(db *sqlx.DB, l hclog.Logger) *Amazon {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	return &Amazon{
		ec2Client: ec2.NewFromConfig(cfg),
		ecsClient: ecs.NewFromConfig(cfg),
		efsClient: efs.NewFromConfig(cfg),
		vpci: platform.VPCInstanceSQLImpl{
			DB: db,
		},
		efsi: platform.EFSInstanceSQLImpl{
			DB: db,
		},
		cluster: platform.ECSClusterSQLImpl{
			DB: db,
		},
		l: l,
	}
}

func (a *Amazon) InitForDeployment(id int) error {
	a.l.Info("Init Deployment", "id", id)

	// Create the VPC
	a.l.Info("Create VPC", "deployment_id", id)
	vpc, err := a.CreateVPC(id)
	if err != nil {
		return err
	}

	_, err = a.vpci.Create(*vpc)
	if err != nil {
		a.l.Error("failed to insert vpc", "err", err, "vpc", vpc)
		return err
	}

	// Create EFS
	a.l.Info("Create EFS", "deployment_id", id)
	efsi, err := a.CreateEFS(id)
	if err != nil {
		return err
	}

	_, err = a.efsi.Create(*efsi)
	if err != nil {
		a.l.Error("failed to insert efs", "err", err, "vpc", vpc)
		return err
	}

	// Create the ECS Cluster
	a.l.Info("Create ECS Cluster", "deployment_id", id)
	cluster, err := a.CreateECSCluster(id)
	if err != nil {
		return err
	}

	_, err = a.cluster.Create(*cluster)
	if err != nil {
		a.l.Error("failed to insert cluster", "err", err, "cluster", cluster)
	}

	return nil
}

func (a *Amazon) CreateVPC(id int) (*models.VPCInstance, error) {
	vpc := models.NewVPCInstance(id)
	ctx := context.Background()

	// Create the VPC
	vpcOutput, err := a.ec2Client.CreateVpc(ctx, &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
	})
	if err != nil {
		a.l.Error("CreateVPC failed", "err", err, "deployment_id", id)
		return nil, err
	}
	a.l.Info("CreateVPC success", "vpc_id", *vpcOutput.Vpc.VpcId)

	// Create the InternetGateway
	igwOutput, err := a.ec2Client.CreateInternetGateway(ctx, &ec2.CreateInternetGatewayInput{})
	if err != nil {
		a.l.Error("CreateInternetGateway failed", "err", err, "deployment_id", id)
		return nil, err
	}
	a.l.Info("CreateVPC success", "igw_id", *igwOutput.InternetGateway.InternetGatewayId)

	// Attach the InternetGateway to VPC
	_, err = a.ec2Client.AttachInternetGateway(ctx, &ec2.AttachInternetGatewayInput{
		InternetGatewayId: igwOutput.InternetGateway.InternetGatewayId,
		VpcId:             vpcOutput.Vpc.VpcId,
	})
	if err != nil {
		a.l.Error("AttachInternetGateway failed", "err", err, "deployment_id", id, "vpc_id", *vpcOutput.Vpc.VpcId, "igw_id", *igwOutput.InternetGateway.InternetGatewayId)
		return nil, err
	}
	a.l.Info("AttachInternetGateway success", "deployment_id", id)

	// Create the subnet
	subnetOutput, err := a.ec2Client.CreateSubnet(ctx, &ec2.CreateSubnetInput{
		VpcId:     vpcOutput.Vpc.VpcId,
		CidrBlock: vpcOutput.Vpc.CidrBlock,
	})
	if err != nil {
		a.l.Error("CreateSubnet failed", "err", err, "deployment_id", id)
		return nil, err
	}
	a.l.Info("CreateSubnet success", "subnet_id", *subnetOutput.Subnet.SubnetId)

	// Get VPC Security Details
	sgOutput, err := a.ec2Client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{*vpcOutput.Vpc.VpcId},
			},
		},
	})
	if err != nil {
		a.l.Error("DescribeSecurityGroups failed", "err", err, "deployment_id", id, "vpc_id", *vpcOutput.Vpc.VpcId)
		return nil, err
	}
	a.l.Info("DescribeSecurityGroups success", "deployment_id", id, "vpc_id", *vpcOutput.Vpc.VpcId)

	// Get Route Tables
	rtOutput, err := a.ec2Client.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{*vpcOutput.Vpc.VpcId},
			},
		},
	})
	if err != nil {
		a.l.Error("DescribeRouteTables failed", "err", err, "deployment_id", id, "vpc_id", *vpcOutput.Vpc.VpcId)
		return nil, err
	}
	a.l.Info("DescribeRouteTables success", "deployment_id", id, "vpc_id", *vpcOutput.Vpc.VpcId)

	// Allow I/O from the Internet
	_, err = a.ec2Client.CreateRoute(ctx, &ec2.CreateRouteInput{
		RouteTableId:         rtOutput.RouteTables[0].RouteTableId,
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            igwOutput.InternetGateway.InternetGatewayId,
	})
	if err != nil {
		a.l.Error("CreateRoute failed", "err", err, "deployment_id", id, "routetable_id", *rtOutput.RouteTables[0].RouteTableId)
		return nil, err
	}
	a.l.Info("CreateRoute success")

	vpc.AwsResourceId = *vpcOutput.Vpc.VpcId
	vpc.SubnetID = *subnetOutput.Subnet.SubnetId
	vpc.SecurityGroupID = *sgOutput.SecurityGroups[0].GroupId
	vpc.InternetGatewayID = *igwOutput.InternetGateway.InternetGatewayId
	vpc.State = vpc_instance.Available

	return vpc, nil
}

func (a *Amazon) CreateEFS(id int) (*models.EFSInstance, error) {
	ctx := context.Background()
	efsi := models.NewEFSInstance(id)

	fsOutput, err := a.efsClient.CreateFileSystem(ctx, &efs.CreateFileSystemInput{
		CreationToken:   aws.String(uuid.NewString()),
		Backup:          aws.Bool(false),
		Encrypted:       aws.Bool(false),
		PerformanceMode: types2.PerformanceModeGeneralPurpose,
		ThroughputMode:  types2.ThroughputModeElastic,
	})
	if err != nil {
		a.l.Error("CreateFileSystem failed", "err", err, "deployment_id", id)
		return nil, err
	}

	switch fsOutput.LifeCycleState {
	case types2.LifeCycleStateError:
		fallthrough
	case types2.LifeCycleStateDeleting:
		fallthrough
	case types2.LifeCycleStateDeleted:
		a.l.Error("FileSystem is failing", "state", fsOutput.LifeCycleState)
	}

	efsi.AwsResourceId = *fsOutput.FileSystemArn
	efsi.AWSFileSystemId = *fsOutput.FileSystemId
	efsi.State = fsOutput.LifeCycleState

	a.l.Info("FileSystem created", "fs_id", efsi.AWSFileSystemId, "state", efsi.State)

	return efsi, nil
}

func (a *Amazon) CreateECSCluster(id int) (*models.ECSCluster, error) {
	ctx := context.Background()
	cluster := models.NewECSCluster(id)

	output, err := a.ecsClient.CreateCluster(ctx, &ecs.CreateClusterInput{
		ClusterName: aws.String(uuid.NewString()),
	})
	if err != nil {
		a.l.Error("CreateCluster failed", "err", err, "deployment_id", id)
		return nil, err
	}

	cluster.ClusterName = *output.Cluster.ClusterName
	cluster.AwsArn = *output.Cluster.ClusterArn
	cluster.Status = ecs_cluster.Status(*output.Cluster.Status)

	a.l.Info("Cluster created", "name", cluster.ClusterName, "status", cluster.Status)

	return cluster, nil
}

/*
func (a *Amazon) CreateInfrastructure(payload *payloads.TaskDefinitionCreatePayload) {
	taskdef := models.NewECSTaskDefinition()

	// Construct input from payload.
	// Collect container definitions.
	var containerDefs []types.ContainerDefinition
	for _, container := range payload.Containers {
		def := types.ContainerDefinition{
			Image:     aws.String(container.Image),
			Name:      aws.String(uuid.NewString()),
			Essential: container.Essential,

			// Logging - CloudWatch displays as: <taskdef.FamilyID>:<container_image>
			LogConfiguration: &types.LogConfiguration{
				LogDriver: types.LogDriverAwslogs,
				Options: map[string]string{
					"awslogs-create-group":  "true",
					"awslogs-group":         taskdef.FamilyId.String(),
					"awslogs-region":        "us-east-1",
					"awslogs-stream-prefix": container.Image,
				},
			},
		}

		// Populate environment variables
		for _, envvar := range container.EnvironmentVars {
			def.Environment = append(def.Environment, types.KeyValuePair{
				Name:  aws.String(envvar.Key),
				Value: aws.String(envvar.Value),
			})
		}

		// Populate ports
		for _, port := range container.Ports {
			protocol := types.TransportProtocolTcp
			if port.Protocol != nil && *port.Protocol == "udp" {
				protocol = types.TransportProtocolUdp
			}

			def.PortMappings = append(def.PortMappings, types.PortMapping{
				AppProtocol:   types.ApplicationProtocolHttp,
				ContainerPort: aws.Int32(port.ContainerPort),
				HostPort:      port.HostPort,
				Name:          aws.String(port.Name),
				Protocol:      protocol,
			})
		}

		// Populate mount points
		for _, volume := range container.Volumes {
			def.MountPoints = append(def.MountPoints, types.MountPoint{
				ContainerPath: aws.String(volume.Path),
				ReadOnly:      volume.ReadOnly,
				SourceVolume:  aws.String(volume.Source),
			})
		}

		containerDefs = append(containerDefs, def)
	}

	// Collect EFS volumes.
	var volumes []types.Volume
	for _, volume := range payload.Volumes {
		volumes = append(volumes, types.Volume{
			EfsVolumeConfiguration: &types.EFSVolumeConfiguration{
				FileSystemId: aws.String(volume.ID),
			},
			Name: aws.String(volume.Name),
		})
	}

	// Prepare aws input.
	taskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    containerDefs,
		Family:                  aws.String(taskdef.FamilyId.String()),
		Cpu:                     aws.String(payload.CPU),
		Memory:                  aws.String(payload.Memory),
		NetworkMode:             types.NetworkModeAwsvpc,
		RequiresCompatibilities: []types.Compatibility{types.CompatibilityFargate},
		ExecutionRoleArn:        aws.String("arn:aws:iam::588285845198:role/ecsTaskExecutionRole"),
		TaskRoleArn:             aws.String("arn:aws:iam::588285845198:role/ecsTaskExecutionRole"),
		Volumes:                 volumes,
	}
	taskDefOutput, err := a.ecsClient.RegisterTaskDefinition(context.Background(), taskDefInput)

	// Create the Instance
}
*/
