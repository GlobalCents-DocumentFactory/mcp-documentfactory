// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pipeline, err := UnmarshalPipeline(bytes)
//    bytes, err = pipeline.Marshal()

package model

import (
	"encoding/json"
	"time"
)

func UnmarshalPipeline(data []byte) (Pipeline, error) {
	var r Pipeline
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Pipeline) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Pipeline struct {
	Customdata                        map[string]interface{}      `json:"customdata,omitempty"`
	Defaultdocumentsavestrategy       *DocumentSaveStrategy       `json:"defaultdocumentsavestrategy,omitempty"`
	Defaultstoresettingscontentserver *StoreSettingsContentServer `json:"defaultstoresettingscontentserver,omitempty"`
	Defaultstoresettingsnats          *StoreSettingsNATS          `json:"defaultstoresettingsnats,omitempty"`
	Defaultstoresettingss3            *StoreSettingsS3            `json:"defaultstoresettingss3,omitempty"`
	Defaultstoresettingssharepoint    *StoreSettingsSharepoint    `json:"defaultstoresettingssharepoint,omitempty"`
	Dum                               *Dum                        `json:"dum,omitempty"`
	ID                                *string                     `json:"id,omitempty"`
	Inputs                            *Artifact                   `json:"inputs,omitempty"`
	Licencekey                        *string                     `json:"licencekey,omitempty"`
	Merge                             *bool                       `json:"merge,omitempty"`
	Mergesettings                     *SettingsMerge              `json:"mergesettings,omitempty"`
	Metadata                          *string                     `json:"metadata,omitempty"`
	Output                            *Artifact                   `json:"output,omitempty"`
	Pipelinesettings                  *SettingsPipeline           `json:"pipelinesettings,omitempty"`
	Postprocess                       []*DocBuilder               `json:"postprocess,omitempty"`
	Processors                        []*DocBuilder               `json:"processors,omitempty"`
	Signatures                        []*Signature                `json:"signatures,omitempty"`
	Workflowtype                      *Workflowtype               `json:"workflowtype,omitempty"`
}

type DocumentSaveStrategy struct {
	Keepfolderstructure      *bool   `json:"keepfolderstructure,omitempty"`
	Namingpattern            *string `json:"namingpattern,omitempty"`
	Overwrite                *bool   `json:"overwrite,omitempty"`
	Replaceoriginalextension *bool   `json:"replaceoriginalextension,omitempty"`
	Samefolderasoriginal     *bool   `json:"samefolderasoriginal,omitempty"`
	Serverside               *bool   `json:"serverside,omitempty"`
}

type StoreSettingsContentServer struct {
	Convertfilenamestocstags *bool       `json:"convertfilenamestocstags,omitempty"`
	Httptimeout              *int64      `json:"httptimeout,omitempty"`
	Password                 *string     `json:"password,omitempty"`
	Sessiontimeout           *int64      `json:"sessiontimeout,omitempty"`
	Tlsversion               *Tlsversion `json:"tlsversion,omitempty"`
	URL                      *string     `json:"url,omitempty"`
	Userlogin                *string     `json:"userlogin,omitempty"`
}

type StoreSettingsNATS struct {
	Inmemory   *bool   `json:"inmemory,omitempty"`
	Privatekey *string `json:"privatekey,omitempty"`
	Publickey  *string `json:"publickey,omitempty"`
	Servername *string `json:"servername,omitempty"`
	URL        *string `json:"url,omitempty"`
}

type StoreSettingsS3 struct {
	Accesskey *string `json:"accesskey,omitempty"`
	Region    *string `json:"region,omitempty"`
	Secretkey *string `json:"secretkey,omitempty"`
	URL       *string `json:"url,omitempty"`
	Usessl    *bool   `json:"usessl,omitempty"`
}

type StoreSettingsSharepoint struct {
	Chunksize    *int64  `json:"chunksize,omitempty"`
	Clientid     *string `json:"clientid,omitempty"`
	Clientsecret *string `json:"clientsecret,omitempty"`
	Siteurl      *string `json:"siteurl,omitempty"`
}

type Dum struct {
	Pipelinestatus *PipelineStatus `json:"pipelinestatus,omitempty"`
	Taskdef        *TaskDef        `json:"taskdef,omitempty"`
	Wfrun          *WorkflowRun    `json:"wfrun,omitempty"`
}

type PipelineStatus struct {
	Customdata map[string]interface{} `json:"customdata,omitempty"`
	ID         *string                `json:"id,omitempty"`
	Output     *Artifact              `json:"output,omitempty"`
	Runid      *string                `json:"runid,omitempty"`
	Statusinfo *StatusInfo            `json:"statusinfo,omitempty"`
}

type Artifact struct {
	Asset      *Asset      `json:"asset,omitempty"`
	Children   []*Artifact `json:"children,omitempty"`
	Clientinfo *ClientInfo `json:"clientinfo,omitempty"`
	ID         *string     `json:"id,omitempty"`
	Level      *int64      `json:"level,omitempty"`
	Metadata   *string     `json:"metadata,omitempty"`
	Ordering   *int64      `json:"ordering,omitempty"`
	Pplid      *string     `json:"pplid,omitempty"`
	Split      *bool       `json:"split,omitempty"`
	Statusinfo *StatusInfo `json:"statusinfo,omitempty"`
}

type Asset struct {
	Artifactcs         *ArtifactCS         `json:"artifactcs,omitempty"`
	Artifactfs         *ArtifactFS         `json:"artifactfs,omitempty"`
	Artifactnats       *ArtifactNATS       `json:"artifactnats,omitempty"`
	Artifacts3         *ArtifactS3         `json:"artifacts3,omitempty"`
	Artifactsharepoint *ArtifactSharepoint `json:"artifactsharepoint,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	Mimetype           *string             `json:"mimetype,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Shouldbeprocessed  *bool               `json:"shouldbeprocessed,omitempty"`
	Storetype          *Storetype          `json:"storetype,omitempty"`
}

type ArtifactCS struct {
	Dataid                     *int64                      `json:"dataid,omitempty"`
	Mimetype                   *string                     `json:"mimetype,omitempty"`
	Name                       *string                     `json:"name,omitempty"`
	Parentid                   *int64                      `json:"parentid,omitempty"`
	Storesettingscontentserver *StoreSettingsContentServer `json:"storesettingscontentserver,omitempty"`
	Subtype                    *int64                      `json:"subtype,omitempty"`
	Vernum                     *int64                      `json:"vernum,omitempty"`
}

type ArtifactFS struct {
	Fullname *string `json:"fullname,omitempty"`
}

type ArtifactNATS struct {
	Bucket            *string            `json:"bucket,omitempty"`
	Natsid            *string            `json:"natsid,omitempty"`
	Storesettingsnats *StoreSettingsNATS `json:"storesettingsnats,omitempty"`
}

type ArtifactS3 struct {
	Bucket          *string          `json:"bucket,omitempty"`
	S3ID            *string          `json:"s3id,omitempty"`
	Storesettingss3 *StoreSettingsS3 `json:"storesettingss3,omitempty"`
}

type ArtifactSharepoint struct {
	Sharepointid            *string                  `json:"sharepointid,omitempty"`
	Storesettingssharepoint *StoreSettingsSharepoint `json:"storesettingssharepoint,omitempty"`
}

type ClientInfo struct {
	Fullname    *string `json:"fullname,omitempty"`
	ID          *string `json:"id,omitempty"`
	Iscontainer *bool   `json:"iscontainer,omitempty"`
	Metadata    *string `json:"metadata,omitempty"`
	Mimetype    *string `json:"mimetype,omitempty"`
	Parentid    *string `json:"parentid,omitempty"`
	Recursive   *bool   `json:"recursive,omitempty"`
	Shortname   *string `json:"shortname,omitempty"`
}

type StatusInfo struct {
	Datasize     *int64    `json:"datasize,omitempty"`
	Errordetails []*string `json:"errordetails,omitempty"`
	Status       *Status   `json:"status,omitempty"`
	Statusmsg    *string   `json:"statusmsg,omitempty"`
	Timing       *int64    `json:"timing,omitempty"`
}

type TaskDef struct {
	Builder      *DocBuilder   `json:"builder,omitempty"`
	ID           *string       `json:"id,omitempty"`
	Input        *Artifact     `json:"input,omitempty"`
	Output       *Artifact     `json:"output,omitempty"`
	Pplid        *string       `json:"pplid,omitempty"`
	Runid        *string       `json:"runid,omitempty"`
	Statusinfo   *StatusInfo   `json:"statusinfo,omitempty"`
	Tasktype     *Tasktype     `json:"tasktype,omitempty"`
	Workflowtype *Workflowtype `json:"workflowtype,omitempty"`
}

type DocBuilder struct {
	Actioncompress             *bool                       `json:"actioncompress,omitempty"`
	Actionconvert              *bool                       `json:"actionconvert,omitempty"`
	Actiondeletewatermark      *bool                       `json:"actiondeletewatermark,omitempty"`
	Actionflattensignatures    *bool                       `json:"actionflattensignatures,omitempty"`
	Actionmapproperties        *bool                       `json:"actionmapproperties,omitempty"`
	Actionmerge                *bool                       `json:"actionmerge,omitempty"`
	Actionmergeheaderfooter    *bool                       `json:"actionmergeheaderfooter,omitempty"`
	Actionmergemeta            *bool                       `json:"actionmergemeta,omitempty"`
	Actionocr                  *bool                       `json:"actionocr,omitempty"`
	Actionsecurepdf            *bool                       `json:"actionsecurepdf,omitempty"`
	Actionsign                 *bool                       `json:"actionsign,omitempty"`
	Actionsplit                *bool                       `json:"actionsplit,omitempty"`
	Actionwatermark            *bool                       `json:"actionwatermark,omitempty"`
	Extensionsfilter           []*string                   `json:"extensionsfilter,omitempty"`
	Settingscad                *SettingsCAD                `json:"settingscad,omitempty"`
	Settingsconvert            *SettingsConvert            `json:"settingsconvert,omitempty"`
	Settingsdeletewatermarks   []*string                   `json:"settingsdeletewatermarks,omitempty"`
	Settingsemaildocument      *SettingsEMailDocument      `json:"settingsemaildocument,omitempty"`
	Settingshtml               *SettingsHTML               `json:"settingshtml,omitempty"`
	Settingsmergeheaderfooters *SettingsMergeHeaderFooters `json:"settingsmergeheaderfooters,omitempty"`
	Settingsmergemeta          *SettingsMergeMeta          `json:"settingsmergemeta,omitempty"`
	Settingsocr                *SettingsOCR                `json:"settingsocr,omitempty"`
	Settingspagesetup          *SettingsPageSetup          `json:"settingspagesetup,omitempty"`
	Settingspdf                *SettingsPDF                `json:"settingspdf,omitempty"`
	Settingspropmappings       *SettingsPropMappings       `json:"settingspropmappings,omitempty"`
	Settingssignature          *Signature                  `json:"settingssignature,omitempty"`
	Settingswatermarks         []*Watermark                `json:"settingswatermarks,omitempty"`
	Settingsword               *SettingsWord               `json:"settingsword,omitempty"`
}

type SettingsCAD struct {
	Cadbgcolor         *string  `json:"cadbgcolor,omitempty"`
	Cadblackwhite      *bool    `json:"cadblackwhite,omitempty"`
	Cadforcelineweight *float64 `json:"cadforcelineweight,omitempty"`
}

type SettingsConvert struct {
	Saveformat *Saveformat `json:"saveformat,omitempty"`
}

type SettingsEMailDocument struct {
	Attachmentseparator                *Artifact `json:"attachmentseparator,omitempty"`
	Excludedemailattachmentsextensions []*string `json:"excludedemailattachmentsextensions,omitempty"`
	Removeextensionsfrombookmarks      *bool     `json:"removeextensionsfrombookmarks,omitempty"`
}

type SettingsHTML struct {
	Additionalmarginwidth *int64  `json:"additionalmarginwidth,omitempty"`
	Border                *int64  `json:"border,omitempty"`
	Extprefix             *string `json:"extprefix,omitempty"`
	Saveassinglefile      *bool   `json:"saveassinglefile,omitempty"`
	Splitpages            *bool   `json:"splitpages,omitempty"`
	Zipoutput             *bool   `json:"zipoutput,omitempty"`
}

type SettingsMergeHeaderFooters struct {
	Appendheaders   *bool     `json:"appendheaders,omitempty"`
	Appenfooters    *bool     `json:"appenfooters,omitempty"`
	Footersdocument *Artifact `json:"footersdocument,omitempty"`
	Headersdocument *Artifact `json:"headersdocument,omitempty"`
}

type SettingsMergeMeta struct {
	Mergenullorempty     *bool `json:"mergenullorempty,omitempty"`
	Removeemptylines     *bool `json:"removeemptylines,omitempty"`
	Usestrictreplacement *bool `json:"usestrictreplacement,omitempty"`
	Wordtrackchanges     *bool `json:"wordtrackchanges,omitempty"`
}

type SettingsOCR struct {
	Autorotateimages *bool      `json:"autorotateimages,omitempty"`
	Checknotextonly  *bool      `json:"checknotextonly,omitempty"`
	Downscalefactor  *int64     `json:"downscalefactor,omitempty"`
	Ocrdevice        *Ocrdevice `json:"ocrdevice,omitempty"`
	Resolution       *int64     `json:"resolution,omitempty"`
}

type SettingsPageSetup struct {
	Forcepagesetup *bool        `json:"forcepagesetup,omitempty"`
	Height         *float64     `json:"height,omitempty"`
	Marginbottom   *float64     `json:"marginbottom,omitempty"`
	Marginleft     *float64     `json:"marginleft,omitempty"`
	Marginright    *float64     `json:"marginright,omitempty"`
	Margintop      *float64     `json:"margintop,omitempty"`
	Orientation    *Orientation `json:"orientation,omitempty"`
	Papersize      *Papersize   `json:"papersize,omitempty"`
	Width          *float64     `json:"width,omitempty"`
}

type SettingsPDF struct {
	Baseurl                 *string                 `json:"baseurl,omitempty"`
	Bookmarkmergeddocs      *bool                   `json:"bookmarkmergeddocs,omitempty"`
	Bookmarksoutlinelevel   *int64                  `json:"bookmarksoutlinelevel,omitempty"`
	Cbcustomtext            *string                 `json:"cbcustomtext,omitempty"`
	Createbookmarksfromword *bool                   `json:"createbookmarksfromword,omitempty"`
	Embedfullfonts          *bool                   `json:"embedfullfonts,omitempty"`
	Expandedoutlinelevels   *int64                  `json:"expandedoutlinelevels,omitempty"`
	Flattenpdf              *bool                   `json:"flattenpdf,omitempty"`
	Greylevels              *bool                   `json:"greylevels,omitempty"`
	Headingsoutlinelevels   *bool                   `json:"headingsoutlinelevels,omitempty"`
	Inheritzoom             *bool                   `json:"inheritzoom,omitempty"`
	Initialmagnification    *Initialmagnification   `json:"initialmagnification,omitempty"`
	Initialpagelayout       *Initialpagelayout      `json:"initialpagelayout,omitempty"`
	Openbookmarks           *bool                   `json:"openbookmarks,omitempty"`
	Optimizeforweb          *bool                   `json:"optimizeforweb,omitempty"`
	Passwords               []*string               `json:"passwords,omitempty"`
	Pdfformat               *Pdfformat              `json:"pdfformat,omitempty"`
	Settingspdfcompression  *SettingsPDFCompression `json:"settingspdfcompression,omitempty"`
	Settingspdfsecurity     *SettingsPDFSecurity    `json:"settingspdfsecurity,omitempty"`
}

type SettingsPDFCompression struct {
	Compress    *bool  `json:"compress,omitempty"`
	Downsample  *int64 `json:"downsample,omitempty"`
	Jpegquality *int64 `json:"jpegquality,omitempty"`
}

type SettingsPDFSecurity struct {
	Cryptoalgorithm   *Cryptoalgorithm   `json:"cryptoalgorithm,omitempty"`
	Documentprivilege *DocumentPrivilege `json:"documentprivilege,omitempty"`
	Ownerpassword     *string            `json:"ownerpassword,omitempty"`
	Usepdf20          *bool              `json:"usepdf20,omitempty"`
	Userpassword      *string            `json:"userpassword,omitempty"`
}

type DocumentPrivilege struct {
	Allowassembly          *bool             `json:"allowassembly,omitempty"`
	Allowcopy              *bool             `json:"allowcopy,omitempty"`
	Allowdegradedprinting  *bool             `json:"allowdegradedprinting,omitempty"`
	Allowfillin            *bool             `json:"allowfillin,omitempty"`
	Allowmodifyannotations *bool             `json:"allowmodifyannotations,omitempty"`
	Allowmodifycontents    *bool             `json:"allowmodifycontents,omitempty"`
	Allowprint             *bool             `json:"allowprint,omitempty"`
	Allowscreenreaders     *bool             `json:"allowscreenreaders,omitempty"`
	Changeallowlevel       *Changeallowlevel `json:"changeallowlevel,omitempty"`
	Copyallowlevel         *Copyallowlevel   `json:"copyallowlevel,omitempty"`
	Printallowlevel        *Printallowlevel  `json:"printallowlevel,omitempty"`
}

type SettingsPropMappings struct {
	Convertutf8   *bool   `json:"convertutf8,omitempty"`
	Mapfields     *string `json:"mapfields,omitempty"`
	Mapproperties *string `json:"mapproperties,omitempty"`
}

type Signature struct {
	Height            *int64             `json:"height,omitempty"`
	Numsigns          *int64             `json:"numsigns,omitempty"`
	Pageno            *int64             `json:"pageno,omitempty"`
	Pdfownerpassword  *string            `json:"pdfownerpassword,omitempty"`
	Settingssignature *SettingsSignature `json:"settingssignature,omitempty"`
	SignComment       *string            `json:"sign_comment,omitempty"`
	SignDate          *time.Time         `json:"sign_date,omitempty"`
	Signinguser       *SignUser          `json:"signinguser,omitempty"`
	Signinguserbehalf *SignUser          `json:"signinguserbehalf,omitempty"`
	Wfholders         []*Signature       `json:"wfholders,omitempty"`
	Width             *int64             `json:"width,omitempty"`
	Wsid              *string            `json:"wsid,omitempty"`
	X                 *int64             `json:"x,omitempty"`
	Y                 *int64             `json:"y,omitempty"`
}

type SettingsSignature struct {
	Appendpageifneeded *bool               `json:"appendpageifneeded,omitempty"`
	Certificationlevel *Certificationlevel `json:"certificationlevel,omitempty"`
	Font               *string             `json:"font,omitempty"`
	Fontrtol           *bool               `json:"fontrtol,omitempty"`
	Fontsize           *int64              `json:"fontsize,omitempty"`
	Height             *int64              `json:"height,omitempty"`
	Horizontalalign    *Horizontalalign    `json:"horizontalalign,omitempty"`
	Pagemarginbottom   *int64              `json:"pagemarginbottom,omitempty"`
	Pagemarginleft     *int64              `json:"pagemarginleft,omitempty"`
	Pagemarginright    *int64              `json:"pagemarginright,omitempty"`
	Pagemargintop      *int64              `json:"pagemargintop,omitempty"`
	Showcontact        *bool               `json:"showcontact,omitempty"`
	Showdept           *bool               `json:"showdept,omitempty"`
	Showfunction       *bool               `json:"showfunction,omitempty"`
	Showlocation       *bool               `json:"showlocation,omitempty"`
	Showreason         *bool               `json:"showreason,omitempty"`
	Showsigndate       *bool               `json:"showsigndate,omitempty"`
	Showtitle          *bool               `json:"showtitle,omitempty"`
	Signatureimage     *bool               `json:"signatureimage,omitempty"`
	Signaturemarginh   *int64              `json:"signaturemarginh,omitempty"`
	Signaturemarginv   *int64              `json:"signaturemarginv,omitempty"`
	Signatureprovider  *string             `json:"signatureprovider,omitempty"`
	Signdateformat     *string             `json:"signdateformat,omitempty"`
	Signlastpage       *bool               `json:"signlastpage,omitempty"`
	Signpagenumber     *int64              `json:"signpagenumber,omitempty"`
	Textjustify        *Horizontalalign    `json:"textjustify,omitempty"`
	Textpositionv      *Textpositionv      `json:"textpositionv,omitempty"`
	Twofactorsprovider *string             `json:"twofactorsprovider,omitempty"`
	Verticalalign      *Textpositionv      `json:"verticalalign,omitempty"`
	Width              *int64              `json:"width,omitempty"`
	X                  *int64              `json:"x,omitempty"`
	Y                  *int64              `json:"y,omitempty"`
}

type SignUser struct {
	Department   *string `json:"department,omitempty"`
	Email        *string `json:"email,omitempty"`
	Firstname    *string `json:"firstname,omitempty"`
	Initials     *string `json:"initials,omitempty"`
	Lastname     *string `json:"lastname,omitempty"`
	Location     *string `json:"location,omitempty"`
	Otp          *string `json:"otp,omitempty"`
	SignPassword *string `json:"sign_password,omitempty"`
	Timezone     *int64  `json:"timezone,omitempty"`
	Title        *string `json:"title,omitempty"`
}

type Watermark struct {
	Barcodetype               *Barcodetype     `json:"barcodetype,omitempty"`
	Barcodezoom               *float64         `json:"barcodezoom,omitempty"`
	Border                    *int64           `json:"border,omitempty"`
	Borderadius               *int64           `json:"borderadius,omitempty"`
	Color                     *string          `json:"color,omitempty"`
	Displaybarcodetext        *bool            `json:"displaybarcodetext,omitempty"`
	Fontbold                  *bool            `json:"fontbold,omitempty"`
	Fontfamily                *Fontfamily      `json:"fontfamily,omitempty"`
	Fontitalic                *bool            `json:"fontitalic,omitempty"`
	Fontsize                  *int64           `json:"fontsize,omitempty"`
	Height                    *int64           `json:"height,omitempty"`
	Horizontalalign           *Horizontalalign `json:"horizontalalign,omitempty"`
	Image                     *Artifact        `json:"image,omitempty"`
	Layername                 *string          `json:"layername,omitempty"`
	Opacity                   *float64         `json:"opacity,omitempty"`
	Rotation                  *int64           `json:"rotation,omitempty"`
	Text                      *string          `json:"text,omitempty"`
	Texthorizontalalign       *Horizontalalign `json:"texthorizontalalign,omitempty"`
	Textmargin                *int64           `json:"textmargin,omitempty"`
	Textverticalposition      *Textpositionv   `json:"textverticalposition,omitempty"`
	Transparentbackground     *bool            `json:"transparentbackground,omitempty"`
	Underlayout               *bool            `json:"underlayout,omitempty"`
	Verticalalign             *Textpositionv   `json:"verticalalign,omitempty"`
	Watermarkfrompage         *string          `json:"watermarkfrompage,omitempty"`
	Watermarkmarginhorizontal *int64           `json:"watermarkmarginhorizontal,omitempty"`
	Watermarkmarginvertical   *int64           `json:"watermarkmarginvertical,omitempty"`
	Watermarkon               *Watermarkon     `json:"watermarkon,omitempty"`
	Watermarktopage           *string          `json:"watermarktopage,omitempty"`
	Watermarktype             *Watermarktype   `json:"watermarktype,omitempty"`
	Width                     *int64           `json:"width,omitempty"`
}

type SettingsWord struct {
	Acceptrevisions      *bool   `json:"acceptrevisions,omitempty"`
	Deletecomments       *bool   `json:"deletecomments,omitempty"`
	Disabletrackchanges  *bool   `json:"disabletrackchanges,omitempty"`
	Inlineimagewithtext  *bool   `json:"inlineimagewithtext,omitempty"`
	Lockedfields         *string `json:"lockedfields,omitempty"`
	Lockselectedfields   *bool   `json:"lockselectedfields,omitempty"`
	Removemacros         *bool   `json:"removemacros,omitempty"`
	Selectedfields       *string `json:"selectedfields,omitempty"`
	Splitsections        *bool   `json:"splitsections,omitempty"`
	Trackchanges         *bool   `json:"trackchanges,omitempty"`
	Updateallfields      *bool   `json:"updateallfields,omitempty"`
	Updatepagelayout     *bool   `json:"updatepagelayout,omitempty"`
	Updateselectedfields *bool   `json:"updateselectedfields,omitempty"`
	Updatesummaries      *bool   `json:"updatesummaries,omitempty"`
}

type WorkflowRun struct {
	ID    *string `json:"id,omitempty"`
	Pplid *string `json:"pplid,omitempty"`
}

type SettingsMerge struct {
	Intermediatepage  *Artifact          `json:"intermediatepage,omitempty"`
	Namingpattern     *string            `json:"namingpattern,omitempty"`
	Saveformat        *Saveformat        `json:"saveformat,omitempty"`
	Settingswordmerge *SettingsWordMerge `json:"settingswordmerge,omitempty"`
}

type SettingsWordMerge struct {
	Columnset                       *Columnset    `json:"columnset,omitempty"`
	Endpagesectionstart             *Sectionstart `json:"endpagesectionstart,omitempty"`
	Inheritpagesetup                *bool         `json:"inheritpagesetup,omitempty"`
	Linebetweencols                 *bool         `json:"linebetweencols,omitempty"`
	Mergetrackchanges               *bool         `json:"mergetrackchanges,omitempty"`
	Mergetrackchangescharacterlevel *bool         `json:"mergetrackchangescharacterlevel,omitempty"`
	Sectionstart                    *Sectionstart `json:"sectionstart,omitempty"`
}

type SettingsPipeline struct {
	Continueonerror             *bool `json:"continueonerror,omitempty"`
	Deleteoriginalonsuccess     *bool `json:"deleteoriginalonsuccess,omitempty"`
	Ignoreunknownfileextensions *bool `json:"ignoreunknownfileextensions,omitempty"`
}

type Tlsversion string

const (
	Ssl3          Tlsversion = "Ssl3"
	SystemDefault Tlsversion = "SystemDefault"
	TLS           Tlsversion = "Tls"
	Tls11         Tlsversion = "Tls11"
	Tls12         Tlsversion = "Tls12"
	Tls13         Tlsversion = "Tls13"
)

type Storetype string

const (
	Contentserver Storetype = "contentserver"
	FS            Storetype = "fs"
	FTP           Storetype = "ftp"
	Nats          Storetype = "nats"
	S3            Storetype = "s3"
	Sharepoint    Storetype = "sharepoint"
	StoretypeNone Storetype = "none"
	URL           Storetype = "url"
)

type Status string

const (
	Canceled   Status = "Canceled"
	Canceling  Status = "Canceling"
	Completed  Status = "Completed"
	Errored    Status = "Errored"
	Pending    Status = "Pending"
	Processing Status = "Processing"
	Unchanged  Status = "Unchanged"
)

type Saveformat string

const (
	Doc   Saveformat = "doc"
	Docx  Saveformat = "docx"
	GIF   Saveformat = "gif"
	HTML  Saveformat = "html"
	JPEG  Saveformat = "jpeg"
	Mhtml Saveformat = "mhtml"
	PDF   Saveformat = "pdf"
	PNG   Saveformat = "png"
	Ppt   Saveformat = "ppt"
	Pptx  Saveformat = "pptx"
	Text  Saveformat = "text"
	Tiff  Saveformat = "tiff"
	Xls   Saveformat = "xls"
	Xlsx  Saveformat = "xlsx"
)

type Ocrdevice string

const (
	Pdfocr24 Ocrdevice = "pdfocr24"
	Pdfocr32 Ocrdevice = "pdfocr32"
	Pdfocr8  Ocrdevice = "pdfocr8"
)

type Orientation string

const (
	Landscape Orientation = "Landscape"
	Portrait  Orientation = "Portrait"
)

type Papersize string

const (
	A0        Papersize = "a0"
	A1        Papersize = "a1"
	A2        Papersize = "a2"
	A3        Papersize = "a3"
	A4        Papersize = "a4"
	A5        Papersize = "a5"
	B4        Papersize = "b4"
	B5        Papersize = "b5"
	Custom    Papersize = "custom"
	Envelope  Papersize = "envelope"
	Executive Papersize = "executive"
	Folio     Papersize = "folio"
	Ledger    Papersize = "ledger"
	Legal     Papersize = "legal"
	Letter    Papersize = "letter"
	P10X14    Papersize = "p10x14"
	Quarto    Papersize = "quarto"
	Statement Papersize = "statement"
	Tabloid   Papersize = "tabloid"
)

type Initialmagnification string

const (
	FitHeight                   Initialmagnification = "FitHeight"
	FitPage                     Initialmagnification = "FitPage"
	FitVisible                  Initialmagnification = "FitVisible"
	FitWidth                    Initialmagnification = "FitWidth"
	InitialmagnificationDefault Initialmagnification = "Default"
	Pct100                      Initialmagnification = "pct100"
	Pct125                      Initialmagnification = "pct125"
	Pct150                      Initialmagnification = "pct150"
	Pct1600                     Initialmagnification = "pct1600"
	Pct200                      Initialmagnification = "pct200"
	Pct25                       Initialmagnification = "pct25"
	Pct3200                     Initialmagnification = "pct3200"
	Pct400                      Initialmagnification = "pct400"
	Pct50                       Initialmagnification = "pct50"
	Pct6400                     Initialmagnification = "pct6400"
	Pct75                       Initialmagnification = "pct75"
	Pct800                      Initialmagnification = "pct800"
)

type Initialpagelayout string

const (
	InitialpagelayoutDefault Initialpagelayout = "Default"
	OneColumn                Initialpagelayout = "OneColumn"
	SinglePage               Initialpagelayout = "SinglePage"
	TwoColumnLeft            Initialpagelayout = "TwoColumnLeft"
	TwoColumnRight           Initialpagelayout = "TwoColumnRight"
	TwoPageLeft              Initialpagelayout = "TwoPageLeft"
	TwoPageRight             Initialpagelayout = "TwoPageRight"
)

type Pdfformat string

const (
	PDFA1A     Pdfformat = "PDF_A_1A"
	PDFA1B     Pdfformat = "PDF_A_1B"
	PDFA2A     Pdfformat = "PDF_A_2A"
	PDFA2B     Pdfformat = "PDF_A_2B"
	PDFA2U     Pdfformat = "PDF_A_2U"
	PDFA3A     Pdfformat = "PDF_A_3A"
	PDFA3B     Pdfformat = "PDF_A_3B"
	PDFA3U     Pdfformat = "PDF_A_3U"
	PDFUa1     Pdfformat = "PDF_UA_1"
	PDFX1A     Pdfformat = "PDF_X_1A"
	PDFX1A2001 Pdfformat = "PDF_X_1A_2001"
	PDFX3      Pdfformat = "PDF_X_3"
	V1_0       Pdfformat = "v_1_0"
	V1_1       Pdfformat = "v_1_1"
	V1_2       Pdfformat = "v_1_2"
	V1_3       Pdfformat = "v_1_3"
	V1_4       Pdfformat = "v_1_4"
	V1_5       Pdfformat = "v_1_5"
	V1_6       Pdfformat = "v_1_6"
	V1_7       Pdfformat = "v_1_7"
	V2_0       Pdfformat = "v_2_0"
	ZUGFeRD    Pdfformat = "ZUGFeRD"
)

type Cryptoalgorithm string

const (
	AESx128 Cryptoalgorithm = "AESx128"
	AESx256 Cryptoalgorithm = "AESx256"
	RC4X128 Cryptoalgorithm = "RC4x128"
	RC4X40  Cryptoalgorithm = "RC4x40"
)

type Changeallowlevel string

const (
	AnyExceptExtract                                            Changeallowlevel = "AnyExceptExtract"
	ChangeallowlevelNone                                        Changeallowlevel = "None"
	CommentingFillingInFormfieldsAndSignExistingSignatureFields Changeallowlevel = "CommentingFillingInFormfieldsAndSignExistingSignatureFields"
	FillingInFormfieldsAndSignExistingSignatureFields           Changeallowlevel = "FillingInFormfieldsAndSignExistingSignatureFields"
	InsertingDeletingRotatingPages                              Changeallowlevel = "InsertingDeletingRotatingPages"
)

type Copyallowlevel string

const (
	CopyallowlevelNone                                           Copyallowlevel = "None"
	EnableCopyingOfTextImagesAndOtherContent                     Copyallowlevel = "EnableCopyingOfTextImagesAndOtherContent"
	EnableTextAccessForScreenReaderDevicesForTheVisuallyImpaired Copyallowlevel = "EnableTextAccessForScreenReaderDevicesForTheVisuallyImpaired"
)

type Printallowlevel string

const (
	HighResolution      Printallowlevel = "HighResolution"
	LowResolution150DPI Printallowlevel = "LowResolution150Dpi"
	PrintallowlevelNone Printallowlevel = "None"
)

type Certificationlevel string

const (
	CertifiedFormFilling               Certificationlevel = "CERTIFIED_FORM_FILLING"
	CertifiedFormFillingAndAnnotations Certificationlevel = "CERTIFIED_FORM_FILLING_AND_ANNOTATIONS"
	CertifiedNoChangesAllowed          Certificationlevel = "CERTIFIED_NO_CHANGES_ALLOWED"
	NotCertified                       Certificationlevel = "NOT_CERTIFIED"
)

type Horizontalalign string

const (
	HorizontalalignCenter Horizontalalign = "center"
	Left                  Horizontalalign = "left"
	Right                 Horizontalalign = "right"
)

type Textpositionv string

const (
	Bottom              Textpositionv = "bottom"
	TextpositionvCenter Textpositionv = "center"
	Top                 Textpositionv = "top"
)

type Barcodetype string

const (
	AustraliaPost                 Barcodetype = "AustraliaPost"
	AustralianPosteParcel         Barcodetype = "AustralianPosteParcel"
	Aztec                         Barcodetype = "Aztec"
	BarcodetypeNone               Barcodetype = "None"
	Codabar                       Barcodetype = "Codabar"
	CodablockF                    Barcodetype = "CodablockF"
	Code11                        Barcodetype = "Code11"
	Code128                       Barcodetype = "Code128"
	Code16K                       Barcodetype = "Code16K"
	Code32                        Barcodetype = "Code32"
	Code39Extended                Barcodetype = "Code39Extended"
	Code39Standard                Barcodetype = "Code39Standard"
	Code93Extended                Barcodetype = "Code93Extended"
	Code93Standard                Barcodetype = "Code93Standard"
	DataLogic2Of5                 Barcodetype = "DataLogic2of5"
	DataMatrix                    Barcodetype = "DataMatrix"
	DatabarExpanded               Barcodetype = "DatabarExpanded"
	DatabarExpandedStacked        Barcodetype = "DatabarExpandedStacked"
	DatabarLimited                Barcodetype = "DatabarLimited"
	DatabarOmniDirectional        Barcodetype = "DatabarOmniDirectional"
	DatabarStacked                Barcodetype = "DatabarStacked"
	DatabarStackedOmniDirectional Barcodetype = "DatabarStackedOmniDirectional"
	DatabarTruncated              Barcodetype = "DatabarTruncated"
	DeutschePostIdentcode         Barcodetype = "DeutschePostIdentcode"
	DeutschePostLeitcode          Barcodetype = "DeutschePostLeitcode"
	DotCode                       Barcodetype = "DotCode"
	DutchKIX                      Barcodetype = "DutchKIX"
	Ean13                         Barcodetype = "EAN13"
	Ean14                         Barcodetype = "EAN14"
	Ean8                          Barcodetype = "EAN8"
	GS1CodablockF                 Barcodetype = "GS1CodablockF"
	GS1Code128                    Barcodetype = "GS1Code128"
	GS1DataMatrix                 Barcodetype = "GS1DataMatrix"
	Gs1Qr                         Barcodetype = "GS1QR"
	IATA2Of5                      Barcodetype = "IATA2of5"
	Interleaved2Of5               Barcodetype = "Interleaved2of5"
	Isbn                          Barcodetype = "ISBN"
	Ismn                          Barcodetype = "ISMN"
	Issn                          Barcodetype = "ISSN"
	ItalianPost25                 Barcodetype = "ItalianPost25"
	Itf14                         Barcodetype = "ITF14"
	Itf6                          Barcodetype = "ITF6"
	MSI                           Barcodetype = "MSI"
	MacroPdf417                   Barcodetype = "MacroPdf417"
	Mailmark                      Barcodetype = "Mailmark"
	Matrix2Of5                    Barcodetype = "Matrix2of5"
	MaxiCode                      Barcodetype = "MaxiCode"
	MicroPdf417                   Barcodetype = "MicroPdf417"
	OneCode                       Barcodetype = "OneCode"
	Opc                           Barcodetype = "OPC"
	PatchCode                     Barcodetype = "PatchCode"
	Pdf417                        Barcodetype = "Pdf417"
	Pharmacode                    Barcodetype = "Pharmacode"
	Planet                        Barcodetype = "Planet"
	Postnet                       Barcodetype = "Postnet"
	Pzn                           Barcodetype = "PZN"
	Qr                            Barcodetype = "QR"
	Rm4Scc                        Barcodetype = "RM4SCC"
	Scc14                         Barcodetype = "SCC14"
	SingaporePost                 Barcodetype = "SingaporePost"
	Sscc18                        Barcodetype = "SSCC18"
	Standard2Of5                  Barcodetype = "Standard2of5"
	SwissPostParcel               Barcodetype = "SwissPostParcel"
	Upca                          Barcodetype = "UPCA"
	UpcaGs1Code128Coupon          Barcodetype = "UpcaGs1Code128Coupon"
	UpcaGs1DatabarCoupon          Barcodetype = "UpcaGs1DatabarCoupon"
	Upce                          Barcodetype = "UPCE"
	Vin                           Barcodetype = "VIN"
)

type Fontfamily string

const (
	Courier      Fontfamily = "COURIER"
	Helvetica    Fontfamily = "HELVETICA"
	Symbol       Fontfamily = "SYMBOL"
	TimesRoman   Fontfamily = "TIMES_ROMAN"
	Undefined    Fontfamily = "UNDEFINED"
	Zapfdingbats Fontfamily = "ZAPFDINGBATS"
)

type Watermarkon string

const (
	Watermarkallpages  Watermarkon = "watermarkallpages"
	Watermarkfirstpage Watermarkon = "watermarkfirstpage"
	Watermarklastpage  Watermarkon = "watermarklastpage"
)

type Watermarktype string

const (
	Image             Watermarktype = "Image"
	Qrcode            Watermarktype = "Qrcode"
	Staticimage       Watermarktype = "Staticimage"
	WatermarktypeText Watermarktype = "Text"
)

type Tasktype string

const (
	Load        Tasktype = "Load"
	Merge       Tasktype = "Merge"
	Postprocess Tasktype = "Postprocess"
	Preprocess  Tasktype = "Preprocess"
	Sign        Tasktype = "Sign"
	TasktypeOcr Tasktype = "Ocr"
	Upload      Tasktype = "Upload"
)

type Workflowtype string

const (
	PDFBook         Workflowtype = "PdfBook"
	Print           Workflowtype = "Print"
	Standard        Workflowtype = "Standard"
	WordBook        Workflowtype = "WordBook"
	WorkflowtypeOcr Workflowtype = "Ocr"
)

type Columnset string

const (
	ColumnsetNone Columnset = "none"
	Twocols       Columnset = "twocols"
)

type Sectionstart string

const (
	Continuous Sectionstart = "Continuous"
	EvenPage   Sectionstart = "EvenPage"
	NewColumn  Sectionstart = "NewColumn"
	NewPage    Sectionstart = "NewPage"
	OddPage    Sectionstart = "OddPage"
)
