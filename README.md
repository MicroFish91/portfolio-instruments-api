# Portfolio Instruments API

The Portfolio Instruments API is a Go-based REST API designed to streamline the management of passive investment portfolios. It allows users to input portfolio data and receive rebalancing guidelines based on user-defined criteria.

### Key Features:
* 📷 <u>Snapshot Tracking</u>: Capture and monitor portfolio snapshots at different points in time.
* 📁 <u>Benchmark Portfolios</u>: Define and reference stable, pre-selected portfolios for easier portfolio management.
* 📈 <u>Rebalance Calculations</u>: Automatically generate rebalancing recommendations based on your chosen benchmark portfolio.

This API automates the process of calculating portfolio rebalances, eliminating the need for manual spreadsheet tracking and providing a more efficient one-stop portfolio solution.

This project is mostly for my personal use, however, you are free to use it as well if you wish.

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

### 🏁 Starting the Server
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