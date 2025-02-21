package bcbp

/*

Mandatory Fields (Fixed Length, 60 Characters)
These fields must be present in every BCBP string.

Field No.	Field Name	Offset Position	Length (Chars)	Notes
1	Format Code	0	1	Usually "M"
2	Number of Legs Encoded	1	1	"1" = Single-leg, "2" = Multi-leg
3	Passenger Name	2	20	Left-justified, trailing spaces
4	Electronic Ticket Indicator	22	1	"E" for e-ticket, " " for paper
5	Operating Carrier PNR Code	23	7	Left-justified, trailing spaces
6	From Airport Code	30	3	IATA 3-letter airport code
7	To Airport Code	33	3	IATA 3-letter airport code
8	Operating Carrier Designator	36	3	IATA airline code, left-justified
9	Flight Number	39	5	Right-justified, leading zeros
10	Date of Flight (Julian Date)	44	3	Format: DDD (001-366)
11	Compartment Code	47	1	Cabin class (e.g., Y = Economy)
12	Seat Number	48	4	Right-justified, leading zeros
13	Check-in Sequence Number	52	5	Right-justified, leading zeros
14	Passenger Status	57	1	"0" = Not checked in, "1" = Checked in

*/

// M|1|DESMARAIS/LUC       |E|ABC123 |YUL|FRA|AC |0834 |226|F|001A|0025 |1

const FORMAT_CODE_OFFSET int = 0
const FORMAT_CODE int = 1

const NUMBER_OF_LEGS_OFFSET int = 1
const NUMBER_OF_LEGS int = 1

const PASSENGER_NAME_OFFSET int = 2
const PASSENGER_NAME int = 20

const ELECTRONIC_TICKET_INDICATOR_OFFSET int = 22
const ELECTRONIC_TICKET_INDICATOR int = 1

const OPERATING_CARRIER_PNR_OFFSET int = 23
const OPERATING_CARRIER_PNR int = 7

const DEPARTURE_AIRPORT_OFFSET int = 30
const DEPARTURE_AIRPORT int = 3

const ARRIVAL_AIRPORT_OFFSET int = 33
const ARRIVAL_AIRPORT int = 3

const OPERATING_CARRIER_DESIGNATOR_OFFSET int = 36
const OPERATING_CARRIER_DESIGNATOR int = 3

const FLIGHT_NUMBER_OFFSET int = 39
const FLIGHT_NUMBER int = 5

const FLIGHT_DATE_OFFSET int = 44
const FLIGHT_DATE int = 3

const COMPARTMENT_CODE_OFFSET int = 47
const COMPARTMENT_CODE int = 1

const SEAT_NUMBER_OFFSET int = 48
const SEAT_NUMBER int = 4

const CHECK_IN_SEQUENCE_NUMBER_OFFSET int = 52
const CHECK_IN_SEQUENCE_NUMBER int = 5

const PASSENGER_STATUS_OFFSET int = 57
const PASSENGER_STATUS int = 1
