name: {{ .Name }}-build-test
displayName: build-test
triggers:
  push:
    branches:
      - master
steps:
  - step: golangbuild@1
    name: golang-build
    displayName: golang-build
    inputs:
      golangVersion: 1.18
      goals: "sed -i '/replace /d' go.mod && mkdir dist && GOOS=linux GOARCH=amd64 go build -ldflags \"-s -w\" -o dist/{{ .Name }}_linux_amd64.amd64 main.go"
      uploadArtifact: false
      uploadArtifactOptions:
        artifactPath: "./dist"
        artifactRepository: "{{ .Name}}"
        artifactName: "{{ .Name }}_build_releases"
