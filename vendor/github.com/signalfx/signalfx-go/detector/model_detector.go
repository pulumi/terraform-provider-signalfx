/*
 * Detectors API
 *
 * **Detectors** define rules for identifying conditions of interest to the customer, and the notifications to send when the conditions occur or stop occurring.
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package detector

type Detector struct {
	AuthorizedWriters *AuthorizedWriters `json:"authorizedWriters,omitempty"`
	// The time the detector was created in milliseconds (UTC); this corresponds to the receipt of the first request to send an invitation to this email address. This value is always set by the system.
	Created int64 `json:"created,omitempty"`
	// User ID of the initial creator
	Creator string `json:"creator,omitempty"`
	// User-defined JSON object containing metadata
	CustomProperties *interface{} `json:"customProperties,omitempty"`
	// Description of the detector. It appears in the Detector window displayed from the web UI Actions menu
	Description string `json:"description,omitempty"`
	// System-defined identifier for the detector
	Id              string          `json:"id,omitempty"`
	LabelResolution LabelResolution `json:"labelResolution,omitempty"`
	// The last time the detector was updated, in milliseconds (UTC) relative to the Unix epoch.
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// The user ID of the last person who updated the object. If the last update was by the system, the value is \"AAAAAAAAAA\" This value is always set by the system.
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// Detector modification state. If `true`, the detector can't be modified in anyway; otherwise, anyone can edit it.
	Locked bool `json:"locked,omitempty"`
	// The number of milliseconds to wait for late datapoints before rejecting them for inclusion in the detector analysis. The default is to detect and apply a sensible value automatically (this option can also be explicitly chosen by setting the property to 0).
	MaxDelay *int32 `json:"maxDelay,omitempty"`
	// The displayed name of the detector in the dashboard
	Name string `json:"name,omitempty"`
	// If true one, or more statements in the detector matched too many time series and the matched set was forcefully limited. In this case, the detector is most likely looking at incomplete data or an incomplete aggregation. If this flag is set to true, you can use a series of partition_filter functions to split the data set into manageable pieces then use the union function to rejoin the results in a subsequent computation. Note that the union function still observes the time series limit; some type of aggregation on the partial streams must limit the data set prior to recombining the streams for this approach to work.
	OverMTSLimit bool `json:"overMTSLimit,omitempty"`
	// The package spec. Should default to "" so is not `omitempty`
	PackageSpecification string `json:"packageSpecifications"`
	// The SignalFlow program that populates the detector. The program must include one or more detect functions and each detect function must be modified by a publish stream method with a label that's unique across the program. If you wish to support custom notification messages that include input data you must use variables to assign the detect conditions . If more than one line of SignalFlow is included, each line should be separated by either semicolons (\";\") or new line characters (\"\\n\"). See the [Detectors Overview](https://developers.signalfx.com/v2/reference.html#detectors-overview) for more information.
	ProgramText string  `json:"programText,omitempty"`
	Rules       []*Rule `json:"rules,omitempty"`
	// A an array of keywords that filters detectors by one of their fields. Use tags to indicate the state of a detector or its data source (for example, you can label a detector with a \"prod\" tag to indicate that it monitors a production environment).
	Tags []string `json:"tags,omitempty"`
	// IDs of teams associated with this detector. Teams associated with a detector can see the detector and its active alerts on the team's landing page in the web UI. The list of teams associated with a detector is independent of notification settings.  Teams specified in this field don\\'\\'t automatically get notified of new alerts, and teams that choose to get alerts do not have to display the detector on their team landing page in the web application.
	Teams []string `json:"teams,omitempty"`
	// Options that control the appearance of a detector in the SignalFx web UI.
	VisualizationOptions *Visualization `json:"visualizationOptions,omitempty"`
}
