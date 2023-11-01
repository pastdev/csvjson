# csvjson

This project provides utilities for converting data between CSV and JSON.

## Running

See the help documentation:

```bash
# for general help
csvjson --help

# for help on converting from csv to json:
csvjson c2j --help
```

## Examples

You can convert your csv file to a JSON object:

```bash
csvjson c2j --format json-object --has-header <<'EOF'
firstname,lastname
miss,piggy
kermit,thefrog
EOF
# [{"firstname":"miss","lastname":"piggy"},{"firstname":"kermit","lastname":"thefrog"}]
```

Which can then be piped into JSON processing tools like [clconf](https://github.com/pastdev/clconf/blob/master/README.md):

```bash
csvjson c2j --format json-object --has-header <<'EOF' | clconf --pipe jsonpath '$[*].lastname'
firstname,lastname
miss,piggy
kermit,thefrog
EOF
# - piggy
# - thefrog
```
