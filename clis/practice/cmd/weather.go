package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

const (
	forecastURL = "https://api.open-meteo.com/v1/forecast"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if latitude == 0 || longitude == 0 {
			return errors.New("both latitude and longitude are required")
		}

		if timezone == "" {
			return errors.New("timezone is required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := http.NewRequest(http.MethodGet, forecastURL, nil)
		if err != nil {
			cmd.Println(err)
			return
		}
		queries := url.Values{}
		queries.Add("latitude", fmt.Sprintf("%.2f", latitude))
		queries.Add("longitude", fmt.Sprintf("%.2f", longitude))
		queries.Add("daily", "temperature_2m_max")
		queries.Add("timezone", timezone)
		r.URL.RawQuery = queries.Encode()

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			cmd.Println(err)
			return
		}

		var forecastResponse ForecastResponse
		if err := json.NewDecoder(res.Body).Decode(&forecastResponse); err != nil {
			cmd.PrintErrln(err)
			return
		}

		forecastResponse.Pretty()
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	weatherCmd.Flags().Float32Var(&latitude, "latitude", 0, "Latitude of the location")
	weatherCmd.Flags().Float32Var(&longitude, "longitude", 0, "Longitude of the location")
	weatherCmd.Flags().StringVar(&timezone, "timezone", "", "Timezone of the location")
}

var (
	latitude  float32
	longitude float32
	timezone  string
)

type ForecastResponse struct {
	Daily struct {
		Times        []string  `json:"time"`
		Temperatures []float32 `json:"temperature_2m_max"`
	} `json:"daily"`
}

func (r ForecastResponse) Pretty() {
	w := tabwriter.NewWriter(os.Stdout, 10, 8, 1, ' ', tabwriter.Debug|tabwriter.AlignRight)
	defer w.Flush()

	fmt.Println("Result:")

	for _, t := range r.Daily.Times {
		fmt.Fprintf(w, "%s\t", t)
	}
	fmt.Fprintln(w)

	for _, t := range r.Daily.Temperatures {
		fmt.Fprintf(w, "%.2f\t", t)
	}
	fmt.Fprintln(w)
}
