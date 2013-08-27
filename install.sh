set -e

go version

go build jarbuilder.go
go build modid.go

echo
echo "== tools successfully built =="
echo
echo "Unzip latest forge to this directory and cd forge && ./install.sh"
echo "For an existing repo, git clone {repo_url} src"
echo "For a new repo, set up your project in src"
echo "Make sure you have a mcmod.info and build.txt file"
echo "build.txt should set MCVERSION, VERSION, ASSETS, and SOURCE variables"
echo "ASSETS = any folders or files that should go in the packaged jar"
echo "This *includes* mcmod.info.  A manifest.mf file will be automatically detected, however."
echo "SOURCE = the namespace of your project, i.e. the root folder of your code"
echo "Any external code must be already built and can then be added through ASSETS"
echo "Build libs (forge/mcp/lib/*) must be added manually for now."
echo
echo "After that, ./build.sh to build or time ./build.sh to cry at javac perf"
