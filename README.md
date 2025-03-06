# Gator CLI
Gator is a training project from the course BootDev. It is a command-line tool (CLI) that can:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

This README file will guide you through the installation and configuration process.
  
## Prerequisites
Before running Gator, ensure you have the following installed on your machine:
-  **PostgreSQL**: Make sure you have PostgreSQL installed and running.
-  **Go**: You need to have Go installed. You can download it from [the official Go website](https://golang.org/dl/).

  

## Installing Gator CLI
  

To install the Gator CLI, use the following command:
```bash
go install github.com/koldunNomad/gator@latest
```

Setting Up the Configuration

Create a configuration file named .gatorconfig.json in the root of your project.

Add the PostgreSQL connection string in JSON format with the key db_url. Here is an example:

```json

{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}

```

Replace username, password, host, port, and database with your actual PostgreSQL credentials and settings.

  

Running the Program

After setting up the configuration, you can run the Gator CLI. Note that Go programs are statically compiled binaries, so after running go build or go install, you can run your code without needing the Go toolchain.

Available Commands

Gator provides several commands for interacting with your PostgreSQL database. Here are a few examples:

- **`gator register <name>`**: Create a new user.
- **`gator login <name>`**: Switch to another user.
- **`gator users`**: Print the list of users.
- **`gator agg [time_between_reqs]`**: Collect feeds.
- **`gator addfeed <name> <url>`**: Add a feed for the user.
- **`gator feeds`**: Print the list of feeds.
- **`gator follow <url>`**: Follow a feed.
- **`gator following`**: Print the list of follows.
- **`gator unfollow <url>`**: Unfollow a feed.
- **`gator browse [limit]`**: Print the list of posts.
