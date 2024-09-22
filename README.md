# AWS EKS Cluster Pulumi Module

## Introduction

The AWS EKS Cluster Pulumi Module provides a standardized and efficient way to define and deploy Amazon Elastic Kubernetes Service (EKS) clusters on AWS using a Kubernetes-like API resource model. By leveraging our unified APIs, developers can specify their EKS cluster configurations in simple YAML files, which the module then uses to create and manage AWS EKS resources through Pulumi. This approach abstracts the complexity of AWS interactions and streamlines the deployment process, enabling consistent infrastructure management across multi-cloud environments.

## Key Features

- **Kubernetes-Like API Resource Model**: Utilizes a familiar structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, making it intuitive for developers accustomed to Kubernetes to define AWS EKS resources.

- **Unified API Structure**: Ensures consistency across different resources and cloud providers by adhering to a standardized API resource model.

- **Pulumi Integration**: Employs Pulumi for infrastructure provisioning, enabling the use of real programming languages and providing robust state management and automation capabilities.

- **Comprehensive EKS Cluster Configuration**: Supports detailed specification of EKS cluster attributes, including AWS region, VPC integration, and worker node management mode.

- **Region Specification**: Allows the cluster to be deployed in any valid AWS region by specifying the `region` field in the `spec`. This provides flexibility in choosing the geographical location of the cluster.

- **VPC Integration**: Offers the option to deploy the EKS cluster into an existing VPC by specifying the `vpcId`, or to create a new VPC if not specified. This ensures seamless integration with existing network infrastructures.

- **Worker Node Management Modes**: Supports different worker node management modes via the `workersManagementMode` field, including self-managed nodes, managed node groups, and Fargate profiles. This provides flexibility in how compute resources are provisioned and managed.

- **Credential Management**: Securely handles AWS credentials via the `awsCredentialId` field, ensuring authenticated and authorized resource deployments without exposing sensitive information.

- **Status Reporting**: Captures and stores outputs such as the cluster endpoint, certificate authority data, and VPC ID in `status.stackOutputs`. This facilitates easy reference and integration with other systems, such as Kubernetes clients or additional automation tools.

- **Scalability and High Availability**: Enables the creation of highly available clusters by deploying worker nodes across multiple availability zones.

## Architecture

The module operates by accepting an AWS EKS Cluster API resource definition as input. It interprets the resource definition and uses Pulumi to interact with AWS, creating the specified EKS resources. The main components involved are:

- **API Resource Definition**: A YAML file that includes all necessary information to define an EKS cluster, following the standard API structure. Developers specify the cluster's desired state in this file, including the AWS region, VPC ID, and worker node management mode.

- **Pulumi Module**: Written in Go, the module reads the API resource and uses Pulumi's AWS SDK to provision EKS resources based on the provided specifications. It abstracts the complexity of resource creation, update, and deletion.

- **AWS Provider Initialization**: The module initializes the AWS provider within Pulumi using the credentials specified by `awsCredentialId`. This ensures that all AWS resource operations are authenticated and authorized.

- **Resource Creation**: Provisions the EKS cluster and associated resources as defined in the `spec`, including VPC, subnets, security groups, IAM roles, and worker nodes. If a VPC ID is provided, the cluster is deployed into that VPC; otherwise, a new VPC is created.

- **Worker Node Management**: Depending on the `workersManagementMode` specified, the module sets up worker nodes using different strategies:

  - **Self-Managed Nodes**: Nodes are managed by the user, providing full control over the EC2 instances.

  - **Managed Node Groups**: AWS manages the worker nodes, simplifying the provisioning and lifecycle management.

  - **Fargate Profiles**: Serverless compute for containers, allowing you to run pods without managing servers.

- **Status Outputs**: Outputs from the Pulumi deployment, such as the cluster endpoint, certificate authority data, and VPC ID, are captured and stored in `status.stackOutputs`. This information is crucial for connecting to the cluster and deploying applications.

## Usage

Refer to the example section for usage instructions.

## Limitations

- **Advanced Features**: Certain advanced features of EKS, such as custom networking configurations, advanced IAM roles, or specific add-ons, that are not specified in the current API resource definition may not be supported. Future updates may include additional capabilities based on user needs.

- **Region and VPC Changes**: Updating the `region` or `vpcId` fields in the `spec` may result in the recreation of the EKS cluster, as these are critical properties that affect the cluster's deployment.

## Contributing

We welcome contributions to enhance the functionality of this module. Please submit pull requests or open issues to help improve the module and its documentation.
