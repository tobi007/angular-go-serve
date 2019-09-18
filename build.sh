#@IgnoreInspection BashAddShebang

if [[ -d releases ]]
then
  rm -rf releases
fi

if [[ -d dist ]]
then
  rm -rf dist
fi

cd ui && ng build --prod && cd ..

go-bindata-assetfs dist/...

mkdir releases \
  && GOOS=windows GOARCH=amd64 go build && mv ./venus-third-party-app.exe ./releases/venus-third-party-ui.exe
