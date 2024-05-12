# Pokédex - Golang CLI Tool
The Pokédex Tool is a command-line application that allows users to explore various Pokémon locations and discover the Pokémon species available in those locations. The tool interacts with the Pokémon API (pokeapi.co) to fetch data about different regions and the Pokémon encounters within them.

## Features

- Command-line Interface (CLI): Users interact with the tool via a command-line interface, where they can input commands to perform various actions.
- Commands: The tool supports several commands, including:
    - help: Displays a list of available commands and their descriptions.
    - exit: Exits the Pokémon Exploration Tool.
    - map: Shows the next 20 regions available for exploration.
    - mapb: Shows the previous 20 regions.
    - explore {location}: Allows users to explore a specific location by providing its name.
    - inspect {pokemon}: Allows users to inspect a specific pokémon by providing its name.
- Caching Mechanism: To optimize API calls and reduce response time, the tool implements a caching mechanism. Cached data is stored locally to avoid redundant requests to the Pokémon API.
- Testing: The tool includes unit tests to ensure the correctness and reliability of its functionalities. Tests cover features such as text cleaning, cache management, and API interactions.

## Implementation
 - The tool is implemented in Go programming language.
 - It utilizes the pokeapi package to interact with the Pokémon API and fetch data about locations and Pokémon encounters.
 - A caching package is used to store and manage cached responses from the API, improving performance.
 - The command-line interface is powered by a set of commands defined within the main package, enabling users to navigate and explore Pokémon regions efficiently.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/doc/install): Make sure you have Go Programming language installed on your system.

## Installation

1. Clone this repository to your local machine using:

   ```bash
   git clone https://github.com/Yash-sudo-web/Pokedex_Golang_CLI-

2. Navigate to the project directory:

    ```bash
    cd Pokedex_Golang_CLI-

3. Build the main.go file:

    ```bash
    go build .

## Usage

1. Execute the built file:

   ```bash
   ./pokedex.exe

2. Command-line Interface: Run the command 'help' in the terminal to get started with the Pokémon Exploration Tool.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow the usual GitHub fork and pull request workflow.

## License

This project is licensed under the [MIT License](https://github.com/Yash-sudo-web/Pokedex_Golang_CLI-/blob/main/LICENSE).
