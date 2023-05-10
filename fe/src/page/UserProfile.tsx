import Avatar from "@mui/material/Avatar";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import { LoaderFunction, LoaderFunctionArgs, json, useLoaderData } from "react-router-dom";
import ProfileTable from "../components/ProfileTable";
import MainLayout from "../templates/MainLayout";


export const UserProfileLoader: LoaderFunction = async ({ params }: LoaderFunctionArgs ) => {
  let results;

  try {
    const response = await fetch(`/apiURL/users/${params.name}`);
    results = await response.json();
  } catch (e) {
    return json(
      { data: 'error' },
      { status: 400 }
    );
  }

  return json(
    {
      username: results.profileData.profileInformation.username,
      bio: results.profileData.profileInformation.username,
      address: results.profileData.address,
      avatar: results.profileData.profileInformation.avatarUrl,
      displayName: results.profileData.profileInformation.displayName,
      followers: results.profileData.profileInformation.followers,
      following: results.profileData.profileInformation.following,
    },
    { status: 200 }
  );
};

export interface Row {
  field: string;
  value: string;
}

interface UserProfileData {
  username: string;
  bio: string;
  address: string;
  avatar: string;
  displayName: string;
  followers: string;
  following: string;
}

function createData(field: string, value: string): Row {
  return { field, value };
}

function UserProfile(): JSX.Element {
  const data = useLoaderData() as UserProfileData;

  const rows = [
    createData("Username", data?.username),
    createData("Bio", data?.bio),
    createData("Address", data?.address),
    createData("# Following", data?.following),
    createData("# Followers", data?.followers),
  ];

  return (
    <MainLayout title="Profile Information">
      <Stack direction="row" spacing={2}>
        <Avatar alt="Remy Sharp" src={data?.avatar} />
        <Typography variant="h4" paddingBottom={2}>
          {data?.displayName}
        </Typography>
      </Stack>
      <ProfileTable rows={rows} />
    </MainLayout>
  );
}

export default UserProfile;
