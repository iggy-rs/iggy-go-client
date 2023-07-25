<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
<!-- [![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url] -->



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/eldpcn/iggy-go">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">iggy-go</h3>

  <p align="center">
    SDK for iggy written using go language
    <!-- <br />
    <a href="https://github.com/eldpcn/iggy-go"><strong>Explore the docs »</strong></a>
    <br /> -->
    <br />
    <a href="https://github.com/eldpcn/iggy-go/tree/dev/samples">View Samples</a>
    ·
    <a href="https://github.com/eldpcn/iggy-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/eldpcn/iggy-go/issues">Request Feature</a>
    ·
    <a href="https://docs.iggy.rs/">iggy documentation</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

<!-- Here's a blank template to get started: To avoid retyping too much info. Do a search and replace with your text editor for the following: `eldpcn`, `iggy-go`, `twitter_handle`, `linkedin_username`, `email_client`, `email`, `project_title`, `project_description` -->

`iggy-go` is a golang SDK for <a href="https://docs.iggy.rs/">iggy</a> - persistent message streaming platform written in Rust.


### Built With
 [![Go][Golang]][Golang-url]


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

In order to use this SDK you need to install golang on your enviroment. Here's a <a hfref="https://go.dev/doc/install">link to official go documentation</a> explaining how you can do that!

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/eldpcn/iggy-go.git
   ```
2. Verify that the solution builds correctly
    ```sh
    cd iggy-go
    go build
    ``` 
<!-- USAGE EXAMPLES -->
## Usage

If you want to use this sdk as a CLI tool, you can do that by following these steps:
1. Clone iggy repo and run it in background
    ```sh
    git clone https://github.com/iggy-rs/iggy.git
    cd iggy
    cargo r --bin server -r
    ```
1. Open new terminal instance and enter `iggy-go` root folder
    ```sh
    cd iggy-go
    ```
2. Run your command
    ```sh
    go run ./cli <commandname>
    ```

    You can run `help` command if would like to see avaiable commands:

    ```sh
      $ go run ./cli help

    Usage:
        getstream -url <url> -port <port> -streamId <streamId>
        createstream -url <url> -port <port> -streamId <streamId> -name <name>
        deletestream -url <url> -port <port> -streamId <streamId>
        gettopic -url <url> -port <port> -streamId <streamId> -topicId <topicId>
        createtopic -url <url> -port <port> -streamId <streamId> -topicId <topicId> -name <name> -partitionsCount <partitionsCount>
        deletetopic -url <url> -port <port> -streamId <streamId> -topicId <topicId>

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

- [x] Basic CLI tool
- [x] Samples
- [ ] Rewrite the CLI tool using CLI frameworks (like Cobra/Viper)
- [ ] Full SDK implementation
    - [ ] TCP
    - [ ] HTTP
    - [ ] QUIC

See the [open issues](https://github.com/eldpcn/iggy-go/issues) for a full list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

If you believe that you can improve this SDK feel free to contribute. Here's how you can do it:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
<!-- ## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
 -->

<!-- CONTACT -->
<!-- ## Contact -->

<!-- Your Name - [@twitter_handle](https://twitter.com/twitter_handle) - email@email_client.com -->
<!-- 
Project Link: [https://github.com/eldpcn/iggy-go](https://github.com/eldpcn/iggy-go)

 -->

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [iggy-rs repository](https://github.com/iggy-rs/iggy)
* [iggy-rs documentation](https://docs.iggy.rs/)
* [iggy-dotnet-client](https://github.com/iggy-rs/iggy-dotnet-client)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/eldpcn/iggy-go.svg?style=for-the-badge
[contributors-url]: https://github.com/eldpcn/iggy-go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/eldpcn/iggy-go.svg?style=for-the-badge
[forks-url]: https://github.com/eldpcn/iggy-go/network/members
[stars-shield]: https://img.shields.io/github/stars/eldpcn/iggy-go.svg?style=for-the-badge
[stars-url]: https://github.com/eldpcn/iggy-go/stargazers
[issues-shield]: https://img.shields.io/github/issues/eldpcn/iggy-go.svg?style=for-the-badge
[issues-url]: https://github.com/eldpcn/iggy-go/issues
[license-shield]: https://img.shields.io/github/license/eldpcn/iggy-go.svg?style=for-the-badge
[license-url]: https://github.com/eldpcn/iggy-go/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[Golang-url]: https://go.dev/
[Golang]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Iggy-docs-url]: https://docs.iggy.rs/
[Iggy-repo-url]: https://github.com/iggy-rs/iggy
[Iggy-dotnet-repo-url]: https://github.com/iggy-rs/iggy-dotnet-client
