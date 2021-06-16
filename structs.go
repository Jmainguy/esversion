package main

type beat struct {
	BeatVersion  string `json:"beatVersion"`
	AgentVersion string `json:"agentVersion"`
}

type aggResponse struct {
	Shards struct {
		Failed     int64 `json:"failed"`
		Skipped    int64 `json:"skipped"`
		Successful int64 `json:"successful"`
		Total      int64 `json:"total"`
	} `json:"_shards"`
	Aggregations struct {
		SixX_hostver struct {
			AfterKey struct {
				Hostname string `json:"hostname"`
				Version  string `json:"version"`
			} `json:"after_key"`
			Buckets []struct {
				DocCount int64 `json:"doc_count"`
				Key      struct {
					Hostname string `json:"hostname"`
					Version  string `json:"version"`
				} `json:"key"`
			} `json:"buckets"`
		} `json:"6x-hostver"`
		SevenX_hostver struct {
			AfterKey struct {
				Hostname string `json:"hostname"`
				Version  string `json:"version"`
			} `json:"after_key"`
			Buckets []struct {
				DocCount int64 `json:"doc_count"`
				Key      struct {
					Hostname string `json:"hostname"`
					Version  string `json:"version"`
				} `json:"key"`
			} `json:"buckets"`
		} `json:"7x-hostver"`
	} `json:"aggregations"`
	Hits struct {
		Hits     []interface{} `json:"hits"`
		MaxScore interface{}   `json:"max_score"`
		Total    struct {
			Relation string `json:"relation"`
			Value    int64  `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool  `json:"timed_out"`
	Took     int64 `json:"took"`
}

type Term struct {
	Terms struct {
		Field string `json:"field,omitempty"`
	} `json:"terms,omitempty"`
}

type Source struct {
	Hostname *Term `json:"hostname,omitempty"`
	Version  *Term `json:"version,omitempty"`
}

type aggQuery struct {
	Aggs struct {
		SixX_hostver struct {
			Composite struct {
				Size    int64    `json:"size"`
				Sources []Source `json:"sources"`
			} `json:"composite"`
		} `json:"6x-hostver"`
		SevenX_hostver struct {
			Composite struct {
				Size    int64    `json:"size"`
				Sources []Source `json:"sources"`
			} `json:"composite"`
		} `json:"7x-hostver"`
	} `json:"aggs"`
	Query struct {
		Bool struct {
			Filter struct {
				Range struct {
					Timestamp struct {
						Gte string `json:"gte"`
						Lte string `json:"lte"`
					} `json:"@timestamp"`
				} `json:"range"`
			} `json:"filter"`
		} `json:"bool"`
	} `json:"query"`
	Size int64 `json:"size"`
}
