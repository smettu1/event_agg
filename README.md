

## Input

Stream of objects
{"id":121509,"market":5773,"price":1.234,"volume":1234.56,"is_buy":true}
{"id":121510,"market":5774,"price":2.345,"volume":2345.67,"is_buy":false}
{"id":121511,"market":5775,"price":3.456,"volume":3456.78,"is_buy":true}


## Output 
2022/07/17 08:28:44 {"market":3329,"total_volume":2293338.151998,"mean_price":1.502500,"mean_volume":2403.918398,"volume_weighted_average_price":1.495162,"percentage_buy":100.000000}
2022/07/17 08:28:44 {"market":6186,"total_volume":2337355.054670,"mean_price":50.515525,"mean_volume":2452.628599,"volume_weighted_average_price":50.519985,"percentage_buy":100.000000}


## Design
Main code is split into few modules
    Models:
        Has all the definitions of the input and output data structures
    Parse:
        Parses all the input streams and validates the input
    Metrics:
        Computes the metrics.

## Usage
Stream data to binary.
/stdoutinator_amd64_darwin.bin | ./event


