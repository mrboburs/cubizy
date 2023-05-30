package sendgrid

import (
	"cubizy/keys"
	"cubizy/model"
	"cubizy/util"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var SendgridAPIKey string
var senderName string
var senderEmail string
var api_ready bool

func init() {
	isApiReady()
}

func isApiReady() bool {
	if !api_ready {
		SendgridAPIKey = model.GetSetting(keys.SendgridAPIKey, "")
		senderName = model.GetSetting(keys.SendgridSenderName, "")
		senderEmail = model.GetSetting(keys.SendgridSenderEmail, "")
		if SendgridAPIKey == "" || senderName == "" || senderEmail == "" {
			api_ready = false
		} else {
			api_ready = true
		}
	}
	return api_ready
}

func Reset() {
	api_ready = false
	isApiReady()
}

// SendTestMessage will senttest message , you dont need to give any parameter to this function
func SendTestMessage() {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "tryout405@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	//client := sendgrid.NewSendClient(os.Getenv(SendgridAPIKey))
	client := sendgrid.NewSendClient(SendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		util.Log(err)
	} else {
		util.Log(response.StatusCode)
		util.Log(response.Body)
		util.Log(response.Headers)
	}
}

//SendEmail will sent email according to given details
func SendEmail(Name, Email, Subject, plainTextContent, htmlContent string) error {

	if !isApiReady() {
		return errors.New("api not set yet")
	}
	from := mail.NewEmail(senderName, senderEmail)
	to := mail.NewEmail(Name, Email)

	message := mail.NewSingleEmail(from, Subject, to, plainTextContent, htmlContent)
	//client := sendgrid.NewSendClient(os.Getenv("SendgridAPIKey"))
	client := sendgrid.NewSendClient(SendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		util.Log(err)
		return err
	}

	util.Log(response.StatusCode)
	util.Log(response.Body)
	util.Log(response.Headers)
	return nil
}

// ResetPasswordTemplate will give email content to send password reset code
func ResetPasswordTemplate(Name, Email, ResetCode, ResetLink string) string {

	content, err := ioutil.ReadFile("emailtemplates/resetPassword.html")
	if err != nil {
		log.Fatal(err)
	}
	emailTemplate := string(content)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{username}}", Name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{useremail}}", Email)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{resetcode}}", ResetCode)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{resetlink}}", ResetLink)

	return emailTemplate
}

// VerificationEmailTemplate will give email content to verify email account
func VerificationEmailTemplate(Name, Email, Emailcode, autologinlink string) string {

	content, err := ioutil.ReadFile("emailtemplates/verificationEmail.html")
	if err != nil {
		log.Fatal(err)
	}
	emailTemplate := string(content)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{username}}", Name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{useremail}}", Email)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{emailcode}}", Emailcode)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{autologinlink}}", autologinlink)

	return emailTemplate
}

// InviteStudentTemplate will give email content to invent student
func InviteStudentTemplate(student_name, sender_name, course_name, teacher_photo, teacher_name, academy_logo, academy_name, login_link string) string {

	content, err := ioutil.ReadFile("emailtemplates/inviteStudentEmail.html")
	if err != nil {
		log.Fatal(err)
	}
	emailTemplate := string(content)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{student_name}}", student_name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{sender_name}}", sender_name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{course_name}}", course_name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{teacher_photo}}", teacher_photo)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{teacher_name}}", teacher_name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{academy_logo}}", academy_logo)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{academy_name}}", academy_name)
	emailTemplate = strings.ReplaceAll(emailTemplate, "{{login_link}}", login_link)

	return emailTemplate
}
