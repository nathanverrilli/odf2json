# odf2json

Convert spreadsheet to JSON documents.
Intended to support ODS, but currently supports only
Microsoft Office Open XML (.xlsx files).
Looking for a decent way to support ODS without 
significant pain.

* Assumes data header is in first row.
 
## Usage

 * __--debug__            _log data to STDERR as well as log file_
 * __--infile string__    _filename to process_
 * __--outfile string__   _Output JSON file_
 * __--type string__      _Name of record type conversion (default "record")_
 * __--verbose__          _Lots of information (default true)_
 
`odf2json --infile zodmo.xlxs --outfile zodmo.json`
 
 Log data is written to `odf2json.log`