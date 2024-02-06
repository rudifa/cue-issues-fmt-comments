#!/bin/bash

# Use a heredoc for the HTML preamble
cat <<EOF
<!DOCTYPE html>
<html>
<head>
  <title>Converted Text</title>
  <style>
    body {
      font-family: "Consolas", monospace;
      font-size: 80%;
    }
    .node {
      background-color: green;
    }
    .closeNode {
      background-color: lightgreen;
    }
  </style>
</head>
<body>
EOF

# Use awk to replace newlines with </br> and apply color highlighting
awk '{
  gsub(/node: 0x[[:xdigit:]]+/, "<span class=\"node\">&</span>");
  gsub(/closeNode/, "<span class=\"closeNode\">&</span>");
  printf "%s</br>", $0;
}'

# Print HTML closing tags
echo '</body>'
echo '</html>'
