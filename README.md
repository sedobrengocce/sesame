# sesame
    ![GitHub Release Date](https://img.shields.io/github/release-date-pre/sedobrengocce/sesame?style=for-the-badge)

### Build With
    ![Golang](https://img.shields.io/badge/Golang-gray?style=for-the-badge&logo=go)

## Contents
    - [About the Project](#about-the-project)
    - [Getting Started](#getting-started)
    - [Usage](#usage)
    - [To Do](#to-do)
    - [License](#license)
    - [Contacts](#contacts)
    - [Change Log](https://github.com/seDobrengocce/sesame/blob/main/CHANGELOG.md)

## About the Project
    ### What is sesame?
        Sesame is a simple port knocking tool written in Go.
    ### What is port knocking?
        Port knocking is a method used to secure connections to network services. The idea is to send a sequence of packets to a specific set of ports to open a connection to a service that is protected by a firewall.
    ### How does sesame work?
        Sesame sends a sequence of packets to a specific set of ports to open a connection to a service that is protected by a firewall.
    ### Why sesame?
        Because all the other port knocking tools I tried doesn't save the sequence of ports used to open the connection. Sesame does. (Not now, but it will)
        Also, all the other port knocking tools show the sequence of ports used to open the connection. Sesame doesn't. (For your security)
## Getting Started
    To get a local copy up and running follow these simple steps.
    At the moment there is only a prelease version of sesame. You can download it from the [releases page](https://github.com/sedobrengocce/sesame/releases).

    But you can also build it from source.
    ### Prerequisites
        - Go
        - Git
    ### Installation
        1. Clone the repo
            ```bash
            git clone git@github.com:sedobrengocce/sesame.git
            ```
        2. Build the project
            ```bash
            go build -o sesame main.go
            ```
        3. Run the project and check the help
            ```bash
            ./sesame -h
            ```

## Usage
   ```bash
   ./sesame -H <host> -p <ports>
   ```
   The list of the ports are coma separated. 

## To Do
    - [ ] Use database to store data (crypted)
    - [ ] Add a TUI
    - [ ] Add a GUI
    - [ ] Add a git integration to sync data
    - [ ] Add a fake verbose mode to hide data
    - [ ] Add Makefile to build and install the project

## License
    This project is licensed under the MIT License
    <a href="https://github.com/sedobrengocce/sesame/blob/main/LICENSE"><img src="https://img.shields.io/github/license/sedobrengocce/sesame?style=for-the-badge"/></a>

## Contacts
    Giuseppe Tufo - [LinkedIn](https://www.linkedin.com/in/giuseppe-tufo-3513a224) - giuseppe.tufo@gmail.com
