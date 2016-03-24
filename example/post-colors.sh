#!/bin/sh

curl -i -F "color[name]=purple" -F "color[color_code]=#9c27b0" -F "color[text_code]=#000000" plog.link/api/v1/colors
