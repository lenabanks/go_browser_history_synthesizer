### Purpose:
  This project exists as a simple multi-platform application to read and write large amounts of typically user-driven browser activity.

  The use case for this is for faking and testing various forensic applications.

### Dependencies:
  This project depends on Go, vgo and some small selection of supporting tools.

  - Go: It is recommended you install Go `1.11.0` via `goenv` Instructions for installing can be found [here](https://github.com/syndbg/goenv) or by running the following commands:
    ```
    git clone https://github.com/syndbg/goenv.git ~/.goenv

    echo `export GOENV_ROOT="$HOME/.goenv"
          export PATH="$GOENV_ROOT/bin:$PATH"
          export PATH="$GOENV_ROOT/shims:$PATH"
          eval "$(goenv init -)"
          export GOPATH="$HOME/workspace/go"` > ~/.bash_profile
    source ~/.bash_profile

    goenv install 1.10.3
    goenv global  1.10.3
    ```

  - vgo: Dependencies are managed via the Go 1.11 built-in dependency management tool vgo to download dependencies needed for building:
    ```
    go mod init
    go mod tidy
    go mod download
    ```
    
  - OPTIONAL: Docker (I am working on a build system for this)  

### Configuration:
  Minimal configuration is required for this project, just download the dependencies (as stated above) and go.
  
  All of the configuration for runtime variables is, at this time, compiled in for simplicity. This will change in time.