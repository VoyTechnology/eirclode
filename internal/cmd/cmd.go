package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"eirclode.voy.technology/data"
)

func Run(cfg Config) error {
	m := make(map[string]Data)

	r := csv.NewReader(bytes.NewReader(data.Data))
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to read all records: %w", err)
	}
	for _, line := range records {
		lat, err := strconv.ParseFloat(line[1], 32)
		if err != nil {
			return fmt.Errorf("unable to parse lattitude: %w", err)
		}
		lon, err := strconv.ParseFloat(line[2], 32)
		if err != nil {
			return fmt.Errorf("unable to parse longitude: %w", err)
		}
		m[line[0]] = Data{
			Eircode:   line[0],
			Lattitude: float32(lat),
			Longitude: float32(lon),
			Address:   line[3],
		}
	}

	// Clear the data variable as we will no longer need it and we can recover
	// some memory
	data.Data = []byte{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		want := strings.TrimPrefix(r.URL.Path, "/")

		d, exists := m[want]
		if !exists {
			http.NotFound(w, r)
			return
		}

		ct := r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", ct)

		switch ct {
		case "application/json":
			json.NewEncoder(w).Encode(d)
		case "":
			fmt.Fprint(w, d.String())
		default:
			http.Error(w, "unknown Content-Type", http.StatusBadRequest)
		}
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}

type Config struct {
	Port int `enconfig:"PORT" default:"8080"`
}

type Data struct {
	Eircode   string  `json:"eircode" yaml:"eircode"`
	Lattitude float32 `json:"lattitide" yaml:"lattitide"`
	Longitude float32 `json:"longitude" yaml:"longitude"`
	Address   string  `json:"address" yaml:"address"`
}

func (d Data) String() string {
	return fmt.Sprintf("%s\n%.5f\n%.5f\n%s",
		d.Eircode,
		d.Lattitude,
		d.Longitude,
		d.Address,
	)
}
