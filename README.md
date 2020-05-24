# fileShrinker
Randomly select a percentage of lines of a given file to generate a new file

# Run

`./fileShrinker -i=fn2fn.csv -p=0.6 -h -s`

- `i`: input file is "fn2fn.csv"
- `p`: shrink percentage is 0.6
- `h`: keep headers. If true, first line would be appeared in the generated file(s)
- `s`: keep both files. If true, keep both fn2fn.csv.0.6 and fn2fn.csv.0.4 
