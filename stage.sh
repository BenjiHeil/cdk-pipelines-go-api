#!/bin/bash
export CODE_IMAGE_REPO_NAME=go-api
export VERSION=$(cat semver)
echo building $VERSION of $CODE_IMAGE_REPO_NAME
echo Build started on `date`

set -e

echo Building the lambda docker image...          
docker build --target stage -t $CODE_IMAGE_REPO_NAME:$VERSION ./src
docker run -v $PWD/out:/out $CODE_IMAGE_REPO_NAME:$VERSION find . -name '*.out' -exec mv -i {} /out/ \;
cat $PWD/out/test.out

#check if benchmark over 200 ns/op
while read p; do
    if [[ $p -gt 200 ]] 
    then 
        >&2 echo $p ns/op greater than threshold of 200, failing...
        exit 1
    fi
done < $PWD/out/benchmark.out

set +e
