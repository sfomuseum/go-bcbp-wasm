package bcbp

import (
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/skrushinsky/scaliger/julian"
)

type Leg struct {
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
	OptionalDataSize           string `json:"optional_data_size"`
	OptionalData               string `json:"optional_data"`
}

func (l *Leg) MonthDay() (int, int, error) {

	jd, err := strconv.ParseFloat(l.DateOfFlight, 64)

	if err != nil {
		return -1, -1, err
	}

	cd := julian.JulianToCivil(jd)
	return cd.Month, int(math.Floor(cd.Day)), nil
}

func (l *Leg) String() string {

	parts := []string{
		l.FormatCode,
		l.NumberOfLegs,
		rightPad(l.PassengerName, " ", PASSENGER_NAME),
		l.ElectronicTicketIndicator,
		rightPad(l.OperatingCarrierPNR, " ", OPERATING_CARRIER_PNR),
		l.FromAirport,
		l.ToAirport,
		rightPad(l.OperatingCarrierDesignator, " ", OPERATING_CARRIER_DESIGNATOR),
		leftPad(l.FlightNumber, "0", FLIGHT_NUMBER),
		l.DateOfFlight,
		l.CompartmentCode,
		leftPad(l.SeatNumber, "0", SEAT_NUMBER),
		leftPad(l.CheckInSequenceNumber, "0", CHECK_IN_SEQUENCE_NUMBER),
		l.PassengerStatus,
		leftPad(l.OptionalDataSize, "0", OPTIONAL_DATA_SIZE),
		l.OptionalData,
	}

	return strings.Join(parts, "")
}

func ParseLeg(raw string) (*Leg, error) {

	leg := &Leg{
		FormatCode:                 getField(raw, FORMAT_CODE_OFFSET, FORMAT_CODE),
		NumberOfLegs:               getField(raw, NUMBER_OF_LEGS_OFFSET, NUMBER_OF_LEGS),
		PassengerName:              strings.TrimSpace(getField(raw, PASSENGER_NAME_OFFSET, PASSENGER_NAME)),
		ElectronicTicketIndicator:  getField(raw, ELECTRONIC_TICKET_INDICATOR_OFFSET, ELECTRONIC_TICKET_INDICATOR),
		OperatingCarrierPNR:        strings.TrimSpace(getField(raw, OPERATING_CARRIER_PNR_OFFSET, OPERATING_CARRIER_PNR)),
		FromAirport:                strings.TrimSpace(getField(raw, DEPARTURE_AIRPORT_OFFSET, DEPARTURE_AIRPORT)),
		ToAirport:                  strings.TrimSpace(getField(raw, ARRIVAL_AIRPORT_OFFSET, ARRIVAL_AIRPORT)),
		OperatingCarrierDesignator: strings.TrimSpace(getField(raw, OPERATING_CARRIER_DESIGNATOR_OFFSET, OPERATING_CARRIER_DESIGNATOR)),
		FlightNumber:               strings.TrimLeft(getField(raw, FLIGHT_NUMBER_OFFSET, FLIGHT_NUMBER), "0"),
		DateOfFlight:               getField(raw, FLIGHT_DATE_OFFSET, FLIGHT_DATE),
		CompartmentCode:            getField(raw, COMPARTMENT_CODE_OFFSET, COMPARTMENT_CODE),
		SeatNumber:                 strings.TrimLeft(getField(raw, SEAT_NUMBER_OFFSET, SEAT_NUMBER), "0"),
		CheckInSequenceNumber:      strings.TrimLeft(getField(raw, CHECK_IN_SEQUENCE_NUMBER_OFFSET, CHECK_IN_SEQUENCE_NUMBER), "0"),
		PassengerStatus:            getField(raw, PASSENGER_STATUS_OFFSET, PASSENGER_STATUS),
		OptionalDataSize:           getField(raw, OPTIONAL_DATA_SIZE_OFFSET, OPTIONAL_DATA_SIZE),
	}

	// Note: There is apparently no requirement that the length of the optional data match
	// the value of the optional data size field

	if len(raw) > 60 {
		leg.OptionalData = raw[60:len(raw)]
	}

	return leg, nil
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
