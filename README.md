# Loggr

Loggr is a simple SSR CRUD website built with Go and HTMX. It serves as a tool for gardeners to manage their gardens, plants, tools, and seeds online. Users can sign up, create gardens, and add items to them, allowing for easy visualization and storage of metadata.

## Features

- **User Authentication**: Secure user authentication system allows users to sign up, log in, and manage their accounts.
- **Garden Management**: Users can create gardens, edit garden details, and delete gardens as needed.
- **Item Tracking**: Gardens can contain items such as plants, tools, and seeds, allowing users to store metadata about each item.
- **CRUD Functionality**: Users can perform CRUD operations on gardens and items, enabling easy management of gardening data.
- **Server-Side Rendering (SSR)**: The application utilizes server-side rendering for improved performance and SEO.
- **In-Memory Storage**: The application currently uses in-memory storage for simplicity, with plans to integrate a backend for persistent storage.

## Technologies Used

- **Go**: Backend development is done entirely in Go, providing a robust and efficient server-side framework.
- **HTMX**: Frontend development is powered by HTMX for seamless, AJAX-driven interactions without the need for a complex frontend framework.
- **HTML/CSS/JavaScript**: Standard web technologies are used for frontend development.
- **In-Memory Storage**: Data is currently stored in memory for simplicity, with plans to integrate backend storage (Turso) in the future.

## Getting Started

To get started with Loggr, follow these steps:

1. **Clone the Repository**: `git clone https://github.com/linkinlog/loggr.git`
2. **Copy and edit the configuration file**: `cp .env.sample .env`
3. **Install Dependencies**: No external dependencies are required.
4. **Run the Application**: Make commands are provided for building and running the application:
   - `make build`: Builds the application.
   - `make dev`: Runs the application.
   - `make clean`: Cleans up the build directory.
5. **Access the Application**: Open your web browser and navigate to `http://localhost` at the port specified in `.env` to access Loggr.

## Contributing

Contributions are welcome! If you'd like to contribute to Loggr, please fork the repository, make your changes, and submit a pull request. Make sure to follow the existing coding style and conventions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Special thanks to the developers of [HTMX](https://htmx.org) for providing a powerful tool for building AJAX-driven web applications without the complexity of frontend frameworks.
- Hat tip to all the contributors who have helped improve Loggr through their feedback and contributions.
