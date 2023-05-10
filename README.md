# Project Name

Take Home Assignment paretolabs

## Getting Started

### Backend

1. To start the backend server, navigate to the `/be` folder and run the following command: `go run .`
2. The first time you start the server, a new file with the user data list is created inside the db folder.
3. If the file exists, the user data is not fetched again.

### Frontend

To start the frontend, navigate to the `/fe` folder and follow these steps:

1. Install `pnpm` if you haven't already. You can install it by running the following command:

```shell
npm install -g pnpm
```
2. Install the project dependencies by running the following command: `pnpm install`
3. Start the development server by running the following command: `pnpm run dev`
4. The frontend project uses the URL to display the profile information. `/users/:pid`
5. When accessing the root URL ("/"), the application will automatically redirect to the profile page of the first user ("/users/1").
6. It also implements an infinite scrolling pagination strategy to load the list of users in the sidebar.
