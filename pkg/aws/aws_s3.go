package aws

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Permission string

const (
	PermissionREAD Permission = "READ"

	PermissionWRITE Permission = "WRITE"

	PermissionREADACP Permission = "READACP"

	PermissionWRITEACP Permission = "WRITEACP"

	PermissionFULLCONTROL Permission = "FULLCONTROL"
)

type StorageClass string

const (
	StorageClassSTANDARD StorageClass = "STANDARD"

	StorageClassREDUCEDREDUNDANCY StorageClass = "REDUCEDREDUNDANCY"

	StorageClassGLACIER StorageClass = "GLACIER"

	StorageClassUNKNOWN StorageClass = "UNKNOWN"
)

type MetadataDirective string

const (
	MetadataDirectiveCOPY MetadataDirective = "COPY"

	MetadataDirectiveREPLACE MetadataDirective = "REPLACE"
)

type Payer string

const (
	PayerBucketOwner Payer = "BucketOwner"

	PayerRequester Payer = "Requester"
)

type MfaDeleteStatus string

const (
	MfaDeleteStatusEnabled MfaDeleteStatus = "Enabled"

	MfaDeleteStatusDisabled MfaDeleteStatus = "Disabled"
)

type VersioningStatus string

const (
	VersioningStatusEnabled VersioningStatus = "Enabled"

	VersioningStatusSuspended VersioningStatus = "Suspended"
)

type CreateBucket struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CreateBucket"`

	Bucket            string             `xml:"Bucket,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
	AWSAccessKeyId    string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp         time.Time          `xml:"Timestamp,omitempty"`
	Signature         string             `xml:"Signature,omitempty"`
}

type CreateBucketResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CreateBucketResponse"`

	CreateBucketReturn *CreateBucketResult `xml:"CreateBucketReturn,omitempty"`
}

type DeleteBucket struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ DeleteBucket"`

	Bucket         string    `xml:"Bucket,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type DeleteBucketResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ DeleteBucketResponse"`

	DeleteBucketResponse *Status `xml:"DeleteBucketResponse,omitempty"`
}

type GetBucketLoggingStatus struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetBucketLoggingStatus"`

	Bucket         string    `xml:"Bucket,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type GetBucketLoggingStatusResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetBucketLoggingStatusResponse"`

	GetBucketLoggingStatusResponse *BucketLoggingStatus `xml:"GetBucketLoggingStatusResponse,omitempty"`
}

type SetBucketLoggingStatus struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetBucketLoggingStatus"`

	Bucket              string               `xml:"Bucket,omitempty"`
	AWSAccessKeyId      string               `xml:"AWSAccessKeyId,omitempty"`
	Timestamp           time.Time            `xml:"Timestamp,omitempty"`
	Signature           string               `xml:"Signature,omitempty"`
	Credential          string               `xml:"Credential,omitempty"`
	BucketLoggingStatus *BucketLoggingStatus `xml:"BucketLoggingStatus,omitempty"`
}

type SetBucketLoggingStatusResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetBucketLoggingStatusResponse"`
}

type GetObjectAccessControlPolicy struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectAccessControlPolicy"`

	Bucket         string    `xml:"Bucket,omitempty"`
	Key            string    `xml:"Key,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type GetObjectAccessControlPolicyResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectAccessControlPolicyResponse"`

	GetObjectAccessControlPolicyResponse *AccessControlPolicy `xml:"GetObjectAccessControlPolicyResponse,omitempty"`
}

type GetBucketAccessControlPolicy struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetBucketAccessControlPolicy"`

	Bucket         string    `xml:"Bucket,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type GetBucketAccessControlPolicyResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetBucketAccessControlPolicyResponse"`

	GetBucketAccessControlPolicyResponse *AccessControlPolicy `xml:"GetBucketAccessControlPolicyResponse,omitempty"`
}

type SetObjectAccessControlPolicy struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetObjectAccessControlPolicy"`

	Bucket            string             `xml:"Bucket,omitempty"`
	Key               string             `xml:"Key,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
	AWSAccessKeyId    string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp         time.Time          `xml:"Timestamp,omitempty"`
	Signature         string             `xml:"Signature,omitempty"`
	Credential        string             `xml:"Credential,omitempty"`
}

type SetObjectAccessControlPolicyResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetObjectAccessControlPolicyResponse"`
}

type SetBucketAccessControlPolicy struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetBucketAccessControlPolicy"`

	Bucket            string             `xml:"Bucket,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
	AWSAccessKeyId    string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp         time.Time          `xml:"Timestamp,omitempty"`
	Signature         string             `xml:"Signature,omitempty"`
	Credential        string             `xml:"Credential,omitempty"`
}

type SetBucketAccessControlPolicyResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ SetBucketAccessControlPolicyResponse"`
}

type GetObject struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObject"`

	Bucket         string    `xml:"Bucket,omitempty"`
	Key            string    `xml:"Key,omitempty"`
	GetMetadata    bool      `xml:"GetMetadata,omitempty"`
	GetData        bool      `xml:"GetData,omitempty"`
	InlineData     bool      `xml:"InlineData,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type GetObjectResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectResponse"`

	GetObjectResponse *GetObjectResult `xml:"GetObjectResponse,omitempty"`
}

type GetObjectExtended struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectExtended"`

	Bucket                                 string    `xml:"Bucket,omitempty"`
	Key                                    string    `xml:"Key,omitempty"`
	GetMetadata                            bool      `xml:"GetMetadata,omitempty"`
	GetData                                bool      `xml:"GetData,omitempty"`
	InlineData                             bool      `xml:"InlineData,omitempty"`
	ByteRangeStart                         int64     `xml:"ByteRangeStart,omitempty"`
	ByteRangeEnd                           int64     `xml:"ByteRangeEnd,omitempty"`
	IfModifiedSince                        time.Time `xml:"IfModifiedSince,omitempty"`
	IfUnmodifiedSince                      time.Time `xml:"IfUnmodifiedSince,omitempty"`
	IfMatch                                string    `xml:"IfMatch,omitempty"`
	IfNoneMatch                            string    `xml:"IfNoneMatch,omitempty"`
	ReturnCompleteObjectOnConditionFailure bool      `xml:"ReturnCompleteObjectOnConditionFailure,omitempty"`
	AWSAccessKeyId                         string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp                              time.Time `xml:"Timestamp,omitempty"`
	Signature                              string    `xml:"Signature,omitempty"`
	Credential                             string    `xml:"Credential,omitempty"`
}

type GetObjectExtendedResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectExtendedResponse"`

	GetObjectResponse *GetObjectResult `xml:"GetObjectResponse,omitempty"`
}

type PutObject struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PutObject"`

	Bucket            string             `xml:"Bucket,omitempty"`
	Key               string             `xml:"Key,omitempty"`
	Metadata          *MetadataEntry     `xml:"Metadata,omitempty"`
	ContentLength     int64              `xml:"ContentLength,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
	StorageClass      *StorageClass      `xml:"StorageClass,omitempty"`
	AWSAccessKeyId    string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp         time.Time          `xml:"Timestamp,omitempty"`
	Signature         string             `xml:"Signature,omitempty"`
	Credential        string             `xml:"Credential,omitempty"`
}

type PutObjectResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PutObjectResponse"`

	PutObjectResponse *PutObjectResult `xml:"PutObjectResponse,omitempty"`
}

type PutObjectInline struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PutObjectInline"`

	Bucket            string             `xml:"Bucket,omitempty"`
	Key               string             `xml:"Key,omitempty"`
	Metadata          *MetadataEntry     `xml:"Metadata,omitempty"`
	Data              []byte             `xml:"Data,omitempty"`
	ContentLength     int64              `xml:"ContentLength,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
	StorageClass      *StorageClass      `xml:"StorageClass,omitempty"`
	AWSAccessKeyId    string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp         time.Time          `xml:"Timestamp,omitempty"`
	Signature         string             `xml:"Signature,omitempty"`
	Credential        string             `xml:"Credential,omitempty"`
}

type PutObjectInlineResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PutObjectInlineResponse"`

	PutObjectInlineResponse *PutObjectResult `xml:"PutObjectInlineResponse,omitempty"`
}

type DeleteObject struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ DeleteObject"`

	Bucket         string    `xml:"Bucket,omitempty"`
	Key            string    `xml:"Key,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type DeleteObjectResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ DeleteObjectResponse"`

	DeleteObjectResponse *Status `xml:"DeleteObjectResponse,omitempty"`
}

type ListBucket struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucket"`

	Bucket         string    `xml:"Bucket,omitempty"`
	Prefix         string    `xml:"Prefix,omitempty"`
	Marker         string    `xml:"Marker,omitempty"`
	MaxKeys        int32     `xml:"MaxKeys,omitempty"`
	Delimiter      string    `xml:"Delimiter,omitempty"`
	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
	Credential     string    `xml:"Credential,omitempty"`
}

type ListBucketResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResponse"`

	ListBucketResponse *ListBucketResult `xml:"ListBucketResponse,omitempty"`
}

type ListVersionsResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListVersionsResponse"`

	ListVersionsResponse *ListVersionsResult `xml:"ListVersionsResponse,omitempty"`
}

type ListAllMyBuckets struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBuckets"`

	AWSAccessKeyId string    `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time `xml:"Timestamp,omitempty"`
	Signature      string    `xml:"Signature,omitempty"`
}

type ListAllMyBucketsResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResponse"`

	ListAllMyBucketsResponse *ListAllMyBucketsResult `xml:"ListAllMyBucketsResponse,omitempty"`
}

type PostResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PostResponse"`

	Bucket string `xml:"Bucket,omitempty"`
	Key    string `xml:"Key,omitempty"`
	ETag   string `xml:"ETag,omitempty"`
}

type CopyObject struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CopyObject"`

	SourceBucket                string             `xml:"SourceBucket,omitempty"`
	SourceKey                   string             `xml:"SourceKey,omitempty"`
	DestinationBucket           string             `xml:"DestinationBucket,omitempty"`
	DestinationKey              string             `xml:"DestinationKey,omitempty"`
	MetadataDirective           *MetadataDirective `xml:"MetadataDirective,omitempty"`
	Metadata                    *MetadataEntry     `xml:"Metadata,omitempty"`
	AccessControlList           *AccessControlList `xml:"AccessControlList,omitempty"`
	CopySourceIfModifiedSince   time.Time          `xml:"CopySourceIfModifiedSince,omitempty"`
	CopySourceIfUnmodifiedSince time.Time          `xml:"CopySourceIfUnmodifiedSince,omitempty"`
	CopySourceIfMatch           string             `xml:"CopySourceIfMatch,omitempty"`
	CopySourceIfNoneMatch       string             `xml:"CopySourceIfNoneMatch,omitempty"`
	StorageClass                *StorageClass      `xml:"StorageClass,omitempty"`
	AWSAccessKeyId              string             `xml:"AWSAccessKeyId,omitempty"`
	Timestamp                   time.Time          `xml:"Timestamp,omitempty"`
	Signature                   string             `xml:"Signature,omitempty"`
	Credential                  string             `xml:"Credential,omitempty"`
}

type CopyObjectResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CopyObjectResponse"`

	CopyObjectResult *CopyObjectResult `xml:"CopyObjectResult,omitempty"`
}

type MetadataEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ MetadataEntry"`

	Name  string `xml:"Name,omitempty"`
	Value string `xml:"Value,omitempty"`
}

type Status struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ Status"`

	Code        int32  `xml:"Code,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type Result struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ Result"`

	Status *Status `xml:"Status,omitempty"`
}

type CreateBucketResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CreateBucketResult"`

	BucketName string `xml:"BucketName,omitempty"`
}

type BucketLoggingStatus struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ BucketLoggingStatus"`

	LoggingEnabled *LoggingSettings `xml:"LoggingEnabled,omitempty"`
}

type LoggingSettings struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ LoggingSettings"`

	TargetBucket string             `xml:"TargetBucket,omitempty"`
	TargetPrefix string             `xml:"TargetPrefix,omitempty"`
	TargetGrants *AccessControlList `xml:"TargetGrants,omitempty"`
}

type Grantee struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ Grantee"`
}

type User struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ User"`

	*Grantee
}

type AmazonCustomerByEmail struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ AmazonCustomerByEmail"`

	*User

	EmailAddress string `xml:"EmailAddress,omitempty"`
}

type CanonicalUser struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CanonicalUser"`

	*User

	ID          string `xml:"ID,omitempty"`
	DisplayName string `xml:"DisplayName,omitempty"`
}

type Group struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ Group"`

	*Grantee

	URI string `xml:"URI,omitempty"`
}

type Grant struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ Grant"`

	Grantee    *Grantee    `xml:"Grantee,omitempty"`
	Permission *Permission `xml:"Permission,omitempty"`
}

type AccessControlList struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ AccessControlList"`

	Grant *Grant `xml:"Grant,omitempty"`
}

type CreateBucketConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CreateBucketConfiguration"`

	LocationConstraint *LocationConstraint `xml:"LocationConstraint,omitempty"`
}

type LocationConstraint struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ LocationConstraint"`

	Value string
}

type AccessControlPolicy struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ AccessControlPolicy"`

	Owner             *CanonicalUser     `xml:"Owner,omitempty"`
	AccessControlList *AccessControlList `xml:"AccessControlList,omitempty"`
}

type GetObjectResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ GetObjectResult"`

	*Result

	Metadata     []*MetadataEntry `xml:"Metadata,omitempty"`
	Data         []byte           `xml:"Data,omitempty"`
	LastModified time.Time        `xml:"LastModified,omitempty"`
	ETag         string           `xml:"ETag,omitempty"`
}

type PutObjectResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PutObjectResult"`

	ETag         string    `xml:"ETag,omitempty"`
	LastModified time.Time `xml:"LastModified,omitempty"`
}

type ListEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListEntry"`

	Key          string         `xml:"Key,omitempty"`
	LastModified time.Time      `xml:"LastModified,omitempty"`
	ETag         string         `xml:"ETag,omitempty"`
	Size         int64          `xml:"Size,omitempty"`
	Owner        *CanonicalUser `xml:"Owner,omitempty"`
	StorageClass *StorageClass  `xml:"StorageClass,omitempty"`
}

type VersionEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ VersionEntry"`

	Key          string         `xml:"Key,omitempty"`
	VersionId    string         `xml:"VersionId,omitempty"`
	IsLatest     bool           `xml:"IsLatest,omitempty"`
	LastModified time.Time      `xml:"LastModified,omitempty"`
	ETag         string         `xml:"ETag,omitempty"`
	Size         int64          `xml:"Size,omitempty"`
	Owner        *CanonicalUser `xml:"Owner,omitempty"`
	StorageClass *StorageClass  `xml:"StorageClass,omitempty"`
}

type DeleteMarkerEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ DeleteMarkerEntry"`

	Key          string         `xml:"Key,omitempty"`
	VersionId    string         `xml:"VersionId,omitempty"`
	IsLatest     bool           `xml:"IsLatest,omitempty"`
	LastModified time.Time      `xml:"LastModified,omitempty"`
	Owner        *CanonicalUser `xml:"Owner,omitempty"`
}

type PrefixEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ PrefixEntry"`

	Prefix string `xml:"Prefix,omitempty"`
}

type ListBucketResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResult"`

	Metadata       []*MetadataEntry `xml:"Metadata,omitempty"`
	Name           string           `xml:"Name,omitempty"`
	Prefix         string           `xml:"Prefix,omitempty"`
	Marker         string           `xml:"Marker,omitempty"`
	NextMarker     string           `xml:"NextMarker,omitempty"`
	MaxKeys        int32            `xml:"MaxKeys,omitempty"`
	Delimiter      string           `xml:"Delimiter,omitempty"`
	IsTruncated    bool             `xml:"IsTruncated,omitempty"`
	Contents       []*ListEntry     `xml:"Contents,omitempty"`
	CommonPrefixes []*PrefixEntry   `xml:"CommonPrefixes,omitempty"`
}

type ListVersionsResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListVersionsResult"`

	Metadata            []*MetadataEntry `xml:"Metadata,omitempty"`
	Name                string           `xml:"Name,omitempty"`
	Prefix              string           `xml:"Prefix,omitempty"`
	KeyMarker           string           `xml:"KeyMarker,omitempty"`
	VersionIdMarker     string           `xml:"VersionIdMarker,omitempty"`
	NextKeyMarker       string           `xml:"NextKeyMarker,omitempty"`
	NextVersionIdMarker string           `xml:"NextVersionIdMarker,omitempty"`
	MaxKeys             int32            `xml:"MaxKeys,omitempty"`
	Delimiter           string           `xml:"Delimiter,omitempty"`
	IsTruncated         bool             `xml:"IsTruncated,omitempty"`
	CommonPrefixes      []*PrefixEntry   `xml:"CommonPrefixes,omitempty"`

	Version      *VersionEntry      `xml:"Version,omitempty"`
	DeleteMarker *DeleteMarkerEntry `xml:"DeleteMarker,omitempty"`
}

type ListAllMyBucketsEntry struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsEntry"`

	Name         string    `xml:"Name,omitempty"`
	CreationDate time.Time `xml:"CreationDate,omitempty"`
}

type ListAllMyBucketsResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResult"`

	Owner   *CanonicalUser        `xml:"Owner,omitempty"`
	Buckets *ListAllMyBucketsList `xml:"Buckets,omitempty"`
}

type ListAllMyBucketsList struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsList"`

	Bucket []*ListAllMyBucketsEntry `xml:"Bucket,omitempty"`
}

type CopyObjectResult struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CopyObjectResult"`

	LastModified time.Time `xml:"LastModified,omitempty"`
	ETag         string    `xml:"ETag,omitempty"`
}

type RequestPaymentConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ RequestPaymentConfiguration"`

	Payer *Payer `xml:"Payer,omitempty"`
}

type VersioningConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ VersioningConfiguration"`

	Status    *VersioningStatus `xml:"Status,omitempty"`
	MfaDelete *MfaDeleteStatus  `xml:"MfaDelete,omitempty"`
}

type NotificationConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ NotificationConfiguration"`

	TopicConfiguration []*TopicConfiguration `xml:"TopicConfiguration,omitempty"`
}

type TopicConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ TopicConfiguration"`

	Topic string   `xml:"Topic,omitempty"`
	Event []string `xml:"Event,omitempty"`
}

type AmazonS3 struct {
	client *SOAPClient
}

func NewAmazonS3(url string, tls bool, auth *BasicAuth) *AmazonS3 {
	if url == "" {
		url = "https://s3.amazonaws.com/soap"
	}
	client := NewSOAPClient(url, tls, auth)

	return &AmazonS3{
		client: client,
	}
}

func (service *AmazonS3) CreateBucket(request *CreateBucket) (*CreateBucketResponse, error) {
	response := new(CreateBucketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) DeleteBucket(request *DeleteBucket) (*DeleteBucketResponse, error) {
	response := new(DeleteBucketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) GetObjectAccessControlPolicy(request *GetObjectAccessControlPolicy) (*GetObjectAccessControlPolicyResponse, error) {
	response := new(GetObjectAccessControlPolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) GetBucketAccessControlPolicy(request *GetBucketAccessControlPolicy) (*GetBucketAccessControlPolicyResponse, error) {
	response := new(GetBucketAccessControlPolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) SetObjectAccessControlPolicy(request *SetObjectAccessControlPolicy) (*SetObjectAccessControlPolicyResponse, error) {
	response := new(SetObjectAccessControlPolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) SetBucketAccessControlPolicy(request *SetBucketAccessControlPolicy) (*SetBucketAccessControlPolicyResponse, error) {
	response := new(SetBucketAccessControlPolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) GetObject(request *GetObject) (*GetObjectResponse, error) {
	response := new(GetObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) GetObjectExtended(request *GetObjectExtended) (*GetObjectExtendedResponse, error) {
	response := new(GetObjectExtendedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) PutObject(request *PutObject) (*PutObjectResponse, error) {
	response := new(PutObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) PutObjectInline(request *PutObjectInline) (*PutObjectInlineResponse, error) {
	response := new(PutObjectInlineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) DeleteObject(request *DeleteObject) (*DeleteObjectResponse, error) {
	response := new(DeleteObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) ListBucket(request *ListBucket) (*ListBucketResponse, error) {
	response := new(ListBucketResponse)
	err := service.client.Call("http://s3.amazonaws.com/doc/2006-03-01/ListBucket", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) ListAllMyBuckets(request *ListAllMyBuckets) (*ListAllMyBucketsResponse, error) {
	response := new(ListAllMyBucketsResponse)
	err := service.client.Call("http://s3.amazonaws.com/doc/2006-03-01/ListAllMyBuckets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) GetBucketLoggingStatus(request *GetBucketLoggingStatus) (*GetBucketLoggingStatusResponse, error) {
	response := new(GetBucketLoggingStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) SetBucketLoggingStatus(request *SetBucketLoggingStatus) (*SetBucketLoggingStatusResponse, error) {
	response := new(SetBucketLoggingStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *AmazonS3) CopyObject(request *CopyObject) (*CopyObjectResponse, error) {
	response := new(CopyObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
