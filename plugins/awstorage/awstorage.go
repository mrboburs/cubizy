package awstorage

import (
	"archive/zip"
	"bytes"
	"cubizy/keys"
	"cubizy/model"
	"cubizy/util"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess *session.Session
var svc *s3.S3
var downloader *s3manager.Downloader
var uploader *s3manager.Uploader

var S3Bucket string
var Region string
var allset = false

func init() {
	util.Log("s3 module initiating")
	defer util.Log("s3 module initiated")
	sets3()
}

func sets3() bool {
	if !allset {
		S3Bucket = model.GetSetting(keys.S3Bucket, "")
		Region = model.GetSetting(keys.S3Region, "")
		S3AwsAccessKeyID := model.GetSetting(keys.S3AwsAccessKeyID, "")
		S3AwsSecretAccessKey := model.GetSetting(keys.S3AwsSecretAccessKey, "")

		if S3Bucket == "" || S3AwsAccessKeyID == "" || S3AwsSecretAccessKey == "" {
			return false
		}
		var err error
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String(Region),
			Credentials: credentials.NewStaticCredentials(S3AwsAccessKeyID, S3AwsSecretAccessKey, ""),
		})

		if err != nil {
			return false
		} else {
			// Create S3 service client
			svc = s3.New(sess)
			uploader = s3manager.NewUploader(sess)
			downloader = s3manager.NewDownloader(sess)
			allset = true // TestS3()
		}
	}
	return true
}

// ResetS3 will reset all s3 settings
func ResetS3() {
	allset = false
	sets3()
}

// TestS3 will test working of conected s3 apis
func TestS3() bool {
	return getBucketRegion(Region)
	/*
				files, err := ListBucketItems("", "")
				if err == nil {
					util.Log(files)
					return true
				} else {
					util.Log(err.Error())
					return false
				}
		}

		func listBuckets() {

			result, err := svc.ListBuckets(nil)
			if err != nil {
				util.Log("Unable to list buckets,", err)
			}

			util.Log("Buckets:")

			for _, b := range result.Buckets {
				util.Log("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
			}
	*/
}

func getBucketRegion(bucket string) bool {
	ctx := aws.BackgroundContext()
	region, err := s3manager.GetBucketRegion(ctx, sess, bucket, "us-west-2")
	if err != nil {
		util.Log(err)
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			util.Log("unable to find bucket region for ", bucket)
		}
		return false
	}
	util.Log("Bucket : ", bucket, " region : ", region, "\n")
	return true
}

// GetPresignedGetURL url will return url to view file for just 2 min
func GetPresignedGetURL(key string) (string, error) {
	//var output *s3.PutObjectOutput
	var req *request.Request
	var url string
	var err error
	if sets3() {
		req, _ = svc.GetObjectRequest(&s3.GetObjectInput{Bucket: aws.String(S3Bucket), Key: aws.String(key)})
		url, err = req.Presign(time.Hour * 24 * 7)
		if err != nil {
			util.Log("Unable to put item url in bucket ", S3Bucket, ", ", err)
		}
	} else {
		err = errors.New("storage api not set")
	}
	return url, err
}

// GetPresignedPutURL url will return url to upload file for just 2 min
func GetPresignedPutURL(key, contenttype string) (string, error) {
	//var output *s3.PutObjectOutput
	var req *request.Request
	var url string
	var err error
	if sets3() {
		req, _ = svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket:      aws.String(S3Bucket),
			Key:         aws.String(key),
			ACL:         aws.String("public-read"),
			ContentType: aws.String(contenttype),
		})
		url, err = req.Presign(time.Minute * 10)
		if err != nil {
			util.Log("Unable to put item url in bucket ", S3Bucket, ", ", err)
		}
	} else {
		err = errors.New("storage api not set")
	}
	return url, err
}

// ListBucketItems will list all items in s3 Bucket
func ListBucketItems(prefix, delimiter string) ([]*s3.Object, error) {
	var files = make([]*s3.Object, 0)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(S3Bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String(delimiter)})
	if err != nil {
		util.Log("Unable to list items in bucket ", S3Bucket, ", ", err)
	} else {
		files = resp.Contents
		/*
			for _, item := range resp.Contents {
				util.Log("Name:         ", *item.Key)
				util.Log("Last modified:", *item.LastModified)
				util.Log("Size:         ", *item.Size)
				util.Log("Storage class:", *item.StorageClass)
				util.Log("")
			}
		*/
	}
	return files, err
}

// GetSize calculet size of folder
func GetSize(prefix string) (int64, error) {
	var size int64
	var delimiter = ""
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(S3Bucket), Prefix: aws.String(prefix), Delimiter: aws.String(delimiter)})
	if err != nil {
		util.Log("Unable to list items in bucket ", S3Bucket, ", ", err)
	} else {
		for _, item := range resp.Contents {
			size += *item.Size
		}
	}
	return size, err
}

// Download file
func Download(item string, newFileName string) error {
	file, err := os.Create(newFileName)
	if err != nil {
		util.Log("Unable to open file ", item, ", ", err)
	}
	defer file.Close()
	//var numBytes int64
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(S3Bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		util.Log("Unable to download item ", item, ", ", err)
	}
	//else {
	//util.Log("Downloaded", file.Name(), numBytes, "bytes")
	//}
	return err
	//upload(bucket, newFileName)
}

// Read file
func ReadFile(item string) (string, error) {
	var myFileContentAsString string
	rawObject, err := svc.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(S3Bucket),
			Key:    aws.String(item),
		})
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(rawObject.Body)
		myFileContentAsString = buf.String()
	}
	return myFileContentAsString, err
	//upload(bucket, newFileName)
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(files []string, folder string, filename string) (string, error) {
	err := os.MkdirAll(folder, 0755)
	newfilePath := folder + "/" + filename + ".zip"
	if err == nil {

		newZipFile, err := os.Create(newfilePath)
		if err == nil {
			defer newZipFile.Close()

			zipWriter := zip.NewWriter(newZipFile)
			defer zipWriter.Close()

			// Add files to zip
			for _, file := range files {
				filenameparts := strings.Split(file, "/")
				newFilName := folder + "/" + filenameparts[len(filenameparts)-1]
				Download(file, newFilName)
				if err = AddFileToZip(zipWriter, newFilName); err != nil {
					util.Log(err)
				}
				os.Remove(newFilName)
			}
		}
	}
	return newfilePath, err
}

// AddFileToZip will add given file to zip file
func AddFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

// Upload will Upload file to given s3 bucket
func Upload(filename string, key string) error {
	file, err := os.Open(filename)
	if err != nil {
		util.Log("Unable to open file ", filename, ", ", err)
	}
	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		util.Log("Unable to upload  ", filename, "to ", S3Bucket, ", ", err)
	}
	return err
	//util.Log("Successfully uploaded ", filename , "to ", bucket)
}

func CopyFile(source, item string) error {
	var err error
	// Copy the item
	_, err = svc.CopyObject(
		&s3.CopyObjectInput{
			Bucket:     aws.String(S3Bucket),
			CopySource: aws.String(S3Bucket + "/" + source),
			Key:        aws.String(item),
			ACL:        aws.String("public-read"),
		},
	)
	if err != nil {
		util.Log("Unable to copy item from ", S3Bucket+"/"+source, " to ", item, ", ", err)
	} else {
		// Wait to see if the item got copied
		err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String(S3Bucket), Key: aws.String(item)})
		if err != nil {
			util.Log("Error occurred while waiting for item ", source, " to be copied to bucket ", item, ", ", err)
		}
		// else {
		//	util.Log("Item ", item, " successfully copied from source ", source)
		//}
	}
	return err
}

// UploadKey will creat empty file for given name
func UploadKey(filename string) error {
	if filename == "" {
		return errors.New("empty folder name")
	}
	if strings.Contains(filename, ".") {
		return errors.New("Folder name can not have . given name is " + filename)
	}
	file := strings.NewReader("")

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		util.Log("Unable to upload ", filename, ", ", err)
	}
	return err
}

// UpdatePublicKey will make file public for given name
func UpdatePublicKey(key string) error {
	if key == "" {
		return errors.New("empty folder name")
	}
	_, err := svc.PutObjectAcl(&s3.PutObjectAclInput{
		AccessControlPolicy: &s3.AccessControlPolicy{},
		Bucket:              aws.String(S3Bucket),
		ACL:                 aws.String("public-read"),
		Key:                 aws.String("/" + key),
	})
	if err != nil {
		// Print the error and exit.
		util.Log("Unable to update", key, err)
	} else {
		util.Log("File updated ", key)

	}
	return err
}

// SetFolders folders in given path
func SetFolders(path string) {
	util.Log("setting folder for path ", path)
	fullpath := ""
	folders := strings.Split(path, "/")
	for _, folder := range folders {
		if !strings.Contains(folder, ".") {
			if folder != "" {
				folder += "/"
			}
			fullpath += folder
			err := UploadKey(fullpath)
			if err != nil {
				util.Log(err)
			} else {
				util.Log("Folder created by name ", folder)
			}
		}
	}
}

// Delete will delete file from S3 bucket
func Delete(key string) error {
	newVal := key[len(key)-1:]
	if newVal == "/" {

		Objects, err := ListBucketItems(key, "")
		objectsToDelete := make([]*s3.ObjectIdentifier, 0, 1000)
		for _, object := range Objects {
			obj := s3.ObjectIdentifier{
				Key: object.Key,
			}
			objectsToDelete = append(objectsToDelete, &obj)
		}

		if err == nil {
			_, err = svc.DeleteObjects(&s3.DeleteObjectsInput{Bucket: aws.String(S3Bucket), Delete: &s3.Delete{Objects: objectsToDelete}})
		}
		return err
	}
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(S3Bucket), Key: aws.String(key)})
	if err != nil {
		util.Log("Unable to delete object ", key, "  from bucket ", S3Bucket, ", ", err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(key),
	})
	//if err == nil {
	// from bucket("Object ", key, " successfully deleted")
	//}
	return err
}

// GetAccessURL will give url to access files on s3 server
func GetAccessURL() string {
	return "https://" + S3Bucket + ".s3." + Region + ".amazonaws.com/"
}
