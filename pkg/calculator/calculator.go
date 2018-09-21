package calculator

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

type Add struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Add"`

	IntA int32 `xml:"intA,omitempty"`
	IntB int32 `xml:"intB,omitempty"`
}

type AddResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AddResponse"`

	AddResult int32 `xml:"AddResult,omitempty"`
}

type Subtract struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Subtract"`

	IntA int32 `xml:"intA,omitempty"`
	IntB int32 `xml:"intB,omitempty"`
}

type SubtractResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SubtractResponse"`

	SubtractResult int32 `xml:"SubtractResult,omitempty"`
}

type Multiply struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Multiply"`

	IntA int32 `xml:"intA,omitempty"`
	IntB int32 `xml:"intB,omitempty"`
}

type MultiplyResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ MultiplyResponse"`

	MultiplyResult int32 `xml:"MultiplyResult,omitempty"`
}

type Divide struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Divide"`

	IntA int32 `xml:"intA,omitempty"`
	IntB int32 `xml:"intB,omitempty"`
}

type DivideResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ DivideResponse"`

	DivideResult int32 `xml:"DivideResult,omitempty"`
}

type CalculatorSoap struct {
	client *SOAPClient
}

func NewCalculatorSoap(url string, tls bool, auth *BasicAuth) *CalculatorSoap {
	if url == "" {
		url = "http://www.dneonline.com/calculator.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &CalculatorSoap{
		client: client,
	}
}

/* Adds two integers. This is a test WebService. ©DNE Online */
func (service *CalculatorSoap) Add(request *Add) (*AddResponse, error) {
	response := new(AddResponse)
	err := service.client.Call("http://tempuri.org/Add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CalculatorSoap) Subtract(request *Subtract) (*SubtractResponse, error) {
	response := new(SubtractResponse)
	err := service.client.Call("http://tempuri.org/Subtract", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CalculatorSoap) Multiply(request *Multiply) (*MultiplyResponse, error) {
	response := new(MultiplyResponse)
	err := service.client.Call("http://tempuri.org/Multiply", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CalculatorSoap) Divide(request *Divide) (*DivideResponse, error) {
	response := new(DivideResponse)
	err := service.client.Call("http://tempuri.org/Divide", request, response)
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

	Body SOAPBody
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
	envelope := SOAPEnvelope{
		//Header:        SoapHeader{},
	}

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

	//log.Println(buffer.String())

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

	//log.Println(string(rawbody))
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
