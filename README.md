# The source code has been moved to [Apache Iggy](https://github.com/apache/iggy/)

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Go][Golang]][Golang-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/iggy-rs/iggy-go-client">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">iggy-go</h3>

  <p align="center">
    SDK for iggy written using go language
    <!-- <br />
    <a href="https://github.com/iggy-rs/iggy-go-client"><strong>Explore the docs »</strong></a>
    <br /> -->
    <br />
    <a href="https://github.com/iggy-rs/iggy-go-client/tree/dev/samples">View Samples</a>
    ·
    <a href="https://github.com/iggy-rs/iggy-go-client/issues">Report Bug</a>
    ·
    <a href="https://github.com/iggy-rs/iggy-go-client/issues">Request Feature</a>
    ·
    <a href="https://docs.iggy.rs/">iggy documentation</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

`iggy-go` is a golang SDK for <a href="https://github.com/iggy-rs/iggy">iggy</a> - persistent message streaming platform written in Rust.

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

In order to use this SDK you need to install golang on your environment. Here's a <a hfref="https://go.dev/doc/install">link to official go documentation</a> explaining how you can do that!

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/iggy-rs/iggy-go-client.git
   ```
2. Verify that the solution builds correctly
    ```sh
    cd iggy-go
    go build
    ``` 
<!-- USAGE EXAMPLES -->
## Usage (OBSOLETE; will be updated soon)

If you want to use this sdk as a CLI tool, you can do that by following these steps:
1. Clone iggy repo and run it in background
    ```sh
    git clone https://github.com/iggy-rs/iggy.git
    cd iggy
    cargo r --bin iggy-server -r
    ```
1. Open new terminal instance and enter `iggy-go` root folder
    ```sh
    cd iggy-go
    ```
2. Run your command
    ```sh
    go run ./cli <commandname>
    ```

    You can run `help` command if you would like to see available commands:

    ```sh
      $ go run ./cli help

    Usage:
        getstream    -url <url> -port <port> -streamId <streamId>
        createstream -url <url> -port <port> -streamId <streamId> -name <name>
        deletestream -url <url> -port <port> -streamId <streamId>
        gettopic     -url <url> -port <port> -streamId <streamId> -topicId <topicId>
        createtopic  -url <url> -port <port> -streamId <streamId> -topicId <topicId> -name <name> -partitionsCount <partitionsCount>
        deletetopic  -url <url> -port <port> -streamId <streamId> -topicId <topicId>

    ```

    Some parameters don't have default values so you have to define them manually:

    ```sh
        $ go run ./cli createstream

    Error: Name flag is required.
    -n string
            Stream name
    -name string
            Stream name
    -port string
            Iggy server port (default "8090")
    -s int
            Alias for Stream Id (default 1)
    -sid int
            Alias for Stream Id (default 1)
    -streamid int
            Stream Id (default 1)
    -url string
            Iggy server url (default "127.0.0.1")
    ```
    This is how `createstream` command is called correctly

    ```sh
    $ go run ./cli createstream -s 1 -n stream_name
    ```
    
<!-- ROADMAP -->
## Roadmap

- [x] Basic CLI tool (already obsolete)
- [x] Samples
- [ ] Rewrite the CLI tool using CLI frameworks (like Cobra/Viper)
- [ ] Implementing all iggy features in the SDK
    - [x] TCP
    - [ ] HTTP (can be picked up at any moment)
    - [ ] QUIC
- [ ] Implementing benchmarks
- [ ] Implementing optional logging
- [ ] Implementing tests
  - [ ] Unit tests for binary serialization
  - [ ] BDD integration tests
    - [ ] Improve BDD tests assertions, so they can detect breaking changes in iggy-server
- [X] Human friendly error handling
- [ ] Documentation
- [ ] GitHub actions CI/CD
- [ ] Publishing to official golang packages repository

See the [open issues](https://github.com/iggy-rs/iggy-go-client/issues) for a full list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

If you believe that you can improve this SDK feel free to contribute. Here's how you can do it:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [iggy-rs repository](https://github.com/iggy-rs/iggy)
* [iggy-rs documentation](https://docs.iggy.rs/)

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/iggy-rs/iggy-go-client.svg?style=for-the-badge
[contributors-url]: https://github.com/iggy-rs/iggy-go-client/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/iggy-rs/iggy-go-client.svg?style=for-the-badge
[forks-url]: https://github.com/iggy-rs/iggy-go-client/network/members
[stars-shield]: https://img.shields.io/github/stars/iggy-rs/iggy-go-client.svg?style=for-the-badge
[stars-url]: https://github.com/iggy-rs/iggy-go-client/stargazers
[issues-shield]: https://img.shields.io/github/issues/iggy-rs/iggy-go-client.svg?style=for-the-badge
[issues-url]: https://github.com/iggy-rs/iggy-go-client/issues
[license-shield]: https://img.shields.io/github/license/iggy-rs/iggy-go-client.svg?style=for-the-badge
[license-url]: https://github.com/iggy-rs/iggy-go-client/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[Golang-url]: https://go.dev/
[Golang]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Iggy-docs-url]: https://docs.iggy.rs/
[Iggy-repo-url]: https://github.com/iggy-rs/iggy
[Iggy-dotnet-repo-url]: https://github.com/iggy-rs/iggy-dotnet-client
