# <img src="resources/pi_logo.png" width="35" /> Portfolio Instruments API

Portfolio Instruments API is a Go-based REST API built to simplify the management of passive investment portfolios. It enables users to input portfolio data and receive automated rebalancing recommendations based on benchmark criteria, eliminating the need for manual spreadsheet tracking. This API streamlines the process of calculating rebalancing requirements, offering a comprehensive solution for portfolio management. Additionally, users can efficiently query their accounts and assets to gain insights into their holdings across different tax shelters, financial institutions, liquidity levels, and more.

### Key Features:
* 📷 <u>Snapshot Tracking</u>: Capture and monitor portfolio snapshots at different points in time.
* 📁 <u>Benchmark Portfolios</u>: Provide portfolio benchmarks to monitor your portfolio snapshots against.
* 📈 <u>Rebalance Calculations</u>: Automatically generate rebalancing recommendations based on your chosen benchmarks.

<b>Note</b>: This project is mostly for my personal use, however, you are free to use it as well if you find it to be useful!

## 📖 Getting Started

### 🛠 ️Local Database Setup
1. Spin up a local Postgres database container by running:
    > make config-up

2. Apply database migrations:
    > make migrate-up

    #### Note:
    * To rollback migrations, run:
        > make migrate-down

    * To stop the database container, use:
        > make config-down

### 💻 Starting the Server
To get the API up and running, simply execute:
> make run

You can test the API routes using an API client like `Postman` or `ThunderClient`. To quickly test endpoints using `ThunderClient`, follow the steps outlined [here](docs/ThunderClient.md).

### 🏃 Running Tests
To run the test suite, use:
> make test

The integration tests will automatically run using a Postgres test container, so no additional configuration is needed.

### 🚀 Deployment
Before deploying, ensure that you've set the correct production environment variables. Example `.env` files are included at the project root.

For deployment:

* 🐳 Docker: A Dockerfile is provided, ensuring a simple and consistent containerized deployment process.
* ️☁️ Azure: If deploying to Azure Container Apps (my preferred cloud solution), the VS Code extension makes deployment fast and efficient. A more detailed guide for this process will be provided soon.

## ✏️ Other API Docs

* [Routes](docs/Swagger.md)
* [Database Design](docs/DatabaseDesign.md)
* [Migrations](docs/Migrations.md)
* [ThunderClient](docs/ThunderClient.md)

## 🔗 Useful Links

* [Modern Portfolio Theory](https://en.wikipedia.org/wiki/Modern_portfolio_theory)
* [Portfolio Charts](https://portfoliocharts.com/)