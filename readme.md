# MCP/Forge Build Skeleton

Go is required for the tools: <http://golang.org/doc/install>

`./install.sh` to build the tools, then read the following

    Unzip latest forge to this directory and cd forge && ./install.sh
    For an existing repo, git clone {repo_url} src
    For a new repo, set up your project in src
    Make sure you have a mcmod.info and build.txt file
    build.txt should set MCVERSION, VERSION, ASSETS, and SOURCE variables
    ASSETS = any folders or files that should go in the packaged jar
    This *includes* mcmod.info.  A manifest.mf file will be automatically detected, however.
    SOURCE = the namespace of your project, i.e. the root folder of your code
    Any external code must be already built and can then be added through ASSETS
    Build libs (forge/mcp/lib/*) must be added manually for now.
    
    After that, ./build.sh to build or time ./build.sh to cry at javac perf
