import {
  createBrowserRouter,
  Navigate
} from "react-router-dom";
import UserProfile, { UserProfileLoader } from "./page/UserProfile";

export const router = createBrowserRouter([
  {
    path: "/users/:name",
    loader: UserProfileLoader,
    element: <UserProfile />,
  },
  {
    path: "*",
    element: <Navigate to="/users/1" replace />
  },
]);