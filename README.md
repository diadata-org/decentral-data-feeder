# decentral-data-feeder

<p align="center">
    <img src="./assets/DIA_logo.png" alt="Dia logo" width="200" height="auto" style="padding: 20px;">
</p>

The node setup instructions are available in our [Wiki](https://github.com/diadata-org/decentral-data-feeder/wiki) page!

This repository hosts a self-contained containerized application for running a data feeder in the Lumina oracle network. It comprises of a scraper that collects data from various sources including Real-World Asset data (TwelveData for Stocks, Forex, Commodities, ETFs), and verifiable random numbers (Randamu).

## Requirements

- Docker or Docker Compose installed

- Container has minimal resource requirements, making it suitable for most machines, including Windows, macOS, Linux, and Raspberry Pi, as long as Docker is installed.

- A private key from MetaMask or any other EVM wallet. Alternatively to generate a private key effortlesly, [eth-keys](https://github.com/ethereum/eth-keys) tool can be used for this.

- DIA tokens in your wallet. For running a node in testnet, you can request tokens from the [faucet](https://faucet.diadata.org).

## Quick Start

> **NOTE:** This guide is using docker compose for running a feeder node. You can explore alternative deployment methods [here](https://github.com/diadata-org/decentral-feeder/wiki/Node-deployment#alternative-deployment-methods).

### Navigate to the Docker Compose Folder

- Locate the `docker-compose` folder in this repository.
- Inside, you will find a file named `docker-compose.yaml`.

### Configure Environment Variables

- Create a `.env` file in the same directory as `docker-compose.yaml`. This file should contain the following variables:

```
SOURCE=TwelveData

# Node operator configuration (MUST be filled by node operators)
PRIVATE_KEY=
NODE_OPERATOR_NAME=""
DEPLOYED_CONTRACT=""
BLOCKCHAIN_NODE=""
BACKUP_NODE=""
IMAGE_TAG=""

# Monitoring
PUSHGATEWAY_USER=
PUSHGATEWAY_PASSWORD=
PUSHGATEWAY_URL=

# Chain configuration
# https://github.com/diadata-org/decentral-data-feeder/wiki/Chain-Info
CHAIN_ID=
UPDATE_SECONDS=80

# Asset symbols to track
STOCK_SYMBOLS=
FX_TICKERS=
COMMODITIES=
ETF=
```

Please refer to the [.env.example](./docker-compose/.env.example) file for the exact environment variables to set.

- Open a terminal in the `docker-compose` folder and start the deployment by running:
  ```bash
  docker-compose up
  ```

### Retrieve Deployed Contract

- Once the container is deployed with `DEPLOYED_CONTRACT` env variable empty the logs will display the deployed contract address in the following format:
  ```plaintext
  │ time="2024-11-25T11:30:08Z" level=info msg="Contract pending deploy: 0xxxxxxxxxxxxxxxxxxxxxxxxxx."
  ```
- Copy the displayed contract address (e.g., `0xxxxxxxxxxxxxxxxxxxxxxxxxx`) and stop the container with `docker rm -f <container_name>`.

- Update your `.env` file with `DEPLOYED_CONTRACT` variable mentioned above. Redeployed the container with `docker-compose up -d`

  ```plaintext
  DEPLOYED_CONTRACT=0xxxxxxxxxxxxxxxxxxxxxxxxxx
  ```

- Check if the container is running correctly by viewing the logs. Run the following command:

  ```bash
  docker-compose logs -f
  ```

- Expected Logs: Look for logs similar to the example below, which indicate a successful startup:

```
app-1  | time="2025-08-19T05:59:23Z" level=info msg="Executing ETF data update for 19 symbols"
app-1  | time="2025-08-19T05:59:24Z" level=info msg="got rwa data: {URTH iShares MSCI World ETF 175.6 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
app-1  | time="2025-08-19T05:59:24Z" level=info msg="got rwa data: {EEM iShares MSCI Emerging Markets ETF 50.235 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
app-1  | time="2025-08-19T05:59:24Z" level=info msg="got rwa data: {VOO Vanguard S&P 500 ETF 591.44 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
app-1  | time="2025-08-19T05:59:25Z" level=info msg="got rwa data: {SPY SPDR S&P 500 ETF Trust 643.2901 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
app-1  | time="2025-08-19T05:59:25Z" level=info msg="got rwa data: {VTI Vanguard Total Stock Market ETF 316.54001 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
app-1  | time="2025-08-19T05:59:25Z" level=info msg="got rwa data: {QQQ Invesco QQQ Trust 577.099976 2025-08-18 19:59:00 +0000 UTC false false ETF TwelveData}"
.
.
```

- You can optionally cleanup the deployment once you're done by running:

  ```
  docker rm -f <container_name>
  ```

- Verify the container has been removed:
  ```
  docker ps -a
  ```

## Error Handling

> **NOTE:** This guide is specific to docker compose. You can check how to handle errors based on your deployment method [here](https://github.com/diadata-org/decentral-feeder/wiki/Error-Handling).

If any issues arise during deployment, follow these steps:

#### Check Logs:

- Run `docker-compose logs -f`

#### Verify Environment Variables:

- Ensure all required variables (`PRIVATE_KEY`, `DEPLOYED_CONTRACT`,...etc) are correctly set by checking the `.env` file.

#### Restart Deployment:

- ```bash
  docker-compose down && docker-compose up -d
  ```

#### Update or Rebuild:

- Ensure you're using the correct image version:
  ```bash
  docker pull diadata/decentralized-data-feeder:<VERSION>
  ```
- Apply fixes and redeploy.

## Documentation

For the full node deployment instructions, you can visit our [Wiki](https://github.com/diadata-org/decentral-data-feeder/wiki) page.

To learn about DIA's oracle stacks, you can visit our documentation [here](https://docs.diadata.org/).

## Issues

To report bugs or suggest enhancements, you can create a [Github Issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/using-issues/creating-an-issue) in the repository.

## Contribution Guidelines

Coming soon...

## Community

You can find our team on the following channels:

- [Discord](https://discord.com/invite/RjHBcZ9mEH)
- [Telegram](https://t.me/diadata_org)
- [X](https://x.com/DIAdata_org)
