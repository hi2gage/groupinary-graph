#!/bin/zsh

# Check if a path is provided as an argument
if [[ $# -eq 0 ]]; then
    echo "Usage: $0 <path>"
    exit 1
fi

# Change to the specified directory
cd "$1" || exit 1


# Get a list of all files ending with .graphql
files=(**/*.graphql)
line_length=80
col_length=5

# Functions

print_full_row() {
    local line_length=$1
    printf '%.0s#' $(seq 1 $line_length)
    echo
}

print_blank_row() {
    local line_length=$1
    local col_len=$2

    printf '%.0s#' $(seq 1 $col_len)
    echo -n
    printf '%.0s ' $(seq 1 $((line_length - col_len * 2)))
    printf '%.0s#' $(seq 1 $col_len)
    echo
}

print_dashed_line() {
    local col_length=$1
    local line_length=$2
    local file=$3
    local spaces=$(( (line_length - ${#file}) / 2 - col_length + 1))

    # Start #'s
    printf '%.0s#' $(seq 1 $col_length)
    echo -n
    

    # Spaces then string
    printf '%.0s ' $(seq 1 $spaces)
    echo -n "$file"   # Output filename as a comment
    printf '%.0s ' $(seq 1 $((spaces - 1)))

    # End #'s
    printf '%.0s#' $(seq 1 $col_length)
    echo
}


# Redirect the output to schema.graphqls
# Loop through each file and concatenate its contents
{
for file in $files; do
    # Output the dashed line with centered filename
    print_full_row $line_length 

    print_blank_row $line_length $col_length
    print_blank_row $line_length $col_length

    print_dashed_line $col_length $line_length $file

    print_blank_row $line_length $col_length
    print_blank_row $line_length $col_length

    print_full_row $line_length

    echo

    # Concatenate file contents
    cat "$file"

    echo "\n\n\n"
done
} > combined_schema.graphqls


