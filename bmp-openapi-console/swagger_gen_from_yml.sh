#!/bin/bash

# install goswagger: https://goswagger.io/install.html#Installing%20from%20source

swagger version
# swagger validate ./swagger.yml goswagger有缺陷，且报错无法提供有效信息，建议使用https://editor.swagger.io/来校验
# validate 和 generate命令在0.25.0版本有panic，0.30.3版本已修复
swagger validate swagger.yml
\rm -rf ../bmp-openapi-console-sdk/client && \rm -rf ../bmp-openapi-console-sdk/models
swagger generate client -f swagger.yml --target=../bmp-openapi-console-sdk
