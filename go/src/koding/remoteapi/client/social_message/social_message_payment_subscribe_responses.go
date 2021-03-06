package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialMessagePaymentSubscribeReader is a Reader for the SocialMessagePaymentSubscribe structure.
type SocialMessagePaymentSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessagePaymentSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessagePaymentSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessagePaymentSubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessagePaymentSubscribeOK creates a SocialMessagePaymentSubscribeOK with default headers values
func NewSocialMessagePaymentSubscribeOK() *SocialMessagePaymentSubscribeOK {
	return &SocialMessagePaymentSubscribeOK{}
}

/*SocialMessagePaymentSubscribeOK handles this case with default header values.

Request processed successfully
*/
type SocialMessagePaymentSubscribeOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessagePaymentSubscribeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.paymentSubscribe][%d] socialMessagePaymentSubscribeOK  %+v", 200, o.Payload)
}

func (o *SocialMessagePaymentSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessagePaymentSubscribeUnauthorized creates a SocialMessagePaymentSubscribeUnauthorized with default headers values
func NewSocialMessagePaymentSubscribeUnauthorized() *SocialMessagePaymentSubscribeUnauthorized {
	return &SocialMessagePaymentSubscribeUnauthorized{}
}

/*SocialMessagePaymentSubscribeUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessagePaymentSubscribeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessagePaymentSubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.paymentSubscribe][%d] socialMessagePaymentSubscribeUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessagePaymentSubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
