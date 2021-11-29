// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package chroot

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-amazon/builder/common"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName         *string                           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType       *string                           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion       *string                           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug             *bool                             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce             *bool                             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError           *string                           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars          map[string]string                 `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars     []string                          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AMIName                 *string                           `mapstructure:"ami_name" required:"true" cty:"ami_name" hcl:"ami_name"`
	AMIDescription          *string                           `mapstructure:"ami_description" required:"false" cty:"ami_description" hcl:"ami_description"`
	AMIVirtType             *string                           `mapstructure:"ami_virtualization_type" required:"false" cty:"ami_virtualization_type" hcl:"ami_virtualization_type"`
	AMIUsers                []string                          `mapstructure:"ami_users" required:"false" cty:"ami_users" hcl:"ami_users"`
	AMIGroups               []string                          `mapstructure:"ami_groups" required:"false" cty:"ami_groups" hcl:"ami_groups"`
	AMIOrgArns              []string                          `mapstructure:"ami_org_arns" required:"false" cty:"ami_org_arns" hcl:"ami_org_arns"`
	AMIOuArns               []string                          `mapstructure:"ami_ou_arns" required:"false" cty:"ami_ou_arns" hcl:"ami_ou_arns"`
	AMIProductCodes         []string                          `mapstructure:"ami_product_codes" required:"false" cty:"ami_product_codes" hcl:"ami_product_codes"`
	AMIRegions              []string                          `mapstructure:"ami_regions" required:"false" cty:"ami_regions" hcl:"ami_regions"`
	AMISkipRegionValidation *bool                             `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation" hcl:"skip_region_validation"`
	AMITags                 map[string]string                 `mapstructure:"tags" required:"false" cty:"tags" hcl:"tags"`
	AMITag                  []config.FlatKeyValue             `mapstructure:"tag" required:"false" cty:"tag" hcl:"tag"`
	AMIENASupport           *bool                             `mapstructure:"ena_support" required:"false" cty:"ena_support" hcl:"ena_support"`
	AMISriovNetSupport      *bool                             `mapstructure:"sriov_support" required:"false" cty:"sriov_support" hcl:"sriov_support"`
	AMIForceDeregister      *bool                             `mapstructure:"force_deregister" required:"false" cty:"force_deregister" hcl:"force_deregister"`
	AMIForceDeleteSnapshot  *bool                             `mapstructure:"force_delete_snapshot" required:"false" cty:"force_delete_snapshot" hcl:"force_delete_snapshot"`
	AMIEncryptBootVolume    *bool                             `mapstructure:"encrypt_boot" required:"false" cty:"encrypt_boot" hcl:"encrypt_boot"`
	AMIKmsKeyId             *string                           `mapstructure:"kms_key_id" required:"false" cty:"kms_key_id" hcl:"kms_key_id"`
	AMIRegionKMSKeyIDs      map[string]string                 `mapstructure:"region_kms_key_ids" required:"false" cty:"region_kms_key_ids" hcl:"region_kms_key_ids"`
	AMISkipBuildRegion      *bool                             `mapstructure:"skip_save_build_region" cty:"skip_save_build_region" hcl:"skip_save_build_region"`
	SnapshotTags            map[string]string                 `mapstructure:"snapshot_tags" required:"false" cty:"snapshot_tags" hcl:"snapshot_tags"`
	SnapshotTag             []config.FlatKeyValue             `mapstructure:"snapshot_tag" required:"false" cty:"snapshot_tag" hcl:"snapshot_tag"`
	SnapshotUsers           []string                          `mapstructure:"snapshot_users" required:"false" cty:"snapshot_users" hcl:"snapshot_users"`
	SnapshotGroups          []string                          `mapstructure:"snapshot_groups" required:"false" cty:"snapshot_groups" hcl:"snapshot_groups"`
	AccessKey               *string                           `mapstructure:"access_key" required:"true" cty:"access_key" hcl:"access_key"`
	AssumeRole              *common.FlatAssumeRoleConfig      `mapstructure:"assume_role" required:"false" cty:"assume_role" hcl:"assume_role"`
	CustomEndpointEc2       *string                           `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2" hcl:"custom_endpoint_ec2"`
	CredsFilename           *string                           `mapstructure:"shared_credentials_file" required:"false" cty:"shared_credentials_file" hcl:"shared_credentials_file"`
	DecodeAuthZMessages     *bool                             `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages" hcl:"decode_authorization_messages"`
	InsecureSkipTLSVerify   *bool                             `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify" hcl:"insecure_skip_tls_verify"`
	MaxRetries              *int                              `mapstructure:"max_retries" required:"false" cty:"max_retries" hcl:"max_retries"`
	MFACode                 *string                           `mapstructure:"mfa_code" required:"false" cty:"mfa_code" hcl:"mfa_code"`
	ProfileName             *string                           `mapstructure:"profile" required:"false" cty:"profile" hcl:"profile"`
	RawRegion               *string                           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	SecretKey               *string                           `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	SkipMetadataApiCheck    *bool                             `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check" hcl:"skip_metadata_api_check"`
	SkipCredsValidation     *bool                             `mapstructure:"skip_credential_validation" cty:"skip_credential_validation" hcl:"skip_credential_validation"`
	Token                   *string                           `mapstructure:"token" required:"false" cty:"token" hcl:"token"`
	VaultAWSEngine          *common.FlatVaultAWSEngineOptions `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine" hcl:"vault_aws_engine"`
	PollingConfig           *common.FlatAWSPollingConfig      `mapstructure:"aws_polling" required:"false" cty:"aws_polling" hcl:"aws_polling"`
	AMIMappings             []common.FlatBlockDevice          `mapstructure:"ami_block_device_mappings" hcl2-schema-generator:"ami_block_device_mappings,direct" required:"false" cty:"ami_block_device_mappings" hcl:"ami_block_device_mappings"`
	ChrootMounts            [][]string                        `mapstructure:"chroot_mounts" required:"false" cty:"chroot_mounts" hcl:"chroot_mounts"`
	CommandWrapper          *string                           `mapstructure:"command_wrapper" required:"false" cty:"command_wrapper" hcl:"command_wrapper"`
	CopyFiles               []string                          `mapstructure:"copy_files" required:"false" cty:"copy_files" hcl:"copy_files"`
	DevicePath              *string                           `mapstructure:"device_path" required:"false" cty:"device_path" hcl:"device_path"`
	NVMEDevicePath          *string                           `mapstructure:"nvme_device_path" required:"false" cty:"nvme_device_path" hcl:"nvme_device_path"`
	FromScratch             *bool                             `mapstructure:"from_scratch" required:"false" cty:"from_scratch" hcl:"from_scratch"`
	MountOptions            []string                          `mapstructure:"mount_options" required:"false" cty:"mount_options" hcl:"mount_options"`
	MountPartition          *string                           `mapstructure:"mount_partition" required:"false" cty:"mount_partition" hcl:"mount_partition"`
	MountPath               *string                           `mapstructure:"mount_path" required:"false" cty:"mount_path" hcl:"mount_path"`
	PostMountCommands       []string                          `mapstructure:"post_mount_commands" required:"false" cty:"post_mount_commands" hcl:"post_mount_commands"`
	PreMountCommands        []string                          `mapstructure:"pre_mount_commands" required:"false" cty:"pre_mount_commands" hcl:"pre_mount_commands"`
	RootDeviceName          *string                           `mapstructure:"root_device_name" required:"false" cty:"root_device_name" hcl:"root_device_name"`
	RootVolumeSize          *int64                            `mapstructure:"root_volume_size" required:"false" cty:"root_volume_size" hcl:"root_volume_size"`
	RootVolumeType          *string                           `mapstructure:"root_volume_type" required:"false" cty:"root_volume_type" hcl:"root_volume_type"`
	SourceAmi               *string                           `mapstructure:"source_ami" required:"true" cty:"source_ami" hcl:"source_ami"`
	SourceAmiFilter         *common.FlatAmiFilterOptions      `mapstructure:"source_ami_filter" required:"false" cty:"source_ami_filter" hcl:"source_ami_filter"`
	RootVolumeTags          map[string]string                 `mapstructure:"root_volume_tags" required:"false" cty:"root_volume_tags" hcl:"root_volume_tags"`
	RootVolumeTag           []config.FlatKeyValue             `mapstructure:"root_volume_tag" required:"false" cty:"root_volume_tag" hcl:"root_volume_tag"`
	RootVolumeEncryptBoot   *bool                             `mapstructure:"root_volume_encrypt_boot" required:"false" cty:"root_volume_encrypt_boot" hcl:"root_volume_encrypt_boot"`
	RootVolumeKmsKeyId      *string                           `mapstructure:"root_volume_kms_key_id" required:"false" cty:"root_volume_kms_key_id" hcl:"root_volume_kms_key_id"`
	Architecture            *string                           `mapstructure:"ami_architecture" required:"false" cty:"ami_architecture" hcl:"ami_architecture"`
	BootMode                *string                           `mapstructure:"boot_mode" required:"false" cty:"boot_mode" hcl:"boot_mode"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":           &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"ami_name":                      &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"ami_description":               &hcldec.AttrSpec{Name: "ami_description", Type: cty.String, Required: false},
		"ami_virtualization_type":       &hcldec.AttrSpec{Name: "ami_virtualization_type", Type: cty.String, Required: false},
		"ami_users":                     &hcldec.AttrSpec{Name: "ami_users", Type: cty.List(cty.String), Required: false},
		"ami_groups":                    &hcldec.AttrSpec{Name: "ami_groups", Type: cty.List(cty.String), Required: false},
		"ami_org_arns":                  &hcldec.AttrSpec{Name: "ami_org_arns", Type: cty.List(cty.String), Required: false},
		"ami_ou_arns":                   &hcldec.AttrSpec{Name: "ami_ou_arns", Type: cty.List(cty.String), Required: false},
		"ami_product_codes":             &hcldec.AttrSpec{Name: "ami_product_codes", Type: cty.List(cty.String), Required: false},
		"ami_regions":                   &hcldec.AttrSpec{Name: "ami_regions", Type: cty.List(cty.String), Required: false},
		"skip_region_validation":        &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"tags":                          &hcldec.AttrSpec{Name: "tags", Type: cty.Map(cty.String), Required: false},
		"tag":                           &hcldec.BlockListSpec{TypeName: "tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"ena_support":                   &hcldec.AttrSpec{Name: "ena_support", Type: cty.Bool, Required: false},
		"sriov_support":                 &hcldec.AttrSpec{Name: "sriov_support", Type: cty.Bool, Required: false},
		"force_deregister":              &hcldec.AttrSpec{Name: "force_deregister", Type: cty.Bool, Required: false},
		"force_delete_snapshot":         &hcldec.AttrSpec{Name: "force_delete_snapshot", Type: cty.Bool, Required: false},
		"encrypt_boot":                  &hcldec.AttrSpec{Name: "encrypt_boot", Type: cty.Bool, Required: false},
		"kms_key_id":                    &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"region_kms_key_ids":            &hcldec.AttrSpec{Name: "region_kms_key_ids", Type: cty.Map(cty.String), Required: false},
		"skip_save_build_region":        &hcldec.AttrSpec{Name: "skip_save_build_region", Type: cty.Bool, Required: false},
		"snapshot_tags":                 &hcldec.AttrSpec{Name: "snapshot_tags", Type: cty.Map(cty.String), Required: false},
		"snapshot_tag":                  &hcldec.BlockListSpec{TypeName: "snapshot_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"snapshot_users":                &hcldec.AttrSpec{Name: "snapshot_users", Type: cty.List(cty.String), Required: false},
		"snapshot_groups":               &hcldec.AttrSpec{Name: "snapshot_groups", Type: cty.List(cty.String), Required: false},
		"access_key":                    &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"assume_role":                   &hcldec.BlockSpec{TypeName: "assume_role", Nested: hcldec.ObjectSpec((*common.FlatAssumeRoleConfig)(nil).HCL2Spec())},
		"custom_endpoint_ec2":           &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"shared_credentials_file":       &hcldec.AttrSpec{Name: "shared_credentials_file", Type: cty.String, Required: false},
		"decode_authorization_messages": &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"insecure_skip_tls_verify":      &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"max_retries":                   &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"mfa_code":                      &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                       &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                    &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_metadata_api_check":       &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"skip_credential_validation":    &hcldec.AttrSpec{Name: "skip_credential_validation", Type: cty.Bool, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":              &hcldec.BlockSpec{TypeName: "vault_aws_engine", Nested: hcldec.ObjectSpec((*common.FlatVaultAWSEngineOptions)(nil).HCL2Spec())},
		"aws_polling":                   &hcldec.BlockSpec{TypeName: "aws_polling", Nested: hcldec.ObjectSpec((*common.FlatAWSPollingConfig)(nil).HCL2Spec())},
		"ami_block_device_mappings":     &hcldec.BlockListSpec{TypeName: "ami_block_device_mappings", Nested: hcldec.ObjectSpec((*common.FlatBlockDevice)(nil).HCL2Spec())},
		"chroot_mounts":                 &hcldec.AttrSpec{Name: "chroot_mounts", Type: cty.List(cty.List(cty.String)), Required: false},
		"command_wrapper":               &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"copy_files":                    &hcldec.AttrSpec{Name: "copy_files", Type: cty.List(cty.String), Required: false},
		"device_path":                   &hcldec.AttrSpec{Name: "device_path", Type: cty.String, Required: false},
		"nvme_device_path":              &hcldec.AttrSpec{Name: "nvme_device_path", Type: cty.String, Required: false},
		"from_scratch":                  &hcldec.AttrSpec{Name: "from_scratch", Type: cty.Bool, Required: false},
		"mount_options":                 &hcldec.AttrSpec{Name: "mount_options", Type: cty.List(cty.String), Required: false},
		"mount_partition":               &hcldec.AttrSpec{Name: "mount_partition", Type: cty.String, Required: false},
		"mount_path":                    &hcldec.AttrSpec{Name: "mount_path", Type: cty.String, Required: false},
		"post_mount_commands":           &hcldec.AttrSpec{Name: "post_mount_commands", Type: cty.List(cty.String), Required: false},
		"pre_mount_commands":            &hcldec.AttrSpec{Name: "pre_mount_commands", Type: cty.List(cty.String), Required: false},
		"root_device_name":              &hcldec.AttrSpec{Name: "root_device_name", Type: cty.String, Required: false},
		"root_volume_size":              &hcldec.AttrSpec{Name: "root_volume_size", Type: cty.Number, Required: false},
		"root_volume_type":              &hcldec.AttrSpec{Name: "root_volume_type", Type: cty.String, Required: false},
		"source_ami":                    &hcldec.AttrSpec{Name: "source_ami", Type: cty.String, Required: false},
		"source_ami_filter":             &hcldec.BlockSpec{TypeName: "source_ami_filter", Nested: hcldec.ObjectSpec((*common.FlatAmiFilterOptions)(nil).HCL2Spec())},
		"root_volume_tags":              &hcldec.AttrSpec{Name: "root_volume_tags", Type: cty.Map(cty.String), Required: false},
		"root_volume_tag":               &hcldec.BlockListSpec{TypeName: "root_volume_tag", Nested: hcldec.ObjectSpec((*config.FlatKeyValue)(nil).HCL2Spec())},
		"root_volume_encrypt_boot":      &hcldec.AttrSpec{Name: "root_volume_encrypt_boot", Type: cty.Bool, Required: false},
		"root_volume_kms_key_id":        &hcldec.AttrSpec{Name: "root_volume_kms_key_id", Type: cty.String, Required: false},
		"ami_architecture":              &hcldec.AttrSpec{Name: "ami_architecture", Type: cty.String, Required: false},
		"boot_mode":                     &hcldec.AttrSpec{Name: "boot_mode", Type: cty.String, Required: false},
	}
	return s
}
