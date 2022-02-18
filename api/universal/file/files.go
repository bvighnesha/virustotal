package file

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type Files interface {
	// Upload and analyse a files
	Upload(file os.File)
	// UploadFromPath Upload and analyse a files
	UploadFromPath(file string)
	// UploadFromStream Upload and analyse a files
	UploadFromStream(filename string, file io.Reader)
	// UploadURLForLargerThan32MB Get a URL for uploading files larger than 32MB
	UploadURLForLargerThan32MB() (string, error)
	// RetrieveInformation Retrieve information about a files
	RetrieveInformation(id string)
	//ReAnalyse a files already in VirusTotal
	/*
		This API endpoint has the potential to produce a denial of service on the scanning infrastructure if abused.
		Please contact us if you are going to be rescanning more than 50K files per day.
	*/
	ReAnalyse(id string)
	// Comments Get comments on a files
	Comments(id, cursor string, limit int32)
	/*
		With this endpoint you can post a comment for a given files. The body for the POST request must be the JSON representation of a comment object.
		Notice however that you don't need to provide an ID for the object, as they are automatically generated for new comments.

		Any word starting with # in your comment's text will be considered a tag, and added to the comment's tag attribute.
	*/
	Comment(id, data string)
	// Votes Get votes on a files
	Votes(id, cursor string, limit int32)
	/*
		With this endpoint you can post a vote for a given files. The body for the POST request must be the JSON representation of a vote object.
		Notice however that you don't need to provide an ID for the object, as they are automatically generated for new votes.

		The verdict attribute must have be either harmless or malicious.
	*/
	Vote(id, data string)
	/*
		Files objects have many relationships to other files and objects.
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URL.

		Some relationships are accessible only to users who have access to VirusTotal Enterprise package.

		Available relationships are described in the Files object documentation.
	*/
	ObjectsRelated(id, cursor, relationship string, limit int32)
	/*
		This endpoint is the same as /files/{id}/{relationship}
		except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptorsRelated(id, relationship, limit, cursor string)
	/*
		This endpoint returns a summary with behavioural information about the files.
		The summary consists in merging together the reports produced by the multiple sandboxes we have integrated in VirusTotal.

		This API call returns all fields contained in the Files behaviour object, except the ones that make sense only for individual sandboxes:

		analysis_date
		rehash
		has_html_report
		has_pcap
		last_modification_date
		sandbox_name
	*/
	BehaviourSummary(id string)
	/*
		A Sandbox report ID has two main components: the analysed files's SHA256 and the sandbox name.
		These two components are joined by a _ character.
		For example, ID 5353e23f3653402339c93a8565307c6308ff378e03fcf23a4378f31c434030b0_VirusTotal Jujubox fetches
		the sandbox report for a files having a SHA256 5353e23f3653402339c93a8565307c6308ff378e03fcf23a4378f31c434030b0 analysed in the VirusTotal Jujubox sandbox.
	*/
	BehaviourReportFromSandbox(sandboxId string)
	/*
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URL.

		Available relationships are described in the Files behaviour object documentation.
	*/
	ObjectsRelatedToBehaviourReport(sandboxId, relationship, cursor string, limit int32)
	/*
		This endpoint is the same as /file_behaviours/{sandbox_id}/{relationship} except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptorsRelatedToBehaviourReport(sandboxId, relationship, cursor string, limit int32)
	// DetailedHTMLBehaviourReport Returns a Files behaviour object as an HTML report.
	DetailedHTMLBehaviourReport(sandboxId string)
	// SigmaAnalyses Get a files’s Sigma analysis
	SigmaAnalyses(id string)
	// ObjectsRelatedToSigmaAnalysis Get objects related to a Sigma analysis
	ObjectsRelatedToSigmaAnalysis(id, relationship, cursor string, limit int32)
	// ObjectDescriptorsRelatedToSigmaAnalysis This endpoint is the same as /sigma_analyses/{id}/{relationship}
	ObjectDescriptorsRelatedToSigmaAnalysis(id, relationship, cursor string, limit int32)
	// SigmaRuleObject Get a Sigma rule object
	SigmaRuleObject(id string)
	// YARARuleset Yara Ruleset used in crowdsourced YARA results.
	YARARuleset(id string)
}

type files struct {
	path   string
	apiKey string
	log    *log.Logger
	client *http.Client
}

func New(apiKey string, log *log.Logger, client *http.Client) Files {
	return files{
		path:   "api/v3/files",
		apiKey: apiKey,
		log:    log,
		client: client,
	}
}

func (f files) Upload(file os.File) {
	f.upload(file.Name(), &file)
}

func (f files) UploadFromPath(file string) {
	openFile, err := os.Open(file)
	if err != nil {
		return
	}

	f.upload(openFile.Name(), openFile)
}

func (f files) UploadFromStream(fileName string, file io.Reader) {
	f.upload(fileName, file)
}

func (f files) upload(fileName string, file io.Reader) {
	buffer := bytes.Buffer{}

	me := multipart.NewWriter(&buffer)
	me.SetBoundary("011000010111000001101001")

	mw, _ := me.CreateFormFile("file", fileName)

	wl, _ := io.Copy(mw, file)

	me.Close()

	url := "https://www.virustotal.com/" + f.path
	if wl > 32*1024*1024 {
		result, _ := f.UploadURLForLargerThan32MB()
		url = gojsonq.New().FromString(result).Find("data").(string)
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, url, &buffer)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)
	headers.Add("user-agent", "vighnesha")
	headers.Set("Content-Type", me.FormDataContentType())

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) UploadURLForLargerThan32MB() (string, error) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/upload_url", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	//fmt.Println(string(body))
	return string(body), err
}

func (f files) RetrieveInformation(id string) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id, nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
	//return string(body), err
}

func (f files) ReAnalyse(id string) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://www.virustotal.com/"+f.path+"/"+id+"/analyse", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) Comments(id, cursor string, limit int32) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id+"/comments", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) Comment(id, data string) {

	type Comment struct {
		Data struct {
			Type       string `json:"type"`
			Attributes struct {
				Text string `json:"text"`
			} `json:"attributes"`
		} `json:"data"`
	}

	comment := Comment{}
	comment.Data.Type = "comment"

	comment.Data.Attributes.Text = data

	commentPayload, err := json.Marshal(comment)
	if err != nil {
		return
	}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://www.virustotal.com/"+f.path+"/"+id+"/comments", strings.NewReader(string(commentPayload)))
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) Votes(id, cursor string, limit int32) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id+"/votes", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) Vote(id, data string) {

	type Vote struct {
		Data struct {
			Type       string `json:"type"`
			Attributes struct {
				Verdict string `json:"verdict"`
			} `json:"attributes"`
		} `json:"data"`
	}

	vote := Vote{}
	vote.Data.Type = "vote"

	vote.Data.Attributes.Verdict = data

	votePayload, err := json.Marshal(vote)
	if err != nil {
		return
	}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://www.virustotal.com/"+f.path+"/"+id+"/votes", strings.NewReader(string(votePayload)))
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) ObjectsRelated(id, cursor, relationship string, limit int32) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id+"/relationship", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) ObjectDescriptorsRelated(id, relationship, limit, cursor string) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id+"/relationships/relationship", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) BehaviourSummary(id string) {

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.virustotal.com/"+f.path+"/"+id+"/behaviour_summary", nil)
	headers := req.Header
	headers.Set("Accept", "application/json")
	headers.Add("x-apikey", f.apiKey)

	res, err := f.client.Do(req)
	if err != nil {
		f.log.Println("error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func (f files) BehaviourReportFromSandbox(sandboxId string) {
	//TODO implement me
	panic("implement me")
}

func (f files) ObjectsRelatedToBehaviourReport(sandboxId, relationship, cursor string, limit int32) {
	//TODO implement me
	panic("implement me")
}

func (f files) ObjectDescriptorsRelatedToBehaviourReport(sandboxId, relationship, cursor string, limit int32) {
	//TODO implement me
	panic("implement me")
}

func (f files) DetailedHTMLBehaviourReport(sandboxId string) {
	//TODO implement me
	panic("implement me")
}

func (f files) SigmaAnalyses(id string) {
	//TODO implement me
	panic("implement me")
}

func (f files) ObjectsRelatedToSigmaAnalysis(id, relationship, cursor string, limit int32) {
	//TODO implement me
	panic("implement me")
}

func (f files) ObjectDescriptorsRelatedToSigmaAnalysis(id, relationship, cursor string, limit int32) {
	//TODO implement me
	panic("implement me")
}

func (f files) SigmaRuleObject(id string) {
	//TODO implement me
	panic("implement me")
}

func (f files) YARARuleset(id string) {
	//TODO implement me
	panic("implement me")
}
