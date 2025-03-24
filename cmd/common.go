package cmd

import (
	"time"
)

// Version information
var (
	Version   = "dev"
	BuildDate = "unknown"
	BuildHash = "unknown"
)

// Global flags
var (
	Debug        bool
	Uint32Flag   bool
	URL          string
	FilePath     string
	Base64Path   string
	UserAgent    string
	FofaFormat   bool
	ShodanFormat bool
	SkipVerify   bool
	Timeout      time.Duration
	OutputFormat string
)

// Server flags
var (
	Host         string
	Port         int
	AuthToken    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

// MonitorData stores favicon monitoring information
type MonitorData struct {
	URL          string
	CurrentHash  string
	PreviousHash string
	FirstHash    string
	FirstSeen    time.Time
	LastChanged  time.Time
	IconData     []byte
	ChangeCount  int
}

// Shared structs for command options
var (
	// Batch command options
	BatchOptions = struct {
		InputFile     string
		OutputFile    string
		Delimiter     string
		Format        string
		ErrorHandling string
	}{}

	// Compare command options
	CompareOptions = struct {
		FirstSource  string
		SecondSource string
		Threshold    float64
	}{}

	// Search command options
	SearchOptions = struct {
		Hash        string
		Query       string
		Engine      string
		OpenBrowser bool
	}{}

	// Screenshot command options
	ScreenshotOptions = struct {
		URL        string
		OutputFile string
		Size       string
		Timeout    int
	}{}

	// Convert command options
	ConvertOptions = struct {
		Hash       string
		FromFormat string
		ToFormat   string
		WithSyntax bool
	}{}

	// Stats command options
	StatsOptions = struct {
		InputFile  string
		OutputFile string
		Format     string
		Visualize  bool
		GroupBy    string
	}{}

	// Monitor command options
	MonitorOptions = struct {
		Targets         []string
		TargetsFile     string
		Interval        time.Duration
		MaxRuns         int
		OutputDir       string
		Notify          bool
		ChangeThreshold float64
	}{}

	// Scan command options
	ScanOptions = struct {
		Targets       []string
		TargetsFile   string
		IPRange       string
		DomainSuffix  string
		Threads       int
		Timeout       time.Duration
		OutputFile    string
		OnlyWithIcons bool
		PortList      string
	}{}
)
