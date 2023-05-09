import { json, useLoaderData } from "react-router-dom";
import MainLayout from "../templates/MainLayout";
import ProfileTable from "../components/ProfileTable";
import Avatar from "@mui/material/Avatar";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

export const UserProfileLoader = async ({ params }) => {
  let results;

  try {
    const response = await fetch(`/apiURL/users/${params.name}`);
    results = await response.json();
  } catch (e) {
    return json(
      { data: e.response.data, statusText: e.statusText },
      { status: e.response.status }
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

function createData(field: string, value: string) {
  return { field, value };
}

function UserProfile() {
  const data = useLoaderData();

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
