package main

type beat struct {
	BeatVersion  string `json:"beatVersion"`
	AgentVersion string `json:"agentVersion"`
}

type Response struct {
	Shards struct {
		Failed     int64 `json:"failed"`
		Skipped    int64 `json:"skipped"`
		Successful int64 `json:"successful"`
		Total      int64 `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Hits []struct {
			ID     string `json:"_id"`
			Index  string `json:"_index"`
			Score  int64  `json:"_score"`
			Source struct {
				Timestamp string `json:"@timestamp"`
				Version   string `json:"@version"`
				Beat      struct {
					Hostname string `json:"hostname"`
					Name     string `json:"name"`
					Version  string `json:"version"`
				} `json:"beat"`
				Agent struct {
					EphemeralID string `json:"ephemeral_id"`
					Hostname    string `json:"hostname"`
					ID          string `json:"id"`
					Name        string `json:"name"`
					Type        string `json:"type"`
					Version     string `json:"version"`
				} `json:"agent"`
				Host struct {
					Name string `json:"name"`
				} `json:"host"`
				Metricset struct {
					Module string `json:"module"`
					Name   string `json:"name"`
					Rtt    int64  `json:"rtt"`
				} `json:"metricset"`
				System struct {
					Core struct {
						ID   int64 `json:"id"`
						Idle struct {
							Pct float64 `json:"pct"`
						} `json:"idle"`
						Iowait struct {
							Pct int64 `json:"pct"`
						} `json:"iowait"`
						Irq struct {
							Pct int64 `json:"pct"`
						} `json:"irq"`
						Nice struct {
							Pct int64 `json:"pct"`
						} `json:"nice"`
						Softirq struct {
							Pct int64 `json:"pct"`
						} `json:"softirq"`
						Steal struct {
							Pct int64 `json:"pct"`
						} `json:"steal"`
						System struct {
							Pct float64 `json:"pct"`
						} `json:"system"`
						User struct {
							Pct float64 `json:"pct"`
						} `json:"user"`
					} `json:"core"`
					CPU struct {
						Cores int64 `json:"cores"`
						Idle  struct {
							Pct float64 `json:"pct"`
						} `json:"idle"`
						Iowait struct {
							Pct int64 `json:"pct"`
						} `json:"iowait"`
						Irq struct {
							Pct int64 `json:"pct"`
						} `json:"irq"`
						Nice struct {
							Pct int64 `json:"pct"`
						} `json:"nice"`
						Softirq struct {
							Pct int64 `json:"pct"`
						} `json:"softirq"`
						Steal struct {
							Pct float64 `json:"pct"`
						} `json:"steal"`
						System struct {
							Pct float64 `json:"pct"`
						} `json:"system"`
						Total struct {
							Pct float64 `json:"pct"`
						} `json:"total"`
						User struct {
							Pct float64 `json:"pct"`
						} `json:"user"`
					} `json:"cpu"`
					Diskio struct {
						Io struct {
							Time int64 `json:"time"`
						} `json:"io"`
						Iostat struct {
							Await float64 `json:"await"`
							Busy  float64 `json:"busy"`
							Queue struct {
								AvgSize float64 `json:"avg_size"`
							} `json:"queue"`
							Read struct {
								Await  int64 `json:"await"`
								PerSec struct {
									Bytes float64 `json:"bytes"`
								} `json:"per_sec"`
								Request struct {
									MergesPerSec int64   `json:"merges_per_sec"`
									PerSec       float64 `json:"per_sec"`
								} `json:"request"`
							} `json:"read"`
							Request struct {
								AvgSize float64 `json:"avg_size"`
							} `json:"request"`
							ServiceTime float64 `json:"service_time"`
							Write       struct {
								Await  float64 `json:"await"`
								PerSec struct {
									Bytes float64 `json:"bytes"`
								} `json:"per_sec"`
								Request struct {
									MergesPerSec int64   `json:"merges_per_sec"`
									PerSec       float64 `json:"per_sec"`
								} `json:"request"`
							} `json:"write"`
						} `json:"iostat"`
						Name string `json:"name"`
						Read struct {
							Bytes int64 `json:"bytes"`
							Count int64 `json:"count"`
							Time  int64 `json:"time"`
						} `json:"read"`
						Write struct {
							Bytes int64 `json:"bytes"`
							Count int64 `json:"count"`
							Time  int64 `json:"time"`
						} `json:"write"`
					} `json:"diskio"`
					Filesystem struct {
						Available  int64  `json:"available"`
						DeviceName string `json:"device_name"`
						Files      int64  `json:"files"`
						Free       int64  `json:"free"`
						FreeFiles  int64  `json:"free_files"`
						MountPoint string `json:"mount_point"`
						Total      int64  `json:"total"`
						Type       string `json:"type"`
						Used       struct {
							Bytes int64   `json:"bytes"`
							Pct   float64 `json:"pct"`
						} `json:"used"`
					} `json:"filesystem"`
					Fsstat struct {
						Count      int64 `json:"count"`
						TotalFiles int64 `json:"total_files"`
						TotalSize  struct {
							Free  int64 `json:"free"`
							Total int64 `json:"total"`
							Used  int64 `json:"used"`
						} `json:"total_size"`
					} `json:"fsstat"`
					Load struct {
						One   float64 `json:"1"`
						One5  float64 `json:"15"`
						Five  float64 `json:"5"`
						Cores int64   `json:"cores"`
						Norm  struct {
							One  float64 `json:"1"`
							One5 float64 `json:"15"`
							Five float64 `json:"5"`
						} `json:"norm"`
					} `json:"load"`
					Memory struct {
						Actual struct {
							Free int64 `json:"free"`
							Used struct {
								Bytes int64   `json:"bytes"`
								Pct   float64 `json:"pct"`
							} `json:"used"`
						} `json:"actual"`
						Free      int64 `json:"free"`
						Hugepages struct {
							DefaultSize int64 `json:"default_size"`
							Free        int64 `json:"free"`
							Reserved    int64 `json:"reserved"`
							Surplus     int64 `json:"surplus"`
							Total       int64 `json:"total"`
							Used        struct {
								Bytes int64 `json:"bytes"`
								Pct   int64 `json:"pct"`
							} `json:"used"`
						} `json:"hugepages"`
						Swap struct {
							Free  int64 `json:"free"`
							Total int64 `json:"total"`
							Used  struct {
								Bytes int64 `json:"bytes"`
								Pct   int64 `json:"pct"`
							} `json:"used"`
						} `json:"swap"`
						Total int64 `json:"total"`
						Used  struct {
							Bytes int64   `json:"bytes"`
							Pct   float64 `json:"pct"`
						} `json:"used"`
					} `json:"memory"`
					Network struct {
						In struct {
							Bytes   int64 `json:"bytes"`
							Dropped int64 `json:"dropped"`
							Errors  int64 `json:"errors"`
							Packets int64 `json:"packets"`
						} `json:"in"`
						Name string `json:"name"`
						Out  struct {
							Bytes   int64 `json:"bytes"`
							Dropped int64 `json:"dropped"`
							Errors  int64 `json:"errors"`
							Packets int64 `json:"packets"`
						} `json:"out"`
					} `json:"network"`
					Process struct {
						Cmdline string `json:"cmdline"`
						CPU     struct {
							StartTime string `json:"start_time"`
							Total     struct {
								Norm struct {
									Pct float64 `json:"pct"`
								} `json:"norm"`
								Pct   float64 `json:"pct"`
								Value int64   `json:"value"`
							} `json:"total"`
						} `json:"cpu"`
						Cwd string `json:"cwd"`
						Fd  struct {
							Limit struct {
								Hard int64 `json:"hard"`
								Soft int64 `json:"soft"`
							} `json:"limit"`
							Open int64 `json:"open"`
						} `json:"fd"`
						Memory struct {
							Rss struct {
								Bytes int64   `json:"bytes"`
								Pct   float64 `json:"pct"`
							} `json:"rss"`
							Share int64 `json:"share"`
							Size  int64 `json:"size"`
						} `json:"memory"`
						Name     string `json:"name"`
						Pgid     int64  `json:"pgid"`
						Pid      int64  `json:"pid"`
						Ppid     int64  `json:"ppid"`
						State    string `json:"state"`
						Username string `json:"username"`
					} `json:"process"`
				} `json:"system"`
			} `json:"_source"`
			Type string `json:"_type"`
		} `json:"hits"`
		MaxScore int64 `json:"max_score"`
		Total    struct {
			Relation string `json:"relation"`
			Value    int64  `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool  `json:"timed_out"`
	Took     int64 `json:"took"`
}
