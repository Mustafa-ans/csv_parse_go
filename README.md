# CSV Parser in Golang

This is a simple CSV parser written in Golang. It reads in a CSV file and prints out the contents in a tabular format.

## Getting Started

To use this parser, you must have Golang installed on your machine. Once you have Golang installed, you can clone this repository and run the following command to execute the code:

`go run main.go` 

This will parse the data.csv file and print out the contents in a table.

## Input File Format

The input file must be in CSV format. The first row of the file must contain the column headers, and the subsequent rows must contain the data. The columns must be separated by commas. The file must have at least one row of data.

## Output Format

The output of the parser is a table that displays the data in a readable format. The table is formatted using the tabwriter package in Golang.
License

This code is released under the MIT License. Feel free to use it for any purpose.
