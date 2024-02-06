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
      font-family: 'Source Code Pro', monospace;

      font-size: 80%; /* Adjust the percentage based on your preference */
    }
  </style>
</head>
<body>
EOF

# Use awk to replace newlines with </br>
awk '{printf "%s</br>", $0}'

# Print HTML closing tags
echo '</body>'
echo '</html>'
