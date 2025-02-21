package bcbp

import (
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/skrushinsky/scaliger/julian"
)

type BCBP struct {
	FormatCode                 string `json:"format_code"`
	NumberOfLegs               string `json:"number_of_legs"`
	PassengerName              string `json:"passenger_name"`
	ElectronicTicketIndicator  string `json:"electronic_ticket_indicator"`
	OperatingCarrierPNR        string `json:"operating_carrier_pnr"`
	FromAirport                string `json:"from_airport"`
	ToAirport                  string `json:"to_airport"`
	OperatingCarrierDesignator string `json:"operating_carrier_designator"`
	FlightNumber               string `json:"flight_number"`
	DateOfFlight               string `json:"date_of_flight"`
	CompartmentCode            string `json:"compartment_code"`
	SeatNumber                 string `json:"seat_number"`
	CheckInSequenceNumber      string `json:"checkin_sequence_number"`
	PassengerStatus            string `json:"passenger_status"`
}

func (b *BCBP) MonthDay() (int, int, error) {

	jd, err := strconv.ParseFloat(b.DateOfFlight, 64)

	if err != nil {
		return -1, -1, err
	}

	cd := julian.JulianToCivil(jd)
	return cd.Month, int(math.Floor(cd.Day)), nil
}

func (b *BCBP) String() string {

	parts := []string{
		b.FormatCode,
		b.NumberOfLegs,
		rightPad(b.PassengerName, " ", PASSENGER_NAME),
		b.ElectronicTicketIndicator,
		rightPad(b.OperatingCarrierPNR, " ", OPERATING_CARRIER_PNR),
		b.FromAirport,
		b.ToAirport,
		rightPad(b.OperatingCarrierDesignator, " ", OPERATING_CARRIER_DESIGNATOR),
		leftPad(b.FlightNumber, "0", FLIGHT_NUMBER),
		b.DateOfFlight,
		b.CompartmentCode,
		leftPad(b.SeatNumber, "0", SEAT_NUMBER),
		leftPad(b.CheckInSequenceNumber, "0", CHECK_IN_SEQUENCE_NUMBER),
		b.PassengerStatus,
	}

	return strings.Join(parts, "")
}

func Parse(bcbp string) (*BCBP, error) {

	bcbpData := &BCBP{
		FormatCode:                 getField(bcbp, FORMAT_CODE_OFFSET, FORMAT_CODE),
		NumberOfLegs:               getField(bcbp, NUMBER_OF_LEGS_OFFSET, NUMBER_OF_LEGS),
		PassengerName:              strings.TrimSpace(getField(bcbp, PASSENGER_NAME_OFFSET, PASSENGER_NAME)),
		ElectronicTicketIndicator:  getField(bcbp, ELECTRONIC_TICKET_INDICATOR_OFFSET, ELECTRONIC_TICKET_INDICATOR),
		OperatingCarrierPNR:        strings.TrimSpace(getField(bcbp, OPERATING_CARRIER_PNR_OFFSET, OPERATING_CARRIER_PNR)),
		FromAirport:                strings.TrimSpace(getField(bcbp, DEPARTURE_AIRPORT_OFFSET, DEPARTURE_AIRPORT)),
		ToAirport:                  strings.TrimSpace(getField(bcbp, ARRIVAL_AIRPORT_OFFSET, ARRIVAL_AIRPORT)),
		OperatingCarrierDesignator: strings.TrimSpace(getField(bcbp, OPERATING_CARRIER_DESIGNATOR_OFFSET, OPERATING_CARRIER_DESIGNATOR)),
		FlightNumber:               strings.TrimLeft(getField(bcbp, FLIGHT_NUMBER_OFFSET, FLIGHT_NUMBER), "0"),
		DateOfFlight:               getField(bcbp, FLIGHT_DATE_OFFSET, FLIGHT_DATE),
		CompartmentCode:            getField(bcbp, COMPARTMENT_CODE_OFFSET, COMPARTMENT_CODE),
		SeatNumber:                 strings.TrimLeft(getField(bcbp, SEAT_NUMBER_OFFSET, SEAT_NUMBER), "0"),
		CheckInSequenceNumber:      strings.TrimLeft(getField(bcbp, CHECK_IN_SEQUENCE_NUMBER_OFFSET, CHECK_IN_SEQUENCE_NUMBER), "0"),
		PassengerStatus:            getField(bcbp, PASSENGER_STATUS_OFFSET, PASSENGER_STATUS),
	}

	return bcbpData, nil
}

func leftPad(raw string, pad string, length int) string {

	for len(raw) < length {
		raw = pad + raw
	}

	return raw
}

func rightPad(raw string, pad string, length int) string {

	for len(raw) < length {
		raw = raw + pad
	}

	return raw
}

func getField(raw string, offset int, length int) string {
	v := raw[offset : offset+length]
	slog.Debug("field", "raw", raw, "offset", offset, "length", length, "v", v)
	return v
}
