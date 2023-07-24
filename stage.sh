export CODE_IMAGE_REPO_NAME=go-api
export VERSION=$(cat semver)
echo building $VERSION of $CODE_IMAGE_REPO_NAME
echo Build started on `date`

echo Building the lambda docker image...          
docker build --target stage -t $CODE_IMAGE_REPO_NAME:$VERSION ./src
docker run -v $PWD/out:/out $CODE_IMAGE_REPO_NAME:$VERSION find . -name '*.out' -exec mv -i {} /out/ \;
cat $PWD/out/test.out

#check if benchmark over 500 ns/op
while read p; do
    if [[ $p -gt 500 ]] 
    then 
        >&2 echo $p ns/op greater than threshold of 500, failing...
    fi
done < $PWD/out/benchmark.out
