# go-bcbp-wasm

## parse_bcbp.wasm

```
var bcbp_str = "M1DESMARAIS/LUC       EABC123 LASSFOUA 0574 419J001A0025 100";

parse_bcbp(bcbp_str).then(rsp => {
	// Do something with rsp	
}).catch(err => {
	console.error("Failed to initialize parse_bcbp.wasm", err)
});
```

Where `rsp` looks like this:

```
{
  "raw": "M1DESMARAIS/LUC       EABC123 LASSFOUA 0574 419J001A0025 100",
  "fields": {
    "format_code": "M",
    "number_of_legs": "1",
    "passenger_name": "DESMARAIS/LUC",
    "electronic_ticket_indicator": "E",
    "operating_carrier_pnr": "ABC123",
    "from_airport": "LAS",
    "to_airport": "SFO",
    "operating_carrier_designator": "UA",
    "flight_number": "574 ",
    "date_of_flight": "419",
    "compartment_code": "J",
    "seat_number": "1A",
    "checkin_sequence_number": "25 ",
    "passenger_status": "1"
  },
  "month": 2,
  "day": 23
}
```

## Example

```
$> make debug
fileserver \
		-root ./www \
		-server-uri http://localhost:8080 \
		-mimetype js=text/javascript \
		-mimetype wasm=application/wasm \
		-enable-cors
2025/02/21 15:36:14 Serving ./www and listening for requests on http://localhost:8080

```

## See also

* https://github.com/sfomuseum/go-bcbp