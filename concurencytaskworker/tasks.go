package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

// Define a struct to hold graduate data
type Graduate struct {
	Year  int    `json:"year"`
	Major string `json:"major"`
	// Add other relevant fields
}

func main() {
	// Parse command-line flags
	concurrentLimit := flag.Int("concurrent_limit", 2, "Limit of concurrent processes")
	outputDir := flag.String("output", ".", "Output directory for CSV files")
	flag.Parse()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a buffered channel to limit the number of concurrent processes
	concurrentLimitChan := make(chan struct{}, *concurrentLimit)

	// Iterate over years
	for year := 2020; year <= 2022; year++ {
		wg.Add(1)
		go func(year int) {
			defer wg.Done()
			// Acquire a token from the channel to limit concurrency
			concurrentLimitChan <- struct{}{}
			defer func() {
				// Release the token when the goroutine finishes
				<-concurrentLimitChan
			}()

			// Fetch graduate data for the current year
			graduates, err := fetchGraduateData(year)
			if err != nil {
				fmt.Printf("Error fetching data for year %d: %v\n", year, err)
				return
			}

			// Process and save the data to a CSV file
			filePath := fmt.Sprintf("%s/%d.csv", *outputDir, year)
			err = saveToCSV(filePath, graduates)
			if err != nil {
				fmt.Printf("Error saving data to CSV for year %d: %v\n", year, err)
				return
			}

			fmt.Printf("Data for year %d saved to %s\n", year, filePath)
		}(year)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Function to fetch graduate data for a given year from the API
func fetchGraduateData(year int) ([]Graduate, error) {
	// Make HTTP request to the API
	resp, err := http.Get(fmt.Sprintf("https://beta.data.gov.sg/datasets?query=test", year))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a slice of Graduate structs
	var graduates []Graduate
	err = json.NewDecoder(resp.Body).Decode(&graduates)
	if err != nil {
		return nil, err
	}

	return graduates, nil
}

// Function to save graduate data to a CSV file
func saveToCSV(filePath string, graduates []Graduate) error {
	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"Year", "Major"} // Add other fields if needed
	err = writer.Write(header)
	if err != nil {
		return err
	}

	// Write data rows
	for _, graduate := range graduates {
		row := []string{fmt.Sprint(graduate.Year), graduate.Major} // Add other fields if needed
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
